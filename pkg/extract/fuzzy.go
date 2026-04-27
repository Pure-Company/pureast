// pkg/extract/fuzzy.go
//
// Fuzzy symbol search.
//
// This replaces the old pkg/index abstraction, which built a separate
// SymbolEntry data structure on every call only to throw it away after
// one query. Pureast invocations are one-shot — there's no persistent
// session that benefits from a pre-built index — so we filter the
// already-discovered []SymbolInfo directly.
//
// The five matching strategies (exact, prefix, contains, subsequence,
// initials) and their score values are preserved verbatim from the
// previous implementation so search results are unchanged.

package extract

import (
	"sort"
	"strings"
)

// Match is a search result with its relevance score. Higher is better.
type Match struct {
	Symbol SymbolInfo
	Score  int
}

// FuzzySearch returns symbols matching pattern, optionally filtered by
// kind. Results are sorted by score (best first) and capped at maxResults.
// Pass maxResults <= 0 for no cap.
//
// Empty pattern matches everything (returns all symbols, optionally
// kind-filtered) — this is convenient for "list with kind filter" use
// cases that don't want to also enumerate via DiscoverAllSymbols + a
// separate filter step.
func FuzzySearch(symbols []SymbolInfo, pattern, kind string, maxResults int) []Match {
	pattern = strings.ToLower(pattern)
	matches := make([]Match, 0, len(symbols))

	for _, sym := range symbols {
		if kind != "" && sym.Kind != kind {
			continue
		}
		score, ok := scoreSymbol(pattern, sym.Name)
		if !ok {
			continue
		}
		matches = append(matches, Match{Symbol: sym, Score: score})
	}

	// Stable sort by score desc, then name asc — name tie-breaker keeps
	// output deterministic for prompt caching when scores collide.
	sort.SliceStable(matches, func(i, j int) bool {
		if matches[i].Score != matches[j].Score {
			return matches[i].Score > matches[j].Score
		}
		return matches[i].Symbol.Name < matches[j].Symbol.Name
	})

	if maxResults > 0 && len(matches) > maxResults {
		matches = matches[:maxResults]
	}
	return matches
}

// scoreSymbol returns the best match score across all strategies, or
// (0, false) if none matched. Strategies are tried in cheap-first order
// so we can short-circuit on exact/prefix matches before spending time
// on subsequence walks.
//
// Score ranges (preserved from the previous implementation so existing
// users don't see different rankings):
//
//	exact match        1000
//	prefix match        800
//	contains match      400-600  (decreases with index in target)
//	subsequence match   100-300  (decreases with gap count)
//	initials match       50      (e.g. "MS" → "MyService")
func scoreSymbol(pattern, name string) (int, bool) {
	if pattern == "" {
		// Empty pattern: everything matches with neutral score so the
		// caller's kind filter and ordering still apply.
		return 0, true
	}
	target := strings.ToLower(name)

	if pattern == target {
		return 1000, true
	}
	if strings.HasPrefix(target, pattern) {
		return 800, true
	}
	if idx := strings.Index(target, pattern); idx >= 0 {
		score := 600 - idx*10
		if score < 400 {
			score = 400
		}
		return score, true
	}
	if isSubsequence(pattern, target) {
		score := 300 - countGaps(pattern, target)*5
		if score < 100 {
			score = 100
		}
		return score, true
	}
	if matchesInitials(pattern, name) {
		return 50, true
	}
	return 0, false
}

// isSubsequence reports whether every character of pattern appears in
// target in order (not necessarily contiguous). "Hndl" is a subsequence
// of "Handler".
func isSubsequence(pattern, target string) bool {
	pi := 0
	for i := 0; i < len(target) && pi < len(pattern); i++ {
		if target[i] == pattern[pi] {
			pi++
		}
	}
	return pi == len(pattern)
}

// countGaps measures how spread-out the subsequence match is. "Hndl"
// in "Handler" has 1 gap (between 'd' and 'l'); tighter matches score
// higher than scattered ones.
func countGaps(pattern, target string) int {
	pi := 0
	gaps := 0
	last := -1
	for ti := 0; ti < len(target) && pi < len(pattern); ti++ {
		if target[ti] == pattern[pi] {
			if last >= 0 {
				gaps += ti - last - 1
			}
			last = ti
			pi++
		}
	}
	return gaps
}

// matchesInitials matches a pattern against the uppercase letters in a
// name. "MS" matches "MyService" because the initials are "MS". Pattern
// is uppercased before comparison; name is scanned for capitals as-is
// (so initial-matching is case-aware in the way Go identifiers are).
func matchesInitials(pattern, name string) bool {
	if pattern == "" {
		return false
	}
	var initials []byte
	for i := 0; i < len(name); i++ {
		c := name[i]
		if c >= 'A' && c <= 'Z' {
			initials = append(initials, c)
		}
	}
	return strings.Contains(string(initials), strings.ToUpper(pattern))
}
