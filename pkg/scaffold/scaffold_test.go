// pkg/scaffold/scaffold_test.go
package scaffold

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScaffold_CreatesGenFiles(t *testing.T) {
	root := t.TempDir()
	m := &Manifest{
		Packages: []Package{
			{
				Path:    "internal/repo",
				Package: "repo",
				Files: []File{
					{Output: "user.go", Task: "Define User and UserRepo"},
				},
			},
			{
				Path:    "internal/cache",
				Package: "cache",
				Files: []File{
					{
						Output: "user_cache.go",
						Task:   "Implement Redis cache for UserRepo",
						Sources: []Source{
							{Pkg: "../repo"},
							{Module: "github.com/redis/go-redis/v9"},
						},
						Kind:  "interface",
						Model: "opus",
					},
				},
			},
		},
	}

	res, err := Scaffold(m, root)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Created) != 2 {
		t.Errorf("want 2 created, got %d: %v", len(res.Created), res.Created)
	}
	if len(res.Skipped) != 0 || len(res.Updated) != 0 {
		t.Errorf("first run should have no skipped/updated, got %+v", res)
	}

	// Check both gen.go files exist and have the expected structure.
	repoGen := filepath.Join(root, "internal", "repo", "gen.go")
	checkGenFile(t, repoGen, []string{
		"//go:build ignore",
		"//go:generate pureast claude-edit",
		`--task "Define User and UserRepo"`,
		`--output user.go`,
		"package repo",
	})

	cacheGen := filepath.Join(root, "internal", "cache", "gen.go")
	checkGenFile(t, cacheGen, []string{
		"//go:build ignore",
		"--model opus",
		`--task "Implement Redis cache for UserRepo"`,
		"--pkg ../repo",
		"--module github.com/redis/go-redis/v9",
		"--kind interface",
		"--output user_cache.go",
		"package cache",
	})
}

func TestScaffold_Idempotent(t *testing.T) {
	root := t.TempDir()
	m := &Manifest{
		Packages: []Package{
			{Path: "a", Package: "a", Files: []File{{Output: "x.go", Task: "x"}}},
		},
	}

	// First run creates.
	res1, err := Scaffold(m, root)
	if err != nil {
		t.Fatal(err)
	}
	if len(res1.Created) != 1 {
		t.Fatalf("first run: want 1 created, got %v", res1)
	}

	// Capture mtime so we can verify the second run doesn't touch the file.
	genPath := filepath.Join(root, "a", "gen.go")
	st1, _ := os.Stat(genPath)

	// Second run: identical manifest -> skipped, file untouched.
	res2, err := Scaffold(m, root)
	if err != nil {
		t.Fatal(err)
	}
	if len(res2.Skipped) != 1 {
		t.Errorf("second run: want 1 skipped, got %+v", res2)
	}
	if len(res2.Created) != 0 || len(res2.Updated) != 0 {
		t.Errorf("second run: want no creates/updates, got %+v", res2)
	}
	st2, _ := os.Stat(genPath)
	if !st1.ModTime().Equal(st2.ModTime()) {
		t.Errorf("idempotent run modified file: mtime %v -> %v", st1.ModTime(), st2.ModTime())
	}
}

func TestScaffold_DetectsUpdate(t *testing.T) {
	root := t.TempDir()
	m1 := &Manifest{
		Packages: []Package{
			{Path: "a", Package: "a", Files: []File{{Output: "x.go", Task: "original"}}},
		},
	}
	if _, err := Scaffold(m1, root); err != nil {
		t.Fatal(err)
	}

	// Change the task -> directive changes -> file should be rewritten.
	m2 := &Manifest{
		Packages: []Package{
			{Path: "a", Package: "a", Files: []File{{Output: "x.go", Task: "modified"}}},
		},
	}
	res, err := Scaffold(m2, root)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Updated) != 1 {
		t.Errorf("want 1 updated, got %+v", res)
	}

	// Verify new content is on disk.
	data, _ := os.ReadFile(filepath.Join(root, "a", "gen.go"))
	if !strings.Contains(string(data), "modified") {
		t.Errorf("update did not persist; got:\n%s", data)
	}
	if strings.Contains(string(data), "original") {
		t.Errorf("old content leaked through; got:\n%s", data)
	}
}

func TestScaffold_MultilineTaskBecomesSingleLine(t *testing.T) {
	// Multi-line task in YAML must collapse to one line in the
	// directive — Go's //go:generate parser doesn't honor backslash
	// continuation. Verify the rendered directive is one line and
	// preserves the words.
	root := t.TempDir()
	m := &Manifest{
		Packages: []Package{
			{
				Path:    "a",
				Package: "a",
				Files: []File{
					{
						Output: "x.go",
						Task:   "Line one.\nLine two.\nLine three.",
					},
				},
			},
		},
	}
	if _, err := Scaffold(m, root); err != nil {
		t.Fatal(err)
	}

	data, _ := os.ReadFile(filepath.Join(root, "a", "gen.go"))
	lines := strings.Split(string(data), "\n")

	var directiveCount int
	for _, ln := range lines {
		if strings.HasPrefix(ln, "//go:generate ") {
			directiveCount++
			if strings.Contains(ln, "\n") {
				t.Errorf("directive contains embedded newline: %q", ln)
			}
			if !strings.Contains(ln, "Line one.") || !strings.Contains(ln, "Line three.") {
				t.Errorf("directive missing task content: %q", ln)
			}
		}
	}
	if directiveCount != 1 {
		t.Errorf("want 1 directive, got %d", directiveCount)
	}
}

func TestScaffold_PackageDocCommentRendered(t *testing.T) {
	root := t.TempDir()
	m := &Manifest{
		Packages: []Package{
			{
				Path:    "a",
				Package: "a",
				Doc:     "Package a does the thing.\nIt also does another thing.",
				Files:   []File{{Output: "x.go", Task: "x"}},
			},
		},
	}
	if _, err := Scaffold(m, root); err != nil {
		t.Fatal(err)
	}
	data, _ := os.ReadFile(filepath.Join(root, "a", "gen.go"))
	s := string(data)
	if !strings.Contains(s, "// Package a does the thing.") {
		t.Errorf("doc line 1 missing; got:\n%s", s)
	}
	if !strings.Contains(s, "// It also does another thing.") {
		t.Errorf("doc line 2 missing; got:\n%s", s)
	}
}

func TestScaffold_RejectsMissingProjectRoot(t *testing.T) {
	m := &Manifest{Packages: []Package{{Path: "a", Package: "a", Files: []File{{Output: "x.go", Task: "x"}}}}}
	_, err := Scaffold(m, "/no/such/dir/anywhere/please")
	if err == nil {
		t.Error("expected error for missing root, got nil")
	}
}

func TestRenderDirective_QuotesSpaces(t *testing.T) {
	// A task with spaces must become a quoted argument so go's
	// directive parser keeps it as one token.
	d := renderDirective(File{
		Output: "x.go",
		Task:   "implement the thing with spaces",
	})
	if !strings.Contains(d, `--task "implement the thing with spaces"`) {
		t.Errorf("task not quoted: %q", d)
	}
}

func TestRenderDirective_KindAllOmitted(t *testing.T) {
	// kind: all is the default — don't emit it in the directive.
	d := renderDirective(File{Output: "x.go", Task: "t", Kind: "all"})
	if strings.Contains(d, "--kind") {
		t.Errorf("--kind all should be omitted; got: %q", d)
	}
}

// checkGenFile asserts that all of `must` substrings appear in the
// file at `path`. Useful for asserting the rendered shape without
// pinning every byte.
func checkGenFile(t *testing.T, path string, must []string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	for _, m := range must {
		if !strings.Contains(string(data), m) {
			t.Errorf("file %s missing substring %q\nfull content:\n%s", path, m, data)
		}
	}
}
