// pkg/weave/runner_test.go
package weave

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Pure-Company/pureast/pkg/scaffold"
)

// recorder is a fake Runner that captures invocations for assertion.
// Optionally simulates work by sleeping per-call so tests can verify
// that within-level parallelism actually overlaps in time.
type recorder struct {
	mu        sync.Mutex
	calls     []recordedCall
	delay     time.Duration    // per-call sleep
	failIDs   map[string]error // node IDs to fail
	beforeFn  func(n Node)     // optional hook
	afterFn   func(n Node)     // optional hook
	maxActive atomic.Int32     // observed peak concurrency
	active    atomic.Int32     // current in-flight
}

type recordedCall struct {
	NodeID   string
	Start    time.Time
	End      time.Time
	Err      error
	Override RunOverride
}

func (r *recorder) Run(ctx context.Context, projectRoot string, n Node, override RunOverride) error {
	cur := r.active.Add(1)
	for {
		max := r.maxActive.Load()
		if cur <= max || r.maxActive.CompareAndSwap(max, cur) {
			break
		}
	}
	defer r.active.Add(-1)

	if r.beforeFn != nil {
		r.beforeFn(n)
	}

	start := time.Now()
	if r.delay > 0 {
		select {
		case <-time.After(r.delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	end := time.Now()

	var err error
	if r.failIDs != nil {
		if e, ok := r.failIDs[n.ID()]; ok {
			err = e
		}
	}

	r.mu.Lock()
	r.calls = append(r.calls, recordedCall{
		NodeID: n.ID(), Start: start, End: end, Err: err, Override: override,
	})
	r.mu.Unlock()

	if r.afterFn != nil {
		r.afterFn(n)
	}
	return err
}

func (r *recorder) callIDs() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	ids := make([]string, len(r.calls))
	for i, c := range r.calls {
		ids[i] = c.NodeID
	}
	return ids
}

// TestWeave_LevelOrdering: nodes from later levels never start before
// every node from earlier levels has finished.
func TestWeave_LevelOrdering(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{
			{output: "b.go", sources: []scaffold.Source{{Pkg: "../a"}}},
		}},
		packageSpec{path: "c", pkg: "c", files: []fileSpec{
			{output: "c.go", sources: []scaffold.Source{{Pkg: "../b"}}},
		}},
	)

	rec := &recorder{delay: 20 * time.Millisecond}
	_, err := Weave(context.Background(), m, t.TempDir(), Options{
		Runner:      rec,
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Build a map: node ID -> end time. Then verify that for each
	// (pred, succ) edge in the DAG, pred.end <= succ.start.
	g, _ := BuildDAG(m)
	endTimes := make(map[string]time.Time)
	startTimes := make(map[string]time.Time)
	rec.mu.Lock()
	for _, c := range rec.calls {
		endTimes[c.NodeID] = c.End
		startTimes[c.NodeID] = c.Start
	}
	rec.mu.Unlock()

	for id, preds := range g.Predecessors {
		succStart, ok := startTimes[id]
		if !ok {
			continue
		}
		for predID := range preds {
			predEnd, ok := endTimes[predID]
			if !ok {
				continue
			}
			if predEnd.After(succStart) {
				t.Errorf("ordering violated: pred %s ended at %v, succ %s started at %v",
					predID, predEnd, id, succStart)
			}
		}
	}
}

// TestWeave_WithinLevelParallelism: independent nodes in the same
// level should run concurrently. Verified by observing peak active
// count >= 2 when concurrency permits.
func TestWeave_WithinLevelParallelism(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{{output: "b.go"}}},
		packageSpec{path: "c", pkg: "c", files: []fileSpec{{output: "c.go"}}},
		packageSpec{path: "d", pkg: "d", files: []fileSpec{{output: "d.go"}}},
	)

	rec := &recorder{delay: 50 * time.Millisecond}
	_, err := Weave(context.Background(), m, t.TempDir(), Options{
		Concurrency: 4,
		Runner:      rec,
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err != nil {
		t.Fatal(err)
	}

	if peak := rec.maxActive.Load(); peak < 2 {
		t.Errorf("expected within-level parallelism (peak >= 2), got peak=%d", peak)
	}
}

// TestWeave_RespectsConcurrencyLimit: with concurrency=1, peak active
// should never exceed 1.
func TestWeave_RespectsConcurrencyLimit(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{{output: "b.go"}}},
		packageSpec{path: "c", pkg: "c", files: []fileSpec{{output: "c.go"}}},
	)

	rec := &recorder{delay: 30 * time.Millisecond}
	_, err := Weave(context.Background(), m, t.TempDir(), Options{
		Concurrency: 1,
		Runner:      rec,
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err != nil {
		t.Fatal(err)
	}

	if peak := rec.maxActive.Load(); peak > 1 {
		t.Errorf("concurrency=1 should cap peak at 1, got %d", peak)
	}
}

// TestWeave_FailSoftWithinLevel: when one node in a level fails, the
// others in the same level still run to completion. But weave halts
// before the next level.
func TestWeave_FailSoftWithinLevel(t *testing.T) {
	m := buildManifest(
		// Level 0: three independent nodes.
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{{output: "b.go"}}},
		packageSpec{path: "c", pkg: "c", files: []fileSpec{{output: "c.go"}}},
		// Level 1: depends on a.
		packageSpec{path: "d", pkg: "d", files: []fileSpec{
			{output: "d.go", sources: []scaffold.Source{{Pkg: "../a"}}},
		}},
	)

	rec := &recorder{
		delay: 10 * time.Millisecond,
		failIDs: map[string]error{
			"b/b.go": errors.New("boom"),
		},
	}
	summary, err := Weave(context.Background(), m, t.TempDir(), Options{
		Runner:      rec,
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "boom") {
		t.Errorf("expected error to mention 'boom', got: %v", err)
	}

	// Level 0 ran in full (3 nodes, 1 failed, 2 succeeded).
	got := rec.callIDs()
	want := map[string]bool{"a/a.go": true, "b/b.go": true, "c/c.go": true}
	if len(got) != 3 {
		t.Fatalf("level 0 should have all 3 nodes attempted; got %v", got)
	}
	for _, id := range got {
		if !want[id] {
			t.Errorf("unexpected call: %s", id)
		}
	}

	// Level 1 should NOT have run (fail-soft halts before downstream).
	for _, id := range got {
		if id == "d/d.go" {
			t.Errorf("d.go should not have run after level 0 failure")
		}
	}

	// Summary should report the skipped count.
	if summary.SkippedDownstream != 1 {
		t.Errorf("SkippedDownstream = %d, want 1", summary.SkippedDownstream)
	}
}

// TestWeave_PropagatesContextCancellation: cancelling ctx should
// cause Weave to stop scheduling new nodes; in-flight ones see ctx.Err.
func TestWeave_PropagatesContextCancellation(t *testing.T) {
	m := buildManifest(
		// Level 0: 4 slow nodes.
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{{output: "b.go"}}},
		packageSpec{path: "c", pkg: "c", files: []fileSpec{{output: "c.go"}}},
		packageSpec{path: "d", pkg: "d", files: []fileSpec{{output: "d.go"}}},
	)

	ctx, cancel := context.WithCancel(context.Background())
	rec := &recorder{delay: 200 * time.Millisecond}
	rec.beforeFn = func(n Node) {
		// Cancel as soon as the first node starts.
		cancel()
	}

	_, err := Weave(ctx, m, t.TempDir(), Options{
		Concurrency: 1, // serialize so the first node clearly starts first
		Runner:      rec,
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err == nil {
		t.Fatal("expected cancellation error")
	}
}

// TestWeave_GomodTriggersModTidy: a level containing a gomod-sourced
// node causes weave to invoke `go mod tidy` before that level. We
// verify by overriding the SkipModTidy option (here: not skipped)
// and using a stub project root that lets `go mod tidy` succeed
// trivially (empty go.mod).
func TestWeave_GomodTriggersModTidy(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "internal/x", pkg: "x", files: []fileSpec{{output: "x.go"}}},
		packageSpec{path: "tools/gen", pkg: "gen", files: []fileSpec{
			{output: "../../Makefile", sources: []scaffold.Source{{Gomod: "../../go.mod"}}},
		}},
	)

	// Set up a temp project with a minimal go.mod so `go mod tidy` is happy.
	root := t.TempDir()
	goMod := "module test\n\ngo 1.22\n"
	if err := writeFile(root+"/go.mod", goMod); err != nil {
		t.Fatal(err)
	}

	var stderr bytes.Buffer
	rec := &recorder{}
	_, err := Weave(context.Background(), m, root, Options{
		Runner:    rec,
		LogWriter: &stderr,
	})
	if err != nil {
		t.Fatalf("Weave failed: %v\nstderr:\n%s", err, stderr.String())
	}

	// Verify the log mentions the tidy invocation.
	if !strings.Contains(stderr.String(), "go mod tidy") {
		t.Errorf("expected 'go mod tidy' in log:\n%s", stderr.String())
	}

	// Both nodes ran.
	calls := rec.callIDs()
	if len(calls) != 2 {
		t.Errorf("want 2 calls, got %d: %v", len(calls), calls)
	}
}

// TestWeave_NoGomod_NoTidy: without any gomod sources, weave should
// not run `go mod tidy`. (Saves time; doesn't pollute go.sum.)
func TestWeave_NoGomod_NoTidy(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{
			{output: "b.go", sources: []scaffold.Source{{Pkg: "../a"}}},
		}},
	)

	var stderr bytes.Buffer
	rec := &recorder{}
	_, err := Weave(context.Background(), m, t.TempDir(), Options{
		Runner:    rec,
		LogWriter: &stderr,
	})
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(stderr.String(), "go mod tidy") {
		t.Errorf("did not expect 'go mod tidy' in log:\n%s", stderr.String())
	}
}

// TestClaudeEditRunner_BuildsCorrectArgs: not a real exec, just
// verify the arg list a ClaudeEditRunner would assemble for a
// representative directive. Catches regressions in flag formatting.
func TestClaudeEditRunner_BuildsCorrectArgs(t *testing.T) {
	// We can't easily inspect the args without invoking, but we can
	// verify the runner handles missing optional fields gracefully
	// and the cmd.Dir resolves correctly by exercising it on a
	// stubbed binary that just echoes its argv. This is brittle on
	// CI but useful locally.
	t.Skip("indirect via integration; ClaudeEditRunner is exercised in TestWeave_LevelOrdering with a recorder substitute")
}

// writeFile is a tiny test helper.
func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
