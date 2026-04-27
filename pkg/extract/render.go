// pkg/extract/render.go
//
// Shared renderers for symbol signatures and group-by-kind layouts.
//
// The CLI's `dump` and the MCP `dump_package` tool need almost the
// same output: group every symbol by kind in a canonical order, then
// emit a per-symbol body. Before this file existed, both verbs had
// their own copy of the heading map, kind ordering, and grouping
// loop. The MCP version skipped per-symbol signatures entirely
// (just printed names) because the renderers lived in the wrong
// package — cmd/pureast/commands/dump.go can't be imported by
// pkg/mcp.
//
// Now both consumers call into here. The CLI version uses
// RenderSignature for the signatures-only path and printNode for
// --bodies. The MCP version gets to use the same RenderSignature,
// which means dump_package now emits proper Go signatures rather
// than bare "name (kind)" tuples.

package extract

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"sort"
	"strings"
)

// KindOrder is the canonical order in which symbol kinds are emitted
// in dump-style output. Types first (the vocabulary an LLM needs to
// build a mental model), then funcs/methods (the verbs), then values
// (consts/vars — least information-dense per byte). This order is
// shared between CLI and MCP so the LLM sees the same shape from
// both paths.
//
// Note that the names here are the canonical kinds returned by
// DiscoverAllSymbols ("function", not "func"). The CLI's dump verb
// uses "func" as a flag alias for user ergonomics; map between the
// two with NormalizeKind / dump's normalizeDumpKind helper.
var KindOrder = []string{
	"struct", "interface", "type", "function", "method", "const", "var",
}

// KindHeadings is the heading text for each kind, used in dump-style
// section breaks. Kept as a map so callers can look up by canonical
// kind without coupling to KindOrder's slice indices.
var KindHeadings = map[string]string{
	"struct":    "// === structs ===",
	"interface": "// === interfaces ===",
	"type":      "// === type aliases ===",
	"function":  "// === functions ===",
	"method":    "// === methods ===",
	"const":     "// === constants ===",
	"var":       "// === variables ===",
}

// GroupSymbolsForDump groups symbols by canonical kind and sorts each
// bucket alphabetically. Returns a map keyed by kind; iterate via
// KindOrder for deterministic, prompt-cache-friendly output.
//
// Symbols whose kind isn't in KindOrder land in the result map under
// their reported kind name; iterating only KindOrder will skip them
// silently. That's the right behavior for canonical dump output —
// unknown kinds are extracted but not in the canonical taxonomy.
func GroupSymbolsForDump(symbols []SymbolInfo) map[string][]SymbolInfo {
	groups := map[string][]SymbolInfo{}
	for _, s := range symbols {
		groups[s.Kind] = append(groups[s.Kind], s)
	}
	for k := range groups {
		sort.Slice(groups[k], func(i, j int) bool {
			return groups[k][i].Name < groups[k][j].Name
		})
	}
	return groups
}

// RenderSignature returns the Go source signature of a symbol — no
// body, just the declaration line. Used by both CLI dump (signatures
// mode) and MCP dump_package. For func bodies use RenderWithBody.
//
// Returns "" if the symbol's Decl is nil (synthesized symbol) or the
// decl shape is unrecognized. Callers should treat empty return as
// "skip this symbol" rather than an error — fset positions are still
// valid for skipped symbols, just not their source.
func RenderSignature(fset *token.FileSet, s SymbolInfo) string {
	if s.Decl == nil {
		return ""
	}
	switch d := s.Decl.(type) {
	case *ast.FuncDecl:
		return renderFuncDecl(fset, d, false)
	case *ast.GenDecl:
		if len(d.Specs) == 0 {
			return ""
		}
		switch spec := d.Specs[0].(type) {
		case *ast.TypeSpec:
			return renderTypeSpec(spec)
		case *ast.ValueSpec:
			kind := "var"
			if d.Tok == token.CONST {
				kind = "const"
			}
			return renderValueSpec(kind, s.Name, spec)
		}
	}
	return ""
}

// RenderWithBody renders a symbol with its body included. For funcs
// and methods this runs go/printer over the whole declaration, which
// preserves comments and formatting. For types/values it falls back
// to RenderSignature — there's no separate "body" for those.
func RenderWithBody(fset *token.FileSet, s SymbolInfo) string {
	if s.Decl == nil {
		return ""
	}
	if fd, ok := s.Decl.(*ast.FuncDecl); ok && fd.Body != nil {
		return printNode(fset, fd)
	}
	return RenderSignature(fset, s)
}

// printNode renders an ast.Node back to source via go/printer. Used
// by RenderWithBody and by callers that need the full source of an
// arbitrary node. Configured for tab indentation to match gofmt.
//
// On the rare occasion go/printer fails (essentially impossible on
// well-formed AST from go/parser, but defensive) we return a comment
// marker rather than panic.
func printNode(fset *token.FileSet, node ast.Node) string {
	var buf bytes.Buffer
	cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	if err := cfg.Fprint(&buf, fset, node); err != nil {
		return fmt.Sprintf("/* printer error: %v */", err)
	}
	return buf.String()
}

// PrintNode is the exported alias for callers that want raw printer
// output for an arbitrary node — e.g. the CLI dump verb in --bodies
// mode wants this for funcs/methods specifically.
func PrintNode(fset *token.FileSet, node ast.Node) string {
	return printNode(fset, node)
}

// SymbolDoc returns the documentation comment text for a symbol's
// declaration, if any. Handles both FuncDecl (doc directly on the
// decl) and GenDecl (doc may be on the GenDecl or on the first
// spec, depending on whether it's a single-spec or grouped block).
func SymbolDoc(decl ast.Decl) string {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		if d.Doc != nil {
			return d.Doc.Text()
		}
	case *ast.GenDecl:
		if d.Doc != nil {
			return d.Doc.Text()
		}
		if len(d.Specs) > 0 {
			switch s := d.Specs[0].(type) {
			case *ast.TypeSpec:
				if s.Doc != nil {
					return s.Doc.Text()
				}
			case *ast.ValueSpec:
				if s.Doc != nil {
					return s.Doc.Text()
				}
			}
		}
	}
	return ""
}

// --- internal renderers (lowercase; callers go through RenderSignature) ---

func renderFuncDecl(fset *token.FileSet, d *ast.FuncDecl, includeBody bool) string {
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
			lines = append(lines, "  "+renderTypeExpr(m.Type))
			continue
		}
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
