// pkg/extract/gomod_test.go
package extract

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func writeGoMod(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "go.mod")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestParseGoMod_DirectOnly(t *testing.T) {
	path := writeGoMod(t, `
module example.com/test

go 1.22

require (
	github.com/foo/direct1 v1.0.0
	github.com/foo/direct2/v2 v2.5.1
)

require (
	github.com/foo/indirect1 v0.0.1 // indirect
	github.com/foo/indirect2/v3 v3.1.0 // indirect
)
`)
	refs, err := ParseGoMod(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(refs) != 2 {
		t.Fatalf("expected 2 direct deps, got %d: %+v", len(refs), refs)
	}
	want := []ModuleRef{
		{Path: "github.com/foo/direct1", Version: "v1.0.0", OriginalPath: "github.com/foo/direct1"},
		{Path: "github.com/foo/direct2/v2", Version: "v2.5.1", OriginalPath: "github.com/foo/direct2/v2"},
	}
	if !reflect.DeepEqual(refs, want) {
		t.Errorf("got %+v\nwant %+v", refs, want)
	}
}

func TestParseGoMod_ReplaceWithVersion(t *testing.T) {
	path := writeGoMod(t, `
module example.com/test
go 1.22
require github.com/foo/bar v1.0.0
replace github.com/foo/bar => github.com/forked/bar v1.0.1
`)
	refs, err := ParseGoMod(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(refs) != 1 {
		t.Fatalf("expected 1 dep, got %d", len(refs))
	}
	r := refs[0]
	if r.Path != "github.com/forked/bar" || r.Version != "v1.0.1" {
		t.Errorf("replace not applied: got %+v", r)
	}
	if r.OriginalPath != "github.com/foo/bar" {
		t.Errorf("OriginalPath should be the require LHS: got %q", r.OriginalPath)
	}
}

func TestParseGoMod_ReplaceWithLocalPath(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "go.mod")
	content := `
module example.com/test
go 1.22
require github.com/foo/bar v1.0.0
replace github.com/foo/bar => ../local-bar
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	refs, err := ParseGoMod(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(refs) != 1 {
		t.Fatalf("expected 1 dep, got %d", len(refs))
	}
	r := refs[0]
	if r.LocalPath == "" {
		t.Errorf("expected LocalPath to be set for local replace, got %+v", r)
	}
	// Should be resolved relative to the go.mod's directory
	wantPrefix := filepath.Dir(dir)
	if !filepath.IsAbs(r.LocalPath) || filepath.Dir(r.LocalPath) != wantPrefix {
		t.Errorf("LocalPath should be absolute and under %s, got %s", wantPrefix, r.LocalPath)
	}
}

func TestParseGoMod_NoDirect(t *testing.T) {
	// Only indirect — should error rather than silently return empty.
	path := writeGoMod(t, `
module example.com/test
go 1.22
require github.com/foo/x v1.0.0 // indirect
`)
	_, err := ParseGoMod(path)
	if err == nil {
		t.Error("expected error for no direct deps, got nil")
	}
}

func TestParseGoMod_FileNotFound(t *testing.T) {
	_, err := ParseGoMod("/nonexistent/go.mod")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestFilterModules_Skip(t *testing.T) {
	refs := []ModuleRef{
		{Path: "a", OriginalPath: "a"},
		{Path: "b", OriginalPath: "b"},
		{Path: "c", OriginalPath: "c"},
	}
	got := FilterModules(refs, nil, []string{"b"})
	if len(got) != 2 || got[0].Path != "a" || got[1].Path != "c" {
		t.Errorf("skip failed: %+v", got)
	}
}

func TestFilterModules_Only(t *testing.T) {
	refs := []ModuleRef{
		{Path: "a", OriginalPath: "a"},
		{Path: "b", OriginalPath: "b"},
		{Path: "c", OriginalPath: "c"},
	}
	got := FilterModules(refs, []string{"b"}, nil)
	if len(got) != 1 || got[0].Path != "b" {
		t.Errorf("only failed: %+v", got)
	}
}

func TestFilterModules_OnlyAndSkip(t *testing.T) {
	refs := []ModuleRef{
		{Path: "a", OriginalPath: "a"},
		{Path: "b", OriginalPath: "b"},
		{Path: "c", OriginalPath: "c"},
	}
	// only=[a,b], skip=[a] => just b
	got := FilterModules(refs, []string{"a", "b"}, []string{"a"})
	if len(got) != 1 || got[0].Path != "b" {
		t.Errorf("only+skip failed: %+v", got)
	}
}

func TestFilterModules_MatchesOriginalPath(t *testing.T) {
	// User-facing names from go.mod (OriginalPath) should match too,
	// so users can skip "github.com/foo/bar" even if it was replaced.
	refs := []ModuleRef{
		{Path: "github.com/forked/bar", OriginalPath: "github.com/foo/bar"},
	}
	got := FilterModules(refs, nil, []string{"github.com/foo/bar"})
	if len(got) != 0 {
		t.Errorf("skip should match OriginalPath, got %+v", got)
	}
}

func TestModuleRefSpec(t *testing.T) {
	cases := []struct {
		ref  ModuleRef
		want string
	}{
		{ModuleRef{Path: "github.com/foo/bar", Version: "v1.0.0"}, "github.com/foo/bar@v1.0.0"},
		{ModuleRef{Path: "github.com/foo/bar", Version: ""}, "github.com/foo/bar"},
		{ModuleRef{Path: "github.com/foo/bar", LocalPath: "/abs/path"}, "/abs/path"},
	}
	for _, c := range cases {
		if got := c.ref.Spec(); got != c.want {
			t.Errorf("Spec(%+v) = %q, want %q", c.ref, got, c.want)
		}
	}
}
