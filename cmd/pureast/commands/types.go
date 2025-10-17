// cmd/pureast/commands/types.go
package commands

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinodhalaharvi/pureast/pkg/cli"
	"github.com/vinodhalaharvi/pureast/pkg/codegen"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
)

type TypesArgs struct {
	FilePath       string
	OutputFile     string
	StructsOnly    bool
	InterfacesOnly bool
	Functions      bool
	Methods        bool
	Exported       bool
}

type TypeDefinition struct {
	Name       string
	Kind       string // struct, interface, type alias
	FilePath   string
	Line       int
	Definition string
}

type FunctionSignature struct {
	Name      string
	IsMethod  bool
	Receiver  string
	Signature string
	FilePath  string
	Line      int
}

func NewTypesCommand() *cobra.Command {
	cmd := cli.NewCommand[TypesArgs]("types").
		Short("Extract type definitions and function signatures").
		Long(`Extract type definitions, interfaces, and function signatures without implementations.
Perfect for providing clean context to LLMs.

By default extracts:
  - All structs and interfaces
  
Options:
  --structs-only     Extract only struct types
  --interfaces-only  Extract only interface types
  --functions        Extract function signatures
  --methods          Extract method signatures
  --exported         Only exported symbols

Examples:
  pureast types --file ./pkg
  pureast types --file ./pkg --structs-only
  pureast types --file ./pkg --interfaces-only
  pureast types --file ./pkg --functions --exported
  pureast types --file ./pkg --functions --methods`).
		ParseArgs(parseTypesArgs).
		Action(typesAction).
		Build()

	cmd.Flags().StringP("file", "f", "", "Go file or directory (required)")
	cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
	cmd.Flags().Bool("structs-only", false, "Extract only structs")
	cmd.Flags().Bool("interfaces-only", false, "Extract only interfaces")
	cmd.Flags().Bool("functions", false, "Extract function signatures")
	cmd.Flags().Bool("methods", false, "Extract method signatures")
	cmd.Flags().Bool("exported", false, "Only exported symbols")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagsMutuallyExclusive("structs-only", "interfaces-only")

	return cmd
}

func parseTypesArgs(cmd *cobra.Command, args []string) result.Result[TypesArgs] {
	file, _ := cmd.Flags().GetString("file")
	output, _ := cmd.Flags().GetString("output")
	structsOnly, _ := cmd.Flags().GetBool("structs-only")
	interfacesOnly, _ := cmd.Flags().GetBool("interfaces-only")
	functions, _ := cmd.Flags().GetBool("functions")
	methods, _ := cmd.Flags().GetBool("methods")
	exported, _ := cmd.Flags().GetBool("exported")

	return result.Ok(TypesArgs{
		FilePath:       file,
		OutputFile:     output,
		StructsOnly:    structsOnly,
		InterfacesOnly: interfacesOnly,
		Functions:      functions,
		Methods:        methods,
		Exported:       exported,
	})
}

func typesAction(ctx context.Context, args TypesArgs) result.Result[cli.Output] {
	fset := token.NewFileSet()
	pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
	if err != nil {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("Error: %v\n", err),
			ExitCode: 1,
		})
	}

	var output strings.Builder

	// Extract types
	if !args.Functions && !args.Methods {
		var types []extract.TypeDeclaration
		if args.StructsOnly {
			types = extract.ExtractAllStructs(pkgNode)
		} else if args.InterfacesOnly {
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
			return result.Ok(cli.Output{
				Text:     fmt.Sprintf("Error: %v\n", err),
				ExitCode: 1,
			})
		}
		output.WriteString(code)
	}

	// Extract function signatures
	if args.Functions || args.Methods {
		signatures, err := extractSignatures(args.FilePath, args.Exported, args.Functions, args.Methods)
		if err != nil {
			return result.Ok(cli.Output{
				Text:     fmt.Sprintf("Error: %v\n", err),
				ExitCode: 1,
			})
		}

		if output.Len() > 0 {
			output.WriteString("\n\n")
		}

		formatSignaturesOutput(&output, signatures)
	}

	code := output.String()

	if args.OutputFile != "" {
		if err := os.WriteFile(args.OutputFile, []byte(code), 0644); err != nil {
			return result.Ok(cli.Output{
				Text:     fmt.Sprintf("Error writing file: %v\n", err),
				ExitCode: 1,
			})
		}
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("✓ Written to %s\n", args.OutputFile),
			ExitCode: 0,
		})
	}

	return result.Ok(cli.Output{Text: code, ExitCode: 0})
}

func extractSignatures(rootDir string, exportedOnly, includeFunctions, includeMethods bool) ([]FunctionSignature, error) {
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

		relPath, _ := filepath.Rel(rootDir, path)

		for _, decl := range file.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				// Skip if export filter is enabled
				if exportedOnly && (len(funcDecl.Name.Name) == 0 || funcDecl.Name.Name[0] < 'A' || funcDecl.Name.Name[0] > 'Z') {
					continue
				}

				isMethod := funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0

				// Filter by function/method type
				if includeFunctions && isMethod {
					continue
				}
				if includeMethods && !isMethod {
					continue
				}

				sig := FunctionSignature{
					Name:      funcDecl.Name.Name,
					IsMethod:  isMethod,
					FilePath:  relPath,
					Line:      fset.Position(funcDecl.Pos()).Line,
					Signature: buildSignature(file, funcDecl),
				}

				if isMethod {
					sig.Receiver = formatFieldList(file, funcDecl.Recv)
				}

				sigs = append(sigs, sig)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Sort by file, then line
	sort.Slice(sigs, func(i, j int) bool {
		if sigs[i].FilePath != sigs[j].FilePath {
			return sigs[i].FilePath < sigs[j].FilePath
		}
		return sigs[i].Line < sigs[j].Line
	})

	return sigs, nil
}

func buildSignature(file *ast.File, fn *ast.FuncDecl) string {
	var buf strings.Builder

	if fn.Recv != nil {
		buf.WriteString("(")
		buf.WriteString(formatFieldList(file, fn.Recv))
		buf.WriteString(") ")
	} else {
		buf.WriteString("func ")
	}

	buf.WriteString(fn.Name.Name)
	buf.WriteString(formatFieldList(file, fn.Type.Params))

	if fn.Type.Results != nil && len(fn.Type.Results.List) > 0 {
		results := formatFieldList(file, fn.Type.Results)
		// Single unnamed return type
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

func formatFieldList(file *ast.File, fl *ast.FieldList) string {
	if fl == nil || len(fl.List) == 0 {
		return "()"
	}

	var parts []string
	for _, field := range fl.List {
		typeStr := formatType(field.Type)

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

func formatType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return formatType(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + formatType(t.X)
	case *ast.ArrayType:
		return "[]" + formatType(t.Elt)
	case *ast.MapType:
		return "map[" + formatType(t.Key) + "]" + formatType(t.Value)
	case *ast.ChanType:
		return "chan " + formatType(t.Value)
	case *ast.FuncType:
		return "func" + formatFieldListSimple(t.Params) + " " + formatFieldListSimple(t.Results)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.IndexExpr:
		return formatType(t.X) + "[" + formatType(t.Index) + "]"
	case *ast.IndexListExpr:
		var params []string
		for _, idx := range t.Indices {
			params = append(params, formatType(idx))
		}
		return formatType(t.X) + "[" + strings.Join(params, ", ") + "]"
	default:
		return "unknown"
	}
}

func formatFieldListSimple(fl *ast.FieldList) string {
	if fl == nil || len(fl.List) == 0 {
		return "()"
	}

	var parts []string
	for _, field := range fl.List {
		parts = append(parts, formatType(field.Type))
	}

	return "(" + strings.Join(parts, ", ") + ")"
}

func formatSignaturesOutput(buf *strings.Builder, sigs []FunctionSignature) {
	if len(sigs) == 0 {
		return
	}

	buf.WriteString("// Function Signatures\n\n")

	currentFile := ""
	for _, sig := range sigs {
		if sig.FilePath != currentFile {
			currentFile = sig.FilePath
			buf.WriteString(fmt.Sprintf("// %s\n", currentFile))
		}

		buf.WriteString(sig.Signature)
		buf.WriteString("\n")
	}
}
