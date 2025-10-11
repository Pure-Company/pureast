package extract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/functor"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
)

// FileDiscovery finds all Go files in a directory tree
type FileDiscovery struct {
	Root      string
	Recursive bool
}

// NewFileDiscovery creates a file discovery instance
func NewFileDiscovery(root string, recursive bool) FileDiscovery {
	return FileDiscovery{
		Root:      root,
		Recursive: recursive,
	}
}

// DiscoverFiles finds all .go files (pure function - returns paths)
func (d FileDiscovery) DiscoverFiles() ([]string, error) {
	files := []string{}

	err := filepath.Walk(d.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories if not recursive
		if info.IsDir() {
			if !d.Recursive && path != d.Root {
				return filepath.SkipDir
			}
			return nil
		}

		// Only include .go files, skip test files
		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// ParseFileConcurrent creates a concurrent computation for parsing a file
func ParseFileConcurrent(fset *token.FileSet, filename string) functor.Concurrent[astpkg.FileNode] {
	fileNodeMonoid := NewFileNodeMonoid()

	return functor.NewConcurrent(
		fileNodeMonoid,
		func() astpkg.FileNode {
			file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
			if err != nil {
				// Return empty on error
				return astpkg.FileNode{
					Name:    "",
					File:    nil,
					Decls:   []astpkg.DeclNode{},
					Imports: []string{},
					Deps:    astpkg.NewDependencies(),
				}
			}

			return ExtractFile(file)
		},
	)
}

// FileNodeMonoid combines FileNodes
type FileNodeMonoid struct{}

func NewFileNodeMonoid() FileNodeMonoid {
	return FileNodeMonoid{}
}

func (FileNodeMonoid) Empty() astpkg.FileNode {
	return astpkg.FileNode{
		Name:    "",
		File:    nil,
		Decls:   []astpkg.DeclNode{},
		Imports: []string{},
		Deps:    astpkg.NewDependencies(),
	}
}

func (m FileNodeMonoid) Combine(a, b astpkg.FileNode) astpkg.FileNode {
	depMonoid := astpkg.NewDependencyMonoid()

	// Combine declarations
	combinedDecls := append(a.Decls, b.Decls...)

	// Combine imports (use set to deduplicate)
	importSet := monoid.NewSetMonoid[string]()
	for _, imp := range a.Imports {
		importSet = importSet.Insert(imp)
	}
	for _, imp := range b.Imports {
		importSet = importSet.Insert(imp)
	}

	return astpkg.FileNode{
		Name:    selectPackageName(a.Name, b.Name),
		File:    selectASTFile(a.File, b.File),
		Decls:   combinedDecls,
		Imports: importSet.ToSlice(),
		Deps:    depMonoid.Combine(a.Deps, b.Deps),
	}
}

func selectPackageName(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func selectASTFile(a, b *ast.File) *ast.File {
	if a != nil {
		return a
	}
	return b
}

func selectFile(a, b interface{}) interface{} {
	if a != nil {
		return a
	}
	return b
}

// ExtractDirectoryConcurrent extracts all files in directory using concurrent applicatives
func ExtractDirectoryConcurrent(
	fset *token.FileSet,
	root string,
	recursive bool,
	workers int,
) (astpkg.PackageNode, error) {
	// Discover files (pure)
	discovery := NewFileDiscovery(root, recursive)
	filePaths, err := discovery.DiscoverFiles()
	if err != nil {
		return astpkg.PackageNode{}, err
	}

	if len(filePaths) == 0 {
		return astpkg.PackageNode{
			Name:  "",
			Files: []astpkg.FileNode{},
			Deps:  astpkg.NewDependencies(),
		}, nil
	}

	// Use TraverseConcurrent to parse files in parallel
	fileNodeMonoid := NewFileNodeMonoid()

	combinedNode := functor.TraverseConcurrent(
		fileNodeMonoid,
		func(path string) astpkg.FileNode {
			file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				return fileNodeMonoid.Empty()
			}
			return ExtractFile(file)
		},
		filePaths,
		workers, // 0 = use runtime.NumCPU()
	).Value()

	// Convert combined FileNode to PackageNode
	return astpkg.PackageNode{
		Name:  combinedNode.Name,
		Files: []astpkg.FileNode{combinedNode},
		Deps:  combinedNode.Deps,
	}, nil
}

// ExtractDirectoryConcurrentBatched uses batched concurrent processing
func ExtractDirectoryConcurrentBatched(
	fset *token.FileSet,
	root string,
	recursive bool,
	batchSize int,
) (astpkg.PackageNode, error) {
	// Discover files
	discovery := NewFileDiscovery(root, recursive)
	filePaths, err := discovery.DiscoverFiles()
	if err != nil {
		return astpkg.PackageNode{}, err
	}

	if len(filePaths) == 0 {
		return astpkg.PackageNode{
			Name:  "",
			Files: []astpkg.FileNode{},
			Deps:  astpkg.NewDependencies(),
		}, nil
	}

	// Use ConcurrentBatch for processing
	fileNodeMonoid := NewFileNodeMonoid()

	processBatch := func(batch []string) astpkg.FileNode {
		// Process batch sequentially (each batch runs in parallel)
		batchResults := fold.Map(
			func(path string) astpkg.FileNode {
				file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
				if err != nil {
					return fileNodeMonoid.Empty()
				}
				return ExtractFile(file)
			},
			batch,
		)

		// Combine batch results
		return monoid.Reduce(fileNodeMonoid, batchResults)
	}

	combinedNode := functor.ConcurrentBatch(
		fileNodeMonoid,
		batchSize,
		processBatch,
		filePaths,
	).Value()

	return astpkg.PackageNode{
		Name:  combinedNode.Name,
		Files: []astpkg.FileNode{combinedNode},
		Deps:  combinedNode.Deps,
	}, nil
}

// ExtractMultipleFilesConcurrent extracts specific files concurrently
func ExtractMultipleFilesConcurrent(
	fset *token.FileSet,
	filePaths []string,
	workers int,
) ([]astpkg.FileNode, error) {
	// Use ParMap for parallel processing
	fileNodes := functor.ParMapWithWorkers(
		func(path string) astpkg.FileNode {
			file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				return astpkg.FileNode{
					Name:    "",
					File:    nil,
					Decls:   []astpkg.DeclNode{},
					Imports: []string{},
					Deps:    astpkg.NewDependencies(),
				}
			}
			return ExtractFile(file)
		},
		filePaths,
		workers,
	)

	return fileNodes, nil
}

// BuildPackageFromFileNodes constructs PackageNode from multiple FileNodes
func BuildPackageFromFileNodes(fileNodes []astpkg.FileNode) astpkg.PackageNode {
	if len(fileNodes) == 0 {
		return astpkg.PackageNode{
			Name:  "",
			Files: []astpkg.FileNode{},
			Deps:  astpkg.NewDependencies(),
		}
	}

	// Combine all dependencies using monoid
	depMonoid := astpkg.NewDependencyMonoid()
	allDeps := fold.FoldLeft(
		func(acc astpkg.Dependencies, node astpkg.FileNode) astpkg.Dependencies {
			return depMonoid.Combine(acc, node.Deps)
		},
		astpkg.NewDependencies(),
		fileNodes,
	)

	return astpkg.PackageNode{
		Name:  fileNodes[0].Name,
		Files: fileNodes,
		Deps:  allDeps,
	}
}
