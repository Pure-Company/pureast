// pkg/analyze/reverse_test.go
//
// Tests for the reverse dependency graph and bounded traversal.
// These are the queries `pureast deps --reverse` and `pureast deps
// --depth N` use, plus the corresponding MCP tools — getting them
// wrong would give the LLM bad impact-analysis data, which is
// exactly the use case we built them for.
package analyze

import (
	"go/parser"
	"go/token"
	"sort"
	"testing"

	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
)

// buildGraph parses the given source, builds a decl map, and wraps it
// in a DependencyGraph. The src must be a complete Go file.
func buildGraph(t *testing.T, src string) DependencyGraph {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	fileNode := extract.ExtractFile(file)
	pkg := astpkg.PackageNode{
		Name:  file.Name.Name,
		Files: []astpkg.FileNode{fileNode},
	}
	declMap := extract.BuildPackageDeclMap(pkg)
	return NewDependencyGraph(declMap)
}

// names returns sorted dependency-set members for stable assertions.
func sortedDeps(d astpkg.Dependencies) []string {
	all := []string{}
	all = append(all, d.Functions.ToSlice()...)
	all = append(all, d.Types.ToSlice()...)
	all = append(all, d.Structs.ToSlice()...)
	all = append(all, d.Interfaces.ToSlice()...)
	sort.Strings(all)
	return all
}

func contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

const reverseTestSrc = `package demo

type User struct {
	Name string
}

type Profile struct {
	User User
}

func NewUser(name string) User {
	return User{Name: name}
}

func NewProfile(u User) Profile {
	return Profile{User: u}
}

func (u User) Greet() string {
	return "hello " + u.Name
}
`

func TestUsers_Direct(t *testing.T) {
	g := buildGraph(t, reverseTestSrc)
	users := g.Users("User")
	got := sortedDeps(users)

	// Direct users of User: Profile (field), NewUser (returns), NewProfile (param)
	wantAny := []string{"Profile", "NewUser", "NewProfile"}
	for _, w := range wantAny {
		if !contains(got, w) {
			t.Errorf("Users(User) missing %q; got %v", w, got)
		}
	}
}

func TestUsers_NoUsers(t *testing.T) {
	src := `package demo
type Lonely struct{}
type Other struct{}
`
	g := buildGraph(t, src)
	users := g.Users("Lonely")
	got := sortedDeps(users)
	if len(got) != 0 {
		t.Errorf("Lonely should have no users, got %v", got)
	}
}

func TestUsers_NonExistentSymbol(t *testing.T) {
	g := buildGraph(t, reverseTestSrc)
	users := g.Users("DoesNotExist")
	got := sortedDeps(users)
	if len(got) != 0 {
		t.Errorf("non-existent symbol should yield empty users, got %v", got)
	}
}

func TestUsersTransitive_IncludesIndirect(t *testing.T) {
	src := `package demo
type A struct{}

type B struct {
	a A
}

type C struct {
	b B
}
`
	// A → used by B → used by C
	g := buildGraph(t, src)

	direct := sortedDeps(g.Users("A"))
	trans := sortedDeps(g.UsersTransitive("A"))

	// Direct should include B
	if !contains(direct, "B") {
		t.Errorf("Users(A) should include B, got %v", direct)
	}

	// Transitive should include both B and C
	if !contains(trans, "B") {
		t.Errorf("UsersTransitive(A) missing B, got %v", trans)
	}
	if !contains(trans, "C") {
		t.Errorf("UsersTransitive(A) missing C, got %v", trans)
	}

	// Transitive should be a superset of direct
	if len(trans) < len(direct) {
		t.Errorf("transitive smaller than direct: trans=%v direct=%v", trans, direct)
	}
}

func TestResolveBounded_Depth0(t *testing.T) {
	g := buildGraph(t, reverseTestSrc)

	d0 := g.ResolveBounded("Profile", 0)
	got := sortedDeps(d0)

	// Depth 0 = only direct dependencies. Profile directly references User.
	if !contains(got, "User") {
		t.Errorf("depth 0 should include User, got %v", got)
	}
}

func TestResolveBounded_DepthEqualsTransitive(t *testing.T) {
	// At sufficient depth, ResolveBounded should match ResolveTransitive
	g := buildGraph(t, reverseTestSrc)

	transitive := sortedDeps(g.ResolveTransitive("Profile"))
	bounded := sortedDeps(g.ResolveBounded("Profile", 100))

	if len(transitive) != len(bounded) {
		t.Errorf("at depth=100, bounded should equal transitive\n  transitive: %v\n  bounded:    %v",
			transitive, bounded)
	}
	for i := range transitive {
		if i >= len(bounded) || transitive[i] != bounded[i] {
			t.Errorf("mismatch at %d: transitive=%v bounded=%v", i, transitive, bounded)
			break
		}
	}
}

func TestUsers_Deterministic(t *testing.T) {
	// Property: identical input produces byte-identical output.
	// Required for prompt caching to actually hit.
	g := buildGraph(t, reverseTestSrc)

	a := sortedDeps(g.Users("User"))
	b := sortedDeps(g.Users("User"))

	if len(a) != len(b) {
		t.Fatalf("non-deterministic length: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			t.Errorf("non-deterministic at %d: %q vs %q", i, a[i], b[i])
		}
	}
}
