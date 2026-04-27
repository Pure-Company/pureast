// cmd/pureast/commands/deps.go
//
// Dependency analysis for a single symbol.
//
// One verb, one knob: --format chooses the rendering. The old --report and
// --dot booleans were two ways of asking the same question ("how should I
// see these deps") and got mutually-exclusive flags slapped on top, which
// is the textbook redundant-path smell. Replaced with --format text|dot|json.
package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"go/token"
	"sort"

	"github.com/spf13/cobra"
	"github.com/vinodhalaharvi/pureast/pkg/analyze"
	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/pureast/pkg/cli"
	"github.com/vinodhalaharvi/pureast/pkg/codegen"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
)

type DepsArgs struct {
	FilePath string
	Symbol   string
	Format   string // text|dot|json
	Minimal  bool
	Depth    int  // -1 = unbounded (equivalent to current full transitive)
	Reverse  bool // who uses this symbol, instead of what it uses
}

func NewDepsCommand() *cobra.Command {
	cmd := cli.NewCommand[DepsArgs]("deps").
		Short("Analyze dependencies for a symbol").
		Long(`Show what a symbol depends on (or who depends on it, with --reverse).
Output is plain text by default; use --format dot for Graphviz, --format json
for machine consumption.

Examples:
  pureast deps User ./pkg
  pureast deps UserService ./pkg --minimal
  pureast deps Profile ./pkg --depth 1            # one-hop only
  pureast deps Profile ./pkg --reverse            # who uses Profile
  pureast deps Profile ./pkg --reverse --depth 1  # direct callers only
  pureast deps Profile ./pkg --format dot > deps.dot
  dot -Tpng deps.dot -o deps.png`).
		ParseArgs(parseDepsArgs).
		Action(depsAction).
		Build()

	cmd.Flags().String("format", "text", "Output format: text|dot|json")
	cmd.Flags().Bool("minimal", false, "Show only direct (non-transitive) dependencies")
	cmd.Flags().Int("depth", -1, "Max traversal depth (-1 = unbounded, 0 = direct only)")
	cmd.Flags().Bool("reverse", false, "Show who depends on the symbol, not what it depends on")

	// Back-compat: --file kept as deprecated alias for positional PATH
	cmd.Flags().StringP("file", "f", "", "[deprecated] use positional PATH")

	return cmd
}

func parseDepsArgs(cmd *cobra.Command, args []string) result.Result[DepsArgs] {
	if len(args) < 1 {
		return result.Err[DepsArgs](fmt.Errorf("requires SYMBOL [PATH]"))
	}
	if len(args) > 2 {
		return result.Err[DepsArgs](fmt.Errorf("expected SYMBOL [PATH], got %d args", len(args)))
	}

	symbol := args[0]
	path, err := resolvePathFromTail(cmd, args[1:])
	if err != nil {
		return result.Err[DepsArgs](err)
	}

	format, _ := cmd.Flags().GetString("format")
	minimal, _ := cmd.Flags().GetBool("minimal")
	depth, _ := cmd.Flags().GetInt("depth")
	reverse, _ := cmd.Flags().GetBool("reverse")

	switch format {
	case "text", "dot", "json":
	default:
		return result.Err[DepsArgs](fmt.Errorf(
			"invalid --format %q (want: text|dot|json)", format))
	}

	// --minimal and --depth are two ways of saying "less expansion".
	// Rather than silently letting one override the other we reject
	// the combination — explicit beats implicit, in the spirit of
	// "no redundant paths."
	if minimal && depth >= 0 {
		return result.Err[DepsArgs](fmt.Errorf(
			"--minimal and --depth are mutually exclusive"))
	}

	return result.Ok(DepsArgs{
		FilePath: path,
		Symbol:   symbol,
		Format:   format,
		Minimal:  minimal,
		Depth:    depth,
		Reverse:  reverse,
	})
}

func depsAction(ctx context.Context, args DepsArgs) result.Result[cli.Output] {
	fset := token.NewFileSet()
	pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
	if err != nil {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("Error: %v\n", err),
			ExitCode: 1,
		})
	}

	declMap := extract.BuildPackageDeclMap(pkgNode)
	graph := analyze.NewDependencyGraph(declMap)

	// Pick the dependency-resolution strategy based on flags. The four
	// branches are mutually exclusive by construction (we rejected
	// --minimal + --depth in parsing), so we don't need a precedence
	// table — the conditions don't overlap:
	//
	//   --reverse + --depth=0  → Users (one hop reverse)
	//   --reverse              → UsersTransitive (full reverse)
	//   --depth=N (N>=0)       → ResolveBounded (N hops forward)
	//   --minimal              → MinimalDependencies
	//   default                → ResolveWithAssociatedCode
	deps := selectDeps(graph, args)

	switch args.Format {
	case "dot":
		// DOT output uses the existing generator which walks the
		// forward graph; --reverse and --depth are honored only by
		// the text/json paths for now. Note this in a comment so
		// we don't pretend we silently support it.
		gen := codegen.NewGenerator(fset)
		return result.Ok(cli.Output{
			Text:     gen.GenerateDOT(args.Symbol, declMap),
			ExitCode: 0,
		})

	case "json":
		return result.Ok(cli.Output{
			Text:     formatDepsJSON(args.Symbol, deps),
			ExitCode: 0,
		})

	default: // text
		header := "Dependencies for " + args.Symbol + ":"
		if args.Reverse {
			header = "Reverse dependencies (users) of " + args.Symbol + ":"
		}
		out := header + "\n\n" + analyze.FormatDependencies(args.Symbol, deps)
		if !args.Reverse && args.Depth < 0 && !args.Minimal {
			// Stats only make sense for the unbounded forward query —
			// for bounded or reverse queries the numbers don't refer
			// to the data we just rendered.
			stats := graph.ComputeStats(args.Symbol)
			out += fmt.Sprintf("\nMax Depth: %d\n", stats.MaxDepth)
		}
		return result.Ok(cli.Output{Text: out, ExitCode: 0})
	}
}

// selectDeps centralizes the strategy decision so depsAction reads as
// "compute deps, then format them" rather than mixing the two concerns.
func selectDeps(g analyze.DependencyGraph, args DepsArgs) astpkg.Dependencies {
	switch {
	case args.Reverse && args.Depth == 0:
		return g.Users(args.Symbol)
	case args.Reverse:
		return g.UsersTransitive(args.Symbol)
	case args.Depth >= 0:
		return g.ResolveBounded(args.Symbol, args.Depth)
	case args.Minimal:
		return g.MinimalDependencies(args.Symbol)
	default:
		return g.ResolveWithAssociatedCode(args.Symbol)
	}
}

// formatDepsJSON renders dependencies as a stable, structured JSON document.
// We expose the fields a downstream tool actually wants: the symbol being
// analyzed, plus sorted lists of related symbol names by category.
// Stable ordering (alphabetical within each list) is essential for LLM
// caching — identical input must produce byte-identical output.
func formatDepsJSON(symbol string, deps astpkg.Dependencies) string {
	payload := struct {
		Symbol     string   `json:"symbol"`
		Types      []string `json:"types"`
		Functions  []string `json:"functions"`
		Structs    []string `json:"structs"`
		Interfaces []string `json:"interfaces"`
		Constants  []string `json:"constants"`
		Variables  []string `json:"variables"`
		Imports    []string `json:"imports"`
	}{
		Symbol:     symbol,
		Types:      sortedSlice(deps.Types.ToSlice()),
		Functions:  sortedSlice(deps.Functions.ToSlice()),
		Structs:    sortedSlice(deps.Structs.ToSlice()),
		Interfaces: sortedSlice(deps.Interfaces.ToSlice()),
		Constants:  sortedSlice(deps.Constants.ToSlice()),
		Variables:  sortedSlice(deps.Variables.ToSlice()),
		Imports:    sortedSlice(deps.Imports.ToSlice()),
	}

	out, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"symbol":%q,"error":%q}`+"\n", symbol, err.Error())
	}
	return string(out) + "\n"
}

func sortedSlice(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	sort.Strings(out)
	return out
}
