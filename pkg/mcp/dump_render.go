// pkg/mcp/dump_render.go
//
// Rendering for the dump_package MCP tool.
//
// Builds on extract.DiscoverAllSymbols() — the same enumeration the
// CLI's `list` command uses — and produces a compact, signatures-mostly
// view suitable for LLM context. This is intentionally a thinner layer
// than cmd/pureast/commands/dump.go: that file does its own AST walk to
// collect richer per-symbol metadata (file/line, signature reconstruction).
// Consolidating both into a single package is a known follow-up; today
// the MCP renderer reuses what's already in pkg/extract and accepts a
// slightly less detailed output in exchange.
//
// What this means for callers: dump_package output via MCP gives you
// names, kinds, and receivers. It does NOT yet emit per-symbol source
// signatures — that requires the per-decl walking the CLI does. For
// raw signatures Claude should call extract_symbol on individual names.
package mcp

import (
	"fmt"
	"sort"
	"strings"

	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
)

// dumpRenderOptions controls how a package is summarized for LLM context.
//
// The defaults (all false / empty) produce the broadest dump: every kind,
// both exported and unexported symbols, no token budget. That's the right
// default for a fresh "tell me about this package" query — the LLM gets
// full visibility and decides what's relevant.
type dumpRenderOptions struct {
	Kind         string // "" or "all" → no filter; otherwise extract.FilterByKind value
	ExportedOnly bool
	Format       string // "go" (default) | "md"
	MaxTokens    int    // 0 = unbounded
}

// renderDumpForMCP produces the package summary text returned to Claude.
//
// Output format mirrors `pureast dump` for visual consistency: sectioned
// by kind, sorted alphabetically within each section. Determinism matters
// here for prompt caching — the same package must render identically
// across calls so the cached prefix hits.
func renderDumpForMCP(pkg astpkg.PackageNode, opts dumpRenderOptions) string {
	symbols := extract.DiscoverAllSymbols(pkg)

	// Filter by kind if requested. extract.FilterByKind handles unknown
	// kinds by returning an empty slice, which matches MCP's "no results"
	// shape — Claude sees an empty section list and infers the filter
	// matched nothing.
	if opts.Kind != "" && opts.Kind != "all" {
		symbols = extract.FilterByKind(opts.Kind, symbols)
	}

	if opts.ExportedOnly {
		symbols = filterExported(symbols)
	}

	body := renderSymbolSections(pkg.Name, symbols)

	if opts.MaxTokens > 0 {
		// Symbol-aware truncation keeps the dump syntactically complete
		// (no half-emitted struct or function). The bool return is
		// ignored here because the MCP layer doesn't have a stderr
		// channel to surface the notice — that's a per-protocol UX
		// decision; the CLI surfaces it, MCP doesn't.
		body, _ = extract.TruncateSymbols(body, opts.MaxTokens)
	}

	if opts.Format == "md" {
		body = wrapMarkdown(fmt.Sprintf("Package %s", pkg.Name), body)
	}

	return body
}

// filterExported keeps only symbols whose name begins with an uppercase
// letter — Go's standard export rule. We don't use unicode.IsUpper because
// extract.SymbolInfo.Name is already a Go identifier, so ASCII suffices
// for the overwhelming majority of cases and matches what `go doc` does.
func filterExported(symbols []extract.SymbolInfo) []extract.SymbolInfo {
	out := make([]extract.SymbolInfo, 0, len(symbols))
	for _, s := range symbols {
		if len(s.Name) > 0 && s.Name[0] >= 'A' && s.Name[0] <= 'Z' {
			out = append(out, s)
		}
	}
	return out
}

// renderSymbolSections groups symbols by kind and renders each section
// with a header. The kind order here is deliberate: types first (the
// vocabulary), then funcs/methods (the verbs), then values (the constants
// the LLM is least likely to need). An LLM reading top-to-bottom builds
// the right mental model.
func renderSymbolSections(pkgName string, symbols []extract.SymbolInfo) string {
	groups := map[string][]extract.SymbolInfo{}
	for _, s := range symbols {
		groups[s.Kind] = append(groups[s.Kind], s)
	}

	for k := range groups {
		sort.Slice(groups[k], func(i, j int) bool {
			return groups[k][i].Name < groups[k][j].Name
		})
	}

	var b strings.Builder
	fmt.Fprintf(&b, "// pureast dump: package %s\n", pkgName)
	fmt.Fprintf(&b, "// %d symbols\n\n", len(symbols))

	order := []string{"struct", "interface", "type", "function", "method", "const", "var"}
	headings := map[string]string{
		"struct":    "// === structs ===",
		"interface": "// === interfaces ===",
		"type":      "// === type aliases ===",
		"function":  "// === functions ===",
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
		b.WriteString("\n")
		for _, s := range ss {
			if s.Receiver != "" {
				fmt.Fprintf(&b, "  %s.%s (%s)\n", s.Receiver, s.Name, s.Kind)
			} else {
				fmt.Fprintf(&b, "  %s (%s)\n", s.Name, s.Kind)
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}

// truncation logic now lives in pkg/extract/budget.go so CLI and MCP
// share a single implementation.

// wrapMarkdown wraps text in a fenced ```go block with a heading. The
// markdown form parses cleanly in Claude's chat UI and gives the LLM
// explicit document structure (heading + code) instead of a flat blob.
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
