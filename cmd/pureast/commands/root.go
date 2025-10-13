// cmd/pureast-cobra/commands/root.go
package commands

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "pureast",
		Short: "PureAST - Pure functional Go AST analysis tool",
		Long: `PureAST extracts Go symbols, analyzes dependencies, and generates
code using pure functional programming principles.`,
	}

	// Add subcommands
	rootCmd.AddCommand(
		NewExtractCommand(),
		NewSearchCommand(),
		NewTypesCommand(),
		NewDepsCommand(),
		NewListCommand(),
		NewProtoCommand(),
	)

	return rootCmd
}
