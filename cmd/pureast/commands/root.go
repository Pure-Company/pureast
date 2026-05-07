// cmd/pureast/commands/root.go
package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/Pure-Company/pureast/pkg/extract"
)

func NewRootCommand() *cobra.Command {
	// Build the dump subcommand once, then reuse it as the root's
	// default action. `pureast .` and `pureast ./pkg` should behave
	// identically to `pureast dump .` — most users want the kitchen-
	// sink view, and forcing them to type `dump` for the common case
	// is friction. The named verb still works for discoverability and
	// for scripts; it's just no longer required.
	dumpCmd := NewDumpCommand()

	rootCmd := &cobra.Command{
		Use:   "pureast [PATH]",
		Short: "PureAST — compact Go AST extraction for LLM context",
		Long: `PureAST extracts Go symbols, analyzes dependencies, and emits
compact representations suitable for feeding to an LLM as context.

With no subcommand, pureast dumps every symbol in the given package
(equivalent to ` + "`pureast dump PATH`" + `):

  pureast .                         # dump current directory
  pureast ./pkg                     # dump a specific package

You can also point pureast at any public Go module via --module
(uses ` + "`go mod download`" + ` under the hood, so GOPROXY/auth/cache
all work as configured):

  pureast --module github.com/spf13/cobra
  pureast --module github.com/gin-gonic/gin@v1.10.0
  pureast --module github.com/spf13/cobra/doc       # sub-package

Or dump every direct dependency in a project's go.mod at once
(indirect deps are skipped automatically):

  pureast --gomod ./go.mod                          # dump all direct deps
  pureast --gomod ./go.mod --kind interface         # contracts only
  pureast --gomod ./go.mod --skip-module github.com/aws/aws-sdk-go-v2
  pureast --gomod ./go.mod --only-module github.com/redis/go-redis/v9

Common workflows:
  pureast dump ./pkg                # every symbol, signatures only
  pureast extract User ./pkg        # one symbol with transitive deps
  pureast deps User ./pkg           # what does User depend on
  pureast diff main ./pkg           # symbols in files changed since main
  pureast search "Handler" ./pkg    # fuzzy symbol search
  pureast list ./pkg                # enumerate all symbols
  pureast types ./pkg               # type declarations only
`,
		// Accept an optional positional PATH so cobra routes
		// `pureast ./pkg` here instead of complaining about an
		// unknown command.
		Args: cobra.MaximumNArgs(1),
		// With no args, fall back to help (preserves the existing
		// behavior of bare `pureast`) — UNLESS the user passed
		// --module or --gomod, in which case there's no positional
		// path to give but they clearly want us to operate on
		// something. Forward to the appropriate handler.
		RunE: func(cmd *cobra.Command, args []string) error {
			mod, _ := cmd.Flags().GetString("module")
			gomod, _ := cmd.Flags().GetString("gomod")

			// Mutual exclusion: at most one input source.
			sources := 0
			if len(args) > 0 {
				sources++
			}
			if mod != "" {
				sources++
			}
			if gomod != "" {
				sources++
			}
			if sources > 1 {
				return fmt.Errorf("--module, --gomod, and a positional PATH are mutually exclusive")
			}

			if len(args) == 0 && mod == "" && gomod == "" {
				return cmd.Help()
			}

			// --gomod path: parse the file, iterate over direct deps,
			// dump each, concatenate. dumpAction is reused unchanged
			// per module — no special-casing inside dump itself.
			if gomod != "" {
				// Parse dump args using a synthetic positional so
				// parseDumpArgs doesn't trip on the missing PATH.
				// FilePath gets overwritten per-module inside
				// gomodAction, so any value works here.
				dumpArgs, err := parseDumpArgs(cmd, []string{"."})
				if err != nil {
					return err
				}
				only, _ := cmd.Flags().GetStringSlice("only-module")
				skip, _ := cmd.Flags().GetStringSlice("skip-module")

				out, err := gomodAction(cmd.Context(), gomod, dumpArgs, only, skip)
				if err != nil {
					return err
				}
				// Honor -o/--output if the user set one. Otherwise
				// stdout, like every other verb.
				if dumpArgs.OutputFile != "" {
					if err := os.WriteFile(dumpArgs.OutputFile, []byte(out.Text), 0644); err != nil {
						return fmt.Errorf("write %s: %w", dumpArgs.OutputFile, err)
					}
					fmt.Fprintf(os.Stderr, "wrote combined dump to %s\n", dumpArgs.OutputFile)
				} else {
					fmt.Print(out.Text)
				}
				return nil
			}

			// --module path: resolve once, synthesize positional.
			// `--module` is trickier than args because it's a
			// persistent root flag, but cobra only merges persistent
			// flags into a child's flag set during its own dispatch
			// — and we're calling RunE directly here, bypassing
			// dispatch. Rather than teaching dumpCmd about --module
			// separately, we resolve it once at the root and
			// synthesize a positional path, which is how every verb
			// already consumes location info.
			if mod != "" {
				res, err := extract.ResolveModule(mod)
				if err != nil {
					return fmt.Errorf("--module %s: %w", mod, err)
				}
				if res.SubPath != "" {
					fmt.Fprintf(os.Stderr,
						"resolved %s -> %s@%s (sub-package %s) at %s\n",
						mod, res.ModulePath, res.Version, res.SubPath, res.Dir)
				} else {
					fmt.Fprintf(os.Stderr,
						"resolved %s -> %s@%s at %s\n",
						mod, res.ModulePath, res.Version, res.Dir)
				}
				args = []string{res.Dir}
			}

			// Forward to dump's RunE for the regular single-path case.
			dumpCmd.SetContext(cmd.Context())
			return dumpCmd.RunE(dumpCmd, args)
		},
		// Cobra's default is to print "Error: <err>" + the full usage
		// block whenever a RunE returns non-nil. For a tool with
		// expected error paths (file not found, symbol missing, etc.)
		// the usage dump is noise — and we already print the error
		// itself from main.go. Silence both; main.go owns the error
		// message, RunE owns the exit code.
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	rootCmd.AddCommand(
		dumpCmd,
		NewExtractCommand(),
		NewSearchCommand(),
		NewTypesCommand(),
		NewDepsCommand(),
		NewDiffCommand(),
		NewListCommand(),
		NewClaudeEditCommand(),
	)

	// Mirror dump's flags onto the root so `pureast . --kind func`
	// and `pureast . --exported` work without needing the `dump`
	// keyword. We can't share the flag set wholesale (cobra binds
	// flags to a command), so we walk dump's local flags and
	// register equivalents here. The default action above forwards
	// to dumpCmd, which reads its own flags — but cobra also makes
	// inherited/local flags on the parent visible to RunE through
	// cmd.Flags(), so dump.parseDumpArgs sees them either way once
	// we bind them.
	dumpCmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Skip flags cobra adds itself (help) and any already on root.
		if rootCmd.Flags().Lookup(f.Name) != nil {
			return
		}
		rootCmd.Flags().AddFlag(f)
	})

	// --module is a *persistent* flag on root, which means every
	// subcommand inherits it automatically. The shared resolvePath
	// / resolvePathFlag helpers in helpers.go check for it before
	// falling back to positional/--file/CWD logic, so adding it here
	// is enough to give every verb (dump, extract, deps, diff, list,
	// search, types) module-resolution support for free.
	rootCmd.PersistentFlags().String("module", "",
		"Resolve a Go module via `go mod download` (e.g. github.com/foo/bar@v1.2.3). "+
			"Mutually exclusive with positional PATH.")

	// --gomod points at a go.mod file and dumps every direct dependency
	// in one shot. Indirect deps (marked `// indirect`) are skipped.
	// Composes with --kind, --exported, --max-tokens, --format, etc.,
	// which are applied uniformly to each module's dump. Per the design
	// spec, no auto-default to "./go.mod" — the user always says
	// `--gomod ./go.mod` explicitly. This avoids surprising behavior
	// when invoked from inside a Go project where the user really
	// meant to dump CWD.
	rootCmd.PersistentFlags().String("gomod", "",
		"Path to a go.mod file. Dumps every direct dependency. "+
			"Indirect deps are skipped. Mutually exclusive with --module and PATH.")

	// --only-module / --skip-module compose: --only narrows first,
	// --skip excludes from the narrowed set. Both match against the
	// require-line path (e.g. github.com/foo/bar) including replaced
	// targets. Repeatable: pass multiple times or comma-separated.
	rootCmd.PersistentFlags().StringSlice("only-module", nil,
		"With --gomod: include ONLY these modules (repeatable; whitelist).")
	rootCmd.PersistentFlags().StringSlice("skip-module", nil,
		"With --gomod: exclude these modules (repeatable; blacklist).")

	return rootCmd
}
