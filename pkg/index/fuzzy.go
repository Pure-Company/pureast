package index

import (
	"strings"

	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/functor"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
)

// FuzzyScore represents a fuzzy match score (this is a monoid!)
type FuzzyScore struct {
	Matched bool
	Score   int
}

// FuzzyScoreMonoid combines fuzzy scores
type FuzzyScoreMonoid struct{}

func NewFuzzyScoreMonoid() FuzzyScoreMonoid {
	return FuzzyScoreMonoid{}
}

// Empty returns no match
func (FuzzyScoreMonoid) Empty() FuzzyScore {
	return FuzzyScore{Matched: false, Score: 0}
}

// Combine takes the better score (max)
func (FuzzyScoreMonoid) Combine(a, b FuzzyScore) FuzzyScore {
	if !a.Matched && !b.Matched {
		return FuzzyScore{Matched: false, Score: 0}
	}
	if !a.Matched {
		return b
	}
	if !b.Matched {
		return a
	}
	// Both matched - take higher score
	if a.Score > b.Score {
		return a
	}
	return b
}

// MatchStrategy represents a pure matching strategy
type MatchStrategy func(pattern, target string) FuzzyScore

// FuzzyMatch performs fuzzy matching using Concurrent applicative
// This runs all strategies in parallel and combines results using monoid
func FuzzyMatch(pattern, target string) FuzzyScore {
	if pattern == "" {
		return FuzzyScore{Matched: true, Score: 0}
	}

	pattern = strings.ToLower(pattern)
	target = strings.ToLower(target)

	// Define matching strategies as computations
	strategies := []MatchStrategy{
		exactMatch,
		prefixMatch,
		containsMatch,
		subsequenceMatch,
		initialsMatch,
	}

	scoreMonoid := NewFuzzyScoreMonoid()

	// Use TraverseConcurrent to run all strategies in parallel
	// and combine results using the monoid
	result := functor.TraverseConcurrent(
		scoreMonoid,
		func(strategy MatchStrategy) FuzzyScore {
			return strategy(pattern, target)
		},
		strategies,
		0, // Use all available CPUs
	)

	return result.Value()
}

// FuzzyMatchSequential performs fuzzy matching sequentially (for comparison)
func FuzzyMatchSequential(pattern, target string) FuzzyScore {
	if pattern == "" {
		return FuzzyScore{Matched: true, Score: 0}
	}

	pattern = strings.ToLower(pattern)
	target = strings.ToLower(target)

	strategies := []MatchStrategy{
		exactMatch,
		prefixMatch,
		containsMatch,
		subsequenceMatch,
		initialsMatch,
	}

	scoreMonoid := NewFuzzyScoreMonoid()

	// Use fold for sequential evaluation
	return fold.FoldLeft(
		func(acc FuzzyScore, strategy MatchStrategy) FuzzyScore {
			score := strategy(pattern, target)
			return scoreMonoid.Combine(acc, score)
		},
		scoreMonoid.Empty(),
		strategies,
	)
}

// exactMatch - highest score (pure function)
func exactMatch(pattern, target string) FuzzyScore {
	if pattern == target {
		return FuzzyScore{Matched: true, Score: 1000}
	}
	return FuzzyScore{Matched: false, Score: 0}
}

// prefixMatch - high score (pure function)
func prefixMatch(pattern, target string) FuzzyScore {
	if strings.HasPrefix(target, pattern) {
		return FuzzyScore{Matched: true, Score: 800}
	}
	return FuzzyScore{Matched: false, Score: 0}
}

// containsMatch - medium-high score (pure function)
func containsMatch(pattern, target string) FuzzyScore {
	if strings.Contains(target, pattern) {
		index := strings.Index(target, pattern)
		score := 600 - (index * 10)
		if score < 400 {
			score = 400
		}
		return FuzzyScore{Matched: true, Score: score}
	}
	return FuzzyScore{Matched: false, Score: 0}
}

// subsequenceMatch - medium score (pure function)
func subsequenceMatch(pattern, target string) FuzzyScore {
	if isSubsequence(pattern, target) {
		gaps := countGaps(pattern, target)
		score := 300 - (gaps * 5)
		if score < 100 {
			score = 100
		}
		return FuzzyScore{Matched: true, Score: score}
	}
	return FuzzyScore{Matched: false, Score: 0}
}

// initialsMatch - low score (pure function)
func initialsMatch(pattern, target string) FuzzyScore {
	initials := extractInitials(target)
	if strings.Contains(initials, strings.ToUpper(pattern)) {
		return FuzzyScore{Matched: true, Score: 50}
	}
	return FuzzyScore{Matched: false, Score: 0}
}

// isSubsequence checks if pattern is a subsequence of target (pure)
func isSubsequence(pattern, target string) bool {
	pIdx := 0

	for _, tChar := range target {
		if pIdx < len(pattern) && pattern[pIdx] == byte(tChar) {
			pIdx++
		}
	}

	return pIdx == len(pattern)
}

// countGaps counts gaps in subsequence match (pure)
func countGaps(pattern, target string) int {
	pIdx := 0
	gaps := 0
	lastMatch := -1

	for tIdx, tChar := range target {
		if pIdx < len(pattern) && pattern[pIdx] == byte(tChar) {
			if lastMatch >= 0 {
				gaps += (tIdx - lastMatch - 1)
			}
			lastMatch = tIdx
			pIdx++
		}
	}

	return gaps
}

// extractInitials extracts capital letters (pure)
func extractInitials(s string) string {
	initials := []rune{}

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			initials = append(initials, char)
		}
	}

	return string(initials)
}

// ScoredSymbol represents a symbol with its match score
type ScoredSymbol struct {
	Entry SymbolEntry
	Score FuzzyScore
}

// ScoredSymbolMonoid combines scored symbols (for parallel search)
type ScoredSymbolMonoid struct{}

func NewScoredSymbolMonoid() ScoredSymbolMonoid {
	return ScoredSymbolMonoid{}
}

func (ScoredSymbolMonoid) Empty() []ScoredSymbol {
	return []ScoredSymbol{}
}

func (ScoredSymbolMonoid) Combine(a, b []ScoredSymbol) []ScoredSymbol {
	return append(a, b...)
}

// FuzzySearchIndex searches with fuzzy matching using parallel map
func FuzzySearchIndex(pattern SearchPattern, idx Index) []ScoredSymbol {
	symbols := valuesFromMap(idx.Symbols)

	// Use ParMap to score all symbols in parallel
	scored := functor.ParMapWithWorkers(
		func(entry SymbolEntry) monoid.Option[ScoredSymbol] {
			score := FuzzyMatch(pattern.SymbolPattern, entry.Name)

			// Filter during map using Option
			if !score.Matched {
				return monoid.None[ScoredSymbol]()
			}

			// Apply additional filters
			if pattern.Kind != "" && entry.Kind != pattern.Kind {
				return monoid.None[ScoredSymbol]()
			}

			if pattern.PackageName != "" && entry.PackageName != pattern.PackageName {
				return monoid.None[ScoredSymbol]()
			}

			return monoid.Some(ScoredSymbol{Entry: entry, Score: score})
		},
		symbols,
		0, // Use all CPUs
	)

	// Filter out None values and extract Some values
	matches := fold.FoldLeft(
		func(acc []ScoredSymbol, opt monoid.Option[ScoredSymbol]) []ScoredSymbol {
			if opt.Valid {
				return append(acc, opt.Value)
			}
			return acc
		},
		[]ScoredSymbol{},
		scored,
	)

	// Sort by score
	return sortByScore(matches)
}

// FuzzySearchIndexConcurrent uses Concurrent applicative for batch processing
func FuzzySearchIndexConcurrent(pattern SearchPattern, idx Index, batchSize int) []ScoredSymbol {
	symbols := valuesFromMap(idx.Symbols)

	scoreMonoid := NewScoredSymbolMonoid()

	// Process in batches using ConcurrentBatch
	processBatch := func(batch []SymbolEntry) []ScoredSymbol {
		// Map each entry in batch
		return fold.FoldLeft(
			func(acc []ScoredSymbol, entry SymbolEntry) []ScoredSymbol {
				score := FuzzyMatch(pattern.SymbolPattern, entry.Name)

				if !score.Matched {
					return acc
				}

				if pattern.Kind != "" && entry.Kind != pattern.Kind {
					return acc
				}

				if pattern.PackageName != "" && entry.PackageName != pattern.PackageName {
					return acc
				}

				return append(acc, ScoredSymbol{Entry: entry, Score: score})
			},
			[]ScoredSymbol{},
			batch,
		)
	}

	result := functor.ConcurrentBatch(
		scoreMonoid,
		batchSize,
		processBatch,
		symbols,
	)

	matches := result.Value()
	return sortByScore(matches)
}

// sortByScore sorts symbols by score (pure using fold)
func sortByScore(symbols []ScoredSymbol) []ScoredSymbol {
	if len(symbols) <= 1 {
		return symbols
	}

	// Use insertion sort via fold
	return fold.FoldLeft(
		func(acc []ScoredSymbol, current ScoredSymbol) []ScoredSymbol {
			return insertSorted(acc, current)
		},
		[]ScoredSymbol{},
		symbols,
	)
}

// insertSorted inserts symbol in sorted position (pure)
func insertSorted(sorted []ScoredSymbol, item ScoredSymbol) []ScoredSymbol {
	for i, s := range sorted {
		if item.Score.Score > s.Score.Score {
			result := make([]ScoredSymbol, 0, len(sorted)+1)
			result = append(result, sorted[:i]...)
			result = append(result, item)
			result = append(result, sorted[i:]...)
			return result
		}
	}
	return append(sorted, item)
}

// valuesFromMap extracts values from a map (helper)
func valuesFromMap(m map[string]SymbolEntry) []SymbolEntry {
	values := make([]SymbolEntry, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
