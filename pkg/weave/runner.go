// pkg/weave/runner.go
//
// Parallel executor for a topologically-sorted DAG. Walks each level
// in order, fans out the level's nodes onto goroutines bounded by a
// semaphore, waits for all to finish before advancing to the next
// level. Between levels, optionally runs `go mod tidy` if the
// upcoming level contains a gomod-sourced directive.
//
// The actual per-node work is delegated to a Runner interface. The
// default implementation shells out to `pureast claude-edit` with
// the right flag combination derived from the manifest entry; tests
// inject a fake Runner to verify scheduling/parallelism without
// calling Claude.

package weave

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Pure-Company/pureast/pkg/scaffold"
)

// Runner executes a single Node — typically by invoking
// pureast claude-edit. The interface exists primarily so tests can
// substitute a deterministic fake; production uses ClaudeEditRunner
// (below).
type Runner interface {
	// Run produces the node's output file. ctx is honored for
	// cancellation; if ctx.Err() != nil after Run returns, weave
	// treats the run as a failure regardless of return value.
	//
	// projectRoot is the directory containing pureast.yaml. The
	// node's directive paths are relative to its package directory,
	// which is itself relative to projectRoot.
	//
	// override controls per-invocation behavior. The zero value is
	// the standard cache-respecting call. Reconcile passes overrides
	// with NoCache=true and AppendContext set to the build output.
	Run(ctx context.Context, projectRoot string, n Node, override RunOverride) error
}

// RunOverride bundles per-call switches that orchestrators (like the
// reconcile loop) can flip without affecting the steady-state
// behavior. The zero value means "normal cache-respecting call."
type RunOverride struct {
	// NoCache passes --no-cache to claude-edit, forcing Claude to
	// be invoked even when the cache would have hit. The header is
	// still written with the normal (task, context) cache key, so
	// future steady-state runs cache-hit naturally.
	NoCache bool

	// AppendContext is passed via --append-context. Adds extra
	// prompt material without influencing the cache key. Reconcile
	// uses this to feed the build error output back to Claude.
	AppendContext string
}

// Result is the per-node outcome surfaced by Weave.
type Result struct {
	Node     Node
	Err      error
	Duration time.Duration
}

// Summary aggregates Results across the entire run, level by level.
type Summary struct {
	// Levels[i] is the slice of Results for level i, in the same
	// order as DAG.Levels()[i]. (Order is stabilized by Levels()
	// so this is reproducible.)
	Levels [][]Result
	// SkippedDownstream is the count of nodes weave never
	// attempted because an earlier level had failures and weave
	// declined to advance.
	SkippedDownstream int
	// Errors is the flat list of errors from any level (for the
	// caller's quick-summary use).
	Errors []error
	// ReconcileRounds is the count of build-fix rounds attempted
	// (0 when --reconcile not set or normal run already green).
	ReconcileRounds int
	// ReconcileSucceeded is true when the project ended up green
	// (either before reconcile started or after a successful round).
	ReconcileSucceeded bool
}

// Options control runtime behavior. Sensible zero values throughout.
type Options struct {
	// Concurrency caps the number of nodes running in parallel
	// within a single level. <=0 means default (4).
	Concurrency int

	// Runner is the per-node action. nil means use ClaudeEditRunner
	// with default settings.
	Runner Runner

	// SkipModTidy disables the automatic `go mod tidy` invocation
	// before gomod-sourced levels. Useful for tests; in real use
	// you almost always want it enabled (default).
	SkipModTidy bool

	// LogWriter receives human-readable progress lines (one per
	// level transition, one per node start/finish, summary at end).
	// nil = os.Stderr.
	LogWriter io.Writer

	// Reconcile, if true, enables a build-error feedback loop after
	// the normal weave run completes. The loop runs `go build ./...`,
	// captures any failure output, broadcasts it to every directive
	// in topological order via Runner.Run with NoCache=true and
	// AppendContext=<build output>, then re-runs the build. Repeats
	// up to MaxRounds (or 3 if 0). Stops on success, on no-progress
	// (same build error blob across rounds), or when MaxRounds is hit.
	Reconcile bool

	// MaxRounds caps the reconcile loop. <=0 means default (3).
	// Each round re-runs every directive once and re-runs go build,
	// so this directly bounds API calls and wall-clock time.
	MaxRounds int

	// BuildCmd is the command run by the reconcile loop to detect
	// completion. Default: ["go", "build", "./..."]. Tests inject
	// a stub. The command is run from projectRoot.
	BuildCmd []string
}

// Weave runs the manifest's directives in topological order. Within
// each level, nodes execute in parallel up to Options.Concurrency.
// Between levels, weave runs `go mod tidy` if the next level's first
// directive consumes a gomod source.
//
// Failures don't cancel sibling work in the same level — weave waits
// for the level to drain before deciding what to do next. If any
// node in the level failed, weave stops before the next level (a
// fail-soft policy: don't waste API calls on downstream work that
// would consume the failed upstream's output).
func Weave(ctx context.Context, m *scaffold.Manifest, projectRoot string, opts Options) (*Summary, error) {
	g, err := BuildDAG(m)
	if err != nil {
		return nil, err
	}

	concurrency := opts.Concurrency
	if concurrency <= 0 {
		concurrency = 4
	}
	runner := opts.Runner
	if runner == nil {
		runner = &ClaudeEditRunner{}
	}
	logw := opts.LogWriter
	if logw == nil {
		logw = os.Stderr
	}

	levels := g.Levels()
	summary := &Summary{Levels: make([][]Result, 0, len(levels))}

	for i, level := range levels {
		// Run go mod tidy if this level needs a populated go.mod.
		if !opts.SkipModTidy && LevelContainsGomod(level) {
			fmt.Fprintln(logw, "weave: running `go mod tidy` (next level reads project go.mod)")
			if err := runGoModTidy(ctx, projectRoot); err != nil {
				summary.Errors = append(summary.Errors,
					fmt.Errorf("go mod tidy: %w", err))
				skipped := countRemaining(levels[i:])
				summary.SkippedDownstream = skipped
				return summary, fmt.Errorf("go mod tidy failed: %w", err)
			}
		}

		fmt.Fprintf(logw, "weave: level %d/%d (%d node(s))\n",
			i+1, len(levels), len(level))
		results := runLevel(ctx, runner, projectRoot, level, concurrency, RunOverride{}, logw)
		summary.Levels = append(summary.Levels, results)

		// Check for failures in this level. Fail-soft: complete the
		// level, then halt before the next.
		var levelErrs []error
		for _, r := range results {
			if r.Err != nil {
				levelErrs = append(levelErrs, fmt.Errorf("%s: %w", r.Node.ID(), r.Err))
			}
		}
		if len(levelErrs) > 0 {
			summary.Errors = append(summary.Errors, levelErrs...)
			summary.SkippedDownstream = countRemaining(levels[i+1:])
			fmt.Fprintf(logw,
				"weave: level %d had %d failure(s); skipping %d downstream node(s)\n",
				i+1, len(levelErrs), summary.SkippedDownstream)
			return summary, levelErrs[0]
		}
	}

	fmt.Fprintf(logw, "weave: complete (%d level(s), %d node(s))\n",
		len(levels), nodeCount(summary))

	// Reconcile loop. After the normal multi-level run completes
	// successfully, optionally run `go build ./...` and feed any
	// failures back to every directive as appended context. Repeats
	// until the build is green, no progress is being made, or the
	// max-rounds cap is reached.
	if opts.Reconcile {
		// Reuse the already-constructed DAG g.
		rounds, ok, recErr := reconcileLoop(
			ctx, g, runner, projectRoot,
			opts.MaxRounds, concurrency, opts.BuildCmd, logw,
		)
		summary.ReconcileRounds = rounds
		summary.ReconcileSucceeded = ok
		if recErr != nil {
			summary.Errors = append(summary.Errors, recErr)
			return summary, recErr
		}
	}

	return summary, nil
}

// runLevel fans out a level's nodes onto goroutines, bounded by a
// semaphore of `concurrency` size. Returns when every node has
// finished (success or failure).
//
// Results are returned in the same order as the input level so they
// stay aligned with deterministic DAG output.
//
// override is forwarded to every Runner invocation in this level.
// Normal weave runs pass the zero value; reconcile passes
// {NoCache: true, AppendContext: <build output>}.
func runLevel(
	ctx context.Context,
	runner Runner,
	projectRoot string,
	level []Node,
	concurrency int,
	override RunOverride,
	logw io.Writer,
) []Result {
	results := make([]Result, len(level))
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i, n := range level {
		wg.Add(1)
		go func(i int, n Node) {
			defer wg.Done()

			select {
			case sem <- struct{}{}:
				defer func() { <-sem }()
			case <-ctx.Done():
				results[i] = Result{Node: n, Err: ctx.Err()}
				return
			}

			fmt.Fprintf(logw, "  ▸ %s\n", n.ID())
			start := time.Now()
			err := runner.Run(ctx, projectRoot, n, override)
			dur := time.Since(start)
			results[i] = Result{Node: n, Err: err, Duration: dur}
			if err != nil {
				fmt.Fprintf(logw, "  ✗ %s (%s) — %v\n", n.ID(), dur.Round(time.Millisecond), err)
			} else {
				fmt.Fprintf(logw, "  ✓ %s (%s)\n", n.ID(), dur.Round(time.Millisecond))
			}
		}(i, n)
	}
	wg.Wait()
	return results
}

// runGoModTidy shells out to `go mod tidy` in the project root. We
// don't capture stdout/stderr — they go straight to weave's stderr,
// because Go's tidy output is itself a useful log (which deps were
// resolved, any issues).
func runGoModTidy(ctx context.Context, projectRoot string) error {
	cmd := exec.CommandContext(ctx, "go", "mod", "tidy")
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func countRemaining(levels [][]Node) int {
	n := 0
	for _, l := range levels {
		n += len(l)
	}
	return n
}

func nodeCount(s *Summary) int {
	n := 0
	for _, l := range s.Levels {
		n += len(l)
	}
	return n
}

// ClaudeEditRunner is the production Runner: it builds a
// `pureast claude-edit` invocation from the node's manifest entry
// and executes it. Stdout/stderr go straight to weave's own
// stdout/stderr so the user sees claude-edit's progress in-line.
type ClaudeEditRunner struct {
	// PureastBin is the binary to invoke. Default: "pureast" on PATH.
	PureastBin string
}

func (r *ClaudeEditRunner) Run(ctx context.Context, projectRoot string, n Node, override RunOverride) error {
	bin := r.PureastBin
	if bin == "" {
		bin = "pureast"
	}

	args := []string{"claude-edit"}

	if n.File.Model != "" {
		args = append(args, "--model", n.File.Model)
	}
	args = append(args, "--task", n.File.Task)

	for _, src := range n.File.Sources {
		switch {
		case src.Pkg != "":
			args = append(args, "--pkg", src.Pkg)
		case src.Module != "":
			args = append(args, "--module", src.Module)
		case src.Symbol != "":
			args = append(args, "--symbol", src.Symbol)
		case src.Gomod != "":
			args = append(args, "--gomod", src.Gomod)
		}
	}

	if n.File.Kind != "" && n.File.Kind != "all" {
		args = append(args, "--kind", n.File.Kind)
	}
	if n.File.ExportedOnly != nil && !*n.File.ExportedOnly {
		args = append(args, "--exported=false")
	}
	if n.File.MaxTokens > 0 {
		args = append(args, "--max-tokens", strconv.Itoa(n.File.MaxTokens))
	}
	args = append(args, "--output", n.File.Output)

	// Override flags (set by reconcile loop only). The zero-value
	// override is unused on a normal weave run.
	if override.NoCache {
		args = append(args, "--no-cache")
	}
	if override.AppendContext != "" {
		args = append(args, "--append-context", override.AppendContext)
	}

	cmd := exec.CommandContext(ctx, bin, args...)
	// claude-edit's --pkg / --output paths are interpreted relative
	// to the working directory (matching `go generate`'s convention
	// of running each directive in the file's directory). So we cd
	// into the package directory before invoking.
	cmd.Dir = filepath.Join(projectRoot, filepath.FromSlash(n.PackagePath))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s: %w", bin, strings.Join(args, " "), err)
	}
	return nil
}
