// cmd/pureast/commands/root.go
package commands

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "pureast",
		Short: "PureAST — compact Go AST extraction for LLM context",
		Long: `PureAST extracts Go symbols, analyzes dependencies, and emits
compact representations suitable for feeding to an LLM as context.

Common workflows:
  pureast dump ./pkg                # every symbol, signatures only
  pureast extract User ./pkg        # one symbol with transitive deps
  pureast deps User ./pkg           # what does User depend on
  pureast diff main ./pkg           # symbols in files changed since main
  pureast search "Handler" ./pkg    # fuzzy symbol search
  pureast list ./pkg                # enumerate all symbols
  pureast types ./pkg               # type declarations only
`,
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
		NewDumpCommand(),
		NewExtractCommand(),
		NewSearchCommand(),
		NewTypesCommand(),
		NewDepsCommand(),
		NewDiffCommand(),
		NewListCommand(),
	)

	return rootCmd
}
