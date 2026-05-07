// pkg/extract/module_test.go
package extract

import "testing"

func TestSplitSpec(t *testing.T) {
	cases := []struct {
		in              string
		wantPath, wantV string
	}{
		{"github.com/foo/bar", "github.com/foo/bar", ""},
		{"github.com/foo/bar@v1.2.3", "github.com/foo/bar", "v1.2.3"},
		{"github.com/foo/bar@latest", "github.com/foo/bar", "latest"},
		{"github.com/foo/bar/sub@v0.0.0-20240101000000-abcdef", "github.com/foo/bar/sub", "v0.0.0-20240101000000-abcdef"},
		// LastIndex on '@' so a stray earlier '@' (rare but defensible)
		// doesn't cause version truncation.
		{"weird@path/x@v1", "weird@path/x", "v1"},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			gotP, gotV := splitSpec(c.in)
			if gotP != c.wantPath || gotV != c.wantV {
				t.Errorf("splitSpec(%q) = (%q, %q), want (%q, %q)",
					c.in, gotP, gotV, c.wantPath, c.wantV)
			}
		})
	}
}

func TestIsSoftResolutionError(t *testing.T) {
	soft := []string{
		"module foo: not a known dependency",
		"module foo: no matching versions for query \"latest\"",
		"unknown revision v999.0.0",
		"github.com/foo/bar is not a valid module path",
	}
	hard := []string{
		"git ls-remote: connection refused",
		"could not read Username for 'https://github.com'",
		"checksum mismatch",
		"",
	}
	for _, msg := range soft {
		if !isSoftResolutionError(msg) {
			t.Errorf("expected soft: %q", msg)
		}
	}
	for _, msg := range hard {
		if isSoftResolutionError(msg) {
			t.Errorf("expected hard: %q", msg)
		}
	}
}
