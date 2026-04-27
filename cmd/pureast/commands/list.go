// cmd/pureast-cobra/commands/list.go
package commands

import (
	"context"
	"fmt"
	"go/token"

	"github.com/spf13/cobra"
	"github.com/vinodhalaharvi/pureast/pkg/cli"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
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

func parseListArgs(cmd *cobra.Command, args []string) result.Result[ListArgs] {
	path, err := resolvePath(cmd, args)
	if err != nil {
		return result.Err[ListArgs](err)
	}
	grouped, _ := cmd.Flags().GetBool("grouped")

	return result.Ok(ListArgs{
		FilePath:    path,
		GroupByKind: grouped,
	})
}

func listAction(ctx context.Context, args ListArgs) result.Result[cli.Output] {
	fset := token.NewFileSet()
	pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
	if err != nil {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("Error: %v\n", err),
			ExitCode: 1,
		})
	}

	symbols := extract.DiscoverAllSymbols(pkgNode)
	output := fmt.Sprintf("Found %d symbols in package '%s'\n", len(symbols), pkgNode.Name)
	output += extract.FormatSymbolList(symbols, args.GroupByKind)

	return result.Ok(cli.Output{Text: output, ExitCode: 0})
}
