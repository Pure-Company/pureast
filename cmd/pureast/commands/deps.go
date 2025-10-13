// cmd/pureast-cobra/commands/deps.go - Now much cleaner!

package commands

import (
    "context"
    "fmt"
    "go/token"

    "github.com/spf13/cobra"
    "github.com/vinodhalaharvi/pureast/pkg/analyze"
    "github.com/vinodhalaharvi/pureast/pkg/cli"
    "github.com/vinodhalaharvi/pureast/pkg/codegen"
    "github.com/vinodhalaharvi/pureast/pkg/extract"
    "github.com/vinodhalaharvi/purekernels/pkg/result"
)

type DepsArgs struct {
    FilePath string
    Symbol   string
    Report   bool
    Dot      bool
    Minimal  bool // Add minimal flag
}

func NewDepsCommand() *cobra.Command {
    cmd := cli.NewCommand[DepsArgs]("deps").
        Short("Analyze dependencies for a symbol").
        Long(`Show dependencies, generate reports, or create DOT graphs.

Examples:
  pureast deps User --file ./pkg
  pureast deps UserService --file ./pkg --report
  pureast deps Profile --file ./pkg --dot > deps.dot
  pureast deps User --file ./pkg --minimal`).
        ParseArgs(parseDepsArgs).
        Action(depsAction).
        Build()

    cmd.Flags().StringP("file", "f", "", "Go file or directory (required)")
    cmd.Flags().Bool("report", false, "Generate detailed report")
    cmd.Flags().Bool("dot", false, "Generate DOT graph")
    cmd.Flags().Bool("minimal", false, "Show minimal dependencies only")
    
    cmd.MarkFlagRequired("file")
    cmd.MarkFlagsMutuallyExclusive("report", "dot")

    return cmd
}

func parseDepsArgs(cmd *cobra.Command, args []string) result.Result[DepsArgs] {
    if len(args) != 1 {
        return result.Err[DepsArgs](fmt.Errorf("requires symbol name"))
    }

    file, _ := cmd.Flags().GetString("file")
    report, _ := cmd.Flags().GetBool("report")
    dot, _ := cmd.Flags().GetBool("dot")
    minimal, _ := cmd.Flags().GetBool("minimal")

    return result.Ok(DepsArgs{
        FilePath: file,
        Symbol:   args[0],
        Report:   report,
        Dot:      dot,
        Minimal:  minimal,
    })
}

func depsAction(ctx context.Context, args DepsArgs) result.Result[cli.Output] {
    fset := token.NewFileSet()
    pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
    if err != nil {
        return result.Ok(cli.Output{
            Text:     fmt.Sprintf("Error: %v\n", err),
            ExitCode: 1,
        })
    }

    declMap := extract.BuildPackageDeclMap(pkgNode)
    graph := analyze.NewDependencyGraph(declMap)

    // DOT graph generation
    if args.Dot {
        gen := codegen.NewGenerator(fset)
        dot := gen.GenerateDOT(args.Symbol, declMap)
        return result.Ok(cli.Output{Text: dot, ExitCode: 0})
    }

    // Determine query strategy
    queryType := analyze.WithAssociatedCode
    if args.Minimal {
        queryType = analyze.Minimal
    }

    // Full report with stats and cycles
    if args.Report {
        report := graph.AnalyzeSymbol(args.Symbol, queryType)
        output := analyze.FormatReport(report)
        output += fmt.Sprintf("\nMax Depth: %d\n", report.Stats.MaxDepth)
        return result.Ok(cli.Output{Text: output, ExitCode: 0})
    }

    // Simple dependency list (default)
    deps := graph.Query(args.Symbol, queryType)
    output := fmt.Sprintf("Dependencies for %s:\n\n", args.Symbol)
    output += analyze.FormatDependencies(args.Symbol, deps)

    return result.Ok(cli.Output{Text: output, ExitCode: 0})
}


