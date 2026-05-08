// pkg/scaffold/manifest_test.go
package scaffold

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeManifest(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "pureast.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestLoadManifest_Minimal(t *testing.T) {
	path := writeManifest(t, `
packages:
  - path: internal/repo
    package: repo
    files:
      - output: user.go
        task: "Define User struct and UserRepo interface"
`)
	m, err := LoadManifest(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Packages) != 1 {
		t.Fatalf("want 1 package, got %d", len(m.Packages))
	}
	pkg := m.Packages[0]
	if pkg.Path != "internal/repo" || pkg.Package != "repo" {
		t.Errorf("unexpected package: %+v", pkg)
	}
	if len(pkg.Files) != 1 || pkg.Files[0].Output != "user.go" {
		t.Errorf("unexpected files: %+v", pkg.Files)
	}
}

func TestLoadManifest_MultiSourceFile(t *testing.T) {
	path := writeManifest(t, `
packages:
  - path: internal/cache
    package: cache
    files:
      - output: user_cache.go
        task: |
          Multi-line task description that spans
          several lines.
        sources:
          - pkg: ../repo
          - module: github.com/redis/go-redis/v9
        kind: interface
`)
	m, err := LoadManifest(path)
	if err != nil {
		t.Fatal(err)
	}
	f := m.Packages[0].Files[0]
	if !strings.Contains(f.Task, "Multi-line") || !strings.Contains(f.Task, "several lines") {
		t.Errorf("multi-line task not preserved: %q", f.Task)
	}
	if len(f.Sources) != 2 {
		t.Fatalf("want 2 sources, got %d", len(f.Sources))
	}
	if f.Sources[0].Pkg != "../repo" {
		t.Errorf("first source should be pkg ../repo, got %+v", f.Sources[0])
	}
	if f.Sources[1].Module != "github.com/redis/go-redis/v9" {
		t.Errorf("second source should be redis module, got %+v", f.Sources[1])
	}
	if f.Kind != "interface" {
		t.Errorf("kind = %q, want interface", f.Kind)
	}
}

func TestValidate_RejectsCommonMistakes(t *testing.T) {
	cases := []struct {
		name    string
		yaml    string
		wantErr string
	}{
		{
			"no packages",
			`packages: []`,
			"no packages",
		},
		{
			"missing path",
			`
packages:
  - package: repo
    files:
      - output: x.go
        task: "x"
`,
			"path is required",
		},
		{
			"absolute path",
			`
packages:
  - path: /etc/passwd
    package: repo
    files:
      - output: x.go
        task: "x"
`,
			"must be relative",
		},
		{
			"escaping path",
			`
packages:
  - path: ../../../etc
    package: repo
    files:
      - output: x.go
        task: "x"
`,
			"must be relative",
		},
		{
			"duplicate package path",
			`
packages:
  - path: a
    package: a
    files:
      - output: x.go
        task: "x"
  - path: a
    package: a
    files:
      - output: y.go
        task: "y"
`,
			"duplicate path",
		},
		{
			"output is a path not a filename",
			`
packages:
  - path: a
    package: a
    files:
      - output: sub/x.go
        task: "x"
`,
			"output must be a filename",
		},
		{
			"output is gen.go",
			`
packages:
  - path: a
    package: a
    files:
      - output: gen.go
        task: "x"
`,
			"reserved",
		},
		{
			"duplicate output in same package",
			`
packages:
  - path: a
    package: a
    files:
      - output: x.go
        task: "x"
      - output: x.go
        task: "y"
`,
			"duplicate output",
		},
		{
			"missing task",
			`
packages:
  - path: a
    package: a
    files:
      - output: x.go
`,
			"task is required",
		},
		{
			"invalid kind",
			`
packages:
  - path: a
    package: a
    files:
      - output: x.go
        task: "x"
        kind: notakind
`,
			"invalid kind",
		},
		{
			"invalid package identifier",
			`
packages:
  - path: a
    package: "1invalid"
    files:
      - output: x.go
        task: "x"
`,
			"not a valid Go identifier",
		},
		{
			"empty source",
			`
packages:
  - path: a
    package: a
    files:
      - output: x.go
        task: "x"
        sources:
          - {}
`,
			"must set exactly one",
		},
		{
			"two sources in one entry",
			`
packages:
  - path: a
    package: a
    files:
      - output: x.go
        task: "x"
        sources:
          - pkg: ../b
            module: github.com/x/y
`,
			"set exactly one",
		},
		{
			"symbol without colon",
			`
packages:
  - path: a
    package: a
    files:
      - output: x.go
        task: "x"
        sources:
          - symbol: NoColonHere
`,
			"NAME:LOC",
		},
		{
			"unknown field (KnownFields strict)",
			`
packages:
  - path: a
    pacakge: a
    files:
      - output: x.go
        task: "x"
`,
			"field pacakge not found",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			path := writeManifest(t, c.yaml)
			_, err := LoadManifest(path)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !strings.Contains(err.Error(), c.wantErr) {
				t.Errorf("error %q does not contain %q", err.Error(), c.wantErr)
			}
		})
	}
}

func TestValidate_AcceptsExportedOnlyFalse(t *testing.T) {
	// exported_only: false should be respected (pointer field
	// distinguishes unset from explicit false).
	path := writeManifest(t, `
packages:
  - path: a
    package: a
    files:
      - output: x.go
        task: "x"
        exported_only: false
`)
	m, err := LoadManifest(path)
	if err != nil {
		t.Fatal(err)
	}
	f := m.Packages[0].Files[0]
	if f.ExportedOnly == nil {
		t.Fatal("expected ExportedOnly to be set, got nil pointer")
	}
	if *f.ExportedOnly {
		t.Error("expected ExportedOnly == false, got true")
	}
}

func TestSortedPackages(t *testing.T) {
	// Even though the manifest lists "z" first, SortedPackages should
	// return them in path-alphabetical order so scaffold output is
	// deterministic regardless of authoring order.
	path := writeManifest(t, `
packages:
  - path: z
    package: z
    files:
      - {output: z.go, task: z}
  - path: a
    package: a
    files:
      - {output: a.go, task: a}
  - path: m
    package: m
    files:
      - {output: m.go, task: m}
`)
	m, _ := LoadManifest(path)
	sorted := m.SortedPackages()
	want := []string{"a", "m", "z"}
	for i, p := range sorted {
		if p.Path != want[i] {
			t.Errorf("sorted[%d].Path = %q, want %q", i, p.Path, want[i])
		}
	}
}
