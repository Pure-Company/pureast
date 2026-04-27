// cmd/pureast/commands/diff_hunks_test.go
package commands

import (
	"testing"
)

func TestParseHunkHeader(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		wantStart int
		wantEnd   int
		wantOK    bool
	}{
		{"basic", "@@ -10,3 +20,5 @@", 20, 25, true},
		{"no count means 1", "@@ -10 +20 @@ context", 20, 21, true},
		{"new file pure-add", "@@ -0,0 +1,5 @@", 1, 6, true},
		{"deletion only — count zero on new side", "@@ -10,5 +20,0 @@", 0, 0, false},
		{"missing plus", "@@ malformed @@", 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, ok := parseHunkHeader(tt.line)
			if ok != tt.wantOK {
				t.Errorf("ok=%v want %v (line: %q)", ok, tt.wantOK, tt.line)
				return
			}
			if !ok {
				return
			}
			if r.Start != tt.wantStart {
				t.Errorf("start=%d want %d", r.Start, tt.wantStart)
			}
			if r.End != tt.wantEnd {
				t.Errorf("end=%d want %d", r.End, tt.wantEnd)
			}
		})
	}
}

func TestParseGitFilePath(t *testing.T) {
	tests := []struct {
		line, want string
	}{
		{"+++ b/path/to/file.go", "path/to/file.go"},
		{"+++ /dev/null", ""},
		{"+++ b/spaces in path/file.go", "spaces in path/file.go"},
	}
	for _, tt := range tests {
		got := parseGitFilePath(tt.line, "+++ ")
		if got != tt.want {
			t.Errorf("parseGitFilePath(%q) = %q, want %q", tt.line, got, tt.want)
		}
	}
}

func TestRangeOverlaps(t *testing.T) {
	hunks := []hunkRange{
		{Start: 10, End: 15}, // changed lines 10..14
		{Start: 30, End: 31}, // changed line 30
	}

	tests := []struct {
		name             string
		symStart, symEnd int
		want             bool
	}{
		{"symbol entirely before all hunks", 1, 9, false},
		{"symbol overlaps first hunk start", 5, 12, true},
		{"symbol entirely inside first hunk", 11, 13, true},
		{"symbol contains first hunk", 8, 20, true},
		{"symbol between hunks", 16, 25, false},
		{"symbol overlaps single-line hunk", 28, 35, true},
		{"symbol entirely after hunks", 100, 200, false},
		{"symbol single line equal to hunk", 30, 30, true},
		// Edge: hunk is half-open [10,15), symbol [15,20] should NOT overlap
		{"symbol starts at hunk end (exclusive)", 15, 20, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rangeOverlaps(tt.symStart, tt.symEnd, hunks)
			if got != tt.want {
				t.Errorf("overlaps([%d,%d], hunks) = %v, want %v",
					tt.symStart, tt.symEnd, got, tt.want)
			}
		})
	}
}

func TestRangeOverlaps_SwappedRange(t *testing.T) {
	// Defensive: if symEnd < symStart we should still produce a sane
	// answer (the impl normalizes to symStart only).
	got := rangeOverlaps(20, 10, []hunkRange{{Start: 18, End: 22}})
	if !got {
		t.Errorf("expected overlap when symStart > symEnd; impl should normalize")
	}
}
