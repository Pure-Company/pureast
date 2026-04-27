// cmd/pureast/commands/diff.go
//
// Extract symbols in files that have changed since a given git ref.
// Intended use: PR-review and "what's new" LLM context. Instead of
// dumping the whole package, dump only what touched code in this
// branch.
//
// Strategy: shell out to `git diff --name-only <ref> HEAD` to find
// changed Go files, then run the same symbol collection as `dump`
// against just those files (or restricted to their content).
// Symbols outside changed files are excluded entirely.
//
// Limitations to call out:
//   - We do file-level granularity, not line-level. A symbol in a
//     changed file is included even if it wasn't itself modified.
//     Line-level filtering would need diff parsing and AST line-range
//     intersection — useful follow-up, not in this first cut.
//   - We don't handle deleted files: a deleted symbol can't appear
//     in a dump because the AST no longer contains it. That's
//     probably the right behavior for context generation, but would
//     matter for review summaries.

package commands

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinodhalaharvi/pureast/pkg/cli"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
)

type DiffArgs struct {
	FilePath   string
	Ref        string
	OutputFile string
	Format     string // go|md
	Bodies     bool
	MaxTokens  int
}

func NewDiffCommand() *cobra.Command {
	cmd := cli.NewCommand[DiffArgs]("diff").
		Short("Dump symbols from files changed since a git ref").
		Long(`Extract every symbol in Go files that have changed since the given git ref.
The intended workflow is PR review: feed an LLM only the code that's new in
this branch, not the entire repo.

The ref can be any git revision (branch, tag, commit, HEAD~N).

Examples:
  pureast diff main
  pureast diff main ./pkg
  pureast diff HEAD~5 --bodies
  pureast diff origin/main --format md -o pr-context.md`).
		ParseArgs(parseDiffArgs).
		Action(diffAction).
		Build()

	cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
	cmd.Flags().String("format", "go", "Output format: go|md")
	cmd.Flags().Bool("bodies", false, "Include function bodies")
	cmd.Flags().Int("max-tokens", 0, "Truncate output to fit token budget (0 = unbounded)")

	return cmd
}

func parseDiffArgs(cmd *cobra.Command, args []string) result.Result[DiffArgs] {
	if len(args) < 1 {
		return result.Err[DiffArgs](fmt.Errorf("requires REF [PATH]"))
	}
	if len(args) > 2 {
		return result.Err[DiffArgs](fmt.Errorf("expected REF [PATH], got %d args", len(args)))
	}

	ref := args[0]
	path := "."
	if len(args) == 2 {
		path = args[1]
	}

	output, _ := cmd.Flags().GetString("output")
	format, _ := cmd.Flags().GetString("format")
	bodies, _ := cmd.Flags().GetBool("bodies")
	maxTokens, _ := cmd.Flags().GetInt("max-tokens")

	if format != "go" && format != "md" {
		return result.Err[DiffArgs](fmt.Errorf(
			"invalid --format %q (want: go|md)", format))
	}

	return result.Ok(DiffArgs{
		FilePath:   path,
		Ref:        ref,
		OutputFile: output,
		Format:     format,
		Bodies:     bodies,
		MaxTokens:  maxTokens,
	})
}

func diffAction(ctx context.Context, args DiffArgs) result.Result[cli.Output] {
	changed, err := changedGoFiles(ctx, args.Ref, args.FilePath)
	if err != nil {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("Error: %v\n", err),
			ExitCode: 1,
		})
	}
	if len(changed) == 0 {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("No Go files changed since %s.\n", args.Ref),
			ExitCode: 0,
		})
	}

	// Reuse the dump symbol collector, but restricted to the changed
	// files only. This keeps formatting consistent with `dump` so
	// users don't have to learn a second output style.
	symbols, pkgName, err := collectSymbolsFromFiles(changed, args.Bodies)
	if err != nil {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("Error: %v\n", err),
			ExitCode: 1,
		})
	}

	out := renderDiffOutput(pkgName, args.Ref, changed, symbols, args.Bodies)

	if args.MaxTokens > 0 {
		out, _ = truncateToBudget(out, args.MaxTokens)
	}
	if args.Format == "md" {
		title := fmt.Sprintf("Changes since %s", args.Ref)
		out = renderAsMarkdown(title, out)
	}

	if args.OutputFile != "" {
		if err := os.WriteFile(args.OutputFile, []byte(out), 0644); err != nil {
			return result.Ok(cli.Output{
				Text:     fmt.Sprintf("Error writing file: %v\n", err),
				ExitCode: 1,
			})
		}
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("✓ Written to %s\n", args.OutputFile),
			ExitCode: 0,
		})
	}

	return result.Ok(cli.Output{Text: out, ExitCode: 0})
}

// changedGoFiles asks git which files differ between ref and HEAD,
// then narrows the result to .go files inside the requested path.
// We use --name-only on the diff so we don't have to parse the patch.
func changedGoFiles(ctx context.Context, ref, root string) ([]string, error) {
	cmd := exec.CommandContext(ctx, "git", "diff", "--name-only", ref, "HEAD")
	cmd.Dir = root
	output, err := cmd.CombinedOutput()
	if err != nil {
		// git's stderr is often the most informative thing in this
		// failure mode (unknown ref, not a git repo, etc).
		return nil, fmt.Errorf("git diff failed: %s", strings.TrimSpace(string(output)))
	}

	var files []string
	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || !strings.HasSuffix(line, ".go") {
			continue
		}
		// git diff returns paths relative to the repo root, which is
		// what cmd.Dir was set to. Resolve to a real path under root.
		full := filepath.Join(root, line)
		if _, err := os.Stat(full); err == nil {
			files = append(files, full)
		}
	}
	sort.Strings(files)
	return files, nil
}

// collectSymbolsFromFiles parses each given file and yields the same
// dumpedSymbol shape that `dump` uses, so renderDump can format the
// result. We don't reuse collectSymbols directly because it walks a
// directory; here we have an explicit file list.
func collectSymbolsFromFiles(paths []string, includeBodies bool) ([]dumpedSymbol, string, error) {
	args := DumpArgs{
		Bodies:       includeBodies,
		Kind:         "all",
		IncludeDocs:  true,
		IncludeTests: true, // diff respects whatever git reports; user already opted in
	}

	var (
		symbols []dumpedSymbol
		pkgName string
	)

	for _, path := range paths {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			// Skip unparseable files rather than aborting the whole
			// diff — a syntax error elsewhere shouldn't prevent
			// surfacing changes that did parse.
			continue
		}
		if pkgName == "" {
			pkgName = file.Name.Name
		}

		base := filepath.Base(path)
		for _, decl := range file.Decls {
			ss := extractFromDecl(fset, file, decl, base, args)
			symbols = append(symbols, ss...)
		}
	}

	sort.Slice(symbols, func(i, j int) bool {
		if symbols[i].File != symbols[j].File {
			return symbols[i].File < symbols[j].File
		}
		return symbols[i].Line < symbols[j].Line
	})

	return symbols, pkgName, nil
}

func renderDiffOutput(pkgName, ref string, files []string, symbols []dumpedSymbol, bodies bool) string {
	var b strings.Builder
	fmt.Fprintf(&b, "// pureast diff: package %s — changes since %s\n", pkgName, ref)
	fmt.Fprintf(&b, "// %d changed file(s), %d symbol(s)", len(files), len(symbols))
	if !bodies {
		b.WriteString(" (signatures only)")
	}
	b.WriteString("\n\n")

	for _, f := range files {
		fmt.Fprintf(&b, "// changed: %s\n", f)
	}
	b.WriteString("\n")

	// Group by kind, same convention as `dump`, so the output is
	// recognizable to anyone familiar with that command.
	groups := map[string][]dumpedSymbol{}
	order := []string{"struct", "interface", "type", "func", "method", "const", "var"}
	for _, s := range symbols {
		groups[s.Kind] = append(groups[s.Kind], s)
	}
	headings := map[string]string{
		"struct":    "// === structs ===",
		"interface": "// === interfaces ===",
		"type":      "// === type aliases ===",
		"func":      "// === functions ===",
		"method":    "// === methods ===",
		"const":     "// === constants ===",
		"var":       "// === variables ===",
	}
	for _, kind := range order {
		ss := groups[kind]
		if len(ss) == 0 {
			continue
		}
		b.WriteString(headings[kind])
		b.WriteString("\n\n")
		for _, s := range ss {
			docEmittedBySource := bodies && (s.Kind == "func" || s.Kind == "method")
			if s.Doc != "" && !docEmittedBySource {
				for _, line := range strings.Split(strings.TrimRight(s.Doc, "\n"), "\n") {
					b.WriteString("// ")
					b.WriteString(line)
					b.WriteString("\n")
				}
			}
			b.WriteString(s.Source)
			b.WriteString("\n\n")
		}
	}

	return b.String()
}

// _ keeps go/ast imported for symmetry with collectSymbols even if
// extractFromDecl is the only consumer of the package within this file.
// Without this the linter complains; deleting the import means a future
// edit needs to re-add it.
var _ = ast.NewIdent
