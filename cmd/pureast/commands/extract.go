// cmd/pureast/commands/extract.go - Cleaned up
package commands

import (
	"context"
	"fmt"
	"go/token"
	"os"

	"github.com/spf13/cobra"
	"github.com/vinodhalaharvi/pureast/pkg/analyze"
	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/pureast/pkg/cli"
	"github.com/vinodhalaharvi/pureast/pkg/codegen"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
)

type ExtractArgs struct {
	FilePath   string
	Symbol     string
	OutputFile string
	Minimal    bool
	Workers    int
}

func NewExtractCommand() *cobra.Command {
	cmd := cli.NewCommand[ExtractArgs]("extract").
		Short("Extract a symbol with all dependencies").
		Long(`Extract a Go symbol with all its dependencies and associated code.

By default, includes:
  - Type definition
  - All dependencies (transitive)
  - Constructors (NewX functions)
  - Methods

Examples:
  pureast extract User --file ./pkg
  pureast extract Profile --file ./pkg --minimal
  pureast extract UserService --file ./pkg --output service.go`).
		ParseArgs(parseExtractArgs).
		Action(extractAction).
		Build()

	cmd.Flags().StringP("file", "f", "", "Go file or directory (required)")
	cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
	cmd.Flags().Bool("minimal", false, "Extract minimal dependencies only")
	cmd.Flags().IntP("workers", "w", 0, "Number of workers (0 = auto)")

	cmd.MarkFlagRequired("file")

	return cmd
}

func parseExtractArgs(cmd *cobra.Command, args []string) result.Result[ExtractArgs] {
	if len(args) != 1 {
		return result.Err[ExtractArgs](fmt.Errorf("requires symbol name"))
	}

	file, _ := cmd.Flags().GetString("file")
	output, _ := cmd.Flags().GetString("output")
	minimal, _ := cmd.Flags().GetBool("minimal")
	workers, _ := cmd.Flags().GetInt("workers")

	return result.Ok(ExtractArgs{
		FilePath:   file,
		Symbol:     args[0],
		OutputFile: output,
		Minimal:    minimal,
		Workers:    workers,
	})
}

func extractAction(ctx context.Context, args ExtractArgs) result.Result[cli.Output] {
	fset := token.NewFileSet()
	pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, args.Workers)
	if err != nil {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("Error: %v\n", err),
			ExitCode: 1,
		})
	}

	declMap := extract.BuildPackageDeclMap(pkgNode)
	graph := analyze.NewDependencyGraph(declMap)

	// Choose query strategy
	var deps astpkg.Dependencies
	if args.Minimal {
		deps = graph.MinimalDependencies(args.Symbol)
	} else {
		deps = graph.ResolveWithAssociatedCode(args.Symbol)
	}

	gen := codegen.NewGenerator(fset)
	code, err := gen.GenerateMinimal(pkgNode.Name, args.Symbol, declMap, deps)
	if err != nil {
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("Error: %v\n", err),
			ExitCode: 1,
		})
	}

	if args.OutputFile != "" {
		if err := os.WriteFile(args.OutputFile, []byte(code), 0644); err != nil {
			return result.Ok(cli.Output{
				Text:     fmt.Sprintf("Error writing file: %v\n", err),
				ExitCode: 1,
			})
		}
		return result.Ok(cli.Output{
			Text:     fmt.Sprintf("✅ Written to %s\n", args.OutputFile),
			ExitCode: 0,
		})
	}

	return result.Ok(cli.Output{Text: code, ExitCode: 0})
}
