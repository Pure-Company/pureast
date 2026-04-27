// cmd/pureast/commands/search.go
package commands

import (
	"context"
	"fmt"
	"go/token"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinodhalaharvi/pureast/pkg/cli"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
)

type SearchArgs struct {
	FilePath   string
	Pattern    string
	Kind       string
	MaxResults int
}

func NewSearchCommand() *cobra.Command {
	cmd := cli.NewCommand[SearchArgs]("search").
		Short("Search for symbols using fuzzy matching").
		Long(`Search for symbols in a Go package using fuzzy matching.

Examples:
  pureast search "Handler" ./pkg
  pureast search "User" ./pkg --kind struct
  pureast search "Process" ./pkg -n 5`).
		ParseArgs(parseSearchArgs).
		Action(searchAction).
		Build()

	cmd.Flags().String("kind", "", "Filter by kind (struct, interface, function)")
	cmd.Flags().IntP("max-results", "n", 20, "Maximum results")

	// Back-compat: --file deprecated alias for positional PATH
	cmd.Flags().StringP("file", "f", "", "[deprecated] use positional PATH")

	return cmd
}

func parseSearchArgs(cmd *cobra.Command, args []string) result.Result[SearchArgs] {
	if len(args) < 1 {
		return result.Err[SearchArgs](fmt.Errorf("requires PATTERN [PATH]"))
	}
	if len(args) > 2 {
		return result.Err[SearchArgs](fmt.Errorf("expected PATTERN [PATH], got %d args", len(args)))
	}

	path, err := resolvePathFromTail(cmd, args[1:])
	if err != nil {
		return result.Err[SearchArgs](err)
	}

	kind, _ := cmd.Flags().GetString("kind")
	maxResults, _ := cmd.Flags().GetInt("max-results")

	return result.Ok(SearchArgs{
		FilePath:   path,
		Pattern:    args[0],
		Kind:       kind,
		MaxResults: maxResults,
	})
}

// searchAction discovers all symbols, filters with FuzzySearch, and
// renders the ranked list. There is no separate "index build" step:
// pureast invocations are one-shot, so building an index just to
// throw it away after one query was overhead with no payoff.
func searchAction(ctx context.Context, args SearchArgs) result.Result[cli.Output] {
	fset := token.NewFileSet()
	pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
	if err != nil {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("Error: %v\n", err),
			ExitCode: 1,
		})
	}

	symbols := extract.DiscoverAllSymbols(pkgNode)
	matches := extract.FuzzySearch(symbols, args.Pattern, args.Kind, args.MaxResults)

	var b strings.Builder
	fmt.Fprintf(&b, "Found %d matches:\n\n", len(matches))
	for i, m := range matches {
		fmt.Fprintf(&b, "%d. %s (%s) [score: %d]\n",
			i+1, m.Symbol.Name, m.Symbol.Kind, m.Score)
	}

	return result.Ok(cli.Output{Text: b.String(), ExitCode: 0})
}
