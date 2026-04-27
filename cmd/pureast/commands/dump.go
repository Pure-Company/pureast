// cmd/pureast/commands/dump.go
//
// Compact, LLM-friendly dump of every symbol in a package.
//
// Unlike `types` (only type declarations) or `extract` (one symbol + its
// transitive deps), `dump` walks the whole package and emits every top-level
// symbol the parser can see: structs, interfaces, type aliases, functions,
// methods, consts, vars.
//
// The default output is signature-only — bodies are stripped. This is the
// single biggest token-compression win: a 5000-line package collapses to
// a few hundred lines of "what exists, what it returns, what it satisfies."
// Use --bodies if you need the implementations too.
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
	"github.com/vinodhalaharvi/purekernels/pkg/result"
)

type DumpArgs struct {
	FilePath     string
	OutputFile   string
	Kind         string // all|type|struct|interface|func|method|const|var
	Format       string // go|md
	Bodies       bool   // include function bodies (default: signatures only)
	ExportedOnly bool
	IncludeTests bool
	IncludeDocs  bool
	MaxTokens    int // 0 = unbounded
}

type dumpedSymbol struct {
	Kind     string // struct, interface, type, func, method, const, var
	Name     string
	Receiver string // for methods
	Doc      string
	Source   string // signature or full text
	File     string
	Line     int
}

func NewDumpCommand() *cobra.Command {
	cmd := cli.NewCommand[DumpArgs]("dump").
		Short("Dump every symbol in a package (LLM context)").
		Long(`Dump all top-level symbols — types, functions, methods, consts, vars —
in a compact form suitable for feeding to an LLM as context.

By default, function bodies are stripped (signatures only). This typically
gives 5–20× compression versus pasting raw source files.

Examples:
  pureast dump ./pkg                        # everything, signatures only
  pureast dump ./pkg --bodies               # include implementations
  pureast dump ./pkg --kind func            # only functions
  pureast dump ./pkg --exported             # only exported symbols
  pureast dump ./pkg --format md            # markdown for LLM
  pureast dump ./pkg --max-tokens 4000      # fit a token budget
  pureast dump ./pkg -o context.txt`).
		ParseArgs(parseDumpArgs).
		Action(dumpAction).
		Build()

	cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
	cmd.Flags().String("kind", "all", "Filter: all|type|struct|interface|func|method|const|var")
	cmd.Flags().String("format", "go", "Output format: go|md")
	cmd.Flags().Bool("bodies", false, "Include function bodies (default: signatures only)")
	cmd.Flags().Bool("exported", false, "Only exported symbols")
	cmd.Flags().Bool("include-tests", false, "Include _test.go files")
	cmd.Flags().Bool("no-docs", false, "Strip doc comments")
	cmd.Flags().Int("max-tokens", 0, "Truncate output to fit token budget (0 = unbounded)")

	// --file kept as a back-compat alias; positional path is canonical
	cmd.Flags().StringP("file", "f", "", "[deprecated] use positional PATH")

	return cmd
}

func parseDumpArgs(cmd *cobra.Command, args []string) result.Result[DumpArgs] {
	path, err := resolvePath(cmd, args)
	if err != nil {
		return result.Err[DumpArgs](err)
	}

	output, _ := cmd.Flags().GetString("output")
	kind, _ := cmd.Flags().GetString("kind")
	format, _ := cmd.Flags().GetString("format")
	bodies, _ := cmd.Flags().GetBool("bodies")
	exported, _ := cmd.Flags().GetBool("exported")
	tests, _ := cmd.Flags().GetBool("include-tests")
	noDocs, _ := cmd.Flags().GetBool("no-docs")
	maxTokens, _ := cmd.Flags().GetInt("max-tokens")

	if !validDumpKind(kind) {
		return result.Err[DumpArgs](fmt.Errorf(
			"invalid --kind %q (want: all|type|struct|interface|func|method|const|var)", kind))
	}
	if format != "go" && format != "md" {
		return result.Err[DumpArgs](fmt.Errorf(
			"invalid --format %q (want: go|md)", format))
	}

	return result.Ok(DumpArgs{
		FilePath:     path,
		OutputFile:   output,
		Kind:         kind,
		Format:       format,
		Bodies:       bodies,
		ExportedOnly: exported,
		IncludeTests: tests,
		IncludeDocs:  !noDocs,
		MaxTokens:    maxTokens,
	})
}

func validDumpKind(k string) bool {
	switch k {
	case "all", "type", "struct", "interface", "func", "method", "const", "var":
		return true
	}
	return false
}

func dumpAction(ctx context.Context, args DumpArgs) result.Result[cli.Output] {
	symbols, pkgName, err := collectSymbols(args)
	if err != nil {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("Error: %v\n", err),
			ExitCode: 1,
		})
	}

	out := renderDump(pkgName, symbols, args)

	// Token budget is applied to the final rendered text, not the
	// individual symbols. This means truncation is line-aware (we cut
	// at line boundaries) and the header survives — the LLM still
	// orients itself even if the tail is missing.
	if args.MaxTokens > 0 {
		out, _ = truncateToBudget(out, args.MaxTokens)
	}

	// Markdown wrapping is applied last so the fence wraps whatever
	// fit inside the budget, including any truncation marker.
	if args.Format == "md" {
		title := fmt.Sprintf("Package %s", pkgName)
		out = renderAsMarkdown(title, out)
	}

	if args.OutputFile != "" {
		if err := os.WriteFile(args.OutputFile, []byte(out), 0644); err != nil {
			return result.Ok(cli.Output{
				Text:     fmt.Sprintf("Error writing file: %v\n", err),
				ExitCode: 1,
			})
		}
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("✓ Written %d symbols to %s\n", len(symbols), args.OutputFile),
			ExitCode: 0,
		})
	}

	return result.Ok(cli.Output{Text: out, ExitCode: 0})
}

func collectSymbols(args DumpArgs) ([]dumpedSymbol, string, error) {
	var (
		symbols []dumpedSymbol
		pkgName string
	)

	walkErr := filepath.WalkDir(args.FilePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") || name == "vendor" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		if !args.IncludeTests && strings.HasSuffix(path, "_test.go") {
			return nil
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if parseErr != nil {
			// skip unparseable files; don't abort the whole dump
			return nil
		}
		if pkgName == "" {
			pkgName = file.Name.Name
		}

		rel, _ := filepath.Rel(args.FilePath, path)
		for _, decl := range file.Decls {
			ss := extractFromDecl(fset, file, decl, rel, args)
			symbols = append(symbols, ss...)
		}
		return nil
	})

	if walkErr != nil {
		return nil, "", walkErr
	}

	// Stable order: file, then line
	sort.Slice(symbols, func(i, j int) bool {
		if symbols[i].File != symbols[j].File {
			return symbols[i].File < symbols[j].File
		}
		return symbols[i].Line < symbols[j].Line
	})

	return symbols, pkgName, nil
}

func extractFromDecl(fset *token.FileSet, file *ast.File, decl ast.Decl, rel string, args DumpArgs) []dumpedSymbol {
	var out []dumpedSymbol

	switch d := decl.(type) {

	case *ast.FuncDecl:
		isMethod := d.Recv != nil && len(d.Recv.List) > 0
		kind := "func"
		if isMethod {
			kind = "method"
		}
		if !kindAllowed(args.Kind, kind) {
			return nil
		}
		if args.ExportedOnly && !ast.IsExported(d.Name.Name) {
			return nil
		}

		sym := dumpedSymbol{
			Kind: kind,
			Name: d.Name.Name,
			File: rel,
			Line: fset.Position(d.Pos()).Line,
		}
		if args.IncludeDocs && d.Doc != nil {
			sym.Doc = d.Doc.Text()
		}
		if isMethod {
			sym.Receiver = renderFieldList(d.Recv)
		}
		sym.Source = renderFuncDecl(fset, d, args.Bodies)
		out = append(out, sym)

	case *ast.GenDecl:
		// type / const / var blocks may contain multiple specs
		for _, spec := range d.Specs {
			switch s := spec.(type) {

			case *ast.TypeSpec:
				kind := classifyType(s.Type)
				if !kindAllowed(args.Kind, kind) && !kindAllowed(args.Kind, "type") {
					continue
				}
				if args.ExportedOnly && !ast.IsExported(s.Name.Name) {
					continue
				}
				sym := dumpedSymbol{
					Kind: kind,
					Name: s.Name.Name,
					File: rel,
					Line: fset.Position(s.Pos()).Line,
				}
				if args.IncludeDocs {
					sym.Doc = combineDoc(d.Doc, s.Doc)
				}
				sym.Source = renderTypeSpec(s)
				out = append(out, sym)

			case *ast.ValueSpec:
				kind := "var"
				if d.Tok == token.CONST {
					kind = "const"
				}
				if !kindAllowed(args.Kind, kind) {
					continue
				}
				for _, name := range s.Names {
					if args.ExportedOnly && !ast.IsExported(name.Name) {
						continue
					}
					sym := dumpedSymbol{
						Kind: kind,
						Name: name.Name,
						File: rel,
						Line: fset.Position(name.Pos()).Line,
					}
					if args.IncludeDocs {
						sym.Doc = combineDoc(d.Doc, s.Doc)
					}
					sym.Source = renderValueSpec(kind, name.Name, s)
					out = append(out, sym)
				}
			}
		}
	}

	return out
}

func kindAllowed(filter, kind string) bool {
	if filter == "all" {
		return true
	}
	if filter == kind {
		return true
	}
	// "type" matches struct/interface/alias
	if filter == "type" && (kind == "struct" || kind == "interface" || kind == "type") {
		return true
	}
	return false
}

func classifyType(expr ast.Expr) string {
	switch expr.(type) {
	case *ast.StructType:
		return "struct"
	case *ast.InterfaceType:
		return "interface"
	default:
		return "type" // alias or named type
	}
}

func renderFuncDecl(fset *token.FileSet, d *ast.FuncDecl, includeBody bool) string {
	// When bodies are requested, hand the whole declaration to go/printer.
	// Manual reconstruction (used for signature mode) loses comments,
	// blank lines, and formatting nuance — fine for compact output, but
	// wrong when the user explicitly asked for the implementation.
	if includeBody && d.Body != nil {
		return printNode(fset, d)
	}

	var b strings.Builder
	if d.Recv != nil {
		b.WriteString("func ")
		b.WriteString(renderFieldList(d.Recv))
		b.WriteString(" ")
		b.WriteString(d.Name.Name)
	} else {
		b.WriteString("func ")
		b.WriteString(d.Name.Name)
	}
	b.WriteString(renderParams(d.Type.Params))
	if d.Type.Results != nil && len(d.Type.Results.List) > 0 {
		b.WriteString(" ")
		b.WriteString(renderResults(d.Type.Results))
	}
	return b.String()
}

func renderTypeSpec(s *ast.TypeSpec) string {
	return "type " + s.Name.Name + " " + renderTypeExpr(s.Type)
}

func renderValueSpec(kind, name string, s *ast.ValueSpec) string {
	var b strings.Builder
	b.WriteString(kind)
	b.WriteString(" ")
	b.WriteString(name)
	if s.Type != nil {
		b.WriteString(" ")
		b.WriteString(renderTypeExpr(s.Type))
	}
	return b.String()
}

func renderFieldList(fl *ast.FieldList) string {
	if fl == nil || len(fl.List) == 0 {
		return "()"
	}
	var parts []string
	for _, f := range fl.List {
		t := renderTypeExpr(f.Type)
		if len(f.Names) == 0 {
			parts = append(parts, t)
			continue
		}
		var names []string
		for _, n := range f.Names {
			names = append(names, n.Name)
		}
		parts = append(parts, strings.Join(names, ", ")+" "+t)
	}
	return "(" + strings.Join(parts, ", ") + ")"
}

func renderParams(fl *ast.FieldList) string {
	return renderFieldList(fl)
}

func renderResults(fl *ast.FieldList) string {
	if fl == nil || len(fl.List) == 0 {
		return ""
	}
	// Single unnamed result -> no parens
	if len(fl.List) == 1 && len(fl.List[0].Names) == 0 {
		return renderTypeExpr(fl.List[0].Type)
	}
	return renderFieldList(fl)
}

func renderTypeExpr(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return renderTypeExpr(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + renderTypeExpr(t.X)
	case *ast.ArrayType:
		return "[]" + renderTypeExpr(t.Elt)
	case *ast.MapType:
		return "map[" + renderTypeExpr(t.Key) + "]" + renderTypeExpr(t.Value)
	case *ast.ChanType:
		return "chan " + renderTypeExpr(t.Value)
	case *ast.FuncType:
		return "func" + renderFieldList(t.Params) + " " + renderFieldList(t.Results)
	case *ast.InterfaceType:
		return renderInterfaceShort(t)
	case *ast.StructType:
		return renderStructShort(t)
	case *ast.IndexExpr:
		return renderTypeExpr(t.X) + "[" + renderTypeExpr(t.Index) + "]"
	case *ast.IndexListExpr:
		var ps []string
		for _, idx := range t.Indices {
			ps = append(ps, renderTypeExpr(idx))
		}
		return renderTypeExpr(t.X) + "[" + strings.Join(ps, ", ") + "]"
	case *ast.Ellipsis:
		return "..." + renderTypeExpr(t.Elt)
	default:
		return "?"
	}
}

func renderStructShort(s *ast.StructType) string {
	if s.Fields == nil || len(s.Fields.List) == 0 {
		return "struct{}"
	}
	var lines []string
	for _, f := range s.Fields.List {
		t := renderTypeExpr(f.Type)
		if len(f.Names) == 0 {
			lines = append(lines, "  "+t)
			continue
		}
		var names []string
		for _, n := range f.Names {
			names = append(names, n.Name)
		}
		lines = append(lines, "  "+strings.Join(names, ", ")+" "+t)
	}
	return "struct {\n" + strings.Join(lines, "\n") + "\n}"
}

func renderInterfaceShort(i *ast.InterfaceType) string {
	if i.Methods == nil || len(i.Methods.List) == 0 {
		return "interface{}"
	}
	var lines []string
	for _, m := range i.Methods.List {
		if len(m.Names) == 0 {
			// embedded interface
			lines = append(lines, "  "+renderTypeExpr(m.Type))
			continue
		}
		// method
		fn, ok := m.Type.(*ast.FuncType)
		if !ok {
			continue
		}
		sig := m.Names[0].Name + renderFieldList(fn.Params)
		if fn.Results != nil && len(fn.Results.List) > 0 {
			sig += " " + renderResults(fn.Results)
		}
		lines = append(lines, "  "+sig)
	}
	return "interface {\n" + strings.Join(lines, "\n") + "\n}"
}

func combineDoc(a, b *ast.CommentGroup) string {
	switch {
	case a != nil && a.Text() != "":
		return a.Text()
	case b != nil && b.Text() != "":
		return b.Text()
	}
	return ""
}

func renderDump(pkgName string, symbols []dumpedSymbol, args DumpArgs) string {
	var b strings.Builder

	// Header — orientation for the LLM
	fmt.Fprintf(&b, "// pureast dump: package %s\n", pkgName)
	fmt.Fprintf(&b, "// %d symbols", len(symbols))
	if !args.Bodies {
		b.WriteString(" (signatures only)")
	}
	if args.ExportedOnly {
		b.WriteString(", exported only")
	}
	b.WriteString("\n\n")

	// Group by kind for readability — order: types, then funcs/methods, then values
	groups := map[string][]dumpedSymbol{}
	order := []string{"struct", "interface", "type", "func", "method", "const", "var"}
	for _, s := range symbols {
		groups[s.Kind] = append(groups[s.Kind], s)
	}

	headings := map[string]string{
		"struct":    "// === structs ===",
		"interface": "// === interfaces ===",
		"type":      "// === type aliases ===",
		"func":      "// === functions ===",
		"method":    "// === methods ===",
		"const":     "// === constants ===",
		"var":       "// === variables ===",
	}

	for _, kind := range order {
		ss := groups[kind]
		if len(ss) == 0 {
			continue
		}
		b.WriteString(headings[kind])
		b.WriteString("\n\n")
		for _, s := range ss {
			// printNode (used in --bodies mode for funcs/methods) emits
			// the doc comment itself; don't double-print it.
			docEmittedBySource := args.Bodies && (s.Kind == "func" || s.Kind == "method")
			if args.IncludeDocs && s.Doc != "" && !docEmittedBySource {
				for _, line := range strings.Split(strings.TrimRight(s.Doc, "\n"), "\n") {
					b.WriteString("// ")
					b.WriteString(line)
					b.WriteString("\n")
				}
			}
			b.WriteString(s.Source)
			b.WriteString("\n\n")
		}
	}

	return b.String()
}

// resolvePath, resolvePathFromTail, estimateTokens, etc. live in helpers.go.
