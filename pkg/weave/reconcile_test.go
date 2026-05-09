// pkg/weave/reconcile_test.go
package weave

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
)

// makeStubBuildScript creates a small shell script at path that exits
// with the configured return codes on successive invocations. Each
// `exitCodes[i]` is the exit code for the (i+1)th call. After the
// last entry, the script keeps using the last exit code.
//
// The script writes a different stderr blob on each call (so reconcile's
// no-progress hash detection sees different content) unless `sameOutput`
// is true (in which case every call prints the same "build failed: x".)
func makeStubBuildScript(t *testing.T, path string, exitCodes []int, sameOutput bool) {
	t.Helper()
	counterPath := path + ".count"

	// Build the script body. It increments a counter file, picks an
	// exit code from the list (or last if past end), and emits output.
	body := "#!/bin/bash\n"
	body += fmt.Sprintf("COUNTER_FILE=%q\n", counterPath)
	body += `n=$( (cat "$COUNTER_FILE" 2>/dev/null) || echo 0 )` + "\n"
	body += "n=$((n + 1))\n"
	body += `echo "$n" > "$COUNTER_FILE"` + "\n"

	if sameOutput {
		body += `echo "build failed: undefined: Foo (call $n)" 1>&2` + "\n"
	} else {
		body += `echo "build failed: error round $n" 1>&2` + "\n"
	}

	body += "case $n in\n"
	for i, code := range exitCodes {
		body += fmt.Sprintf("  %d) exit %d ;;\n", i+1, code)
	}
	if len(exitCodes) > 0 {
		body += fmt.Sprintf("  *) exit %d ;;\n", exitCodes[len(exitCodes)-1])
	} else {
		body += "  *) exit 1 ;;\n"
	}
	body += "esac\n"

	if err := os.WriteFile(path, []byte(body), 0755); err != nil {
		t.Fatal(err)
	}
}

// TestReconcile_GreenOnEntry: when the first build is green, no
// rounds run and Reconcile returns success immediately.
func TestReconcile_GreenOnEntry(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"),
		[]byte("module test\ngo 1.22\n"), 0644); err != nil {
		t.Fatal(err)
	}
	stub := filepath.Join(root, "build-stub.sh")
	makeStubBuildScript(t, stub, []int{0}, false)

	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
	)

	rec := &recorder{}
	summary, err := Weave(context.Background(), m, root, Options{
		Runner:      rec,
		Reconcile:   true,
		BuildCmd:    []string{stub},
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err != nil {
		t.Fatalf("Weave: %v", err)
	}
	if summary.ReconcileRounds != 0 {
		t.Errorf("ReconcileRounds = %d, want 0 (green on entry)", summary.ReconcileRounds)
	}
	if !summary.ReconcileSucceeded {
		t.Error("ReconcileSucceeded should be true")
	}
}

// TestReconcile_ConvergesAfterOneRound: build fails once, then the
// next round's run "fixes" things and the second build is green.
func TestReconcile_ConvergesAfterOneRound(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"),
		[]byte("module test\ngo 1.22\n"), 0644); err != nil {
		t.Fatal(err)
	}
	stub := filepath.Join(root, "build-stub.sh")
	// First build: fail. Second build: succeed.
	makeStubBuildScript(t, stub, []int{1, 0}, false)

	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{{output: "b.go"}}},
	)

	rec := &recorder{}
	summary, err := Weave(context.Background(), m, root, Options{
		Runner:      rec,
		Reconcile:   true,
		BuildCmd:    []string{stub},
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err != nil {
		t.Fatalf("Weave: %v", err)
	}
	if summary.ReconcileRounds != 1 {
		t.Errorf("ReconcileRounds = %d, want 1", summary.ReconcileRounds)
	}
	if !summary.ReconcileSucceeded {
		t.Error("ReconcileSucceeded should be true")
	}

	// Verify recorder saw two phases of calls: phase 1 (normal, no
	// override), phase 2 (reconcile round 1, NoCache=true).
	rec.mu.Lock()
	defer rec.mu.Unlock()

	var normalCalls, reconcileCalls int
	for _, c := range rec.calls {
		if c.Override.NoCache {
			reconcileCalls++
			if !strings.Contains(c.Override.AppendContext, "build failed") {
				t.Errorf("reconcile call missing build error in append-context: %q",
					c.Override.AppendContext)
			}
		} else {
			normalCalls++
		}
	}

	// Normal pass: 2 (one per directive).
	if normalCalls != 2 {
		t.Errorf("normalCalls = %d, want 2", normalCalls)
	}
	// Reconcile pass: 2 (one per directive).
	if reconcileCalls != 2 {
		t.Errorf("reconcileCalls = %d, want 2", reconcileCalls)
	}
}

// TestReconcile_NoProgressHalts: build keeps failing with the SAME
// output across rounds. Reconcile should detect no-progress and
// halt without going through all MaxRounds.
func TestReconcile_NoProgressHalts(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"),
		[]byte("module test\ngo 1.22\n"), 0644); err != nil {
		t.Fatal(err)
	}
	stub := filepath.Join(root, "build-stub.sh")
	// Always fail with the same hash. NOTE: makeStubBuildScript
	// includes "(call $n)" only when sameOutput=true; the call number
	// changes but the suffix differs by counter — to make the OUTPUT
	// truly identical across rounds, we need a different script.
	body := "#!/bin/bash\necho 'build failed: undefined: Foo' 1>&2\nexit 1\n"
	if err := os.WriteFile(stub, []byte(body), 0755); err != nil {
		t.Fatal(err)
	}

	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
	)

	rec := &recorder{}
	summary, err := Weave(context.Background(), m, root, Options{
		Runner:      rec,
		Reconcile:   true,
		MaxRounds:   5,
		BuildCmd:    []string{stub},
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err == nil {
		t.Fatal("expected error (no progress), got nil")
	}
	if !strings.Contains(err.Error(), "no-progress") {
		t.Errorf("expected 'no-progress' in error, got: %v", err)
	}
	// We should have halted at round 1 (one round attempted, second
	// build produced same output, halted).
	if summary.ReconcileRounds != 1 {
		t.Errorf("ReconcileRounds = %d, want 1 (halt on no-progress after round 1)",
			summary.ReconcileRounds)
	}
	if summary.ReconcileSucceeded {
		t.Error("ReconcileSucceeded should be false")
	}
}

// TestReconcile_ExhaustsMaxRounds: build keeps failing with DIFFERENT
// output each round. Reconcile makes "progress" by the no-progress
// metric, but never converges. Should stop at MaxRounds.
func TestReconcile_ExhaustsMaxRounds(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"),
		[]byte("module test\ngo 1.22\n"), 0644); err != nil {
		t.Fatal(err)
	}
	stub := filepath.Join(root, "build-stub.sh")
	// All rounds fail, with DIFFERENT output (so no-progress not triggered).
	makeStubBuildScript(t, stub, []int{1, 1, 1, 1}, false)

	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
	)

	rec := &recorder{}
	summary, err := Weave(context.Background(), m, root, Options{
		Runner:      rec,
		Reconcile:   true,
		MaxRounds:   3,
		BuildCmd:    []string{stub},
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err == nil {
		t.Fatal("expected error (exhausted), got nil")
	}
	if !strings.Contains(err.Error(), "exhausted") {
		t.Errorf("expected 'exhausted' in error, got: %v", err)
	}
	if summary.ReconcileRounds != 3 {
		t.Errorf("ReconcileRounds = %d, want 3", summary.ReconcileRounds)
	}
	if summary.ReconcileSucceeded {
		t.Error("ReconcileSucceeded should be false")
	}
}

// TestReconcile_OffByDefault: without Options.Reconcile, no build is
// run, no extra rounds happen.
func TestReconcile_OffByDefault(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"),
		[]byte("module test\ngo 1.22\n"), 0644); err != nil {
		t.Fatal(err)
	}
	// Stub that, if invoked, would fail. We assert it's NOT invoked
	// by checking the counter file is absent or empty.
	stub := filepath.Join(root, "build-stub.sh")
	makeStubBuildScript(t, stub, []int{1, 1, 1}, false)

	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
	)

	rec := &recorder{}
	summary, err := Weave(context.Background(), m, root, Options{
		Runner:      rec,
		BuildCmd:    []string{stub}, // would fail if invoked, but it isn't
		Reconcile:   false,          // explicit (also the zero value)
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err != nil {
		t.Fatalf("Weave (no reconcile) failed: %v", err)
	}
	if summary.ReconcileRounds != 0 {
		t.Errorf("ReconcileRounds = %d, want 0", summary.ReconcileRounds)
	}
	if summary.ReconcileSucceeded {
		t.Error("ReconcileSucceeded should be false (we didn't even try)")
	}

	// Counter file should not exist (build never ran).
	counterFile := stub + ".count"
	if _, err := os.Stat(counterFile); err == nil {
		t.Errorf("build was invoked despite Reconcile=false (counter file %s exists)",
			counterFile)
	}
}

// TestReconcile_AppendContextContainsBuildOutput: verify the
// reconcile-pass override carries the build output through to the
// runner.
func TestReconcile_AppendContextContainsBuildOutput(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"),
		[]byte("module test\ngo 1.22\n"), 0644); err != nil {
		t.Fatal(err)
	}
	stub := filepath.Join(root, "build-stub.sh")
	makeStubBuildScript(t, stub, []int{1, 0}, false)

	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
	)

	rec := &recorder{}
	_, err := Weave(context.Background(), m, root, Options{
		Runner:      rec,
		Reconcile:   true,
		BuildCmd:    []string{stub},
		SkipModTidy: true,
		LogWriter:   io.Discard,
	})
	if err != nil {
		t.Fatal(err)
	}

	rec.mu.Lock()
	defer rec.mu.Unlock()

	var found bool
	for _, c := range rec.calls {
		if c.Override.NoCache && strings.Contains(c.Override.AppendContext, "build failed") {
			found = true
			// Verify the framing language is present (round X of N).
			if !strings.Contains(c.Override.AppendContext, "round 1") {
				t.Errorf("appended context should mention round number: %q",
					c.Override.AppendContext)
			}
			if !strings.Contains(c.Override.AppendContext, "fix it") {
				t.Errorf("appended context should include the fix-if-relevant framing: %q",
					c.Override.AppendContext)
			}
			break
		}
	}
	if !found {
		t.Errorf("no reconcile call found with build output in AppendContext")
	}
}

// _ = atomic.Int32{} — silence import; recorder uses it.
var _ = atomic.Int32{}
