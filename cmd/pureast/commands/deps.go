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
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/Pure-Company/pureast/pkg/analyze"
	astpkg "github.com/Pure-Company/pureast/pkg/ast"
	"github.com/Pure-Company/pureast/pkg/cli"
	"github.com/Pure-Company/pureast/pkg/codegen"
	"github.com/Pure-Company/pureast/pkg/extract"
)

type DepsArgs struct {
	FilePath  string
	Symbol    string
	Format    string // text|dot|json
	Minimal   bool
	Depth     int  // -1 = unbounded (equivalent to current full transitive)
	Reverse   bool // who uses this symbol, instead of what it uses
	Locations bool // include file:line for each dep entry (always on for --reverse)
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
  pureast deps Profile ./pkg --locations          # show file:line for each dep
  pureast deps Profile ./pkg --format dot > deps.dot
  dot -Tpng deps.dot -o deps.png`).
		ParseArgs(parseDepsArgs).
		Action(depsAction).
		Build()

	cmd.Flags().String("format", "text", "Output format: text|dot|json")
	cmd.Flags().Bool("minimal", false, "Show only direct (non-transitive) dependencies")
	cmd.Flags().Int("depth", -1, "Max traversal depth (-1 = unbounded, 0 = direct only)")
	cmd.Flags().Bool("reverse", false, "Show who depends on the symbol, not what it depends on")
	cmd.Flags().Bool("locations", false,
		"Include file:line for each dep entry. Always on for --reverse "+
			"(the use case where it's most useful); opt-in for forward.")

	// Back-compat: --file kept as deprecated alias for positional PATH
	cmd.Flags().StringP("file", "f", "", "[deprecated] use positional PATH")

	return cmd
}

func parseDepsArgs(cmd *cobra.Command, args []string) (DepsArgs, error) {
	if len(args) < 1 {
		return DepsArgs{}, fmt.Errorf("requires SYMBOL [PATH]")
	}
	if len(args) > 2 {
		return DepsArgs{}, fmt.Errorf("expected SYMBOL [PATH], got %d args", len(args))
	}

	symbol := args[0]
	path, err := resolvePathFromTail(cmd, args[1:])
	if err != nil {
		return DepsArgs{}, err
	}

	format, _ := cmd.Flags().GetString("format")
	minimal, _ := cmd.Flags().GetBool("minimal")
	depth, _ := cmd.Flags().GetInt("depth")
	reverse, _ := cmd.Flags().GetBool("reverse")
	locations, _ := cmd.Flags().GetBool("locations")

	switch format {
	case "text", "dot", "json":
	default:
		return DepsArgs{}, fmt.Errorf(
			"invalid --format %q (want: text|dot|json)", format)
	}

	// --minimal and --depth are two ways of saying "less expansion".
	// Rather than silently letting one override the other we reject
	// the combination — explicit beats implicit, in the spirit of
	// "no redundant paths."
	if minimal && depth >= 0 {
		return DepsArgs{}, fmt.Errorf(
			"--minimal and --depth are mutually exclusive")
	}

	return DepsArgs{
		FilePath:  path,
		Symbol:    symbol,
		Format:    format,
		Minimal:   minimal,
		Depth:     depth,
		Reverse:   reverse,
		Locations: locations,
	}, nil
}

func depsAction(ctx context.Context, args DepsArgs) (cli.Output, error) {
	fset := token.NewFileSet()
	pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
	if err != nil {
		return cli.Output{}, fmt.Errorf("extract %s: %w", args.FilePath, err)
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

	// Strip parser-leak noise from the Functions set. The forward-dep
	// extractor records every Ident/SelectorExpr it walks — including
	// receiver-variable references like "p" or "p.Address" — because
	// it lacks type-resolution context. CleanDependencies removes
	// those by intersecting against declMap. See pkg/analyze/clean.go.
	deps = analyze.CleanDependencies(deps, declMap)

	switch args.Format {
	case "dot":
		// DOT output uses the existing generator which walks the
		// forward graph; --reverse and --depth are honored only by
		// the text/json paths for now. Note this in a comment so
		// we don't pretend we silently support it.
		gen := codegen.NewGenerator(fset)
		return cli.Output{
			Text:     gen.GenerateDOT(args.Symbol, declMap),
			ExitCode: 0,
		}, nil

	case "json":
		// Locations are emitted for --reverse (where they're most useful —
		// "who calls X" wants jump-to-source) and opt-in for forward via
		// --locations. The two together cover the cases where file:line
		// adds value without changing default forward output.
		withLocations := args.Reverse || args.Locations
		return cli.Output{
			Text:     formatDepsJSON(args.Symbol, deps, withLocations, fset, declMap, args.FilePath),
			ExitCode: 0,
		}, nil

	default: // text
		header := "Dependencies for " + args.Symbol + ":"
		if args.Reverse {
			header = "Reverse dependencies (users) of " + args.Symbol + ":"
		}
		var out string
		withLocations := args.Reverse || args.Locations
		if withLocations {
			// File:line is the headline information for "who calls X" and
			// is opt-in valuable for forward deps too (e.g. when the user
			// wants to know where a transitive dep is declared).
			out = header + "\n\n" + formatDepsWithLocations(deps, fset, declMap, args.FilePath)
		} else {
			out = header + "\n\n" + analyze.FormatDependencies(args.Symbol, deps)
		}
		if !args.Reverse && args.Depth < 0 && !args.Minimal {
			// Stats only make sense for the unbounded forward query —
			// for bounded or reverse queries the numbers don't refer
			// to the data we just rendered.
			stats := graph.ComputeStats(args.Symbol)
			out += fmt.Sprintf("\nMax Depth: %d\n", stats.MaxDepth)
		}
		return cli.Output{Text: out, ExitCode: 0}, nil
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
// When withLocations is true (typically for --reverse), each entry becomes
// an object with name + file + line instead of a bare string. Stable
// ordering (alphabetical) is essential for LLM caching — identical input
// must produce byte-identical output.
func formatDepsJSON(symbol string, deps astpkg.Dependencies, withLocations bool, fset *token.FileSet, declMap map[string]astpkg.DeclNode, basePath string) string {
	if !withLocations {
		// Original flat-string shape preserved for forward-deps callers
		// that already parse this format.
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

	// With-locations variant: each name expands into {name, file, line}.
	// Imports stay as plain strings — they don't have a position in
	// our package's source.
	payload := struct {
		Symbol     string          `json:"symbol"`
		Types      []symbolWithLoc `json:"types"`
		Functions  []symbolWithLoc `json:"functions"`
		Structs    []symbolWithLoc `json:"structs"`
		Interfaces []symbolWithLoc `json:"interfaces"`
		Constants  []symbolWithLoc `json:"constants"`
		Variables  []symbolWithLoc `json:"variables"`
		Imports    []string        `json:"imports"`
	}{
		Symbol:     symbol,
		Types:      withLoc(deps.Types.ToSlice(), fset, declMap, basePath),
		Functions:  withLoc(deps.Functions.ToSlice(), fset, declMap, basePath),
		Structs:    withLoc(deps.Structs.ToSlice(), fset, declMap, basePath),
		Interfaces: withLoc(deps.Interfaces.ToSlice(), fset, declMap, basePath),
		Constants:  withLoc(deps.Constants.ToSlice(), fset, declMap, basePath),
		Variables:  withLoc(deps.Variables.ToSlice(), fset, declMap, basePath),
		Imports:    sortedSlice(deps.Imports.ToSlice()),
	}
	out, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"symbol":%q,"error":%q}`+"\n", symbol, err.Error())
	}
	return string(out) + "\n"
}

type symbolWithLoc struct {
	Name string `json:"name"`
	File string `json:"file,omitempty"`
	Line int    `json:"line,omitempty"`
}

// withLoc looks up file:line for each name in declMap. Names with no
// matching declaration (e.g. cross-package references) are kept with
// empty file/line — better to show the name than drop the entry.
func withLoc(names []string, fset *token.FileSet, declMap map[string]astpkg.DeclNode, basePath string) []symbolWithLoc {
	sort.Strings(names)
	out := make([]symbolWithLoc, 0, len(names))
	for _, n := range names {
		entry := symbolWithLoc{Name: n}
		if decl, ok := declMap[n]; ok && decl.Decl != nil {
			pos := fset.Position(decl.Decl.Pos())
			entry.File = relativizePath(pos.Filename, basePath)
			entry.Line = pos.Line
		}
		out = append(out, entry)
	}
	return out
}

// formatDepsWithLocations renders Dependencies as text with a "name
// (file:line)" suffix on every entry that we can resolve. Names from
// outside the analyzed package (declMap miss) are rendered without the
// suffix rather than dropped — the user still wants to see them.
func formatDepsWithLocations(deps astpkg.Dependencies, fset *token.FileSet, declMap map[string]astpkg.DeclNode, basePath string) string {
	var b strings.Builder

	emit := func(label string, names []string) {
		if len(names) == 0 {
			return
		}
		sort.Strings(names)
		fmt.Fprintf(&b, "%s (%d):\n", label, len(names))
		for _, n := range names {
			if decl, ok := declMap[n]; ok && decl.Decl != nil {
				pos := fset.Position(decl.Decl.Pos())
				rel := relativizePath(pos.Filename, basePath)
				fmt.Fprintf(&b, "  - %s  (%s:%d)\n", n, rel, pos.Line)
			} else {
				fmt.Fprintf(&b, "  - %s\n", n)
			}
		}
		b.WriteString("\n")
	}

	emit("Types", deps.Types.ToSlice())
	emit("Functions", deps.Functions.ToSlice())
	emit("Structs", deps.Structs.ToSlice())
	emit("Interfaces", deps.Interfaces.ToSlice())
	emit("Constants", deps.Constants.ToSlice())
	emit("Variables", deps.Variables.ToSlice())

	if len(deps.Imports.ToSlice()) > 0 {
		emit("Imports", deps.Imports.ToSlice())
	}

	if b.Len() == 0 {
		return "(no dependencies found)\n"
	}
	return b.String()
}

// relativizePath converts an absolute filename returned by fset.Position
// into a path relative to the directory the user passed on the command
// line. We try filepath.Rel; if it fails (different volumes, etc.) or
// produces something with .. that's longer than the absolute, we fall
// back to the absolute path. Either is correct, but the relative form
// is dramatically nicer at the terminal.
func relativizePath(absPath, basePath string) string {
	if absPath == "" {
		return ""
	}
	absBase, err := filepath.Abs(basePath)
	if err != nil {
		return absPath
	}
	rel, err := filepath.Rel(absBase, absPath)
	if err != nil {
		return absPath
	}
	// Don't return "../../../home/..." style traversal — if relativizing
	// makes the path longer or escapes the base, keep absolute.
	if strings.HasPrefix(rel, "..") {
		return absPath
	}
	return rel
}

func sortedSlice(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	sort.Strings(out)
	return out
}
