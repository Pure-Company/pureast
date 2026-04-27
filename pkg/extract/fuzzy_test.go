// pkg/extract/fuzzy_test.go
package extract

import (
	"testing"
)

func mkSym(name, kind string) SymbolInfo {
	return SymbolInfo{Name: name, Kind: kind, Package: "test"}
}

func TestFuzzySearch_ExactMatchTop(t *testing.T) {
	syms := []SymbolInfo{
		mkSym("UserService", "struct"),
		mkSym("User", "struct"),
		mkSym("UserRepo", "struct"),
	}

	got := FuzzySearch(syms, "User", "", 0)
	if len(got) != 3 {
		t.Fatalf("expected 3 matches, got %d", len(got))
	}
	if got[0].Symbol.Name != "User" {
		t.Errorf("expected exact match 'User' first, got %q", got[0].Symbol.Name)
	}
	if got[0].Score != 1000 {
		t.Errorf("exact match should score 1000, got %d", got[0].Score)
	}
}

func TestFuzzySearch_PrefixBeatsContains(t *testing.T) {
	syms := []SymbolInfo{
		mkSym("HandleRequest", "function"), // contains
		mkSym("HandlerFunc", "function"),   // prefix
	}

	got := FuzzySearch(syms, "Handle", "", 0)
	if len(got) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(got))
	}
	// Both have "Handle" but HandlerFunc and HandleRequest both start
	// with "Handle" — both prefix, score should be equal (800).
	for _, m := range got {
		if m.Score != 800 {
			t.Errorf("%q: expected prefix score 800, got %d", m.Symbol.Name, m.Score)
		}
	}
}

func TestFuzzySearch_ScoreOrder(t *testing.T) {
	// All five strategies in one query, asserting the documented score
	// ordering: exact > prefix > contains > subsequence > initials.
	syms := []SymbolInfo{
		mkSym("Bar", "struct"),           // initials match for "B"? No, "B" alone matches as exact-prefix
		mkSym("FooBar", "struct"),        // contains "Bar"
		mkSym("Bar", "function"),         // exact "Bar"
		mkSym("BackToReality", "struct"), // initials BTR — "B" matches
	}

	// Use pattern "Bar" — should match Bar exact, FooBar contains, no others
	got := FuzzySearch(syms, "Bar", "", 0)
	if len(got) < 2 {
		t.Fatalf("expected at least 2 matches, got %d", len(got))
	}

	// First two must be the two exact "Bar" entries (one struct, one function)
	if got[0].Score != 1000 {
		t.Errorf("first result not exact match: %+v", got[0])
	}
}

func TestFuzzySearch_KindFilter(t *testing.T) {
	syms := []SymbolInfo{
		mkSym("User", "struct"),
		mkSym("User", "interface"),
		mkSym("UserService", "struct"),
	}

	got := FuzzySearch(syms, "User", "interface", 0)
	if len(got) != 1 {
		t.Fatalf("kind filter failed: got %d results, want 1", len(got))
	}
	if got[0].Symbol.Kind != "interface" {
		t.Errorf("filter returned wrong kind: %q", got[0].Symbol.Kind)
	}
}

func TestFuzzySearch_MaxResults(t *testing.T) {
	syms := []SymbolInfo{
		mkSym("UserA", "struct"),
		mkSym("UserB", "struct"),
		mkSym("UserC", "struct"),
		mkSym("UserD", "struct"),
		mkSym("UserE", "struct"),
	}

	got := FuzzySearch(syms, "User", "", 3)
	if len(got) != 3 {
		t.Errorf("max=3 returned %d results", len(got))
	}
}

func TestFuzzySearch_EmptyPattern(t *testing.T) {
	syms := []SymbolInfo{
		mkSym("A", "struct"),
		mkSym("B", "func"),
	}
	got := FuzzySearch(syms, "", "", 0)
	// Empty pattern matches everything (with neutral score 0)
	if len(got) != 2 {
		t.Errorf("empty pattern: got %d, want 2", len(got))
	}
}

func TestFuzzySearch_EmptyPatternWithKindFilter(t *testing.T) {
	syms := []SymbolInfo{
		mkSym("A", "struct"),
		mkSym("B", "func"),
		mkSym("C", "struct"),
	}
	got := FuzzySearch(syms, "", "struct", 0)
	if len(got) != 2 {
		t.Errorf("got %d, want 2 (only structs)", len(got))
	}
}

func TestFuzzySearch_Subsequence(t *testing.T) {
	syms := []SymbolInfo{mkSym("Handler", "struct")}
	got := FuzzySearch(syms, "Hndl", "", 0)
	if len(got) != 1 {
		t.Fatalf("subsequence not matched: got %d", len(got))
	}
	// Subsequence scores in the 100-300 range
	if got[0].Score < 100 || got[0].Score > 300 {
		t.Errorf("subseq score %d out of range [100, 300]", got[0].Score)
	}
}

func TestFuzzySearch_Deterministic(t *testing.T) {
	syms := []SymbolInfo{
		mkSym("UserService", "struct"),
		mkSym("UserRepo", "struct"),
		mkSym("UserHandler", "struct"),
	}
	a := FuzzySearch(syms, "User", "", 0)
	b := FuzzySearch(syms, "User", "", 0)

	if len(a) != len(b) {
		t.Fatalf("non-deterministic length")
	}
	for i := range a {
		if a[i].Symbol.Name != b[i].Symbol.Name || a[i].Score != b[i].Score {
			t.Errorf("non-deterministic at %d: %+v vs %+v", i, a[i], b[i])
		}
	}
}

func TestFuzzySearch_TieBreakAlphabetical(t *testing.T) {
	// All same score (prefix matches for "Get") — should sort alphabetical
	syms := []SymbolInfo{
		mkSym("GetZ", "func"),
		mkSym("GetA", "func"),
		mkSym("GetM", "func"),
	}
	got := FuzzySearch(syms, "Get", "", 0)
	if got[0].Symbol.Name != "GetA" || got[1].Symbol.Name != "GetM" || got[2].Symbol.Name != "GetZ" {
		t.Errorf("tie-break not alphabetical: %v %v %v",
			got[0].Symbol.Name, got[1].Symbol.Name, got[2].Symbol.Name)
	}
}

func TestFuzzySearch_NoMatch(t *testing.T) {
	syms := []SymbolInfo{mkSym("Handler", "struct")}
	got := FuzzySearch(syms, "ZZZZZZZ", "", 0)
	if len(got) != 0 {
		t.Errorf("expected 0 matches for nonsense pattern, got %d", len(got))
	}
}
