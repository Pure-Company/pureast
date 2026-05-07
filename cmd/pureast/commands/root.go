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
		// --module, in which case there's no positional path to
		// give but they clearly want us to operate on something.
		// Forward to dump in that case; resolvePathFlag will turn
		// --module into a real directory.
		RunE: func(cmd *cobra.Command, args []string) error {
			mod, _ := cmd.Flags().GetString("module")
			if len(args) == 0 && mod == "" {
				return cmd.Help()
			}
			// Forward to dump's RunE. dumpCmd's flags are mirrored
			// onto the root below, so they share storage:
			// parseDumpArgs reads the same flag values cobra wrote
			// during root parsing.
			//
			// `--module` is trickier: it's a persistent root flag,
			// but cobra only merges persistent flags into a child's
			// flag set during its own dispatch — and we're calling
			// RunE directly here, bypassing dispatch. Rather than
			// teaching dumpCmd about --module separately, we resolve
			// it once at the root and synthesize a positional path,
			// which is how every verb already consumes location info.
			if mod != "" && len(args) == 0 {
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
			} else if mod != "" && len(args) > 0 {
				return fmt.Errorf("--module and a positional PATH are mutually exclusive")
			}
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

	return rootCmd
}
