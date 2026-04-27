package extract

import (
	"go/ast"
	"regexp"
	"sort"
	"strings"

	astpkg "github.com/Pure-Company/pureast/pkg/ast"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
)

// SymbolInfo represents a discovered symbol.
//
// Decl is the underlying go/ast declaration. It's populated from the
// upstream PackageNode and is the entry point for any consumer that
// needs richer information than name/kind — render signatures, look
// up file:line via a FileSet, walk doc comments, etc. Treat it as
// read-only; mutating shared AST nodes would corrupt the input.
type SymbolInfo struct {
	Name     string
	Kind     string // "struct", "interface", "function", "method", "const", "var"
	Package  string
	Receiver string   // For methods: the receiver type
	Decl     ast.Decl // Underlying AST declaration (nil only for synthesized symbols)
}

// DiscoverAllSymbols finds all symbols in a package
func DiscoverAllSymbols(pkgNode astpkg.PackageNode) []SymbolInfo {
	// Fold over all files
	allSymbols := fold.FoldLeft(
		func(acc []SymbolInfo, fileNode astpkg.FileNode) []SymbolInfo {
			fileSymbols := discoverSymbolsInFile(fileNode, pkgNode.Name)
			return append(acc, fileSymbols...)
		},
		[]SymbolInfo{},
		pkgNode.Files,
	)

	// Sort by name for better display
	sort.Slice(allSymbols, func(i, j int) bool {
		return allSymbols[i].Name < allSymbols[j].Name
	})

	return allSymbols
}

// discoverSymbolsInFile extracts symbols from a file
func discoverSymbolsInFile(fileNode astpkg.FileNode, packageName string) []SymbolInfo {
	return fold.FoldLeft(
		func(acc []SymbolInfo, declNode astpkg.DeclNode) []SymbolInfo {
			symbols := extractSymbolInfo(declNode, packageName)
			return append(acc, symbols...)
		},
		[]SymbolInfo{},
		fileNode.Decls,
	)
}

// extractSymbolInfo extracts symbol information from a declaration
func extractSymbolInfo(declNode astpkg.DeclNode, packageName string) []SymbolInfo {
	name := declNode.Name
	if name == "" {
		return []SymbolInfo{}
	}

	// Determine kind
	kind := "unknown"
	receiver := ""

	// Check if it's a method (contains a dot)
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		if len(parts) == 2 {
			receiver = parts[0]
			name = parts[1]
			kind = "method"
		}
	} else {
		// Check in dependencies what kind it is
		if declNode.Deps.Structs.Contains(name) {
			kind = "struct"
		} else if declNode.Deps.Interfaces.Contains(name) {
			kind = "interface"
		} else if declNode.Deps.Functions.Contains(name) {
			kind = "function"
		} else if declNode.Deps.Types.Contains(name) {
			kind = "type"
		}

		// Better heuristic: check the actual declaration
		kind = inferKindFromDecl(declNode)
	}

	return []SymbolInfo{{
		Name:     name,
		Kind:     kind,
		Package:  packageName,
		Receiver: receiver,
		Decl:     declNode.Decl,
	}}
}

// inferKindFromDecl infers the kind from the declaration by inspecting
// the underlying ast.Decl directly. Earlier versions of this function
// inferred from the dependency set, which gave the wrong answer for
// types whose own dependencies don't reference themselves (e.g. the
// struct User, whose Deps doesn't contain "User"). The AST inspection
// is unambiguous: the declaration node either is a struct, interface,
// func, etc., or it isn't.
func inferKindFromDecl(declNode astpkg.DeclNode) string {
	switch d := declNode.Decl.(type) {
	case *ast.FuncDecl:
		// Methods carry a receiver list; bare functions don't. The
		// caller already strips "Type." prefixes from the name and
		// sets kind="method" before reaching this branch in the
		// happy path, but we double-check here for callers (like
		// the MCP renderer) that hand us DeclNodes directly.
		if d.Recv != nil && len(d.Recv.List) > 0 {
			return "method"
		}
		return "function"
	case *ast.GenDecl:
		// GenDecl groups specs (type, const, var, import). For our
		// purposes only the first spec's shape matters because the
		// extractor splits multi-spec blocks into separate DeclNodes.
		if len(d.Specs) == 0 {
			return "type"
		}
		switch s := d.Specs[0].(type) {
		case *ast.TypeSpec:
			switch s.Type.(type) {
			case *ast.StructType:
				return "struct"
			case *ast.InterfaceType:
				return "interface"
			}
			return "type"
		case *ast.ValueSpec:
			if d.Tok.String() == "const" {
				return "const"
			}
			return "var"
		case *ast.ImportSpec:
			return "import"
		}
	}
	return "unknown"
}

// MatchSymbols filters symbols by pattern (regex or exact match)
func MatchSymbols(pattern string, symbols []SymbolInfo) []SymbolInfo {
	// Try to compile as regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		// Not a valid regex, treat as exact match
		return fold.Filter(
			func(s SymbolInfo) bool {
				return s.Name == pattern
			},
			symbols,
		)
	}

	// Use regex matching
	return fold.Filter(
		func(s SymbolInfo) bool {
			return re.MatchString(s.Name)
		},
		symbols,
	)
}

// MatchMultipleSymbols matches multiple patterns (comma-separated)
func MatchMultipleSymbols(patterns string, symbols []SymbolInfo) []SymbolInfo {
	if patterns == "" {
		return []SymbolInfo{}
	}

	// Split by comma
	patternList := strings.Split(patterns, ",")

	// Use SetMonoid to deduplicate results
	matchedNames := monoid.NewSetMonoid[string]()
	matched := []SymbolInfo{}

	for _, pattern := range patternList {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}

		// Match this pattern
		results := MatchSymbols(pattern, symbols)

		// Add to results (deduplicating by name)
		for _, result := range results {
			if !matchedNames.Contains(result.Name) {
				matchedNames = matchedNames.Insert(result.Name)
				matched = append(matched, result)
			}
		}
	}

	return matched
}

// GroupSymbolsByKind groups symbols by their kind
func GroupSymbolsByKind(symbols []SymbolInfo) map[string][]SymbolInfo {
	return fold.FoldLeft(
		func(acc map[string][]SymbolInfo, s SymbolInfo) map[string][]SymbolInfo {
			acc[s.Kind] = append(acc[s.Kind], s)
			return acc
		},
		make(map[string][]SymbolInfo),
		symbols,
	)
}

// FilterByKind filters symbols by kind
func FilterByKind(kind string, symbols []SymbolInfo) []SymbolInfo {
	return fold.Filter(
		func(s SymbolInfo) bool {
			return s.Kind == kind
		},
		symbols,
	)
}

// FormatSymbolList formats symbols for display
func FormatSymbolList(symbols []SymbolInfo, groupByKind bool) string {
	if len(symbols) == 0 {
		return "No symbols found."
	}

	if !groupByKind {
		// Simple list
		lines := fold.Map(
			func(s SymbolInfo) string {
				if s.Receiver != "" {
					return s.Receiver + "." + s.Name + " (" + s.Kind + ")"
				}
				return s.Name + " (" + s.Kind + ")"
			},
			symbols,
		)
		return strings.Join(lines, "\n")
	}

	// Group by kind
	grouped := GroupSymbolsByKind(symbols)

	var result strings.Builder

	// Order: structs, interfaces, functions, methods, other
	order := []string{"struct", "interface", "function", "method", "type", "const", "var"}

	for _, kind := range order {
		items := grouped[kind]
		if len(items) == 0 {
			continue
		}

		result.WriteString("\n")
		result.WriteString(strings.ToUpper(kind))
		result.WriteString("S:\n")

		for _, item := range items {
			result.WriteString("  ")
			if item.Receiver != "" {
				result.WriteString(item.Receiver)
				result.WriteString(".")
			}
			result.WriteString(item.Name)
			result.WriteString("\n")
		}
	}

	return result.String()
}
