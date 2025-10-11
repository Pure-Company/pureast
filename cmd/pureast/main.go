package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"

	"github.com/vinodhalaharvi/pureast/pkg/analyze"
	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/pureast/pkg/codegen"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
)

func main() {
	cfg := parseFlags()

	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func showMethodsFromPackage(typeName string, pkgNode astpkg.PackageNode) error {
	allMethods := []astpkg.MethodNode{}

	// Collect methods from all files
	for _, fileNode := range pkgNode.Files {
		if fileNode.File != nil {
			methods := extract.ExtractMethods(typeName, fileNode.File)
			allMethods = append(allMethods, methods...)
		}
	}

	if len(allMethods) == 0 {
		fmt.Printf("No methods found for type: %s\n", typeName)
		return nil
	}

	fmt.Printf("Methods for %s:\n", typeName)
	for _, method := range allMethods {
		fmt.Printf("  - %s\n", method.MethodName)

		if method.Deps.Types.Size() > 0 {
			fmt.Printf("    Types: %v\n", method.Deps.Types.ToSlice())
		}
		if method.Deps.Functions.Size() > 0 {
			fmt.Printf("    Functions: %v\n", method.Deps.Functions.ToSlice())
		}
	}

	return nil
}

func showDependencies(targetName string, graph analyze.DependencyGraph) error {
	deps := graph.ResolveTransitive(targetName)

	fmt.Printf("Dependencies for %s:\n\n", targetName)

	if deps.Types.Size() > 0 {
		fmt.Printf("Types (%d):\n", deps.Types.Size())
		for _, name := range deps.Types.ToSlice() {
			fmt.Printf("  - %s\n", name)
		}
		fmt.Println()
	}

	if deps.Functions.Size() > 0 {
		fmt.Printf("Functions (%d):\n", deps.Functions.Size())
		for _, name := range deps.Functions.ToSlice() {
			fmt.Printf("  - %s\n", name)
		}
		fmt.Println()
	}

	if deps.Structs.Size() > 0 {
		fmt.Printf("Structs (%d):\n", deps.Structs.Size())
		for _, name := range deps.Structs.ToSlice() {
			fmt.Printf("  - %s\n", name)
		}
		fmt.Println()
	}

	if deps.Interfaces.Size() > 0 {
		fmt.Printf("Interfaces (%d):\n", deps.Interfaces.Size())
		for _, name := range deps.Interfaces.ToSlice() {
			fmt.Printf("  - %s\n", name)
		}
		fmt.Println()
	}

	if deps.Imports.Size() > 0 {
		fmt.Printf("Imports (%d):\n", deps.Imports.Size())
		for _, name := range deps.Imports.ToSlice() {
			fmt.Printf("  - %s\n", name)
		}
		fmt.Println()
	}

	return nil
}

func showReport(targetName string, graph analyze.DependencyGraph, fset *token.FileSet) error {
	deps := graph.ResolveTransitive(targetName)
	stats := graph.ComputeStats(targetName)

	gen := codegen.NewGenerator(fset)
	report := gen.GenerateReport(targetName, deps)

	fmt.Println(report)
	fmt.Printf("\nMax Depth: %d\n", stats.MaxDepth)

	return nil
}

func generateDOT(targetName string, graph analyze.DependencyGraph, fset *token.FileSet) error {
	gen := codegen.NewGenerator(fset)
	dot := gen.GenerateDOT(targetName, graph.Decls)

	fmt.Println(dot)
	return nil
}

func generateCode(
	cfg Config,
	pkgNode astpkg.PackageNode,
	graph analyze.DependencyGraph,
	fset *token.FileSet,
) error {
	gen := codegen.NewGenerator(fset)

	// Resolve dependencies with associated code (functions/methods)
	var deps astpkg.Dependencies
	if cfg.Minimal {
		deps = graph.ResolveWithAssociatedCode(cfg.Symbol)
	} else {
		deps = graph.ResolveWithAssociatedCode(cfg.Symbol)
	}

	// Generate code
	code, err := gen.GenerateMinimal(
		pkgNode.Name,
		cfg.Symbol,
		graph.Decls,
		deps,
	)
	if err != nil {
		return fmt.Errorf("generation error: %w", err)
	}

	// Output
	if cfg.OutputFile != "" {
		return os.WriteFile(cfg.OutputFile, []byte(code), 0644)
	}

	fmt.Println(code)
	return nil
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
	AllTypes       bool // NEW: Extract all types
	StructsOnly    bool // NEW: Extract only structs
	InterfacesOnly bool // NEW: Extract only interfaces
	TypesSummary   bool // NEW: Show summary of types
}

func parseFlags() Config {
	cfg := Config{}

	flag.StringVar(&cfg.InputPath, "file", "", "Input Go file or directory (required)")
	flag.StringVar(&cfg.Symbol, "symbol", "", "Target symbol to extract")
	flag.StringVar(&cfg.OutputFile, "output", "", "Output file (default: stdout)")
	flag.BoolVar(&cfg.ShowMethods, "methods", false, "Show methods for type")
	flag.BoolVar(&cfg.ShowDeps, "deps", false, "Show dependencies only")
	flag.BoolVar(&cfg.ShowReport, "report", false, "Generate dependency report")
	flag.BoolVar(&cfg.GenerateDOT, "dot", false, "Generate DOT graph")
	flag.BoolVar(&cfg.Minimal, "minimal", false, "Extract minimal dependencies")
	flag.BoolVar(&cfg.Recursive, "recursive", true, "Process directories recursively")
	flag.IntVar(&cfg.Workers, "workers", 0, "Number of workers (0 = NumCPU)")
	flag.IntVar(&cfg.BatchSize, "batch", 10, "Batch size for concurrent processing")

	// NEW FLAGS
	flag.BoolVar(&cfg.AllTypes, "all-types", false, "Extract all structs and interfaces (for LLM context)")
	flag.BoolVar(&cfg.StructsOnly, "structs", false, "Extract only struct definitions")
	flag.BoolVar(&cfg.InterfacesOnly, "interfaces", false, "Extract only interface definitions")
	flag.BoolVar(&cfg.TypesSummary, "types-summary", false, "Show summary of all types")

	flag.Parse()

	// Symbol not required for all-types mode
	if cfg.InputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if !cfg.AllTypes && !cfg.StructsOnly && !cfg.InterfacesOnly && !cfg.TypesSummary && cfg.Symbol == "" {
		fmt.Fprintln(os.Stderr, "Error: -symbol required unless using -all-types, -structs, -interfaces, or -types-summary")
		flag.Usage()
		os.Exit(1)
	}

	return cfg
}

func run(cfg Config) error {
	// Check if input is file or directory
	info, err := os.Stat(cfg.InputPath)
	if err != nil {
		return fmt.Errorf("cannot access path: %w", err)
	}

	fset := token.NewFileSet()
	var pkgNode astpkg.PackageNode

	if info.IsDir() {
		// Process directory concurrently
		fmt.Fprintf(os.Stderr, "Processing directory: %s (concurrent with %d workers)\n",
			cfg.InputPath, cfg.Workers)

		pkgNode, err = extract.ExtractDirectoryConcurrent(
			fset,
			cfg.InputPath,
			cfg.Recursive,
			cfg.Workers,
		)
		if err != nil {
			return fmt.Errorf("directory processing error: %w", err)
		}
	} else {
		// Single file
		file, err := parser.ParseFile(fset, cfg.InputPath, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("parse error: %w", err)
		}

		fileNode := extract.ExtractFile(file)
		pkgNode = astpkg.PackageNode{
			Name:  fileNode.Name,
			Files: []astpkg.FileNode{fileNode},
			Deps:  fileNode.Deps,
		}
	}

	// Handle all-types modes
	if cfg.AllTypes || cfg.StructsOnly || cfg.InterfacesOnly || cfg.TypesSummary {
		return handleTypesExtraction(cfg, pkgNode, fset)
	}

	// Build unified declaration map from all files
	declMap := extract.BuildPackageDeclMap(pkgNode)

	// Show methods if requested
	if cfg.ShowMethods {
		return showMethodsFromPackage(cfg.Symbol, pkgNode)
	}

	// Create dependency graph
	graph := analyze.NewDependencyGraph(declMap)

	// Show dependencies if requested
	if cfg.ShowDeps {
		return showDependencies(cfg.Symbol, graph)
	}

	// Generate report if requested
	if cfg.ShowReport {
		return showReport(cfg.Symbol, graph, fset)
	}

	// Generate DOT if requested
	if cfg.GenerateDOT {
		return generateDOT(cfg.Symbol, graph, fset)
	}

	// Default: generate code
	return generateCode(cfg, pkgNode, graph, fset)
}

// handleTypesExtraction handles -all-types, -structs, -interfaces modes
func handleTypesExtraction(
	cfg Config,
	pkgNode astpkg.PackageNode,
	fset *token.FileSet,
) error {
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
		fmt.Println(report)
		return nil
	}

	// Generate code with only type declarations
	gen := codegen.NewGenerator(fset)
	code, err := gen.GenerateTypesOnly(
		pkgNode.Name,
		types,
		pkgNode.Deps.Imports.ToSlice(),
	)
	if err != nil {
		return fmt.Errorf("generation error: %w", err)
	}

	// Output
	if cfg.OutputFile != "" {
		return os.WriteFile(cfg.OutputFile, []byte(code), 0644)
	}

	fmt.Println(code)
	return nil
}
