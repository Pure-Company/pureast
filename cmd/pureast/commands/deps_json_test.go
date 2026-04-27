// cmd/pureast/commands/deps_json_test.go
//
// Determinism test for JSON dependency output. The product story for
// pureast as LLM context relies on byte-identical output across runs
// — that's what makes prompt caching hit. If formatDepsJSON ever
// returns different bytes for the same input (map iteration order
// leaking through, time-dependent fields, etc.), caching breaks
// silently. This test catches that.
package commands

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/Pure-Company/pureast/pkg/analyze"
	astpkg "github.com/Pure-Company/pureast/pkg/ast"
	"github.com/Pure-Company/pureast/pkg/extract"
)

const jsonTestSrc = `package demo

type User struct {
	Name string
}

type Profile struct {
	User    User
	Address string
}

func NewProfile(u User, a string) Profile {
	return Profile{User: u, Address: a}
}
`

func setupGraph(t *testing.T) (analyze.DependencyGraph, *token.FileSet, map[string]astpkg.DeclNode) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", jsonTestSrc, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	pkg := astpkg.PackageNode{
		Name:  file.Name.Name,
		Files: []astpkg.FileNode{extract.ExtractFile(file)},
	}
	declMap := extract.BuildPackageDeclMap(pkg)
	return analyze.NewDependencyGraph(declMap), fset, declMap
}

func TestFormatDepsJSON_Deterministic_ForwardDeps(t *testing.T) {
	g, fset, declMap := setupGraph(t)
	deps := g.ResolveTransitive("Profile")

	a := formatDepsJSON("Profile", deps, false, fset, declMap, ".")
	b := formatDepsJSON("Profile", deps, false, fset, declMap, ".")

	if a != b {
		t.Errorf("non-deterministic JSON output:\n--- A ---\n%s\n--- B ---\n%s", a, b)
	}
}

func TestFormatDepsJSON_Deterministic_ReverseDeps(t *testing.T) {
	g, fset, declMap := setupGraph(t)
	users := g.Users("User")

	a := formatDepsJSON("User", users, true, fset, declMap, ".")
	b := formatDepsJSON("User", users, true, fset, declMap, ".")

	if a != b {
		t.Errorf("non-deterministic reverse JSON:\n--- A ---\n%s\n--- B ---\n%s", a, b)
	}
}

func TestFormatDepsJSON_ValidJSON(t *testing.T) {
	g, fset, declMap := setupGraph(t)
	deps := g.ResolveTransitive("Profile")

	out := formatDepsJSON("Profile", deps, false, fset, declMap, ".")

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v\noutput:\n%s", err, out)
	}
	if parsed["symbol"] != "Profile" {
		t.Errorf("expected symbol=Profile, got %v", parsed["symbol"])
	}
}

func TestFormatDepsJSON_ReverseShape(t *testing.T) {
	// With locations, function entries should be objects with name+file+line,
	// not plain strings.
	g, fset, declMap := setupGraph(t)
	users := g.Users("User")

	out := formatDepsJSON("User", users, true, fset, declMap, ".")

	var parsed struct {
		Symbol    string `json:"symbol"`
		Functions []struct {
			Name string `json:"name"`
			File string `json:"file"`
			Line int    `json:"line"`
		} `json:"functions"`
	}
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("reverse JSON failed to parse with object shape: %v\noutput:\n%s", err, out)
	}
	if parsed.Symbol != "User" {
		t.Errorf("symbol field wrong")
	}
	if len(parsed.Functions) == 0 {
		t.Fatalf("expected at least one function user of User")
	}
	for _, f := range parsed.Functions {
		if f.Name == "" {
			t.Errorf("function entry missing name field: %+v", f)
		}
	}
}

func TestFormatDepsJSON_ForwardShape(t *testing.T) {
	// Without locations, dependency entries should be plain strings —
	// preserved for backward-compat with existing parsers. We assert
	// against `types` here rather than `functions` because Profile's
	// forward deps in our fixture are User (a type), not any function.
	g, fset, declMap := setupGraph(t)
	deps := g.ResolveTransitive("Profile")

	out := formatDepsJSON("Profile", deps, false, fset, declMap, ".")

	var parsed struct {
		Types []string `json:"types"`
	}
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("forward JSON failed string-array shape: %v\noutput:\n%s", err, out)
	}

	if len(parsed.Types) == 0 {
		t.Errorf("expected types in forward deps of Profile, got empty")
	}
	// Profile references User
	found := false
	for _, name := range parsed.Types {
		if name == "User" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected User in Profile's forward deps, got %v", parsed.Types)
	}
}

func TestFormatDepsJSON_SortedAlphabetically(t *testing.T) {
	g, fset, declMap := setupGraph(t)
	deps := g.ResolveTransitive("Profile")

	out := formatDepsJSON("Profile", deps, false, fset, declMap, ".")

	var parsed struct {
		Functions []string `json:"functions"`
		Types     []string `json:"types"`
		Structs   []string `json:"structs"`
	}
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatal(err)
	}

	for _, list := range [][]string{parsed.Functions, parsed.Types, parsed.Structs} {
		for i := 1; i < len(list); i++ {
			if list[i-1] > list[i] {
				t.Errorf("not sorted: %v", list)
				break
			}
		}
	}

	// Also confirm output contains expected sorted ordering visually
	if !strings.Contains(out, "\"functions\"") {
		t.Errorf("output missing functions field: %s", out)
	}
}

// TestFormatDepsJSON_ForwardWithLocations covers the --locations flag
// for forward deps. Same rich shape as --reverse output: each entry
// becomes {name, file, line}.
func TestFormatDepsJSON_ForwardWithLocations(t *testing.T) {
	g, fset, declMap := setupGraph(t)
	deps := g.ResolveTransitive("Profile")

	out := formatDepsJSON("Profile", deps, true, fset, declMap, ".")

	var parsed struct {
		Symbol string `json:"symbol"`
		Types  []struct {
			Name string `json:"name"`
			File string `json:"file"`
			Line int    `json:"line"`
		} `json:"types"`
	}
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("forward+locations should produce object shape: %v\noutput:\n%s", err, out)
	}

	if len(parsed.Types) == 0 {
		t.Fatalf("expected types in Profile's deps")
	}
	// At least one entry should have a non-empty file (the in-package types do)
	hasFile := false
	for _, e := range parsed.Types {
		if e.Name == "" {
			t.Errorf("missing name: %+v", e)
		}
		if e.File != "" {
			hasFile = true
		}
	}
	if !hasFile {
		t.Errorf("no file populated; locations didn't kick in for forward deps")
	}
}
