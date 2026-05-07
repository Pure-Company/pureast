// cmd/pureast/commands/gomod.go
//
// Multi-module dump driven by a go.mod file.
//
// `pureast --gomod ./go.mod` reads the project's direct dependencies
// (skipping indirect ones), resolves each via `go mod download` —
// reusing the same machinery --module already uses — and concatenates
// per-module dumps into a single output. All other flags compose:
// --kind interface, --exported, --max-tokens, --format md, etc., are
// applied uniformly to every module in the result.
//
// This is the killer use case for LLM-context prep: instead of typing
// `-p github.com/x/y -p github.com/a/b -p ...` for every dep, point
// at the project's go.mod and get the whole API surface in one go.
//
// Implementation strategy: call dump's existing action function once
// per module with a synthesized DumpArgs whose FilePath points at the
// resolved $GOMODCACHE directory. The dump action is unchanged — it
// has no idea it's being driven by a higher-level loop. This keeps
// the surgery minimal and means any future improvements to dump
// (caching, format additions, etc.) automatically benefit --gomod.
package commands

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Pure-Company/pureast/pkg/cli"
	"github.com/Pure-Company/pureast/pkg/extract"
)

// gomodAction is the orchestrator. It is invoked from root.go's RunE
// when --gomod is set, and produces the same shape of output as a
// regular dump (cli.Output) so the rest of the CLI machinery doesn't
// need to know it ran differently.
//
// gomodPath is the user-supplied path to a go.mod file (always
// explicit per the design spec; "." is not auto-resolved). dumpArgs
// is a partially-populated template — its FilePath is overwritten
// per module, but everything else (Kind, ExportedOnly, MaxTokens,
// Format, etc.) flows through unchanged.
func gomodAction(
	ctx context.Context,
	gomodPath string,
	dumpArgs DumpArgs,
	only []string,
	skip []string,
) (cli.Output, error) {
	refs, err := extract.ParseGoMod(gomodPath)
	if err != nil {
		return cli.Output{}, fmt.Errorf("--gomod %s: %w", gomodPath, err)
	}
	refs = extract.FilterModules(refs, only, skip)
	if len(refs) == 0 {
		return cli.Output{}, fmt.Errorf(
			"--gomod %s: all direct dependencies were filtered out (check --only-module / --skip-module)",
			gomodPath)
	}

	fmt.Fprintf(os.Stderr, "pureast --gomod: %d direct module(s) to dump\n", len(refs))

	// We accumulate per-module outputs into a single text. A monoid
	// fold would be more elegant (cli.OutputMonoid exists for exactly
	// this), but the per-module section header has to be inserted
	// between elements, so a plain string builder reads more honestly.
	var b strings.Builder
	var combinedExitCode int

	for i, ref := range refs {
		// Resolve the module to an on-disk directory. For local-path
		// replacements we skip the download and use the path directly;
		// for everything else we go through the same ResolveModule
		// path that --module uses, which leverages $GOMODCACHE.
		var dir, displayVersion string
		if ref.LocalPath != "" {
			dir = ref.LocalPath
			displayVersion = "(local)"
			fmt.Fprintf(os.Stderr,
				"  [%d/%d] %s -> %s (replace)\n",
				i+1, len(refs), ref.OriginalPath, ref.LocalPath)
		} else {
			res, err := extract.ResolveModule(ref.Spec())
			if err != nil {
				// One module's failure shouldn't kill the whole dump.
				// Note it on stderr and skip — the user gets partial
				// output rather than nothing, which is much more
				// useful when one private dep is briefly unreachable.
				fmt.Fprintf(os.Stderr,
					"  [%d/%d] %s: SKIP (%v)\n",
					i+1, len(refs), ref.OriginalPath, err)
				continue
			}
			dir = res.Dir
			displayVersion = res.Version
			fmt.Fprintf(os.Stderr,
				"  [%d/%d] %s@%s\n",
				i+1, len(refs), ref.OriginalPath, displayVersion)
		}

		// Synthesize per-module DumpArgs. We have to also clear
		// MaxTokens-like cumulative limits if we wanted the limit
		// to apply globally — but per-module is more useful here:
		// users typically want "no module's dump exceeds N tokens"
		// rather than "the whole bundle stops at N tokens, who
		// knows where." Document this in --gomod help.
		perModuleArgs := dumpArgs
		perModuleArgs.FilePath = dir
		// Suppress per-call output file writing; we collect into the
		// combined output below. If the user passed -o, we write the
		// combined result once at the end (handled by the caller).
		perModuleArgs.OutputFile = ""

		out, err := dumpAction(ctx, perModuleArgs)
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"  [%d/%d] %s: dump failed: %v\n",
				i+1, len(refs), ref.OriginalPath, err)
			continue
		}

		// Section header makes it unambiguous to a downstream LLM
		// (or human) which symbols came from which module. The
		// format mirrors what the claude-with-stdlib script produces
		// by hand, so users who've been doing this manually get
		// identical structure for free.
		fmt.Fprintf(&b, "// ============================================================\n")
		fmt.Fprintf(&b, "// MODULE: %s@%s\n", ref.OriginalPath, displayVersion)
		if ref.Path != ref.OriginalPath {
			fmt.Fprintf(&b, "// (replaced by: %s)\n", ref.Path)
		}
		fmt.Fprintf(&b, "// ============================================================\n\n")
		b.WriteString(out.Text)
		if !strings.HasSuffix(out.Text, "\n") {
			b.WriteString("\n")
		}
		b.WriteString("\n")

		if out.ExitCode != 0 && combinedExitCode == 0 {
			combinedExitCode = out.ExitCode
		}
	}

	return cli.Output{Text: b.String(), ExitCode: combinedExitCode}, nil
}

// gomodArgsFromFlags reads --gomod, --skip-module, --only-module from
// any cobra command (they're persistent root flags so every verb sees
// them). Returns ("", nil, nil) when --gomod is not set, signaling to
// the caller that the regular single-path flow should run.
func gomodArgsFromFlags(cmd *cobra.Command) (path string, only, skip []string) {
	path, _ = cmd.Flags().GetString("gomod")
	only, _ = cmd.Flags().GetStringSlice("only-module")
	skip, _ = cmd.Flags().GetStringSlice("skip-module")
	return
}
