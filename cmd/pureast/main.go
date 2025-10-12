package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/vinodhalaharvi/pureast/pkg/analyze"
	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/pureast/pkg/codegen"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
	"github.com/vinodhalaharvi/pureast/pkg/index"
	"github.com/vinodhalaharvi/purekernels/pkg/functor"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
)

func main() {
	// Parse configuration
	configResult := parseFlags()

	// Execute with error handling
	if !configResult.IsOk() {
		fmt.Fprintf(os.Stderr, "Configuration error: %v\n", configResult.Error())
		flag.Usage()
		os.Exit(1)
	}

	cfg := configResult.Unwrap()

	// Run the application
	runResult := run(cfg)

	if !runResult.IsOk() {
		fmt.Fprintf(os.Stderr, "Error: %v\n", runResult.Error())
		os.Exit(1)
	}

	// Success
	os.Exit(0)
}

// Config holds CLI configuration
type Config struct {
	InputPath      string
	Symbol         string
	OutputFile     string
	ShowMethods    bool
	ShowDeps       bool
	ShowReport     bool
	GenerateDOT    bool
	Minimal        bool
	Recursive      bool
	Workers        int
	BatchSize      int
	AllTypes       bool
	StructsOnly    bool
	InterfacesOnly bool
	TypesSummary   bool
	ListSymbols    bool
	GroupByKind    bool
	FuzzySearch    bool   // NEW: Fuzzy search mode
	SearchPattern  string // NEW: Search pattern for fuzzy search
	BuildIndex     bool   // NEW: Build and cache index
	IndexPath      string // NEW: Path to index file
}

// parseFlags parses command line flags and returns Result[Config]
func parseFlags() result.Result[Config] {
	cfg := Config{}

	flag.StringVar(&cfg.InputPath, "file", "", "Input Go file or directory (required)")
	flag.StringVar(&cfg.Symbol, "symbol", "", "Target symbol(s) to extract (supports regex and comma-separated list)")
	flag.StringVar(&cfg.OutputFile, "output", "", "Output file (default: stdout)")
	flag.BoolVar(&cfg.ShowMethods, "methods", false, "Show methods for type")
	flag.BoolVar(&cfg.ShowDeps, "deps", false, "Show dependencies only")
	flag.BoolVar(&cfg.ShowReport, "report", false, "Generate dependency report")
	flag.BoolVar(&cfg.GenerateDOT, "dot", false, "Generate DOT graph")
	flag.BoolVar(&cfg.Minimal, "minimal", false, "Extract minimal dependencies")
	flag.BoolVar(&cfg.Recursive, "recursive", true, "Process directories recursively")
	flag.IntVar(&cfg.Workers, "workers", 0, "Number of workers (0 = NumCPU)")
	flag.IntVar(&cfg.BatchSize, "batch", 10, "Batch size for concurrent processing")
	flag.BoolVar(&cfg.AllTypes, "all-types", false, "Extract all structs and interfaces (for LLM context)")
	flag.BoolVar(&cfg.StructsOnly, "structs", false, "Extract only struct definitions")
	flag.BoolVar(&cfg.InterfacesOnly, "interfaces", false, "Extract only interface definitions")
	flag.BoolVar(&cfg.TypesSummary, "types-summary", false, "Show summary of all types")
	flag.BoolVar(&cfg.ListSymbols, "list-symbols", false, "List all available symbols")
	flag.BoolVar(&cfg.GroupByKind, "group", true, "Group symbols by kind when listing")
	flag.BoolVar(&cfg.FuzzySearch, "search", false, "Fuzzy search for symbols")
	flag.StringVar(&cfg.SearchPattern, "pattern", "", "Search pattern for fuzzy search")
	flag.BoolVar(&cfg.BuildIndex, "index", false, "Build and save search index")
	flag.StringVar(&cfg.IndexPath, "index-path", ".pureast-index.json", "Path to index file")

	flag.Parse()

	// Validate configuration
	return validateConfig(cfg)
}

// validateConfig validates the configuration
func validateConfig(cfg Config) result.Result[Config] {
	// Input path is always required
	if cfg.InputPath == "" {
		return result.Err[Config](fmt.Errorf("-file is required"))
	}

	// Symbol required for certain modes
	requiresSymbol := !cfg.AllTypes && !cfg.StructsOnly && !cfg.InterfacesOnly &&
		!cfg.TypesSummary && !cfg.ListSymbols && !cfg.FuzzySearch && !cfg.BuildIndex

	if requiresSymbol && cfg.Symbol == "" {
		return result.Err[Config](fmt.Errorf("-symbol required for this operation"))
	}

	// Fuzzy search requires pattern
	if cfg.FuzzySearch && cfg.SearchPattern == "" {
		return result.Err[Config](fmt.Errorf("-pattern required for fuzzy search"))
	}

	return result.Ok(cfg)
}

// run executes the main application logic
func run(cfg Config) result.Result[string] {
	// Load package
	pkgResult := loadPackage(cfg)
	if !pkgResult.IsOk() {
		return result.Err[string](pkgResult.Error())
	}

	pkgNode := pkgResult.Unwrap()

	// Build index if requested
	if cfg.BuildIndex {
		return buildAndSaveIndex(cfg, pkgNode)
	}

	// Fuzzy search mode
	if cfg.FuzzySearch {
		return fuzzySearchMode(cfg, pkgNode)
	}

	// List symbols mode
	if cfg.ListSymbols {
		return listSymbolsMode(cfg, pkgNode)
	}

	// Types extraction modes
	if cfg.AllTypes || cfg.StructsOnly || cfg.InterfacesOnly || cfg.TypesSummary {
		return handleTypesExtraction(cfg, pkgNode)
	}

	// Symbol extraction modes
	return handleSymbolExtraction(cfg, pkgNode)
}

// loadPackage loads the package from file or directory
func loadPackage(cfg Config) result.Result[astpkg.PackageNode] {
	info, err := os.Stat(cfg.InputPath)
	if err != nil {
		return result.Err[astpkg.PackageNode](fmt.Errorf("cannot access path: %w", err))
	}

	fset := token.NewFileSet()

	if info.IsDir() {
		fmt.Fprintf(os.Stderr, "Processing directory: %s (concurrent with %d workers)\n",
			cfg.InputPath, cfg.Workers)

		pkgNode, err := extract.ExtractDirectoryConcurrent(
			fset,
			cfg.InputPath,
			cfg.Recursive,
			cfg.Workers,
		)
		if err != nil {
			return result.Err[astpkg.PackageNode](fmt.Errorf("directory processing error: %w", err))
		}
		return result.Ok(pkgNode)
	}

	// Single file
	file, err := parser.ParseFile(fset, cfg.InputPath, nil, parser.ParseComments)
	if err != nil {
		return result.Err[astpkg.PackageNode](fmt.Errorf("parse error: %w", err))
	}

	fileNode := extract.ExtractFile(file)
	pkgNode := astpkg.PackageNode{
		Name:  fileNode.Name,
		Files: []astpkg.FileNode{fileNode},
		Deps:  fileNode.Deps,
	}

	return result.Ok(pkgNode)
}

// buildAndSaveIndex builds and saves a search index
func buildAndSaveIndex(cfg Config, pkgNode astpkg.PackageNode) result.Result[string] {
	fmt.Fprintf(os.Stderr, "Building search index...\n")

	// Build index
	idx := index.BuildIndex(pkgNode)

	// Save to file
	if err := index.SaveIndex(idx, cfg.IndexPath); err != nil {
		return result.Err[string](fmt.Errorf("failed to save index: %w", err))
	}

	msg := fmt.Sprintf("Index saved to %s\n", cfg.IndexPath)
	fmt.Fprintf(os.Stderr, "%s", msg)
	fmt.Fprintf(os.Stderr, "Indexed %d symbols\n", len(idx.Symbols))

	return result.Ok(msg)
}

// fuzzySearchMode performs fuzzy search using the index
func fuzzySearchMode(cfg Config, pkgNode astpkg.PackageNode) result.Result[string] {
	fmt.Fprintf(os.Stderr, "Fuzzy searching for: %s\n", cfg.SearchPattern)

	// Try to load existing index, otherwise build new one
	idx := loadOrBuildIndex(cfg, pkgNode)

	// Create search pattern
	var pattern = index.SearchPattern{
		SymbolPattern: cfg.SearchPattern,
		Kind:          "", // Could be extended with -kind flag
		PackageName:   "",
	} // Perform fuzzy search using concurrent processing
	matches := index.FuzzySearchIndexConcurrent(pattern, idx, cfg.BatchSize)

	if len(matches) == 0 {
		msg := fmt.Sprintf("No matches found for pattern: %s\n", cfg.SearchPattern)
		fmt.Print(msg)
		return result.Ok(msg)
	}

	// Format and output results
	output := formatFuzzyResults(matches, cfg.SearchPattern)
	fmt.Print(output)

	return result.Ok(output)
}

// loadOrBuildIndex loads index from file or builds a new one
func loadOrBuildIndex(cfg Config, pkgNode astpkg.PackageNode) index.Index {
	// Try to load from file
	if _, err := os.Stat(cfg.IndexPath); err == nil {
		if idx, err := index.LoadIndex(cfg.IndexPath); err == nil {
			fmt.Fprintf(os.Stderr, "Loaded index from %s\n", cfg.IndexPath)
			return idx
		}
	}

	// Build new index
	fmt.Fprintf(os.Stderr, "Building new index...\n")
	return index.BuildIndex(pkgNode)
}

// formatFuzzyResults formats fuzzy search results
func formatFuzzyResults(matches []index.ScoredSymbol, pattern string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Found %d matches for '%s':\n\n", len(matches), pattern))

	for i, match := range matches {
		if i >= 20 { // Limit to top 20 results
			sb.WriteString(fmt.Sprintf("\n... and %d more matches\n", len(matches)-20))
			break
		}

		sb.WriteString(fmt.Sprintf("%d. %s (score: %d)\n",
			i+1, match.Entry.Name, match.Score.Score))
		sb.WriteString(fmt.Sprintf("   Kind: %s\n", match.Entry.Kind))
		sb.WriteString(fmt.Sprintf("   Package: %s\n", match.Entry.PackageName))
		sb.WriteString(fmt.Sprintf("   File: %s\n", match.Entry.File))
		sb.WriteString("\n")
	}

	return sb.String()
}

// listSymbolsMode lists all available symbols
func listSymbolsMode(cfg Config, pkgNode astpkg.PackageNode) result.Result[string] {
	symbols := extract.DiscoverAllSymbols(pkgNode)

	output := fmt.Sprintf("Found %d symbols in package '%s'\n", len(symbols), pkgNode.Name)
	output += extract.FormatSymbolList(symbols, cfg.GroupByKind)

	fmt.Println(output)
	return result.Ok(output)
}

// handleTypesExtraction handles type extraction modes
func handleTypesExtraction(cfg Config, pkgNode astpkg.PackageNode) result.Result[string] {
	fset := token.NewFileSet()

	// Extract types based on mode
	var types []extract.TypeDeclaration

	if cfg.AllTypes {
		types = extract.ExtractAllStructsAndInterfaces(pkgNode)
	} else if cfg.StructsOnly {
		types = extract.ExtractAllStructs(pkgNode)
	} else if cfg.InterfacesOnly {
		types = extract.ExtractAllInterfaces(pkgNode)
	} else if cfg.TypesSummary {
		// Show summary
		allTypes := extract.ExtractAllTypes(pkgNode)
		summary := extract.SummarizeTypes(allTypes)
		report := codegen.GenerateTypesSummaryReport(summary)
		return outputResult(cfg, report)
	}

	// Generate code with only type declarations
	gen := codegen.NewGenerator(fset)
	code, err := gen.GenerateTypesOnly(
		pkgNode.Name,
		types,
		pkgNode.Deps.Imports.ToSlice(),
	)
	if err != nil {
		return result.Err[string](fmt.Errorf("generation error: %w", err))
	}

	return outputResult(cfg, code)
}

// handleSymbolExtraction handles symbol extraction modes
func handleSymbolExtraction(cfg Config, pkgNode astpkg.PackageNode) result.Result[string] {
	fset := token.NewFileSet()
	declMap := extract.BuildPackageDeclMap(pkgNode)

	// Check if symbol contains patterns (regex or comma-separated)
	if strings.Contains(cfg.Symbol, ",") || isRegexPattern(cfg.Symbol) {
		return handleMultipleSymbols(cfg, pkgNode, declMap, fset)
	}

	// Single symbol modes
	if cfg.ShowMethods {
		return showMethodsFromPackage(cfg.Symbol, pkgNode)
	}

	graph := analyze.NewDependencyGraph(declMap)

	if cfg.ShowDeps {
		return showDependencies(cfg.Symbol, graph)
	}

	if cfg.ShowReport {
		return showReport(cfg.Symbol, graph, fset)
	}

	if cfg.GenerateDOT {
		return generateDOT(cfg.Symbol, graph, fset)
	}

	return generateCode(cfg, pkgNode, graph, fset)
}

// handleMultipleSymbols handles regex or comma-separated symbols using concurrent processing
func handleMultipleSymbols(
	cfg Config,
	pkgNode astpkg.PackageNode,
	declMap map[string]astpkg.DeclNode,
	fset *token.FileSet,
) result.Result[string] {
	// Discover all symbols
	allSymbols := extract.DiscoverAllSymbols(pkgNode)

	// Match patterns
	matched := extract.MatchMultipleSymbols(cfg.Symbol, allSymbols)

	if len(matched) == 0 {
		return result.Err[string](fmt.Errorf("no symbols matched pattern: %s", cfg.Symbol))
	}

	fmt.Fprintf(os.Stderr, "Matched %d symbols: ", len(matched))
	for i, s := range matched {
		if i > 0 {
			fmt.Fprintf(os.Stderr, ", ")
		}
		fmt.Fprintf(os.Stderr, "%s", s.Name)
	}
	fmt.Fprintf(os.Stderr, "\n\n")

	// Extract all matched symbols using concurrent processing
	graph := analyze.NewDependencyGraph(declMap)
	gen := codegen.NewGenerator(fset)

	// Use Concurrent applicative to resolve dependencies in parallel
	depMonoid := astpkg.NewDependencyMonoid()

	// Use TraverseConcurrent to resolve all dependencies in parallel
	combinedDeps := functor.TraverseConcurrent(
		depMonoid,
		func(symbol extract.SymbolInfo) astpkg.Dependencies {
			return graph.ResolveWithAssociatedCode(symbol.Name)
		},
		matched,
		cfg.Workers,
	).Value()

	// Generate code
	code, err := gen.GenerateMinimal(
		pkgNode.Name,
		"", // No single target
		declMap,
		combinedDeps,
	)
	if err != nil {
		return result.Err[string](fmt.Errorf("generation error: %w", err))
	}

	return outputResult(cfg, code)
}

// showMethodsFromPackage shows methods for a type
func showMethodsFromPackage(typeName string, pkgNode astpkg.PackageNode) result.Result[string] {
	// Collect methods from all files using parallel processing
	methodMonoid := NewMethodNodeMonoid()

	methods := functor.TraverseConcurrent(
		methodMonoid,
		func(fileNode astpkg.FileNode) []astpkg.MethodNode {
			if fileNode.File != nil {
				return extract.ExtractMethods(typeName, fileNode.File)
			}
			return []astpkg.MethodNode{}
		},
		pkgNode.Files,
		0,
	).Value()

	if len(methods) == 0 {
		msg := fmt.Sprintf("No methods found for type: %s\n", typeName)
		fmt.Print(msg)
		return result.Ok(msg)
	}

	var output strings.Builder
	output.WriteString(fmt.Sprintf("Methods for %s:\n", typeName))
	for _, method := range methods {
		output.WriteString(fmt.Sprintf("  - %s\n", method.MethodName))

		if method.Deps.Types.Size() > 0 {
			output.WriteString(fmt.Sprintf("    Types: %v\n", method.Deps.Types.ToSlice()))
		}
		if method.Deps.Functions.Size() > 0 {
			output.WriteString(fmt.Sprintf("    Functions: %v\n", method.Deps.Functions.ToSlice()))
		}
	}

	fmt.Print(output.String())
	return result.Ok(output.String())
}

// MethodNodeMonoid for combining method lists
type MethodNodeMonoid struct{}

func NewMethodNodeMonoid() MethodNodeMonoid {
	return MethodNodeMonoid{}
}

func (MethodNodeMonoid) Empty() []astpkg.MethodNode {
	return []astpkg.MethodNode{}
}

func (MethodNodeMonoid) Combine(a, b []astpkg.MethodNode) []astpkg.MethodNode {
	return append(a, b...)
}

// showDependencies shows dependencies for a symbol
func showDependencies(targetName string, graph analyze.DependencyGraph) result.Result[string] {
	deps := graph.ResolveTransitive(targetName)

	var output strings.Builder
	output.WriteString(fmt.Sprintf("Dependencies for %s:\n\n", targetName))

	if deps.Types.Size() > 0 {
		output.WriteString(fmt.Sprintf("Types (%d):\n", deps.Types.Size()))
		for _, name := range deps.Types.ToSlice() {
			output.WriteString(fmt.Sprintf("  - %s\n", name))
		}
		output.WriteString("\n")
	}

	if deps.Functions.Size() > 0 {
		output.WriteString(fmt.Sprintf("Functions (%d):\n", deps.Functions.Size()))
		for _, name := range deps.Functions.ToSlice() {
			output.WriteString(fmt.Sprintf("  - %s\n", name))
		}
		output.WriteString("\n")
	}

	if deps.Structs.Size() > 0 {
		output.WriteString(fmt.Sprintf("Structs (%d):\n", deps.Structs.Size()))
		for _, name := range deps.Structs.ToSlice() {
			output.WriteString(fmt.Sprintf("  - %s\n", name))
		}
		output.WriteString("\n")
	}

	if deps.Interfaces.Size() > 0 {
		output.WriteString(fmt.Sprintf("Interfaces (%d):\n", deps.Interfaces.Size()))
		for _, name := range deps.Interfaces.ToSlice() {
			output.WriteString(fmt.Sprintf("  - %s\n", name))
		}
		output.WriteString("\n")
	}

	if deps.Imports.Size() > 0 {
		output.WriteString(fmt.Sprintf("Imports (%d):\n", deps.Imports.Size()))
		for _, name := range deps.Imports.ToSlice() {
			output.WriteString(fmt.Sprintf("  - %s\n", name))
		}
		output.WriteString("\n")
	}

	fmt.Print(output.String())
	return result.Ok(output.String())
}

// showReport generates a dependency report
func showReport(targetName string, graph analyze.DependencyGraph, fset *token.FileSet) result.Result[string] {
	deps := graph.ResolveTransitive(targetName)
	stats := graph.ComputeStats(targetName)

	gen := codegen.NewGenerator(fset)
	report := gen.GenerateReport(targetName, deps)

	output := fmt.Sprintf("%s\n\nMax Depth: %d\n", report, stats.MaxDepth)
	fmt.Print(output)

	return result.Ok(output)
}

// generateDOT generates a DOT graph
func generateDOT(targetName string, graph analyze.DependencyGraph, fset *token.FileSet) result.Result[string] {
	gen := codegen.NewGenerator(fset)
	dot := gen.GenerateDOT(targetName, graph.Decls)

	fmt.Println(dot)
	return result.Ok(dot)
}

// generateCode generates extracted code
func generateCode(
	cfg Config,
	pkgNode astpkg.PackageNode,
	graph analyze.DependencyGraph,
	fset *token.FileSet,
) result.Result[string] {
	gen := codegen.NewGenerator(fset)

	// Resolve dependencies with associated code
	deps := graph.ResolveWithAssociatedCode(cfg.Symbol)

	// Generate code
	code, err := gen.GenerateMinimal(
		pkgNode.Name,
		cfg.Symbol,
		graph.Decls,
		deps,
	)
	if err != nil {
		return result.Err[string](fmt.Errorf("generation error: %w", err))
	}

	return outputResult(cfg, code)
}

// outputResult handles output to file or stdout
func outputResult(cfg Config, content string) result.Result[string] {
	if cfg.OutputFile != "" {
		if err := os.WriteFile(cfg.OutputFile, []byte(content), 0644); err != nil {
			return result.Err[string](fmt.Errorf("failed to write output: %w", err))
		}
		msg := fmt.Sprintf("Output written to %s\n", cfg.OutputFile)
		fmt.Fprintf(os.Stderr, "%s", msg)
		return result.Ok(msg)
	}

	fmt.Print(content)
	return result.Ok(content)
}

// isRegexPattern checks if string looks like a regex
func isRegexPattern(s string) bool {
	regexChars := []string{"*", ".", "+", "?", "[", "]", "^", "$", "|", "(", ")"}
	for _, char := range regexChars {
		if strings.Contains(s, char) {
			return true
		}
	}
	return false
}
