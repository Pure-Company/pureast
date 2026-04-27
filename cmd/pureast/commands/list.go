// cmd/pureast/commands/list.go
package commands

import (
	"context"
	"fmt"
	"go/token"

	"github.com/spf13/cobra"
	"github.com/Pure-Company/pureast/pkg/cli"
	"github.com/Pure-Company/pureast/pkg/extract"
)

type ListArgs struct {
	FilePath    string
	GroupByKind bool
}

func NewListCommand() *cobra.Command {
	cmd := cli.NewCommand[ListArgs]("list").
		Short("List all symbols in a package").
		Long(`List all symbols (structs, interfaces, functions) in a package.

Examples:
  pureast list ./pkg
  pureast list ./pkg --grouped=false`).
		ParseArgs(parseListArgs).
		Action(listAction).
		Build()

	cmd.Flags().Bool("grouped", true, "Group symbols by kind")

	// Back-compat: --file deprecated alias for positional PATH
	cmd.Flags().StringP("file", "f", "", "[deprecated] use positional PATH")

	return cmd
}

func parseListArgs(cmd *cobra.Command, args []string) (ListArgs, error) {
	path, err := resolvePath(cmd, args)
	if err != nil {
		return ListArgs{}, err
	}
	grouped, _ := cmd.Flags().GetBool("grouped")

	return ListArgs{
		FilePath:    path,
		GroupByKind: grouped,
	}, nil
}

func listAction(ctx context.Context, args ListArgs) (cli.Output, error) {
	fset := token.NewFileSet()
	pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
	if err != nil {
		return cli.Output{}, fmt.Errorf("extract %s: %w", args.FilePath, err)
	}

	symbols := extract.DiscoverAllSymbols(pkgNode)
	output := fmt.Sprintf("Found %d symbols in package '%s'\n", len(symbols), pkgNode.Name)
	output += extract.FormatSymbolList(symbols, args.GroupByKind)

	return cli.Output{Text: output, ExitCode: 0}, nil
}
