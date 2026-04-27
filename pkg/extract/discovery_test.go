// pkg/extract/discovery_test.go
//
// End-to-end tests for symbol discovery and kind inference. We feed
// real Go source through the parser and assert the resulting kinds,
// because the bug we fixed earlier (struct mis-classified as "type"
// when its Deps didn't contain itself) only surfaces against actual
// AST output — synthetic DeclNodes can't reproduce it.
package extract

import (
	"go/parser"
	"go/token"
	"sort"
	"testing"

	astpkg "github.com/Pure-Company/pureast/pkg/ast"
)

// parseToPackage is a small helper that turns a Go source string into
// the PackageNode shape that DiscoverAllSymbols expects. Keeps each
// test focused on what it actually asserts.
func parseToPackage(t *testing.T, src string) astpkg.PackageNode {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	fileNode := ExtractFile(file)
	return astpkg.PackageNode{
		Name:  file.Name.Name,
		Files: []astpkg.FileNode{fileNode},
	}
}

func TestDiscover_KindClassification(t *testing.T) {
	src := `package demo

// User is a struct.
type User struct {
	Name string
}

// Reader is an interface.
type Reader interface {
	Read() string
}

// Alias is a type alias.
type Alias = string

// Named is a named type (not struct/interface).
type Named int

// Hello is a function.
func Hello() string { return "hi" }

// Greet is a method on User.
func (u User) Greet() string { return "hello " + u.Name }

// Pi is a constant.
const Pi = 3.14

// Counter is a variable.
var Counter int
`
	pkg := parseToPackage(t, src)
	syms := DiscoverAllSymbols(pkg)

	// Build a name → kind map for assertion. Methods come back with the
	// receiver stripped from the name, but stored in Receiver.
	got := map[string]string{}
	for _, s := range syms {
		key := s.Name
		if s.Receiver != "" {
			key = s.Receiver + "." + s.Name
		}
		got[key] = s.Kind
	}

	want := map[string]string{
		"User":       "struct",
		"Reader":     "interface",
		"Hello":      "function",
		"User.Greet": "method",
		"Pi":         "const",
		"Counter":    "var",
	}

	for name, wantKind := range want {
		if got[name] != wantKind {
			t.Errorf("symbol %q: got kind=%q, want %q", name, got[name], wantKind)
		}
	}
}

// TestDiscover_StructWithoutSelfRef is the explicit regression test
// for the bug we fixed: a struct whose Deps doesn't contain itself
// used to be classified as "type" instead of "struct".
func TestDiscover_StructWithoutSelfRef(t *testing.T) {
	src := `package demo

// Address has no recursion or self-reference.
type Address struct {
	Street string
	City   string
}
`
	pkg := parseToPackage(t, src)
	syms := DiscoverAllSymbols(pkg)

	if len(syms) == 0 {
		t.Fatal("no symbols discovered")
	}

	for _, s := range syms {
		if s.Name == "Address" {
			if s.Kind != "struct" {
				t.Errorf("regression: Address classified as %q, want struct", s.Kind)
			}
			return
		}
	}
	t.Fatal("Address symbol not found")
}

func TestDiscover_FilterByKind(t *testing.T) {
	src := `package demo
type A struct { X int }
type B interface { M() }
type C int
func F() {}
`
	pkg := parseToPackage(t, src)
	all := DiscoverAllSymbols(pkg)

	structs := FilterByKind("struct", all)
	if len(structs) != 1 || structs[0].Name != "A" {
		t.Errorf("FilterByKind(struct): got %v", structs)
	}

	ifaces := FilterByKind("interface", all)
	if len(ifaces) != 1 || ifaces[0].Name != "B" {
		t.Errorf("FilterByKind(interface): got %v", ifaces)
	}

	funcs := FilterByKind("function", all)
	if len(funcs) != 1 || funcs[0].Name != "F" {
		t.Errorf("FilterByKind(function): got %v", funcs)
	}
}

// TestDiscover_StableOrder asserts byte-identical output across runs.
// Critical for prompt caching.
func TestDiscover_StableOrder(t *testing.T) {
	src := `package demo
type Z struct{}
type A struct{}
type M struct{}
func Foo() {}
func Bar() {}
`
	pkg := parseToPackage(t, src)

	a := DiscoverAllSymbols(pkg)
	b := DiscoverAllSymbols(pkg)

	if len(a) != len(b) {
		t.Fatalf("non-stable count: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			t.Errorf("non-stable at %d: %+v vs %+v", i, a[i], b[i])
		}
	}

	// Also assert names sorted: that's the documented contract.
	names := make([]string, len(a))
	for i, s := range a {
		names[i] = s.Name
	}
	sorted := make([]string, len(names))
	copy(sorted, names)
	sort.Strings(sorted)
	for i := range names {
		if names[i] != sorted[i] {
			t.Errorf("not sorted: %v", names)
			break
		}
	}
}
