// pkg/mcp/dump_render.go
//
// Rendering for the dump_package and diff_since MCP tools.
//
// Builds on extract.DiscoverAllSymbols() — the same enumeration the
// CLI's `list` command uses — and produces a compact view suitable for
// LLM context. Both signature-only (default) and full-body modes are
// supported, matching CLI parity with `pureast dump --bodies`.
package mcp

import (
	"fmt"
	"go/token"
	"strings"

	astpkg "github.com/Pure-Company/pureast/pkg/ast"
	"github.com/Pure-Company/pureast/pkg/extract"
)

// dumpRenderOptions controls how a package is summarized for LLM context.
//
// The defaults (all false / empty) produce the broadest dump: every kind,
// both exported and unexported symbols, signatures only, no token budget.
type dumpRenderOptions struct {
	Kind         string // "" or "all" → no filter; otherwise extract.FilterByKind value
	ExportedOnly bool
	Bodies       bool   // include function bodies (default: signatures only)
	IncludeTests bool   // include _test.go symbols
	NoDocs       bool   // strip doc comments
	Format       string // "go" (default) | "md"
	MaxTokens    int    // 0 = unbounded
}

// renderDumpForMCP produces the package summary text returned to Claude.
//
// Output format mirrors `pureast dump` for visual consistency: sectioned
// by kind, sorted alphabetically within each section. Determinism matters
// here for prompt caching — the same package must render identically
// across calls so the cached prefix hits.
func renderDumpForMCP(fset *token.FileSet, pkg astpkg.PackageNode, opts dumpRenderOptions) string {
	symbols := extract.DiscoverAllSymbols(pkg)

	if opts.Kind != "" && opts.Kind != "all" {
		symbols = extract.FilterByKind(opts.Kind, symbols)
	}

	if opts.ExportedOnly {
		symbols = filterExported(symbols)
	}

	if !opts.IncludeTests {
		symbols = filterNonTest(fset, symbols)
	}

	body := renderSymbolSections(fset, pkg.Name, symbols, opts)

	if opts.MaxTokens > 0 {
		body, _ = extract.TruncateSymbols(body, opts.MaxTokens)
	}

	if opts.Format == "md" {
		body = wrapMarkdown(fmt.Sprintf("Package %s", pkg.Name), body)
	}

	return body
}

// filterExported keeps only symbols whose name begins with an uppercase letter.
func filterExported(symbols []extract.SymbolInfo) []extract.SymbolInfo {
	out := make([]extract.SymbolInfo, 0, len(symbols))
	for _, s := range symbols {
		if len(s.Name) > 0 && s.Name[0] >= 'A' && s.Name[0] <= 'Z' {
			out = append(out, s)
		}
	}
	return out
}

// filterNonTest removes symbols declared in _test.go files.
func filterNonTest(fset *token.FileSet, symbols []extract.SymbolInfo) []extract.SymbolInfo {
	out := make([]extract.SymbolInfo, 0, len(symbols))
	for _, s := range symbols {
		if s.Decl == nil {
			out = append(out, s)
			continue
		}
		pos := fset.Position(s.Decl.Pos())
		if !strings.HasSuffix(pos.Filename, "_test.go") {
			out = append(out, s)
		}
	}
	return out
}

// renderSymbolSections groups symbols by kind and renders each section.
func renderSymbolSections(fset *token.FileSet, pkgName string, symbols []extract.SymbolInfo, opts dumpRenderOptions) string {
	groups := extract.GroupSymbolsForDump(symbols)

	var b strings.Builder
	fmt.Fprintf(&b, "// pureast dump: package %s\n", pkgName)
	fmt.Fprintf(&b, "// %d symbols", len(symbols))
	if !opts.Bodies {
		b.WriteString(" (signatures only)")
	}
	b.WriteString("\n\n")

	for _, kind := range extract.KindOrder {
		ss := groups[kind]
		if len(ss) == 0 {
			continue
		}
		b.WriteString(extract.KindHeadings[kind])
		b.WriteString("\n\n")
		for _, s := range ss {
			// Doc comments: emit unless NoDocs, and skip when Bodies already
			// includes them (RenderWithBody emits the doc itself for funcs).
			docEmittedBySource := opts.Bodies && (s.Kind == "function" || s.Kind == "method")
			if !opts.NoDocs && !docEmittedBySource {
				if doc := extract.SymbolDoc(s.Decl); doc != "" {
					for _, line := range strings.Split(strings.TrimRight(doc, "\n"), "\n") {
						b.WriteString("// ")
						b.WriteString(line)
						b.WriteString("\n")
					}
				}
			}

			var sig string
			if opts.Bodies {
				sig = extract.RenderWithBody(fset, s)
			} else {
				sig = extract.RenderSignature(fset, s)
			}

			if sig == "" {
				if s.Receiver != "" {
					fmt.Fprintf(&b, "  %s.%s (%s)\n", s.Receiver, s.Name, s.Kind)
				} else {
					fmt.Fprintf(&b, "  %s (%s)\n", s.Name, s.Kind)
				}
			} else {
				b.WriteString(sig)
				b.WriteString("\n")
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}

// wrapMarkdown wraps text in a fenced ```go block with a heading.
func wrapMarkdown(title, body string) string {
	var b strings.Builder
	if title != "" {
		fmt.Fprintf(&b, "# %s\n\n", title)
	}
	b.WriteString("```go\n")
	b.WriteString(strings.TrimRight(body, "\n"))
	b.WriteString("\n```\n")
	return b.String()
}
