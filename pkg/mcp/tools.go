package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"go/token"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Pure-Company/pureast/pkg/analyze"
	astpkg "github.com/Pure-Company/pureast/pkg/ast"
	"github.com/Pure-Company/pureast/pkg/codegen"
	"github.com/Pure-Company/pureast/pkg/extract"
	"github.com/Pure-Company/purekernels/pkg/functor"
	"github.com/Pure-Company/purekernels/pkg/monoid"
	"github.com/Pure-Company/purekernels/pkg/result"
)

// ToolExecutor executes pureast tools using applicative kernels
type ToolExecutor struct {
	workers int
}

// NewToolExecutor creates a tool executor
func NewToolExecutor(workers int) *ToolExecutor {
	return &ToolExecutor{workers: workers}
}

// SearchSymbolsHandler searches for symbols using fuzzy matching.
func (te *ToolExecutor) SearchSymbolsHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				var params struct {
					Name      string `json:"name"`
					Arguments struct {
						Pattern    string `json:"pattern"`
						Path       string `json:"path"`
						Fuzzy      bool   `json:"fuzzy"`
						Kind       string `json:"kind,omitempty"`
						MaxResults int    `json:"maxResults,omitempty"`
					} `json:"arguments"`
				}

				if err := json.Unmarshal(req.Params, &params); err != nil {
					return ErrorResponse(req.ID, InvalidParams, "Invalid parameters")
				}

				fset := token.NewFileSet()
				pkgResult := loadPackage(fset, params.Arguments.Path, te.workers)
				if !pkgResult.IsOk() {
					return ErrorResponse(req.ID, InternalError, pkgResult.Error().Error())
				}
				pkgNode := pkgResult.Unwrap()

				maxResults := params.Arguments.MaxResults
				if maxResults <= 0 || maxResults > 100 {
					maxResults = 20
				}

				symbols := extract.DiscoverAllSymbols(pkgNode)
				matches := extract.FuzzySearch(
					symbols,
					params.Arguments.Pattern,
					params.Arguments.Kind,
					maxResults,
				)

				if !params.Arguments.Fuzzy {
					filtered := matches[:0]
					for _, m := range matches {
						if m.Score >= 400 {
							filtered = append(filtered, m)
						}
					}
					matches = filtered
				}

				return MCPResponse{
					JSONRPC: "2.0",
					ID:      req.ID,
					Result: map[string]interface{}{
						"content": []map[string]interface{}{
							{
								"type": "text",
								"text": formatSearchResults(matches, pkgNode.Name),
							},
						},
					},
				}
			},
		)
	}
}

// ExtractSymbolHandler extracts a symbol with dependencies.
//
// New params vs original:
//   - format: "go" (default) | "md" — wraps output in a fenced code block
//   - maxTokens: 0 = unbounded; truncates output to fit token budget
func (te *ToolExecutor) ExtractSymbolHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				var params struct {
					Name      string `json:"name"`
					Arguments struct {
						Symbol    string `json:"symbol"`
						Path      string `json:"path"`
						Minimal   bool   `json:"minimal"`
						Format    string `json:"format,omitempty"`    // "go"|"md"
						MaxTokens int    `json:"maxTokens,omitempty"` // 0 = unbounded
					} `json:"arguments"`
				}

				if err := json.Unmarshal(req.Params, &params); err != nil {
					return ErrorResponse(req.ID, InvalidParams, "Invalid parameters")
				}

				fset := token.NewFileSet()
				pkgResult := loadPackage(fset, params.Arguments.Path, te.workers)
				if !pkgResult.IsOk() {
					return ErrorResponse(req.ID, InternalError, pkgResult.Error().Error())
				}
				pkgNode := pkgResult.Unwrap()

				declMap := extract.BuildPackageDeclMap(pkgNode)
				graph := analyze.NewDependencyGraph(declMap)
				deps := graph.ResolveWithAssociatedCode(params.Arguments.Symbol)

				gen := codegen.NewGenerator(fset)
				code, err := gen.GenerateMinimal(
					pkgNode.Name,
					params.Arguments.Symbol,
					declMap,
					deps,
				)
				if err != nil {
					return ErrorResponse(req.ID, InternalError, err.Error())
				}

				if params.Arguments.MaxTokens > 0 {
					code, _ = extract.TruncateSymbols(code, params.Arguments.MaxTokens)
				}
				if params.Arguments.Format == "md" {
					code = "```go\n" + strings.TrimRight(code, "\n") + "\n```\n"
				}

				return textResponse(req.ID, code)
			},
		)
	}
}

// ListSymbolsHandler lists all symbols in a package.
func (te *ToolExecutor) ListSymbolsHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				var params struct {
					Name      string `json:"name"`
					Arguments struct {
						Path        string `json:"path"`
						GroupByKind bool   `json:"groupByKind"`
					} `json:"arguments"`
				}

				if err := json.Unmarshal(req.Params, &params); err != nil {
					return ErrorResponse(req.ID, InvalidParams, "Invalid parameters")
				}

				fset := token.NewFileSet()
				pkgResult := loadPackage(fset, params.Arguments.Path, te.workers)
				if !pkgResult.IsOk() {
					return ErrorResponse(req.ID, InternalError, pkgResult.Error().Error())
				}
				pkgNode := pkgResult.Unwrap()

				symbols := extract.DiscoverAllSymbols(pkgNode)
				text := extract.FormatSymbolList(symbols, params.Arguments.GroupByKind)

				return textResponse(req.ID, text)
			},
		)
	}
}

// ExtractTypesHandler extracts type definitions (structs and interfaces).
func (te *ToolExecutor) ExtractTypesHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				var params struct {
					Name      string `json:"name"`
					Arguments struct {
						Path           string `json:"path"`
						StructsOnly    bool   `json:"structsOnly"`
						InterfacesOnly bool   `json:"interfacesOnly"`
					} `json:"arguments"`
				}

				if err := json.Unmarshal(req.Params, &params); err != nil {
					return ErrorResponse(req.ID, InvalidParams, "Invalid parameters")
				}

				fset := token.NewFileSet()
				pkgResult := loadPackage(fset, params.Arguments.Path, te.workers)
				if !pkgResult.IsOk() {
					return ErrorResponse(req.ID, InternalError, pkgResult.Error().Error())
				}
				pkgNode := pkgResult.Unwrap()

				var types []extract.TypeDeclaration
				if params.Arguments.StructsOnly {
					types = extract.ExtractAllStructs(pkgNode)
				} else if params.Arguments.InterfacesOnly {
					types = extract.ExtractAllInterfaces(pkgNode)
				} else {
					types = extract.ExtractAllStructsAndInterfaces(pkgNode)
				}

				gen := codegen.NewGenerator(fset)
				code, err := gen.GenerateTypesOnly(
					pkgNode.Name,
					types,
					pkgNode.Deps.Imports.ToSlice(),
				)
				if err != nil {
					return ErrorResponse(req.ID, InternalError, err.Error())
				}

				return textResponse(req.ID, code)
			},
		)
	}
}

// ShowDependenciesHandler shows dependencies for a symbol.
//
// New params vs original:
//   - format:    "text" (default) | "json"
//   - depth:     0 = unbounded (default), N = N hops forward
//   - minimal:   direct non-transitive deps only
//   - locations: include file:line in text output
func (te *ToolExecutor) ShowDependenciesHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				var params struct {
					Name      string `json:"name"`
					Arguments struct {
						Symbol    string `json:"symbol"`
						Path      string `json:"path"`
						Format    string `json:"format,omitempty"`    // "text"|"json"
						Depth     int    `json:"depth,omitempty"`     // 0=unbounded, N=N hops
						Minimal   bool   `json:"minimal,omitempty"`   // direct deps only
						Locations bool   `json:"locations,omitempty"` // include file:line
					} `json:"arguments"`
				}

				if err := json.Unmarshal(req.Params, &params); err != nil {
					return ErrorResponse(req.ID, InvalidParams, "Invalid parameters")
				}

				fset := token.NewFileSet()
				pkgResult := loadPackage(fset, params.Arguments.Path, te.workers)
				if !pkgResult.IsOk() {
					return ErrorResponse(req.ID, InternalError, pkgResult.Error().Error())
				}
				pkgNode := pkgResult.Unwrap()

				declMap := extract.BuildPackageDeclMap(pkgNode)
				graph := analyze.NewDependencyGraph(declMap)

				// Select resolution strategy.
				// depth=0 via omitempty means "not set" → unbounded.
				var deps astpkg.Dependencies
				switch {
				case params.Arguments.Minimal:
					deps = graph.MinimalDependencies(params.Arguments.Symbol)
				case params.Arguments.Depth > 0:
					deps = graph.ResolveBounded(params.Arguments.Symbol, params.Arguments.Depth)
				default:
					deps = graph.ResolveWithAssociatedCode(params.Arguments.Symbol)
				}
				deps = analyze.CleanDependencies(deps, declMap)

				var text string
				switch params.Arguments.Format {
				case "json":
					text = formatDepsJSONForMCP(params.Arguments.Symbol, deps, params.Arguments.Locations, fset, declMap, params.Arguments.Path)
				default: // text
					header := "Dependencies for " + params.Arguments.Symbol + ":"
					if params.Arguments.Locations {
						text = header + "\n\n" + formatDepsLocationsForMCP(deps, fset, declMap, params.Arguments.Path)
					} else {
						text = header + "\n\n" + analyze.FormatDependencies(params.Arguments.Symbol, deps)
					}
				}

				return textResponse(req.ID, text)
			},
		)
	}
}

// ── helpers ───────────────────────────────────────────────────────────────────

func loadPackage(fset *token.FileSet, path string, workers int) result.Result[astpkg.PackageNode] {
	pkgNode, err := extract.ExtractDirectoryConcurrent(fset, path, true, workers)
	if err != nil {
		return result.Err[astpkg.PackageNode](err)
	}
	return result.Ok(pkgNode)
}

func formatSearchResults(matches []extract.Match, pkgName string) string {
	if len(matches) == 0 {
		return "No symbols found."
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Found %d symbols:\n\n", len(matches))
	for _, m := range matches {
		fmt.Fprintf(&b, "- %s (%s) [score: %d] in package %s\n",
			m.Symbol.Name, m.Symbol.Kind, m.Score, pkgName)
	}
	return strings.TrimRight(b.String(), "\n")
}

func formatDependencies(symbol string, deps astpkg.Dependencies) string {
	var parts []string

	parts = append(parts, fmt.Sprintf("Dependencies for %s:\n", symbol))

	if deps.Types.Size() > 0 {
		parts = append(parts, fmt.Sprintf("\nTypes (%d):", deps.Types.Size()))
		for _, name := range deps.Types.ToSlice() {
			parts = append(parts, "  - "+name)
		}
	}

	if deps.Functions.Size() > 0 {
		parts = append(parts, fmt.Sprintf("\nFunctions (%d):", deps.Functions.Size()))
		for _, name := range deps.Functions.ToSlice() {
			parts = append(parts, "  - "+name)
		}
	}

	if deps.Imports.Size() > 0 {
		parts = append(parts, fmt.Sprintf("\nImports (%d):", deps.Imports.Size()))
		for _, name := range deps.Imports.ToSlice() {
			parts = append(parts, "  - "+name)
		}
	}

	return monoid.Reduce(monoid.NewStringJoinMonoid("\n"), parts)
}

// formatDepsJSONForMCP renders dependencies as structured JSON.
// Mirrors cmd/pureast/commands/deps.go:formatDepsJSON but lives in pkg/mcp
// so the HTTP/stdio MCP paths share it without importing cmd packages.
func formatDepsJSONForMCP(
	symbol string,
	deps astpkg.Dependencies,
	withLocations bool,
	fset *token.FileSet,
	declMap map[string]astpkg.DeclNode,
	basePath string,
) string {
	type entry struct {
		Name string `json:"name"`
		File string `json:"file,omitempty"`
		Line int    `json:"line,omitempty"`
	}
	type section struct {
		Kind    string  `json:"kind"`
		Entries []entry `json:"entries"`
	}

	sections := []section{}

	emit := func(kind string, names []string) {
		if len(names) == 0 {
			return
		}
		sorted := make([]string, len(names))
		copy(sorted, names)
		sort.Strings(sorted)
		entries := make([]entry, 0, len(sorted))
		for _, n := range sorted {
			e := entry{Name: n}
			if withLocations {
				if loc, line, ok := lookupFileLineForMCP(n, fset, declMap, basePath); ok {
					e.File = loc
					e.Line = line
				}
			}
			entries = append(entries, e)
		}
		sections = append(sections, section{Kind: kind, Entries: entries})
	}

	emit("types", deps.Types.ToSlice())
	emit("functions", deps.Functions.ToSlice())
	emit("imports", deps.Imports.ToSlice())

	out, err := json.MarshalIndent(map[string]interface{}{
		"symbol":   symbol,
		"sections": sections,
	}, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": %q}`, err.Error())
	}
	return string(out)
}

// formatDepsLocationsForMCP renders dependencies as text with file:line annotations.
func formatDepsLocationsForMCP(
	deps astpkg.Dependencies,
	fset *token.FileSet,
	declMap map[string]astpkg.DeclNode,
	basePath string,
) string {
	var b strings.Builder

	emit := func(label string, names []string) {
		if len(names) == 0 {
			return
		}
		sorted := make([]string, len(names))
		copy(sorted, names)
		sort.Strings(sorted)
		fmt.Fprintf(&b, "%s (%d):\n", label, len(sorted))
		for _, n := range sorted {
			if loc, line, ok := lookupFileLineForMCP(n, fset, declMap, basePath); ok {
				fmt.Fprintf(&b, "  - %s  (%s:%d)\n", n, loc, line)
			} else {
				fmt.Fprintf(&b, "  - %s\n", n)
			}
		}
		b.WriteString("\n")
	}

	emit("Types", deps.Types.ToSlice())
	emit("Functions", deps.Functions.ToSlice())
	emit("Imports", deps.Imports.ToSlice())

	return b.String()
}

func lookupFileLineForMCP(
	name string,
	fset *token.FileSet,
	declMap map[string]astpkg.DeclNode,
	basePath string,
) (string, int, bool) {
	decl, ok := declMap[name]
	if !ok || decl.Decl == nil {
		return "", 0, false
	}
	pos := fset.Position(decl.Decl.Pos())
	if pos.Filename == "" {
		return "", 0, false
	}
	abs, err := filepath.Abs(basePath)
	if err != nil {
		return pos.Filename, pos.Line, true
	}
	rel, err := filepath.Rel(abs, pos.Filename)
	if err != nil || strings.HasPrefix(rel, "..") {
		return pos.Filename, pos.Line, true
	}
	return rel, pos.Line, true
}
