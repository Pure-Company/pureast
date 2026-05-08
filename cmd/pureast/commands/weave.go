// cmd/pureast/commands/weave.go
//
// `pureast weave` — parallel orchestrator for the directives in a
// pureast manifest. Reads pureast.yaml, builds a DAG from the
// declared sources, runs the manifest's directives level-by-level
// with within-level parallelism. Equivalent to `go generate -tags
// ignore ./...` but DAG-aware and concurrent.

package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/Pure-Company/pureast/pkg/cli"
	"github.com/Pure-Company/pureast/pkg/scaffold"
	"github.com/Pure-Company/pureast/pkg/weave"
)

type WeaveArgs struct {
	ManifestPath string
	ProjectRoot  string
	Concurrency  int
	SkipModTidy  bool
}

func NewWeaveCommand() *cobra.Command {
	cmd := cli.NewCommand[WeaveArgs]("weave").
		Short("Parallel-execute manifest directives in topological order").
		Long(`weave reads pureast.yaml, builds a dependency DAG from each
File's sources, and runs the directives in topologically-sorted
order with within-level parallelism.

Source-type semantics for ordering:

  pkg / local symbol  : edge to every directive in the target package
  module / remote sym : no edge (resolved via go mod download)
  gomod               : edge to every directive whose output ends
                        in .go; weave runs 'go mod tidy' before any
                        level containing a gomod-sourced directive.

Equivalent to 'go generate -tags ignore ./...' but DAG-aware:
independent directives run in parallel, build files run last after
the project's go.mod is populated.`).
		ParseArgs(parseWeaveArgs).
		Action(weaveAction).
		Build()

	cmd.Flags().StringP("manifest", "m", "pureast.yaml",
		"Path to the manifest YAML file.")
	cmd.Flags().String("root", "",
		"Project root (defaults to the manifest's directory).")
	cmd.Flags().IntP("concurrency", "j", 4,
		"Maximum number of directives to run in parallel within one level.")
	cmd.Flags().Bool("skip-mod-tidy", false,
		"Don't run 'go mod tidy' before gomod-sourced levels.")

	return cmd
}

func parseWeaveArgs(cmd *cobra.Command, args []string) (WeaveArgs, error) {
	if len(args) > 0 {
		return WeaveArgs{}, fmt.Errorf("weave takes no positional arguments")
	}
	manifest, _ := cmd.Flags().GetString("manifest")
	root, _ := cmd.Flags().GetString("root")
	concurrency, _ := cmd.Flags().GetInt("concurrency")
	skipTidy, _ := cmd.Flags().GetBool("skip-mod-tidy")

	if manifest == "" {
		return WeaveArgs{}, fmt.Errorf("--manifest is required (default ./pureast.yaml)")
	}
	abs, err := filepath.Abs(manifest)
	if err != nil {
		return WeaveArgs{}, fmt.Errorf("resolve manifest: %w", err)
	}
	if root == "" {
		root = filepath.Dir(abs)
	} else {
		rabs, err := filepath.Abs(root)
		if err != nil {
			return WeaveArgs{}, fmt.Errorf("resolve root: %w", err)
		}
		root = rabs
	}
	return WeaveArgs{
		ManifestPath: abs,
		ProjectRoot:  root,
		Concurrency:  concurrency,
		SkipModTidy:  skipTidy,
	}, nil
}

func weaveAction(ctx context.Context, args WeaveArgs) (cli.Output, error) {
	m, err := scaffold.LoadManifest(args.ManifestPath)
	if err != nil {
		return cli.Output{}, err
	}

	summary, err := weave.Weave(ctx, m, args.ProjectRoot, weave.Options{
		Concurrency: args.Concurrency,
		SkipModTidy: args.SkipModTidy,
		LogWriter:   os.Stderr,
	})
	if err != nil {
		// Even on error, print a brief summary so the user sees how
		// far we got. Then return the error so the CLI exits non-zero.
		printSummary(summary)
		return cli.Output{}, err
	}

	printSummary(summary)
	return cli.Output{}, nil
}

func printSummary(s *weave.Summary) {
	if s == nil {
		return
	}
	totalNodes := 0
	totalDuration := time.Duration(0)
	for _, lvl := range s.Levels {
		for _, r := range lvl {
			totalNodes++
			totalDuration += r.Duration
		}
	}
	fmt.Fprintf(os.Stderr,
		"\nweave summary: %d level(s), %d node(s), summed time %s",
		len(s.Levels), totalNodes, totalDuration.Round(time.Millisecond))
	if s.SkippedDownstream > 0 {
		fmt.Fprintf(os.Stderr, ", %d skipped after upstream failure", s.SkippedDownstream)
	}
	fmt.Fprintln(os.Stderr)
}
