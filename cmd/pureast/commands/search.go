// cmd/pureast-cobra/commands/search.go
package commands

import (
    "context"
    "fmt"
    "go/token"

    "github.com/spf13/cobra"
    "github.com/vinodhalaharvi/pureast/pkg/cli"
    "github.com/vinodhalaharvi/pureast/pkg/extract"
    "github.com/vinodhalaharvi/pureast/pkg/index"
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
        Long("Search for symbols in a Go package using fuzzy matching.").
        ParseArgs(parseSearchArgs).
        Action(searchAction).
        Build()

    cmd.Flags().StringP("file", "f", "", "Go file or directory (required)")
    cmd.Flags().String("kind", "", "Filter by kind (struct, interface, function)")
    cmd.Flags().IntP("max-results", "n", 20, "Maximum results")
    
    cmd.MarkFlagRequired("file")

    return cmd
}

func parseSearchArgs(cmd *cobra.Command, args []string) result.Result[SearchArgs] {
    if len(args) != 1 {
        return result.Err[SearchArgs](fmt.Errorf("requires search pattern"))
    }

    file, _ := cmd.Flags().GetString("file")
    kind, _ := cmd.Flags().GetString("kind")
    maxResults, _ := cmd.Flags().GetInt("max-results")

    return result.Ok(SearchArgs{
        FilePath:   file,
        Pattern:    args[0],
        Kind:       kind,
        MaxResults: maxResults,
    })
}

func searchAction(ctx context.Context, args SearchArgs) result.Result[cli.Output] {
    fset := token.NewFileSet()
    pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
    if err != nil {
        return result.Ok(cli.Output{
            Text:     fmt.Sprintf("Error: %v\n", err),
            ExitCode: 1,
        })
    }

    idx := index.BuildIndex(pkgNode)
    pattern := index.SearchPattern{
        SymbolPattern: args.Pattern,
        Kind:          args.Kind,
    }

    matches := index.FuzzySearchIndexConcurrent(pattern, idx, 10)
    
    if len(matches) > args.MaxResults {
        matches = matches[:args.MaxResults]
    }

    output := fmt.Sprintf("Found %d matches:\n\n", len(matches))
    for i, match := range matches {
        output += fmt.Sprintf("%d. %s (%s) [score: %d]\n",
            i+1, match.Entry.Name, match.Entry.Kind, match.Score.Score)
    }

    return result.Ok(cli.Output{Text: output, ExitCode: 0})
}


