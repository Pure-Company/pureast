// pkg/analyze/clean_test.go
package analyze

import (
	"testing"

	astpkg "github.com/Pure-Company/pureast/pkg/ast"
)

// fakeDeclMap returns a decl map seeded with the given names. The
// DeclNode itself is empty — CleanDependencies only checks for
// presence in the map, not the contents.
func fakeDeclMap(names ...string) map[string]astpkg.DeclNode {
	m := map[string]astpkg.DeclNode{}
	for _, n := range names {
		m[n] = astpkg.DeclNode{Name: n}
	}
	return m
}

func makeDepsWithFunctions(names ...string) astpkg.Dependencies {
	d := astpkg.NewDependencies()
	for _, n := range names {
		d = d.AddFunction(n)
	}
	return d
}

func functionsOf(d astpkg.Dependencies) []string {
	return d.Functions.ToSlice()
}

func TestCleanDependencies_DropsBareReceiverVar(t *testing.T) {
	deps := makeDepsWithFunctions("NewProfile", "p", "Profile.IsComplete")
	declMap := fakeDeclMap("NewProfile", "Profile.IsComplete")

	cleaned := CleanDependencies(deps, declMap)
	got := functionsOf(cleaned)

	if contains(got, "p") {
		t.Errorf("bare receiver 'p' should be filtered; got %v", got)
	}
	if !contains(got, "NewProfile") {
		t.Errorf("real dep 'NewProfile' should be kept; got %v", got)
	}
}

func TestCleanDependencies_DropsReceiverFieldAccess(t *testing.T) {
	deps := makeDepsWithFunctions("p.Address", "p.User", "NewProfile")
	declMap := fakeDeclMap("NewProfile")

	cleaned := CleanDependencies(deps, declMap)
	got := functionsOf(cleaned)

	if contains(got, "p.Address") || contains(got, "p.User") {
		t.Errorf("receiver field access should be filtered; got %v", got)
	}
}

func TestCleanDependencies_KeepsRealSingleLetterDecls(t *testing.T) {
	// If the package legitimately declares a function named "p", we
	// keep it. (Unusual but allowed.)
	deps := makeDepsWithFunctions("p", "Foo")
	declMap := fakeDeclMap("p", "Foo")

	cleaned := CleanDependencies(deps, declMap)
	got := functionsOf(cleaned)

	if !contains(got, "p") {
		t.Errorf("declared single-letter function should be kept: %v", got)
	}
}

func TestCleanDependencies_KeepsImportQualified(t *testing.T) {
	// fmt.Println is import-qualified — multi-character prefix, not noise.
	deps := makeDepsWithFunctions("fmt.Println", "json.Marshal")
	declMap := fakeDeclMap() // empty — these aren't local decls

	cleaned := CleanDependencies(deps, declMap)
	got := functionsOf(cleaned)

	if !contains(got, "fmt.Println") {
		t.Errorf("import-qualified call should be kept: %v", got)
	}
	if !contains(got, "json.Marshal") {
		t.Errorf("import-qualified call should be kept: %v", got)
	}
}

func TestCleanDependencies_PreservesOtherCategories(t *testing.T) {
	// Cleaning Functions must not lose Types/Structs/etc.
	d := astpkg.NewDependencies()
	d = d.AddFunction("p")
	d = d.AddFunction("RealFunc")
	d = d.AddType("User")
	d = d.AddStruct("Profile")
	d = d.AddInterface("Reader")
	d = d.AddImport("fmt")

	declMap := fakeDeclMap("RealFunc", "User", "Profile", "Reader")
	cleaned := CleanDependencies(d, declMap)

	if cleaned.Types.Size() != 1 {
		t.Errorf("Types lost: %v", cleaned.Types.ToSlice())
	}
	if cleaned.Structs.Size() != 1 {
		t.Errorf("Structs lost: %v", cleaned.Structs.ToSlice())
	}
	if cleaned.Interfaces.Size() != 1 {
		t.Errorf("Interfaces lost: %v", cleaned.Interfaces.ToSlice())
	}
	if cleaned.Imports.Size() != 1 {
		t.Errorf("Imports lost: %v", cleaned.Imports.ToSlice())
	}
}

func TestCleanDependencies_Empty(t *testing.T) {
	d := astpkg.NewDependencies()
	cleaned := CleanDependencies(d, fakeDeclMap())
	if !cleaned.IsEmpty() {
		t.Errorf("empty input should produce empty output")
	}
}
