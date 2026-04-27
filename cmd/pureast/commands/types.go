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
)

type TypesArgs struct {
	FilePath   string
	OutputFile string
	Kind       string // all|struct|interface
	Format     string // go|md
	Functions  bool   // [deprecated] use `pureast dump --kind func`
	Methods    bool   // [deprecated] use `pureast dump --kind method`
	Exported   bool
	MaxTokens  int
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
		Short("Extract type definitions [deprecated — use 'dump --kind']").
		Long(`Extract type definitions (structs, interfaces, aliases) without
function bodies.

DEPRECATED: 'pureast types' is being phased out. Its functionality is
fully covered by 'pureast dump --kind <kind>', which is the canonical
verb for filtering by symbol kind. Migration:

  pureast types ./pkg                   →  pureast dump ./pkg --kind type
  pureast types ./pkg --kind struct     →  pureast dump ./pkg --kind struct
  pureast types ./pkg --kind interface  →  pureast dump ./pkg --kind interface

The 'types' verb still works and prints output as before. A future
release will remove it entirely. Switch your scripts and prompts now.

For functions and methods, prefer pureast dump --kind func or
--kind method, which is the canonical path going forward.

Filtering uses --kind:
  --kind all        (default) structs and interfaces
  --kind struct     only structs
  --kind interface  only interfaces

Examples:
  pureast types ./pkg
  pureast types ./pkg --kind struct
  pureast types ./pkg --kind interface --exported
  pureast types ./pkg --format md
  pureast types ./pkg --max-tokens 4000`).
		ParseArgs(parseTypesArgs).
		Action(typesAction).
		Build()

	cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
	cmd.Flags().String("kind", "all", "Filter: all|struct|interface")
	cmd.Flags().String("format", "go", "Output format: go|md")
	cmd.Flags().Bool("exported", false, "Only exported symbols")
	cmd.Flags().Int("max-tokens", 0, "Truncate output to fit token budget (0 = unbounded)")

	// Deprecated flags — kept working so existing scripts don't break.
	// The mutually-exclusive booleans were exactly the redundant-path
	// smell we set out to fix; --kind is the canonical knob.
	cmd.Flags().Bool("structs-only", false, "[deprecated] use --kind struct")
	cmd.Flags().Bool("interfaces-only", false, "[deprecated] use --kind interface")
	cmd.Flags().Bool("functions", false, "[deprecated] use `pureast dump --kind func`")
	cmd.Flags().Bool("methods", false, "[deprecated] use `pureast dump --kind method`")

	// Back-compat: --file deprecated alias for positional PATH
	cmd.Flags().StringP("file", "f", "", "[deprecated] use positional PATH")

	cmd.MarkFlagsMutuallyExclusive("structs-only", "interfaces-only")

	return cmd
}

func parseTypesArgs(cmd *cobra.Command, args []string) (TypesArgs, error) {
	// Top-level deprecation notice. Prints once per invocation, before
	// any output, so it's visible whether or not the user redirects
	// stdout. The verb still works — this is just a heads-up that the
	// path is going away. Suppress with TYPES_NO_DEPRECATION_WARN=1
	// for users who've migrated and want quiet output during the
	// transition window.
	if os.Getenv("TYPES_NO_DEPRECATION_WARN") == "" {
		fmt.Fprintln(os.Stderr,
			"warning: 'pureast types' is deprecated; use 'pureast dump --kind <kind>' instead")
	}

	path, err := resolvePath(cmd, args)
	if err != nil {
		return TypesArgs{}, err
	}
	output, _ := cmd.Flags().GetString("output")
	kind, _ := cmd.Flags().GetString("kind")
	format, _ := cmd.Flags().GetString("format")
	exported, _ := cmd.Flags().GetBool("exported")
	maxTokens, _ := cmd.Flags().GetInt("max-tokens")

	// Deprecated boolean flags are translated into the canonical --kind.
	// We warn once if either is explicitly set so the user knows they're
	// on a path that's going away.
	structsOnly, _ := cmd.Flags().GetBool("structs-only")
	interfacesOnly, _ := cmd.Flags().GetBool("interfaces-only")
	functions, _ := cmd.Flags().GetBool("functions")
	methods, _ := cmd.Flags().GetBool("methods")

	if structsOnly {
		warnDeprecated("--structs-only", "--kind struct")
		kind = "struct"
	}
	if interfacesOnly {
		warnDeprecated("--interfaces-only", "--kind interface")
		kind = "interface"
	}
	if functions {
		warnDeprecated("--functions", "pureast dump --kind func")
	}
	if methods {
		warnDeprecated("--methods", "pureast dump --kind method")
	}

	switch kind {
	case "all", "struct", "interface":
	default:
		return TypesArgs{}, fmt.Errorf(
			"invalid --kind %q (want: all|struct|interface)", kind)
	}
	if format != "go" && format != "md" {
		return TypesArgs{}, fmt.Errorf(
			"invalid --format %q (want: go|md)", format)
	}

	return TypesArgs{
		FilePath:   path,
		OutputFile: output,
		Kind:       kind,
		Format:     format,
		Functions:  functions,
		Methods:    methods,
		Exported:   exported,
		MaxTokens:  maxTokens,
	}, nil
}

// warnDeprecated emits a one-line stderr notice steering the user to the
// canonical flag. We keep the deprecated path working — the warning is
// the cost of using it.
func warnDeprecated(old, replacement string) {
	fmt.Fprintf(os.Stderr, "warning: %s is deprecated, use %s\n", old, replacement)
}

func typesAction(ctx context.Context, args TypesArgs) (cli.Output, error) {
	fset := token.NewFileSet()
	pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
	if err != nil {
		return cli.Output{}, fmt.Errorf("extract %s: %w", args.FilePath, err)
	}

	var output strings.Builder

	// Type extraction is the canonical job for this verb. The Kind
	// switch is the single knob — the old --structs-only/--interfaces-only
	// pair has been collapsed into one path.
	if !args.Functions && !args.Methods {
		var types []extract.TypeDeclaration
		switch args.Kind {
		case "struct":
			types = extract.ExtractAllStructs(pkgNode)
		case "interface":
			types = extract.ExtractAllInterfaces(pkgNode)
		default: // "all"
			types = extract.ExtractAllStructsAndInterfaces(pkgNode)
		}

		gen := codegen.NewGenerator(fset)
		code, err := gen.GenerateTypesOnly(
			pkgNode.Name,
			types,
			pkgNode.Deps.Imports.ToSlice(),
		)
		if err != nil {
			return cli.Output{}, fmt.Errorf("generate types: %w", err)
		}
		output.WriteString(code)
	}

	// Function/method extraction is on a deprecation path; preserved
	// here so existing scripts keep working until removal.
	if args.Functions || args.Methods {
		signatures, err := extractSignatures(args.FilePath, args.Exported, args.Functions, args.Methods)
		if err != nil {
			return cli.Output{}, fmt.Errorf("extract signatures: %w", err)
		}
		if output.Len() > 0 {
			output.WriteString("\n\n")
		}
		formatSignaturesOutput(&output, signatures)
	}

	code := output.String()

	// Apply budget before format wrapping so a markdown fence always
	// closes. Symbol-aware so output stays compilable.
	if args.MaxTokens > 0 {
		var truncated bool
		code, truncated = extract.TruncateSymbols(code, args.MaxTokens)
		if truncated {
			fmt.Fprintf(os.Stderr,
				"notice: types truncated to fit --max-tokens %d\n", args.MaxTokens)
		}
	}
	if args.Format == "md" {
		title := fmt.Sprintf("%s — types", pkgNode.Name)
		code = renderAsMarkdown(title, code)
	}

	if args.OutputFile != "" {
		if err := os.WriteFile(args.OutputFile, []byte(code), 0644); err != nil {
			return cli.Output{}, fmt.Errorf("write %s: %w", args.OutputFile, err)
		}
		return cli.Output{
			Text:     fmt.Sprintf("✓ Written to %s\n", args.OutputFile),
			ExitCode: 0,
		}, nil
	}

	return cli.Output{Text: code, ExitCode: 0}, nil
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
