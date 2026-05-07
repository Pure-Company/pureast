// cmd/pureast/commands/helpers.go
//
// Shared utilities used by every verb. Centralized here so the CLI
// surface stays consistent: one path-resolution rule, one token estimator,
// one markdown wrapper.
package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Pure-Company/pureast/pkg/extract"
)

// resolvePath returns the target path from either:
//
//  1. a positional argument (canonical: `pureast verb ./pkg`),
//  2. the --module flag (resolves a Go module via `go mod download`),
//  3. the deprecated --file flag (warns to stderr but keeps working),
//  4. the current directory as a default.
//
// --module and a positional path are mutually exclusive — picking
// one source of truth keeps the contract simple. We error rather
// than guess.
//
// Most verbs call this with args == os.Args' positional remainder.
// When a verb takes other positional args before the path, use
// resolvePathFromTail instead (see deps.go).
func resolvePath(cmd *cobra.Command, args []string) (string, error) {
	if len(args) > 1 {
		return "", fmt.Errorf("expected at most one PATH argument, got %d", len(args))
	}
	if len(args) == 1 {
		if mod, _ := cmd.Flags().GetString("module"); mod != "" {
			return "", fmt.Errorf("--module and a positional PATH are mutually exclusive")
		}
		return args[0], nil
	}
	return resolvePathFlag(cmd)
}

// resolvePathFromTail handles `verb POSITIONAL... [PATH]` shapes where
// the path is the optional last argument. tail is the slice after the
// required positionals have been consumed.
func resolvePathFromTail(cmd *cobra.Command, tail []string) (string, error) {
	if len(tail) > 1 {
		return "", fmt.Errorf("expected at most one PATH argument, got %d", len(tail))
	}
	if len(tail) == 1 {
		if mod, _ := cmd.Flags().GetString("module"); mod != "" {
			return "", fmt.Errorf("--module and a positional PATH are mutually exclusive")
		}
		return tail[0], nil
	}
	return resolvePathFlag(cmd)
}

// resolvePathFlag covers the no-positional case: try --module, then
// --file, else ".". --module wins over --file so users migrating from
// scripts that set --file can opt into remote-module mode by adding
// --module without first removing --file.
func resolvePathFlag(cmd *cobra.Command) (string, error) {
	if mod, _ := cmd.Flags().GetString("module"); mod != "" {
		res, err := extract.ResolveModule(mod)
		if err != nil {
			return "", fmt.Errorf("--module %s: %w", mod, err)
		}
		// Diagnostic to stderr so it doesn't pollute the AST output
		// on stdout. Users piping pureast into claude or jq want a
		// clean stdout; the resolution detail belongs in stderr.
		if res.SubPath != "" {
			fmt.Fprintf(os.Stderr,
				"resolved %s -> %s@%s (sub-package %s) at %s\n",
				mod, res.ModulePath, res.Version, res.SubPath, res.Dir)
		} else {
			fmt.Fprintf(os.Stderr,
				"resolved %s -> %s@%s at %s\n",
				mod, res.ModulePath, res.Version, res.Dir)
		}
		return res.Dir, nil
	}
	if f, _ := cmd.Flags().GetString("file"); f != "" {
		fmt.Fprintln(os.Stderr, "warning: --file is deprecated, pass PATH as a positional argument")
		return f, nil
	}
	return ".", nil
}

// Token-budget helpers live in pkg/extract/budget.go.
// AST rendering helpers (printNode, RenderSignature, RenderWithBody,
// SymbolDoc) live in pkg/extract/render.go. Both packages are imported
// by callers inside cmd/pureast/commands/ — see dump.go, extract.go,
// types.go.

// renderAsMarkdown wraps Go output in a fenced code block with a
// header. For the LLM-context use case markdown often parses better
// than raw .go because the model picks up section structure.
func renderAsMarkdown(title, body string) string {
	var b strings.Builder
	if title != "" {
		fmt.Fprintf(&b, "# %s\n\n", title)
	}
	b.WriteString("```go\n")
	b.WriteString(strings.TrimRight(body, "\n"))
	b.WriteString("\n```\n")
	return b.String()
}
