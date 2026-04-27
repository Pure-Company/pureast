package extract

import (
	"go/ast"
	"go/token"

	astpkg "github.com/Pure-Company/pureast/pkg/ast"
	"github.com/Pure-Company/purekernels/pkg/fold"
	"github.com/Pure-Company/purekernels/pkg/monoid"
)

// TypeKind represents the kind of type declaration
type TypeKind int

const (
	TypeKindStruct TypeKind = iota
	TypeKindInterface
	TypeKindAlias
	TypeKindOther
)

// TypeDeclaration represents a type with its kind
type TypeDeclaration struct {
	Name string
	Kind TypeKind
	Decl astpkg.DeclNode
}

// extractTypesFromFile extracts types from a single file
func extractTypesFromFile(fileNode astpkg.FileNode) []TypeDeclaration {
	return fold.FoldLeft(
		func(acc []TypeDeclaration, declNode astpkg.DeclNode) []TypeDeclaration {
			if typeDecl := extractTypeDeclaration(declNode); typeDecl != nil {
				return append(acc, *typeDecl)
			}
			return acc
		},
		[]TypeDeclaration{},
		fileNode.Decls,
	)
}

// extractTypeDeclaration extracts type info from a declaration node
func extractTypeDeclaration(declNode astpkg.DeclNode) *TypeDeclaration {
	genDecl, ok := declNode.Decl.(*ast.GenDecl)
	if !ok || genDecl.Tok != token.TYPE {
		return nil
	}

	if len(genDecl.Specs) == 0 {
		return nil
	}

	typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return nil
	}

	kind := classifyType(typeSpec.Type)

	return &TypeDeclaration{
		Name: typeSpec.Name.Name,
		Kind: kind,
		Decl: declNode,
	}
}

// classifyType determines the kind of type
func classifyType(expr ast.Expr) TypeKind {
	switch expr.(type) {
	case *ast.StructType:
		return TypeKindStruct
	case *ast.InterfaceType:
		return TypeKindInterface
	default:
		return TypeKindOther
	}
}

// FilterStructs filters to only struct types (pure function)
func FilterStructs(types []TypeDeclaration) []TypeDeclaration {
	return fold.Filter(
		func(t TypeDeclaration) bool {
			return t.Kind == TypeKindStruct
		},
		types,
	)
}

// FilterInterfaces filters to only interface types (pure function)
func FilterInterfaces(types []TypeDeclaration) []TypeDeclaration {
	return fold.Filter(
		func(t TypeDeclaration) bool {
			return t.Kind == TypeKindInterface
		},
		types,
	)
}

// FilterStructsAndInterfaces filters to structs and interfaces (pure function)
func FilterStructsAndInterfaces(types []TypeDeclaration) []TypeDeclaration {
	return fold.Filter(
		func(t TypeDeclaration) bool {
			return t.Kind == TypeKindStruct || t.Kind == TypeKindInterface
		},
		types,
	)
}

// ExtractAllStructs extracts only struct declarations
func ExtractAllStructs(pkgNode astpkg.PackageNode) []TypeDeclaration {
	allTypes := ExtractAllTypes(pkgNode)
	return FilterStructs(allTypes)
}

// ExtractAllInterfaces extracts only interface declarations
func ExtractAllInterfaces(pkgNode astpkg.PackageNode) []TypeDeclaration {
	allTypes := ExtractAllTypes(pkgNode)
	return FilterInterfaces(allTypes)
}

// ExtractAllStructsAndInterfaces extracts structs and interfaces
func ExtractAllStructsAndInterfaces(pkgNode astpkg.PackageNode) []TypeDeclaration {
	allTypes := ExtractAllTypes(pkgNode)
	return FilterStructsAndInterfaces(allTypes)
}

// TypeSummary provides a summary of type declarations
type TypeSummary struct {
	TotalTypes     int
	StructCount    int
	InterfaceCount int
	OtherCount     int
	Names          []string
}

// SummarizeTypes creates a summary (pure function using folds)
func SummarizeTypes(types []TypeDeclaration) TypeSummary {
	// Count by kind using fold
	counts := fold.FoldLeft(
		func(acc map[TypeKind]int, t TypeDeclaration) map[TypeKind]int {
			acc[t.Kind]++
			return acc
		},
		make(map[TypeKind]int),
		types,
	)

	// Extract names using map
	names := fold.Map(
		func(t TypeDeclaration) string {
			return t.Name
		},
		types,
	)

	return TypeSummary{
		TotalTypes:     len(types),
		StructCount:    counts[TypeKindStruct],
		InterfaceCount: counts[TypeKindInterface],
		OtherCount:     counts[TypeKindOther],
		Names:          names,
	}
}

// DeduplicateTypes removes duplicate type declarations by name (pure function)
func DeduplicateTypes(types []TypeDeclaration) []TypeDeclaration {
	seen := monoid.NewSetMonoid[string]()

	return fold.FoldLeft(
		func(acc []TypeDeclaration, t TypeDeclaration) []TypeDeclaration {
			if seen.Contains(t.Name) {
				return acc
			}
			seen = seen.Insert(t.Name)
			return append(acc, t)
		},
		[]TypeDeclaration{},
		types,
	)
}

// ExtractAllTypes extracts all type declarations from a package
func ExtractAllTypes(pkgNode astpkg.PackageNode) []TypeDeclaration {
	// Fold over all files to collect types
	allTypes := fold.FoldLeft(
		func(acc []TypeDeclaration, fileNode astpkg.FileNode) []TypeDeclaration {
			fileTypes := extractTypesFromFile(fileNode)
			return append(acc, fileTypes...)
		},
		[]TypeDeclaration{},
		pkgNode.Files,
	)

	// Deduplicate before returning
	return DeduplicateTypes(allTypes)
}
