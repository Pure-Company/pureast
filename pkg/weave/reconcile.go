// pkg/weave/reconcile.go
//
// Build-feedback loop. After a successful weave run, we run
// `go build ./...` and feed any compile errors back to every
// directive as appended context. Each directive's claude-edit gets
// a chance to fix the error if it thinks the error is its problem.
// We re-run the build and repeat until green or until we hit a cap.
//
// The reconcile loop is opt-in via Options.Reconcile. When off,
// Weave behaves exactly as before. When on, the post-weave hook
// kicks in only after the normal multi-level run has completed
// without error — there's no point trying to fix build errors when
// we already know some directive failed to generate.

package weave

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

// reconcileLoop runs go build, on failure broadcasts the build output
// to every directive (in topological order, level-by-level), re-runs
// go build, and repeats until green, no-progress, or MaxRounds.
//
// Returns:
//   - rounds: how many fix rounds we attempted (0 if first build was green)
//   - success: true if the project ended green
//   - err: an error describing why we stopped (nil on success)
func reconcileLoop(
	ctx context.Context,
	g *DAG,
	runner Runner,
	projectRoot string,
	maxRounds int,
	concurrency int,
	buildCmd []string,
	logw io.Writer,
) (rounds int, success bool, err error) {
	if maxRounds <= 0 {
		maxRounds = 3
	}
	if len(buildCmd) == 0 {
		buildCmd = []string{"go", "build", "./..."}
	}

	levels := g.Levels()

	// Initial build. If it's already green, we're done before doing
	// any work.
	fmt.Fprintln(logw, "weave: reconcile starting; running initial build...")
	out, buildErr := runBuild(ctx, projectRoot, buildCmd)
	if buildErr == nil {
		fmt.Fprintln(logw, "weave: build green on entry; nothing to reconcile")
		return 0, true, nil
	}
	fmt.Fprintf(logw, "weave: initial build failed; %d byte(s) of error output\n", len(out))

	prevErrHash := hashBytes(out)

	for round := 1; round <= maxRounds; round++ {
		fmt.Fprintf(logw, "weave: reconcile round %d/%d (broadcasting build errors to %d directive(s))\n",
			round, maxRounds, totalNodes(levels))

		// Re-run every level with NoCache + AppendContext set to the
		// current build output. Levels still run in parallel within
		// themselves, sequential between, just like the normal run.
		// Each directive's claude-edit decides whether the error is
		// its problem (fix it) or not (leave the file essentially
		// unchanged).
		override := RunOverride{
			NoCache:       true,
			AppendContext: formatBuildErrorContext(out, round, maxRounds),
		}

		anyDirectiveError := false
		for i, level := range levels {
			fmt.Fprintf(logw, "  reconcile level %d/%d (%d node(s))\n",
				i+1, len(levels), len(level))
			results := runLevel(ctx, runner, projectRoot, level, concurrency, override, logw)
			for _, r := range results {
				if r.Err != nil {
					anyDirectiveError = true
				}
			}
		}

		if anyDirectiveError {
			// Some directives failed during the fix pass. We don't
			// abort the whole loop — many runs will still produce a
			// valid build because OTHER directives fixed the right
			// thing. But we surface this in the log.
			fmt.Fprintln(logw, "weave: some directives errored during reconcile; continuing to build check")
		}

		// Re-run the build to see if anything's better.
		out, buildErr = runBuild(ctx, projectRoot, buildCmd)
		if buildErr == nil {
			fmt.Fprintf(logw, "weave: reconcile round %d converged; build green\n", round)
			return round, true, nil
		}

		// No-progress check: if the build error is byte-identical to
		// the previous round's, additional rounds won't help.
		curErrHash := hashBytes(out)
		if curErrHash == prevErrHash {
			fmt.Fprintf(logw,
				"weave: reconcile round %d made no progress (build output unchanged); halting\n",
				round)
			return round, false, fmt.Errorf("reconcile no-progress at round %d", round)
		}
		prevErrHash = curErrHash

		fmt.Fprintf(logw, "weave: reconcile round %d build still failing; will try another round\n", round)
	}

	return maxRounds, false, fmt.Errorf("reconcile exhausted %d round(s) without convergence", maxRounds)
}

// runBuild executes the build command from projectRoot and returns
// the combined stdout+stderr output along with the exec error (if any).
//
// We capture combined output (not separate streams) because go's
// compile errors interleave between stderr and stdout in subtle
// ways depending on which sub-process produced them. The combined
// output is what the user sees on a normal `go build`, so it's the
// correct context to feed back to Claude.
func runBuild(ctx context.Context, projectRoot string, buildCmd []string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, buildCmd[0], buildCmd[1:]...)
	cmd.Dir = projectRoot
	var combined bytes.Buffer
	cmd.Stdout = &combined
	cmd.Stderr = &combined
	err := cmd.Run()
	return combined.Bytes(), err
}

// formatBuildErrorContext wraps the raw build output in a header
// that tells Claude what it's looking at. Reconcile-loop semantics
// is "fix it if it's yours, leave alone otherwise" — the header
// reinforces that.
func formatBuildErrorContext(buildOutput []byte, round, maxRounds int) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(
		"The project's `go build ./...` is failing (reconcile round %d of %d).\n\n",
		round, maxRounds))
	sb.WriteString("Build output:\n\n```\n")
	sb.Write(bytes.TrimSpace(buildOutput))
	sb.WriteString("\n```\n\n")
	sb.WriteString("If any line above references THIS file (by name) or describes a problem ")
	sb.WriteString("with code that THIS file produces, fix it. If the errors are in other files ")
	sb.WriteString("and not your concern, return your file essentially unchanged.\n")
	return sb.String()
}

func hashBytes(b []byte) string {
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}

func totalNodes(levels [][]Node) int {
	n := 0
	for _, l := range levels {
		n += len(l)
	}
	return n
}

// reconcileTimeout is a sanity cap so a runaway reconcile (e.g. a
// model that keeps producing the same garbage) can't burn the user's
// API credits indefinitely. Each round can have its own per-directive
// timeouts via ctx; this is the OUTER bound on the whole loop.
//
// Currently unused — we rely on MaxRounds for the bound. Kept here as
// a placeholder for a future per-loop deadline option if it turns out
// MaxRounds is too coarse.
var _ = 30 * time.Minute
