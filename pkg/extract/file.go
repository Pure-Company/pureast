package extract

import (
	"go/ast"
	"go/parser"
	"go/token"

	astpkg "github.com/Pure-Company/pureast/pkg/ast"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
)

// ExtractFile extracts all declarations from a file
// Returns a FileNode with accumulated dependencies
func ExtractFile(file *ast.File) astpkg.FileNode {
	if file == nil {
		return astpkg.FileNode{
			Name:    "",
			File:    nil,
			Decls:   []astpkg.DeclNode{},
			Imports: []string{},
			Deps:    astpkg.NewDependencies(),
		}
	}

	// Extract imports first
	imports := extractImports(file)

	// Map each declaration to a DeclNode
	declNodes := fold.Map(
		func(decl ast.Decl) astpkg.DeclNode {
			return extractDeclNode(decl)
		},
		file.Decls,
	)

	// Combine all dependencies using monoid
	depMonoid := astpkg.NewDependencyMonoid()
	allDeps := fold.FoldLeft(
		func(acc astpkg.Dependencies, node astpkg.DeclNode) astpkg.Dependencies {
			return depMonoid.Combine(acc, node.Deps)
		},
		astpkg.NewDependencies(),
		declNodes,
	)

	return astpkg.FileNode{
		Name:    file.Name.Name,
		File:    file,
		Decls:   declNodes,
		Imports: imports,
		Deps:    allDeps,
	}
}

// extractImports extracts import paths from file
func extractImports(file *ast.File) []string {
	imports := []string{}

	for _, imp := range file.Imports {
		if imp.Path != nil {
			path := imp.Path.Value
			// Remove quotes
			path = path[1 : len(path)-1]
			imports = append(imports, path)
		}
	}

	return imports
}

// extractDeclNode creates a DeclNode from an ast.Decl
func extractDeclNode(decl ast.Decl) astpkg.DeclNode {
	name := extractDeclName(decl)

	// Extract dependencies using our visitor
	visitor := ExtractDecl(decl)
	deps := visitor(astpkg.NewDependencies())

	return astpkg.DeclNode{
		Name: name,
		Decl: decl,
		Deps: deps,
	}
}

// extractDeclName gets the name of a declaration
func extractDeclName(decl ast.Decl) string {
	switch d := decl.(type) {
	case *ast.GenDecl:
		if len(d.Specs) > 0 {
			switch spec := d.Specs[0].(type) {
			case *ast.TypeSpec:
				return spec.Name.Name
			case *ast.ValueSpec:
				if len(spec.Names) > 0 {
					return spec.Names[0].Name
				}
			}
		}
		return ""

	case *ast.FuncDecl:
		if d.Recv != nil {
			// Method: ReceiverType.MethodName
			recvType := extractReceiverType(d.Recv)
			return recvType + "." + d.Name.Name
		}
		return d.Name.Name

	default:
		return ""
	}
}

// extractReceiverType gets receiver type name
func extractReceiverType(recv *ast.FieldList) string {
	if recv == nil || len(recv.List) == 0 {
		return ""
	}

	field := recv.List[0]
	switch t := field.Type.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name
		}
	}
	return ""
}

// ExtractFileFromPath parses and extracts a file from filesystem
func ExtractFileFromPath(fset *token.FileSet, filename string) (astpkg.FileNode, error) {
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return astpkg.FileNode{}, err
	}

	return ExtractFile(file), nil
}

// ExtractPackageFiles extracts multiple files as a package
func ExtractPackageFiles(files []*ast.File) astpkg.PackageNode {
	if len(files) == 0 {
		return astpkg.PackageNode{
			Name:  "",
			Files: []astpkg.FileNode{},
			Deps:  astpkg.NewDependencies(),
		}
	}

	// Map files to FileNodes
	fileNodes := fold.Map(ExtractFile, files)

	// Combine all dependencies using monoid reduction
	depMonoid := astpkg.NewDependencyMonoid()
	allDeps := monoid.Reduce(
		depMonoid,
		fold.Map(func(f astpkg.FileNode) astpkg.Dependencies { return f.Deps }, fileNodes),
	)

	// Get package name from first file
	pkgName := ""
	if len(fileNodes) > 0 {
		pkgName = fileNodes[0].Name
	}

	return astpkg.PackageNode{
		Name:  pkgName,
		Files: fileNodes,
		Deps:  allDeps,
	}
}

// ExtractPackageFromPaths loads and extracts package from file paths
func ExtractPackageFromPaths(fset *token.FileSet, filenames []string) (astpkg.PackageNode, error) {
	// Parse all files
	files := []*ast.File{}
	for _, filename := range filenames {
		file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
		if err != nil {
			return astpkg.PackageNode{}, err
		}
		files = append(files, file)
	}

	return ExtractPackageFiles(files), nil
}

// ExtractPackageFromDir loads all .go files from a directory
func ExtractPackageFromDir(fset *token.FileSet, dir string) (astpkg.PackageNode, error) {
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return astpkg.PackageNode{}, err
	}

	// Get first package (usually only one)
	for _, pkg := range pkgs {
		files := []*ast.File{}
		for _, file := range pkg.Files {
			files = append(files, file)
		}
		return ExtractPackageFiles(files), nil
	}

	return astpkg.PackageNode{}, nil
}

// BuildDeclMap creates a map of declaration names to DeclNodes
// This is useful for dependency resolution
func BuildDeclMap(fileNode astpkg.FileNode) map[string]astpkg.DeclNode {
	declMap := make(map[string]astpkg.DeclNode)

	for _, decl := range fileNode.Decls {
		if decl.Name != "" {
			declMap[decl.Name] = decl
		}
	}

	return declMap
}

// BuildPackageDeclMap creates a map for entire package
func BuildPackageDeclMap(pkgNode astpkg.PackageNode) map[string]astpkg.DeclNode {
	declMap := make(map[string]astpkg.DeclNode)

	for _, file := range pkgNode.Files {
		for _, decl := range file.Decls {
			if decl.Name != "" {
				declMap[decl.Name] = decl
			}
		}
	}

	return declMap
}

// ExtractMethods extracts all methods for a given type
func ExtractMethods(typeName string, file *ast.File) []astpkg.MethodNode {
	if file == nil {
		return []astpkg.MethodNode{}
	}

	// Fold over declarations to find methods
	methods := fold.FoldLeft(
		func(acc []astpkg.MethodNode, decl ast.Decl) []astpkg.MethodNode {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok || funcDecl.Recv == nil {
				return acc
			}

			// Check if receiver matches
			recvType := extractReceiverType(funcDecl.Recv)
			if recvType != typeName {
				return acc
			}

			// Extract method dependencies
			visitor := ExtractFuncDecl(funcDecl)
			deps := visitor(astpkg.NewDependencies())

			method := astpkg.MethodNode{
				ReceiverType: recvType,
				MethodName:   funcDecl.Name.Name,
				Func:         funcDecl,
				Deps:         deps,
			}

			return append(acc, method)
		},
		[]astpkg.MethodNode{},
		file.Decls,
	)

	return methods
}

// FilterDecls filters declarations by predicate
func FilterDecls(pred func(astpkg.DeclNode) bool, decls []astpkg.DeclNode) []astpkg.DeclNode {
	return fold.Filter(pred, decls)
}

// FindDecl finds a declaration by name
func FindDecl(name string, decls []astpkg.DeclNode) *astpkg.DeclNode {
	for _, decl := range decls {
		if decl.Name == name {
			return &decl
		}
	}
	return nil
}
