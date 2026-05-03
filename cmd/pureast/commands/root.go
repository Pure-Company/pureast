// cmd/pureast/commands/root.go
package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
		// behavior of bare `pureast`). With a path, delegate to the
		// dump command — same flag set, same output.
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			// Forward the positional path (and any flags the user
			// passed at the root level — e.g. --kind, --exported)
			// to dump's RunE. dumpCmd's flags are mirrored onto the
			// root below, so they share storage: parseDumpArgs reads
			// the same flag values cobra wrote during root parsing.
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

	return rootCmd
}
