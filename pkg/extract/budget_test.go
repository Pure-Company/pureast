// pkg/extract/budget_test.go
package extract

import (
	"strings"
	"testing"
)

func TestEstimateTokens(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int // approximate; we assert ranges, not exact values
	}{
		{"empty", "", 0},
		{"short", "hello", 2},                  // 5 chars * 2 + 6 = 16; 16/7 = 2
		{"sentence", "the quick brown fox", 6}, // 19*2 + 6 = 44; 44/7 = 6
		{"newlines counted", "a\nb\nc\nd", 2},  // 7*2 + 6 = 20; 20/7 = 2
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EstimateTokens(tt.in)
			if got != tt.want {
				t.Errorf("EstimateTokens(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestEstimateTokens_Monotonic(t *testing.T) {
	// Property: longer input never produces fewer tokens.
	// Catches bugs like off-by-one rounding that could make a longer
	// string round down below a shorter one.
	prev := 0
	for i := 0; i < 100; i++ {
		s := strings.Repeat("x", i)
		got := EstimateTokens(s)
		if got < prev {
			t.Fatalf("non-monotonic: len=%d tokens=%d prev=%d", i, got, prev)
		}
		prev = got
	}
}

func TestTruncateLines_NoBudget(t *testing.T) {
	in := "line1\nline2\nline3\n"
	out, truncated := TruncateLines(in, 0)
	if out != in {
		t.Errorf("with maxTokens=0, expected unchanged input")
	}
	if truncated {
		t.Errorf("truncated=true with maxTokens=0")
	}
}

func TestTruncateLines_FitsWhole(t *testing.T) {
	in := "short\n"
	out, truncated := TruncateLines(in, 1000)
	if out != in {
		t.Errorf("short input changed: got %q want %q", out, in)
	}
	if truncated {
		t.Errorf("short input flagged as truncated")
	}
}

func TestTruncateLines_CutsAtLineBoundary(t *testing.T) {
	// Build input with predictable line lengths
	in := strings.Repeat("0123456789\n", 100) // 1100 chars

	out, truncated := TruncateLines(in, 50) // budget ≈ 175 chars
	if !truncated {
		t.Fatal("expected truncation")
	}
	if !strings.Contains(out, "truncated to fit token budget") {
		t.Errorf("truncation marker missing")
	}
	// The output before the marker should end with \n (line boundary)
	beforeMarker := strings.Split(out, "// ... truncated")[0]
	if !strings.HasSuffix(beforeMarker, "\n") {
		t.Errorf("did not cut at line boundary: %q", beforeMarker[len(beforeMarker)-20:])
	}
}

func TestTruncateSymbols_NoBudget(t *testing.T) {
	in := "func A() {}\n\nfunc B() {}\n"
	out, truncated := TruncateSymbols(in, 0)
	if out != in {
		t.Errorf("changed input with maxTokens=0")
	}
	if truncated {
		t.Errorf("truncated=true with maxTokens=0")
	}
}

func TestTruncateSymbols_DropsTrailingSymbol(t *testing.T) {
	// Three "symbols" separated by blank lines. Budget room for ~1.5
	// of them — should keep 1, drop 2.
	in := "func A() {\n  body\n}\n\nfunc B() {\n  body\n}\n\nfunc C() {\n  body\n}\n"

	out, truncated := TruncateSymbols(in, 5) // budget ≈ 17 chars
	if !truncated {
		t.Fatal("expected truncation")
	}
	if strings.Contains(out, "func B") {
		t.Errorf("partial symbol B leaked into output: %q", out)
	}
	if strings.Contains(out, "func C") {
		t.Errorf("partial symbol C leaked into output: %q", out)
	}
	// We do expect A to be intact (the first symbol)
	if !strings.Contains(out, "func A") {
		t.Errorf("first symbol missing entirely; output too aggressive: %q", out)
	}
}

func TestTruncateSymbols_NeverSplitsMidSymbol(t *testing.T) {
	// Property: the substring before the truncation marker, when split
	// on "\n\n", should never end with a non-empty partial chunk that
	// doesn't appear whole in the input. We approximate by asserting
	// the output minus marker is a prefix of the input ending on "\n\n"
	// (or equals the input).
	in := "type A struct {\n  Field int\n}\n\ntype B struct {\n  Other int\n}\n\ntype C struct{}\n"

	out, truncated := TruncateSymbols(in, 8)
	if !truncated {
		t.Fatal("expected truncation")
	}
	// Strip marker
	idx := strings.Index(out, "// ... truncated")
	if idx < 0 {
		t.Fatal("marker missing")
	}
	prefix := strings.TrimRight(out[:idx], "\n")

	// prefix must appear verbatim in the input
	if !strings.Contains(in, prefix) {
		t.Errorf("output prefix not found verbatim in input — mid-symbol split detected\nprefix=%q", prefix)
	}
}

func TestTruncateSymbols_HugeFirstSymbol(t *testing.T) {
	// Edge case: a single symbol bigger than the budget. Falls back to
	// line truncation rather than emitting nothing.
	huge := "func Big() {\n" + strings.Repeat("  // long comment\n", 50) + "}\n"
	in := huge // single chunk, no \n\n separator

	out, truncated := TruncateSymbols(in, 20)
	if !truncated {
		t.Fatal("expected truncation")
	}
	// Should have produced *something* before the marker
	if !strings.Contains(out, "func Big") {
		t.Errorf("fallback didn't preserve the function header: %q", out[:min(100, len(out))])
	}
}

// Determinism property: same input → byte-identical output across calls.
// Critical for prompt caching.
func TestTruncateSymbols_Deterministic(t *testing.T) {
	in := "func A() {}\n\nfunc B() {}\n\nfunc C() {}\n"
	out1, _ := TruncateSymbols(in, 5)
	out2, _ := TruncateSymbols(in, 5)
	if out1 != out2 {
		t.Errorf("non-deterministic output:\n  call 1: %q\n  call 2: %q", out1, out2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
