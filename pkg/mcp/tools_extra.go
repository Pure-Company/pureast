// pkg/mcp/tools_extra.go
//
// MCP tool handlers added in the LLM-context expansion: dump_package,
// reverse_deps, and diff_since.
//
// These are the same operations the CLI got via the cleanup patch — but
// exposed to Claude through MCP so the LLM can call them directly during
// a conversation rather than relying on the user to paste pre-extracted
// context.
//
// Pattern note: each handler matches the existing tools.go shape exactly
// (functor.Concurrent[MCPResponse] returning a CallToolResult). The
// duplication of boilerplate between handlers is intentional — it keeps
// each handler self-contained and easy to read in isolation, which is
// the right tradeoff for a small set of tools.
package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"go/token"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/vinodhalaharvi/pureast/pkg/analyze"
	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
	"github.com/vinodhalaharvi/purekernels/pkg/functor"
)

// DumpPackageHandler emits a compact, signatures-mostly view of every
// symbol in a package — the LLM-context flagship.
//
// Why this exists separately from extract_symbol: extract_symbol is
// per-symbol and includes transitive deps. dump_package is per-package
// and gives Claude an orientation map. The right call sequence is
// usually dump_package first ("what's in here?"), then extract_symbol
// for whatever the LLM decides to focus on.
func (te *ToolExecutor) DumpPackageHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				var params struct {
					Name      string `json:"name"`
					Arguments struct {
						Path         string `json:"path"`
						Kind         string `json:"kind,omitempty"`
						ExportedOnly bool   `json:"exportedOnly,omitempty"`
						Format       string `json:"format,omitempty"`
						MaxTokens    int    `json:"maxTokens,omitempty"`
					} `json:"arguments"`
				}

				if err := json.Unmarshal(req.Params, &params); err != nil {
					return ErrorResponse(req.ID, InvalidParams, "Invalid parameters")
				}

				fset := token.NewFileSet()
				pkgResult := loadPackage(fset, params.Arguments.Path, te.workers)
				if !pkgResult.IsOk() {
					return ErrorResponse(req.ID, InternalError, pkgResult.Error().Error())
				}
				pkgNode := pkgResult.Unwrap()

				text := renderDumpForMCP(pkgNode, dumpRenderOptions{
					Kind:         params.Arguments.Kind,
					ExportedOnly: params.Arguments.ExportedOnly,
					Format:       params.Arguments.Format,
					MaxTokens:    params.Arguments.MaxTokens,
				})

				return textResponse(req.ID, text)
			},
		)
	}
}

// ReverseDepsHandler answers "who uses this symbol?" — the impact-analysis
// query. Critical when the LLM is reasoning about a refactor: it needs to
// know what calls X before it can say anything sensible about changing X.
func (te *ToolExecutor) ReverseDepsHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				var params struct {
					Name      string `json:"name"`
					Arguments struct {
						Symbol     string `json:"symbol"`
						Path       string `json:"path"`
						Transitive bool   `json:"transitive,omitempty"`
					} `json:"arguments"`
				}

				if err := json.Unmarshal(req.Params, &params); err != nil {
					return ErrorResponse(req.ID, InvalidParams, "Invalid parameters")
				}

				fset := token.NewFileSet()
				pkgResult := loadPackage(fset, params.Arguments.Path, te.workers)
				if !pkgResult.IsOk() {
					return ErrorResponse(req.ID, InternalError, pkgResult.Error().Error())
				}
				pkgNode := pkgResult.Unwrap()

				declMap := extract.BuildPackageDeclMap(pkgNode)
				graph := analyze.NewDependencyGraph(declMap)

				// Direct vs transitive callers. Direct is the more useful
				// default for LLM context: it answers "what immediately
				// breaks if I change this signature?" without flooding the
				// response with everything in the call graph.
				var users = graph.Users(params.Arguments.Symbol)
				if params.Arguments.Transitive {
					users = graph.UsersTransitive(params.Arguments.Symbol)
				}

				text := formatReverseDeps(params.Arguments.Symbol, users, params.Arguments.Transitive)
				return textResponse(req.ID, text)
			},
		)
	}
}

// DiffSinceHandler enumerates symbols in files that have changed since
// the given git ref. Wraps `git diff --name-only <ref> -- '*.go'` and
// dumps just those files.
//
// File-level granularity, not line-level: a symbol in a touched file
// is included even if that specific symbol wasn't changed. Tightening
// to line-level needs hunk parsing intersected with AST line ranges —
// reasonable follow-up but not in this round.
func (te *ToolExecutor) DiffSinceHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				var params struct {
					Name      string `json:"name"`
					Arguments struct {
						Ref       string `json:"ref"`
						Path      string `json:"path"`
						Format    string `json:"format,omitempty"`
						MaxTokens int    `json:"maxTokens,omitempty"`
					} `json:"arguments"`
				}

				if err := json.Unmarshal(req.Params, &params); err != nil {
					return ErrorResponse(req.ID, InvalidParams, "Invalid parameters")
				}
				if params.Arguments.Ref == "" {
					return ErrorResponse(req.ID, InvalidParams, "ref is required")
				}

				path := params.Arguments.Path
				if path == "" {
					path = "."
				}

				changed, err := changedGoFiles(ctx, params.Arguments.Ref, path)
				if err != nil {
					return ErrorResponse(req.ID, InternalError, err.Error())
				}
				if len(changed) == 0 {
					return textResponse(req.ID, fmt.Sprintf(
						"No Go files changed since %s.\n", params.Arguments.Ref))
				}

				// Load only the changed files. We deliberately don't
				// fall back to ExtractDirectoryConcurrent — that would
				// pull in unchanged files and defeat the verb's purpose.
				fset := token.NewFileSet()
				pkgNode, err := extract.ExtractPackageFromPaths(fset, changed)
				if err != nil {
					return ErrorResponse(req.ID, InternalError, err.Error())
				}

				text := renderDumpForMCP(pkgNode, dumpRenderOptions{
					Format:    params.Arguments.Format,
					MaxTokens: params.Arguments.MaxTokens,
				})
				header := fmt.Sprintf("// %d Go file(s) changed since %s:\n",
					len(changed), params.Arguments.Ref)
				for _, f := range changed {
					header += "//   " + f + "\n"
				}
				header += "\n"

				return textResponse(req.ID, header+text)
			},
		)
	}
}

// changedGoFiles runs `git diff --name-only <ref> HEAD` scoped to path,
// filters to Go files that still exist, and returns absolute paths.
//
// We filter for existence because deleted files appear in `git diff`
// output but obviously can't be parsed. For diff-style context this
// is the right behavior: the LLM gets what's *currently* on disk in
// the changed set, not historical content.
func changedGoFiles(ctx context.Context, ref, path string) ([]string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	cmd := exec.CommandContext(ctx, "git", "diff", "--name-only", ref, "HEAD", "--", "*.go")
	cmd.Dir = abs
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git diff: %w", err)
	}

	var files []string
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || !strings.HasSuffix(line, ".go") {
			continue
		}
		full := filepath.Join(abs, line)
		// Skip files that no longer exist (deletions, renames)
		if !fileExists(full) {
			continue
		}
		files = append(files, full)
	}
	sort.Strings(files) // deterministic order for prompt caching
	return files, nil
}

func fileExists(p string) bool {
	cmd := exec.Command("test", "-f", p)
	return cmd.Run() == nil
}

// formatReverseDeps renders Users()/UsersTransitive() output as a flat
// section list. Sorted within each section for caching determinism.
func formatReverseDeps(symbol string, d astpkg.Dependencies, transitive bool) string {
	header := fmt.Sprintf("Reverse dependencies (users) of %s", symbol)
	if transitive {
		header += " — transitive"
	}

	var b strings.Builder
	fmt.Fprintf(&b, "%s:\n\n", header)
	startLen := b.Len()

	emit := func(label string, names []string) {
		if len(names) == 0 {
			return
		}
		sort.Strings(names)
		fmt.Fprintf(&b, "%s (%d):\n", label, len(names))
		for _, n := range names {
			fmt.Fprintf(&b, "  - %s\n", n)
		}
		b.WriteString("\n")
	}

	emit("Types", d.Types.ToSlice())
	emit("Functions", d.Functions.ToSlice())
	emit("Structs", d.Structs.ToSlice())
	emit("Interfaces", d.Interfaces.ToSlice())
	emit("Constants", d.Constants.ToSlice())
	emit("Variables", d.Variables.ToSlice())

	if b.Len() == startLen {
		fmt.Fprintf(&b, "(no users found)\n")
	}

	return b.String()
}

// textResponse is a small convenience wrapper for the CallToolResult
// shape every handler returns. Centralizing it avoids the per-handler
// nested map literal that bloats the existing tools.go.
func textResponse(id interface{}, text string) MCPResponse {
	return MCPResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result: map[string]interface{}{
			"content": []map[string]interface{}{
				{"type": "text", "text": text},
			},
		},
	}
}
