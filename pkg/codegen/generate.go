package codegen

import (
	"go/format"
	"go/printer"
	"go/token"
	"strconv"

	astpkg "github.com/Pure-Company/pureast/pkg/ast"
	"github.com/Pure-Company/purekernels/pkg/fold"
	"github.com/Pure-Company/purekernels/pkg/monoid"
)

// Generator generates Go code from AST nodes
type Generator struct {
	fset *token.FileSet
}

// NewGenerator creates a new code generator
func NewGenerator(fset *token.FileSet) Generator {
	return Generator{fset: fset}
}

// Code represents generated code as a monoid
type Code struct {
	Lines []string
}

// CodeMonoid - monoid for combining code
type CodeMonoid struct{}

func NewCodeMonoid() CodeMonoid {
	return CodeMonoid{}
}

func (CodeMonoid) Empty() Code {
	return Code{Lines: []string{}}
}

func (CodeMonoid) Combine(a, b Code) Code {
	return Code{
		Lines: append(a.Lines, b.Lines...),
	}
}

// ToCode converts a string to Code
func ToCode(s string) Code {
	return Code{Lines: []string{s}}
}

// Join joins code lines with separator
func (c Code) Join(sep string) string {
	m := monoid.NewStringJoinMonoid(sep)
	return monoid.Reduce(m, c.Lines)
}

// GenerateFile generates code for a single file
func (g Generator) GenerateFile(
	packageName string,
	imports []string,
	decls []astpkg.DeclNode,
) (string, error) {
	codeMonoid := NewCodeMonoid()

	// Package declaration
	packageCode := ToCode("package " + packageName)

	// Imports
	importsCode := g.generateImports(imports)

	// Declarations
	declsCode := g.generateDecls(decls)

	// Combine using monoid
	allCode := monoid.Reduce(
		codeMonoid,
		[]Code{packageCode, importsCode, declsCode},
	)

	// Join with double newlines
	result := allCode.Join("\n\n")

	// Format
	formatted, err := format.Source([]byte(result))
	if err != nil {
		return result, err
	}

	return string(formatted), nil
}

// generateImports generates import block as Code
func (g Generator) generateImports(imports []string) Code {
	if len(imports) == 0 {
		return Code{Lines: []string{}}
	}

	if len(imports) == 1 {
		return ToCode("import \"" + imports[0] + "\"")
	}

	// Map imports to quoted strings
	quotedImports := fold.Map(
		func(imp string) string {
			return "\t\"" + imp + "\""
		},
		imports,
	)

	// Join with newlines
	joinedImports := monoid.Reduce(
		monoid.NewStringJoinMonoid("\n"),
		quotedImports,
	)

	return ToCode("import (\n" + joinedImports + "\n)")
}

// generateDecls generates all declarations as Code
func (g Generator) generateDecls(decls []astpkg.DeclNode) Code {
	codeMonoid := NewCodeMonoid()

	// Map each declaration to Code
	declCodes := fold.Map(
		func(node astpkg.DeclNode) Code {
			return g.generateDecl(node)
		},
		decls,
	)

	// Combine all using monoid
	return monoid.Reduce(codeMonoid, declCodes)
}

// generateDecl generates code for a declaration
func (g Generator) generateDecl(node astpkg.DeclNode) Code {
	// Use printer to convert AST to string
	var buf []byte
	tmpBuf := &bytesBuffer{data: buf}

	err := printer.Fprint(tmpBuf, g.fset, node.Decl)
	if err != nil {
		return Code{Lines: []string{}}
	}

	return ToCode(string(tmpBuf.data))
}

// bytesBuffer implements io.Writer for printer
type bytesBuffer struct {
	data []byte
}

func (b *bytesBuffer) Write(p []byte) (n int, err error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

// GenerateMinimal generates minimal extraction for target
func (g Generator) GenerateMinimal(
	packageName string,
	targetName string,
	allDecls map[string]astpkg.DeclNode,
	deps astpkg.Dependencies,
) (string, error) {
	// Collect all declaration names to include
	allNames := collectDeclNames(targetName, deps)

	// Map names to declarations using fold
	includedDecls := fold.FoldLeft(
		func(acc []astpkg.DeclNode, name string) []astpkg.DeclNode {
			if decl, ok := allDecls[name]; ok {
				return append(acc, decl)
			}
			return acc
		},
		[]astpkg.DeclNode{},
		allNames,
	)

	// Generate code
	return g.GenerateFile(
		packageName,
		deps.Imports.ToSlice(),
		includedDecls,
	)
}

// collectDeclNames collects all declaration names from dependencies
func collectDeclNames(targetName string, deps astpkg.Dependencies) []string {
	// Use SetMonoid to collect unique names
	nameSet := monoid.NewSetMonoid[string]()
	nameSet = nameSet.Insert(targetName)

	// Add all dependency types
	for _, name := range deps.Types.ToSlice() {
		nameSet = nameSet.Insert(name)
	}

	// Add all functions
	for _, name := range deps.Functions.ToSlice() {
		nameSet = nameSet.Insert(name)
	}

	// Add all structs
	for _, name := range deps.Structs.ToSlice() {
		nameSet = nameSet.Insert(name)
	}

	// Add all interfaces
	for _, name := range deps.Interfaces.ToSlice() {
		nameSet = nameSet.Insert(name)
	}

	return nameSet.ToSlice()
}

// GenerateOrdered generates code with topological ordering
func (g Generator) GenerateOrdered(
	packageName string,
	order []string,
	allDecls map[string]astpkg.DeclNode,
	imports []string,
) (string, error) {
	// Collect declarations in order using fold
	orderedDecls := fold.FoldLeft(
		func(acc []astpkg.DeclNode, name string) []astpkg.DeclNode {
			if decl, ok := allDecls[name]; ok {
				return append(acc, decl)
			}
			return acc
		},
		[]astpkg.DeclNode{},
		order,
	)

	return g.GenerateFile(packageName, imports, orderedDecls)
}

// GenerateWithComments generates code with dependency comments
func (g Generator) GenerateWithComments(
	packageName string,
	targetName string,
	allDecls map[string]astpkg.DeclNode,
	deps astpkg.Dependencies,
) (string, error) {
	codeMonoid := NewCodeMonoid()

	// Header
	headerCode := generateHeader(targetName)

	// Package
	packageCode := ToCode("package " + packageName)

	// Imports
	importsCode := g.generateImports(deps.Imports.ToSlice())

	// Stats
	statsCode := generateStats(deps)

	// Target declaration
	targetCode := Code{Lines: []string{}}
	if decl, ok := allDecls[targetName]; ok {
		targetCode = g.generateDecl(decl)
	}

	// Dependency declarations with comments
	depsCode := g.generateDeclsWithComments(deps, allDecls)

	// Combine all using monoid
	allCode := monoid.Reduce(
		codeMonoid,
		[]Code{
			headerCode,
			packageCode,
			importsCode,
			statsCode,
			targetCode,
			depsCode,
		},
	)

	result := allCode.Join("\n\n")

	// Format
	formatted, err := format.Source([]byte(result))
	if err != nil {
		return result, err
	}

	return string(formatted), nil
}

// generateHeader generates header comment
func generateHeader(targetName string) Code {
	lines := []string{
		"// Code generated by pureast - DO NOT EDIT.",
		"// Extracted dependencies for: " + targetName,
	}
	return Code{Lines: lines}
}

// generateStats generates statistics comment
func generateStats(deps astpkg.Dependencies) Code {
	lines := []string{
		"// Dependency Stats:",
		"//   Types: " + strconv.Itoa(deps.Types.Size()),
		"//   Functions: " + strconv.Itoa(deps.Functions.Size()),
		"//   Structs: " + strconv.Itoa(deps.Structs.Size()),
		"//   Interfaces: " + strconv.Itoa(deps.Interfaces.Size()),
	}
	return Code{Lines: lines}
}

// generateDeclsWithComments generates declarations with comments
func (g Generator) generateDeclsWithComments(
	deps astpkg.Dependencies,
	allDecls map[string]astpkg.DeclNode,
) Code {
	codeMonoid := NewCodeMonoid()

	// Map type names to commented declarations
	typeCodes := fold.Map(
		func(name string) Code {
			if decl, ok := allDecls[name]; ok {
				comment := ToCode("// Dependency: " + name)
				declCode := g.generateDecl(decl)
				return codeMonoid.Combine(comment, declCode)
			}
			return Code{Lines: []string{}}
		},
		deps.Types.ToSlice(),
	)

	// Combine all using monoid
	return monoid.Reduce(codeMonoid, typeCodes)
}

// GenerateReport generates a text report of dependencies
func (g Generator) GenerateReport(
	targetName string,
	deps astpkg.Dependencies,
) string {
	sections := []Code{
		generateReportHeader(targetName),
		generateReportSummary(deps),
		generateReportSection("Types", deps.Types.ToSlice()),
		generateReportSection("Functions", deps.Functions.ToSlice()),
		generateReportSection("Imports", deps.Imports.ToSlice()),
	}

	codeMonoid := NewCodeMonoid()
	allSections := monoid.Reduce(codeMonoid, sections)

	return allSections.Join("\n\n")
}

// generateReportHeader generates report header
func generateReportHeader(targetName string) Code {
	return ToCode("# Dependency Report for: " + targetName)
}

// generateReportSummary generates summary section
func generateReportSummary(deps astpkg.Dependencies) Code {
	lines := []string{
		"## Summary",
		"- Types: " + strconv.Itoa(deps.Types.Size()),
		"- Functions: " + strconv.Itoa(deps.Functions.Size()),
		"- Structs: " + strconv.Itoa(deps.Structs.Size()),
		"- Interfaces: " + strconv.Itoa(deps.Interfaces.Size()),
		"- Imports: " + strconv.Itoa(deps.Imports.Size()),
	}
	return Code{Lines: lines}
}

// generateReportSection generates a report section
func generateReportSection(title string, items []string) Code {
	if len(items) == 0 {
		return Code{Lines: []string{}}
	}

	// Map items to list items
	listItems := fold.Map(
		func(item string) string {
			return "- " + item
		},
		items,
	)

	lines := append([]string{"## " + title}, listItems...)
	return Code{Lines: lines}
}

// GenerateDOT generates DOT graph format for visualization
func (g Generator) GenerateDOT(
	targetName string,
	allDecls map[string]astpkg.DeclNode,
) string {
	lines := []string{
		"digraph dependencies {",
		"  rankdir=LR;",
		"  node [shape=box];",
		"",
	}

	// Generate edges for each declaration using fold
	edgeLines := fold.FoldLeft(
		func(acc []string, name string) []string {
			decl := allDecls[name]

			// Type edges
			typeEdges := generateEdges(name, decl.Deps.Types.ToSlice(), "type")

			// Function edges
			funcEdges := generateEdges(name, decl.Deps.Functions.ToSlice(), "func")

			return append(acc, append(typeEdges, funcEdges...)...)
		},
		[]string{},
		getKeys(allDecls),
	)

	lines = append(lines, edgeLines...)
	lines = append(lines, "}")

	// Join with newlines
	return monoid.Reduce(monoid.NewStringJoinMonoid("\n"), lines)
}

// generateEdges generates DOT edges
func generateEdges(from string, targets []string, label string) []string {
	return fold.Map(
		func(to string) string {
			return "  \"" + from + "\" -> \"" + to + "\" [label=\"" + label + "\"];"
		},
		targets,
	)
}

// getKeys extracts keys from map
func getKeys(m map[string]astpkg.DeclNode) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
