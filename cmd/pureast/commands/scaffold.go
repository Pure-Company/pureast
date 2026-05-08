// cmd/pureast/commands/scaffold.go
//
// `pureast scaffold` — materialize a project skeleton from a YAML manifest.
//
// The verb takes a single input (--manifest, defaulting to ./pureast.yaml)
// and produces a tree of `gen.go` files, one per package described in
// the manifest. Each gen.go is build-ignored and contains the //go:generate
// directives that, when invoked via `go generate -tags ignore ./...`,
// produce the package's actual source files via claude-edit.
//
// Two clean phases:
//
//	pureast scaffold                              # phase 1: lay down structure
//	go generate -tags ignore ./...                # phase 2: fill in content
//
// The scaffold step is deterministic and LLM-free. The content step is
// LLM-driven via claude-edit. Each phase has its own cache discipline,
// so re-running either is cheap when nothing has changed.
package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/Pure-Company/pureast/pkg/cli"
	"github.com/Pure-Company/pureast/pkg/scaffold"
)

type ScaffoldArgs struct {
	ManifestPath string
	ProjectRoot  string
}

func NewScaffoldCommand() *cobra.Command {
	cmd := cli.NewCommand[ScaffoldArgs]("scaffold").
		Short("Materialize a project skeleton from a YAML manifest").
		Long(`scaffold reads pureast.yaml (or another manifest file passed via
--manifest) and creates one gen.go file per package described in it.
Each gen.go is build-ignored and carries //go:generate directives that
invoke claude-edit to produce the package's actual source files.

Typical workflow:

  1. Author pureast.yaml describing your project's package layout
  2. Run 'pureast scaffold'              -> creates gen.go tree
  3. Run 'go generate -tags ignore ./...' -> fills in source files

The scaffold step is fully deterministic — no LLM involvement. The
manifest is the source of truth; running scaffold twice with the same
manifest is a no-op.

Project root is inferred from the manifest's location (the directory
containing pureast.yaml). Override with --root if needed.`).
		ParseArgs(parseScaffoldArgs).
		Action(scaffoldAction).
		Build()

	cmd.Flags().StringP("manifest", "m", "pureast.yaml",
		"Path to the manifest YAML file.")
	cmd.Flags().String("root", "",
		"Project root (defaults to the manifest's directory).")

	return cmd
}

func parseScaffoldArgs(cmd *cobra.Command, args []string) (ScaffoldArgs, error) {
	if len(args) > 0 {
		return ScaffoldArgs{}, fmt.Errorf("scaffold takes no positional arguments")
	}
	manifest, _ := cmd.Flags().GetString("manifest")
	root, _ := cmd.Flags().GetString("root")

	if manifest == "" {
		return ScaffoldArgs{}, fmt.Errorf("--manifest is required (default ./pureast.yaml)")
	}
	abs, err := filepath.Abs(manifest)
	if err != nil {
		return ScaffoldArgs{}, fmt.Errorf("resolve manifest path: %w", err)
	}
	if root == "" {
		// Infer project root from the manifest's directory. This is
		// the natural default — the manifest sits at the project's
		// top level, so its directory is what package paths are
		// relative to.
		root = filepath.Dir(abs)
	} else {
		rabs, err := filepath.Abs(root)
		if err != nil {
			return ScaffoldArgs{}, fmt.Errorf("resolve root: %w", err)
		}
		root = rabs
	}
	return ScaffoldArgs{
		ManifestPath: abs,
		ProjectRoot:  root,
	}, nil
}

func scaffoldAction(ctx context.Context, args ScaffoldArgs) (cli.Output, error) {
	m, err := scaffold.LoadManifest(args.ManifestPath)
	if err != nil {
		return cli.Output{}, err
	}

	res, err := scaffold.Scaffold(m, args.ProjectRoot)
	if err != nil {
		return cli.Output{}, fmt.Errorf("scaffold: %w", err)
	}

	// Human-readable progress goes to stderr — stdout is reserved for
	// data output (per the project's existing convention so pipelines
	// stay clean).
	for _, p := range res.Created {
		fmt.Fprintf(os.Stderr, "  created  %s\n", p)
	}
	for _, p := range res.Updated {
		fmt.Fprintf(os.Stderr, "  updated  %s\n", p)
	}
	for _, p := range res.Skipped {
		fmt.Fprintf(os.Stderr, "  skipped  %s (already up to date)\n", p)
	}

	fmt.Fprintf(os.Stderr,
		"\npureast scaffold: %d created, %d updated, %d skipped\n",
		len(res.Created), len(res.Updated), len(res.Skipped))

	if len(res.Created)+len(res.Updated) > 0 {
		fmt.Fprintln(os.Stderr,
			"\nNext: run `go generate -tags ignore ./...` to fill in source files.")
	}

	return cli.Output{}, nil
}
