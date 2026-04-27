// pkg/extract/render_test.go
package extract

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"

	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
)

// renderTestPkg parses a snippet and returns symbols + fset suitable
// for exercising the renderers.
func renderTestPkg(t *testing.T, src string) ([]SymbolInfo, *token.FileSet) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	pkg := astpkg.PackageNode{
		Name:  file.Name.Name,
		Files: []astpkg.FileNode{ExtractFile(file)},
	}
	return DiscoverAllSymbols(pkg), fset
}

func findSymbol(syms []SymbolInfo, name string) (SymbolInfo, bool) {
	for _, s := range syms {
		if s.Name == name {
			return s, true
		}
	}
	return SymbolInfo{}, false
}

func TestRenderSignature_Func(t *testing.T) {
	src := `package demo
func Add(a, b int) int { return a + b }
`
	syms, fset := renderTestPkg(t, src)
	add, ok := findSymbol(syms, "Add")
	if !ok {
		t.Fatal("Add not discovered")
	}
	got := RenderSignature(fset, add)
	if !strings.Contains(got, "func Add") {
		t.Errorf("missing 'func Add': %q", got)
	}
	if !strings.Contains(got, "int") {
		t.Errorf("missing return type: %q", got)
	}
	if strings.Contains(got, "return") {
		t.Errorf("signature mode should not include body: %q", got)
	}
}

func TestRenderSignature_Method(t *testing.T) {
	src := `package demo
type T struct{}
func (t T) Greet(name string) string { return "hi " + name }
`
	syms, fset := renderTestPkg(t, src)
	method, ok := findSymbol(syms, "Greet")
	if !ok {
		t.Fatal("Greet not discovered")
	}
	got := RenderSignature(fset, method)
	if !strings.Contains(got, "Greet") {
		t.Errorf("missing method name: %q", got)
	}
	if !strings.Contains(got, "(t T)") {
		t.Errorf("missing receiver: %q", got)
	}
}

func TestRenderSignature_Struct(t *testing.T) {
	src := `package demo
type User struct {
	Name string
	Age  int
}
`
	syms, fset := renderTestPkg(t, src)
	user, ok := findSymbol(syms, "User")
	if !ok {
		t.Fatal("User not discovered")
	}
	got := RenderSignature(fset, user)
	if !strings.HasPrefix(got, "type User struct") {
		t.Errorf("missing type-struct prefix: %q", got)
	}
	if !strings.Contains(got, "Name string") {
		t.Errorf("missing field: %q", got)
	}
}

func TestRenderSignature_Interface(t *testing.T) {
	src := `package demo
type Reader interface {
	Read(p []byte) (int, error)
	Close() error
}
`
	syms, fset := renderTestPkg(t, src)
	r, ok := findSymbol(syms, "Reader")
	if !ok {
		t.Fatal("Reader not discovered")
	}
	got := RenderSignature(fset, r)
	if !strings.HasPrefix(got, "type Reader interface") {
		t.Errorf("missing interface header: %q", got)
	}
	if !strings.Contains(got, "Read") {
		t.Errorf("missing Read method: %q", got)
	}
}

func TestRenderWithBody_FuncIncludesBody(t *testing.T) {
	src := `package demo
func Compute(x int) int {
	return x * 2
}
`
	syms, fset := renderTestPkg(t, src)
	c, ok := findSymbol(syms, "Compute")
	if !ok {
		t.Fatal("Compute not found")
	}
	got := RenderWithBody(fset, c)
	if !strings.Contains(got, "return x * 2") {
		t.Errorf("body missing: %q", got)
	}
}

func TestRenderWithBody_TypeFallsBackToSignature(t *testing.T) {
	src := `package demo
type Marker struct{}
`
	syms, fset := renderTestPkg(t, src)
	m, _ := findSymbol(syms, "Marker")
	got := RenderWithBody(fset, m)
	// Same as RenderSignature — types don't have a separate body
	want := RenderSignature(fset, m)
	if got != want {
		t.Errorf("RenderWithBody and RenderSignature differ for type: %q vs %q", got, want)
	}
}

func TestRenderSignature_NilDecl(t *testing.T) {
	got := RenderSignature(token.NewFileSet(), SymbolInfo{Name: "x", Kind: "function"})
	if got != "" {
		t.Errorf("nil Decl should yield empty string, got %q", got)
	}
}

func TestSymbolDoc_Func(t *testing.T) {
	src := `package demo
// Add returns the sum of two integers.
func Add(a, b int) int { return a + b }
`
	syms, _ := renderTestPkg(t, src)
	add, _ := findSymbol(syms, "Add")
	got := SymbolDoc(add.Decl)
	if !strings.Contains(got, "sum of two integers") {
		t.Errorf("doc not extracted: %q", got)
	}
}

func TestSymbolDoc_TypeViaSpec(t *testing.T) {
	src := `package demo

// User is the application's user type.
type User struct{ ID int }
`
	syms, _ := renderTestPkg(t, src)
	user, _ := findSymbol(syms, "User")
	got := SymbolDoc(user.Decl)
	if !strings.Contains(got, "user type") {
		t.Errorf("type doc not extracted: %q", got)
	}
}

func TestGroupSymbolsForDump_OrderingPreservation(t *testing.T) {
	src := `package demo
type Z struct{}
type A struct{}
func GoZ() {}
func GoA() {}
`
	syms, _ := renderTestPkg(t, src)
	groups := GroupSymbolsForDump(syms)

	// Within each kind, alphabetical
	structs := groups["struct"]
	if len(structs) != 2 || structs[0].Name != "A" || structs[1].Name != "Z" {
		t.Errorf("structs not sorted: %v", structs)
	}
	funcs := groups["function"]
	if len(funcs) != 2 || funcs[0].Name != "GoA" || funcs[1].Name != "GoZ" {
		t.Errorf("functions not sorted: %v", funcs)
	}
}

func TestGroupSymbolsForDump_KindOrderCovers(t *testing.T) {
	// KindOrder should contain every kind that DiscoverAllSymbols
	// might produce. If a future kind is added without adding to
	// KindOrder/KindHeadings, this catches it.
	expected := []string{"struct", "interface", "type", "function", "method", "const", "var"}
	if len(KindOrder) < len(expected) {
		t.Errorf("KindOrder shrunk: %v", KindOrder)
	}
	for _, e := range expected {
		found := false
		for _, k := range KindOrder {
			if k == e {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("kind %q missing from KindOrder", e)
		}
		if _, ok := KindHeadings[e]; !ok {
			t.Errorf("kind %q missing from KindHeadings", e)
		}
	}
}
