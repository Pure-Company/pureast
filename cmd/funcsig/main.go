package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FunctionSignature struct {
	Package  string
	Name     string
	IsMethod bool
	Receiver string
	Params   string
	Returns  string
	FilePath string
	Line     int
	FullSig  string
}

func main() {
	var (
		dir       = flag.String("dir", ".", "Directory to scan")
		pattern   = flag.String("pattern", "", "Filter by function name (supports *)")
		method    = flag.Bool("methods", false, "Only show methods")
		functions = flag.Bool("functions", false, "Only show functions")
		pkg       = flag.String("pkg", "", "Filter by package name")
		exported  = flag.Bool("exported", false, "Only show exported functions")
	)

	flag.Parse()

	sigs, err := extractSignatures(*dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	sigs = filterSignatures(sigs, filterOptions{
		pattern:   *pattern,
		methods:   *method,
		functions: *functions,
		pkg:       *pkg,
		exported:  *exported,
	})

	sort.Slice(sigs, func(i, j int) bool {
		if sigs[i].Package != sigs[j].Package {
			return sigs[i].Package < sigs[j].Package
		}
		return sigs[i].Name < sigs[j].Name
	})

	output(sigs)
}

type filterOptions struct {
	pattern   string
	methods   bool
	functions bool
	pkg       string
	exported  bool
}

func extractSignatures(rootDir string) ([]FunctionSignature, error) {
	var sigs []FunctionSignature

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if strings.HasPrefix(d.Name(), ".") || d.Name() == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil
		}

		pkgName := file.Name.Name
		relPath, _ := filepath.Rel(rootDir, path)

		for _, decl := range file.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				sig := FunctionSignature{
					Package:  pkgName,
					Name:     funcDecl.Name.Name,
					FilePath: relPath,
					Line:     fset.Position(funcDecl.Pos()).Line,
				}

				// Build full signature from source
				sig.FullSig = extractSourceSignature(fset, file, funcDecl)

				if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
					sig.IsMethod = true
					sig.Receiver = formatFieldList(fset, file, funcDecl.Recv)
				}

				sigs = append(sigs, sig)
			}
		}

		return nil
	})

	return sigs, err
}

func extractSourceSignature(fset *token.FileSet, file *ast.File, fn *ast.FuncDecl) string {
	var buf strings.Builder

	if fn.Recv != nil {
		buf.WriteString("(")
		buf.WriteString(formatFieldList(fset, file, fn.Recv))
		buf.WriteString(") ")
	} else {
		buf.WriteString("func ")
	}

	buf.WriteString(fn.Name.Name)
	buf.WriteString(formatFieldList(fset, file, fn.Type.Params))

	if fn.Type.Results != nil && len(fn.Type.Results.List) > 0 {
		results := formatFieldList(fset, file, fn.Type.Results)
		if strings.Count(results, ",") == 0 && !strings.Contains(results, " ") {
			results = strings.TrimPrefix(strings.TrimSuffix(results, ")"), "(")
			buf.WriteString(" ")
			buf.WriteString(results)
		} else {
			buf.WriteString(" ")
			buf.WriteString(results)
		}
	}

	return buf.String()
}

func formatFieldList(fset *token.FileSet, file *ast.File, fl *ast.FieldList) string {
	if fl == nil || len(fl.List) == 0 {
		return "()"
	}

	var parts []string
	for _, field := range fl.List {
		typeStr := formatTypeExpr(fset, file, field.Type)

		if len(field.Names) == 0 {
			parts = append(parts, typeStr)
		} else {
			names := make([]string, len(field.Names))
			for i, n := range field.Names {
				names[i] = n.Name
			}
			parts = append(parts, strings.Join(names, ", ")+" "+typeStr)
		}
	}

	return "(" + strings.Join(parts, ", ") + ")"
}

func formatTypeExpr(fset *token.FileSet, file *ast.File, expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name

	case *ast.SelectorExpr:
		return formatTypeExpr(fset, file, t.X) + "." + t.Sel.Name

	case *ast.StarExpr:
		return "*" + formatTypeExpr(fset, file, t.X)

	case *ast.ArrayType:
		return "[]" + formatTypeExpr(fset, file, t.Elt)

	case *ast.MapType:
		return "map[" + formatTypeExpr(fset, file, t.Key) + "]" + formatTypeExpr(fset, file, t.Value)

	case *ast.ChanType:
		return "chan " + formatTypeExpr(fset, file, t.Value)

	case *ast.FuncType:
		return "func" + formatFieldList(fset, file, t.Params) + " " + formatFieldList(fset, file, t.Results)

	case *ast.InterfaceType:
		return "interface{}"

	case *ast.StructType:
		return "struct{}"

	case *ast.IndexExpr:
		// Generic type like []T or Map[K]V
		return formatTypeExpr(fset, file, t.X) + "[" + formatTypeExpr(fset, file, t.Index) + "]"

	case *ast.IndexListExpr:
		// Multiple type parameters like Map[K, V]
		var params []string
		for _, idx := range t.Indices {
			params = append(params, formatTypeExpr(fset, file, idx))
		}
		return formatTypeExpr(fset, file, t.X) + "[" + strings.Join(params, ", ") + "]"

	default:
		return "unknown"
	}
}

func filterSignatures(sigs []FunctionSignature, opts filterOptions) []FunctionSignature {
	var filtered []FunctionSignature

	for _, sig := range sigs {
		if opts.methods && !sig.IsMethod {
			continue
		}
		if opts.functions && sig.IsMethod {
			continue
		}

		if opts.pkg != "" && sig.Package != opts.pkg {
			continue
		}

		if opts.exported && (len(sig.Name) == 0 || sig.Name[0] < 'A' || sig.Name[0] > 'Z') {
			continue
		}

		if opts.pattern != "" && !matchPattern(sig.Name, opts.pattern) {
			continue
		}

		filtered = append(filtered, sig)
	}

	return filtered
}

func matchPattern(name, pattern string) bool {
	parts := strings.Split(pattern, "*")

	if len(parts) == 1 {
		return name == pattern
	}

	if parts[0] != "" && !strings.HasPrefix(name, parts[0]) {
		return false
	}

	if parts[len(parts)-1] != "" && !strings.HasSuffix(name, parts[len(parts)-1]) {
		return false
	}

	return true
}

func output(sigs []FunctionSignature) {
	if len(sigs) == 0 {
		fmt.Println("No functions found")
		return
	}

	for _, sig := range sigs {
		fmt.Println(sig.FullSig)
	}
}
