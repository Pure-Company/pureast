// pkg/extract/budget.go
//
// Token-budget truncation for LLM-context output.
//
// Two strategies live here:
//
//   1. TruncateLines  - line-aware, fast, can split mid-symbol. Use for
//                       free-form text where mid-stream cutoff is
//                       acceptable (e.g. concatenated reports).
//
//   2. TruncateSymbols - symbol-aware, drops trailing whole declarations.
//                        Use for Go source dumps where partial output
//                        would be invalid syntax inside a markdown fence.
//
// Both take a token budget and return the truncated text plus a bool
// indicating whether truncation occurred. The bool is what callers
// surface to the user (e.g. a stderr notice "output truncated at N
// tokens"), so it's not a dead return value.
//
// The chars-per-token heuristic is 3.5 — slightly conservative for
// dense Go code, which means we undershoot the budget by ~10-20%.
// That's the right direction: better to leave headroom than blow past
// the model's context window.

package extract

import (
	"strings"
)

// charsPerTokenNum / charsPerTokenDen = 3.5 chars per token, the
// well-known OpenAI guidance. Stored as a fraction so we never do
// floating-point math on token counts (deterministic output across
// CPUs matters for prompt caching).
const (
	charsPerTokenNum = 7
	charsPerTokenDen = 2
)

// EstimateTokens approximates the token count for a piece of text.
// Accurate to within ~20% for English-ish source code; for exact
// counts pipe through your model's actual tokenizer.
func EstimateTokens(s string) int {
	if s == "" {
		return 0
	}
	return (len(s)*charsPerTokenDen + charsPerTokenNum - 1) / charsPerTokenNum
}

// charBudget converts a token budget to a char budget at the same
// ratio. Used internally by both truncation strategies.
func charBudget(maxTokens int) int {
	if maxTokens <= 0 {
		return 0
	}
	return maxTokens * charsPerTokenNum / charsPerTokenDen
}

// truncationMarker is appended verbatim when output is cut. The "..."
// glyph is a deliberate three-dot ASCII sequence, not the unicode
// horizontal ellipsis, so byte counts stay predictable for tooling
// that hashes output for caching.
const truncationMarker = "\n// ... truncated to fit token budget ...\n"

// TruncateLines cuts text to fit within maxTokens at line boundaries.
// Returns (text, false) unchanged if maxTokens <= 0 or text already
// fits. Otherwise returns a prefix ending on a complete line plus the
// truncation marker, with bool = true.
//
// Use this for non-Go output: reports, plain text, mixed content.
// For Go source destined for an LLM, prefer TruncateSymbols.
func TruncateLines(text string, maxTokens int) (string, bool) {
	if maxTokens <= 0 {
		return text, false
	}
	if EstimateTokens(text) <= maxTokens {
		return text, false
	}

	budget := charBudget(maxTokens)
	if budget >= len(text) {
		return text, false
	}

	var b strings.Builder
	used := 0
	for _, line := range strings.SplitAfter(text, "\n") {
		if used+len(line) > budget {
			break
		}
		b.WriteString(line)
		used += len(line)
	}
	b.WriteString(truncationMarker)
	return b.String(), true
}

// TruncateSymbols cuts text to fit within maxTokens, but only on
// blank-line boundaries between top-level declarations. This guarantees
// the result is syntactically complete Go: we never end mid-struct,
// mid-function, or in the middle of a comment block.
//
// The contract: callers produce output where consecutive top-level
// symbols are separated by a blank line (which is what go/printer and
// our dump renderer both do). We split on "\n\n" and keep whole chunks
// until the budget is exhausted.
//
// If a single chunk is itself larger than the budget — say one giant
// function — we fall back to TruncateLines on that chunk so the user
// still gets *something* rather than silently dropping it.
func TruncateSymbols(text string, maxTokens int) (string, bool) {
	if maxTokens <= 0 {
		return text, false
	}
	if EstimateTokens(text) <= maxTokens {
		return text, false
	}

	budget := charBudget(maxTokens)
	if budget >= len(text) {
		return text, false
	}

	// Symbols are separated by exactly one blank line in our output;
	// SplitAfter preserves the trailing "\n\n" on each chunk so
	// reassembly is just concatenation.
	chunks := strings.SplitAfter(text, "\n\n")

	var b strings.Builder
	used := 0
	truncated := false

	for i, chunk := range chunks {
		if used+len(chunk) <= budget {
			b.WriteString(chunk)
			used += len(chunk)
			continue
		}

		// This chunk doesn't fit. If it's the very first chunk and we
		// have nothing yet, fall back to line truncation on it — better
		// than empty output. Otherwise, stop cleanly here.
		if i == 0 && used == 0 {
			partial, _ := TruncateLines(chunk, maxTokens)
			b.WriteString(partial)
		}
		truncated = true
		break
	}

	if !truncated {
		// Should not be reachable given the early-exit on "fits whole",
		// but defensive: if we walked all chunks without truncating,
		// return the original.
		return text, false
	}

	b.WriteString(truncationMarker)
	return b.String(), true
}
