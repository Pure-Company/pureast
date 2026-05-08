// pkg/weave/dag_test.go
package weave

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/Pure-Company/pureast/pkg/scaffold"
)

// helper: build a Manifest by concise spec.
func buildManifest(specs ...packageSpec) *scaffold.Manifest {
	m := &scaffold.Manifest{}
	for _, ps := range specs {
		pkg := scaffold.Package{
			Path:    ps.path,
			Package: ps.pkg,
		}
		for _, fs := range ps.files {
			f := scaffold.File{Output: fs.output, Task: "t"}
			for _, src := range fs.sources {
				f.Sources = append(f.Sources, src)
			}
			pkg.Files = append(pkg.Files, f)
		}
		m.Packages = append(m.Packages, pkg)
	}
	return m
}

type packageSpec struct {
	path  string
	pkg   string
	files []fileSpec
}

type fileSpec struct {
	output  string
	sources []scaffold.Source
}

func levelIDs(level []Node) []string {
	ids := make([]string, len(level))
	for i, n := range level {
		ids[i] = n.ID()
	}
	return ids
}

// TestExample1: pkg sources -> internal edges, module source -> no edge.
//
//	internal/domain/link.go      (no sources)        — level 0
//	internal/domain/errors.go    (no sources)        — level 0
//	internal/repo/link_repo.go   (pkg ../domain)     — level 1
//	internal/cache/link_cache.go (pkg ../repo +
//	                              pkg ../domain +
//	                              module redis)      — level 2
func TestBuildDAG_Example1_PkgChain(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "internal/domain", pkg: "domain", files: []fileSpec{
			{output: "link.go"},
			{output: "errors.go"},
		}},
		packageSpec{path: "internal/repo", pkg: "repo", files: []fileSpec{
			{output: "link_repo.go", sources: []scaffold.Source{{Pkg: "../domain"}}},
		}},
		packageSpec{path: "internal/cache", pkg: "cache", files: []fileSpec{
			{output: "link_cache.go", sources: []scaffold.Source{
				{Pkg: "../repo"},
				{Pkg: "../domain"},
				{Module: "github.com/redis/go-redis/v9"},
			}},
		}},
	)

	g, err := BuildDAG(m)
	if err != nil {
		t.Fatal(err)
	}

	levels := g.Levels()
	if len(levels) != 3 {
		t.Fatalf("want 3 levels, got %d: %v", len(levels), levelsAsIDs(levels))
	}

	wantL0 := []string{"internal/domain/errors.go", "internal/domain/link.go"}
	if got := levelIDs(levels[0]); !reflect.DeepEqual(got, wantL0) {
		t.Errorf("level 0: want %v, got %v", wantL0, got)
	}

	wantL1 := []string{"internal/repo/link_repo.go"}
	if got := levelIDs(levels[1]); !reflect.DeepEqual(got, wantL1) {
		t.Errorf("level 1: want %v, got %v", wantL1, got)
	}

	wantL2 := []string{"internal/cache/link_cache.go"}
	if got := levelIDs(levels[2]); !reflect.DeepEqual(got, wantL2) {
		t.Errorf("level 2: want %v, got %v", wantL2, got)
	}
}

// TestExample2: only module sources -> directive sits in level 0,
// independent of everything else.
func TestBuildDAG_Example2_OnlyModuleSources(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "internal/domain", pkg: "domain", files: []fileSpec{
			{output: "user.go"},
		}},
		packageSpec{path: "internal/auth", pkg: "auth", files: []fileSpec{
			{output: "middleware.go", sources: []scaffold.Source{
				{Module: "github.com/golang-jwt/jwt/v5"},
				{Module: "net/http"},
			}},
		}},
	)

	g, err := BuildDAG(m)
	if err != nil {
		t.Fatal(err)
	}

	levels := g.Levels()
	if len(levels) != 1 {
		t.Fatalf("want 1 level (everything independent), got %d: %v",
			len(levels), levelsAsIDs(levels))
	}

	wantL0 := []string{"internal/auth/middleware.go", "internal/domain/user.go"}
	if got := levelIDs(levels[0]); !reflect.DeepEqual(got, wantL0) {
		t.Errorf("level 0: want %v, got %v", wantL0, got)
	}

	// Verify the auth middleware has zero predecessors (modules don't gate).
	authID := "internal/auth/middleware.go"
	if n := len(g.Predecessors[authID]); n != 0 {
		t.Errorf("middleware.go should have 0 predecessors, has %d", n)
	}
}

// TestExample3: gomod source -> directive depends on every .go output
// in the manifest. Build files run last.
func TestBuildDAG_Example3_GomodGatesOnAllGoFiles(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "internal/domain", pkg: "domain", files: []fileSpec{
			{output: "user.go"},
		}},
		packageSpec{path: "internal/repo", pkg: "repo", files: []fileSpec{
			{output: "user_repo.go", sources: []scaffold.Source{{Pkg: "../domain"}}},
		}},
		packageSpec{path: "tools/gen", pkg: "gen", files: []fileSpec{
			{output: "../../Makefile", sources: []scaffold.Source{{Gomod: "../../go.mod"}}},
			{output: "../../Dockerfile", sources: []scaffold.Source{{Gomod: "../../go.mod"}}},
		}},
	)

	g, err := BuildDAG(m)
	if err != nil {
		t.Fatal(err)
	}

	// Both build files should depend on every .go directive.
	// Note: the Node.ID() canonicalizes via filepath.Clean, so
	// "tools/gen" + "../../Makefile" -> "Makefile" (resolved to
	// project root). That's correct — the ID describes where the
	// file actually lands.
	makefileID := "Makefile"
	wantPreds := map[string]bool{
		"internal/domain/user.go":    true,
		"internal/repo/user_repo.go": true,
	}
	gotPreds := g.Predecessors[makefileID]
	if len(gotPreds) != len(wantPreds) {
		t.Errorf("Makefile preds: want %d, got %d (%v)",
			len(wantPreds), len(gotPreds), gotPreds)
	}
	for want := range wantPreds {
		if _, ok := gotPreds[want]; !ok {
			t.Errorf("Makefile missing pred %q", want)
		}
	}

	// Levels: domain (0), repo (1), build files (2).
	levels := g.Levels()
	if len(levels) != 3 {
		t.Fatalf("want 3 levels, got %d: %v", len(levels), levelsAsIDs(levels))
	}

	// Level 2 must contain both build files; LevelContainsGomod should be true.
	if !LevelContainsGomod(levels[2]) {
		t.Errorf("level 2 should be detected as gomod-containing")
	}
	if LevelContainsGomod(levels[0]) {
		t.Errorf("level 0 should NOT be detected as gomod-containing")
	}

	// Build files should NOT depend on each other (they're both gomod-sourced
	// but neither produces a .go file, so no inter-build edge).
	dockerfileID := "Dockerfile"
	if _, depends := g.Predecessors[makefileID][dockerfileID]; depends {
		t.Errorf("Makefile should NOT depend on Dockerfile (peer build files)")
	}
	if _, depends := g.Predecessors[dockerfileID][makefileID]; depends {
		t.Errorf("Dockerfile should NOT depend on Makefile (peer build files)")
	}
}

// TestBuildDAG_RejectsCycle: A -> B -> A in pkg dependencies.
func TestBuildDAG_RejectsCycle(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{
			{output: "x.go", sources: []scaffold.Source{{Pkg: "../b"}}},
		}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{
			{output: "y.go", sources: []scaffold.Source{{Pkg: "../a"}}},
		}},
	)
	_, err := BuildDAG(m)
	if err == nil {
		t.Fatal("expected cycle error, got nil")
	}
	if !strings.Contains(err.Error(), "cycle") {
		t.Errorf("error should mention cycle: %v", err)
	}
}

// TestBuildDAG_SelfPkgIgnored: pkg: . (the package's own directory)
// is a self-reference; should be silently ignored, not become a cycle.
func TestBuildDAG_SelfPkgIgnored(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "internal/repo", pkg: "repo", files: []fileSpec{
			{output: "postgres.go", sources: []scaffold.Source{
				{Pkg: "."}, // self
				{Module: "github.com/jackc/pgx/v5/pgxpool"},
			}},
		}},
	)
	g, err := BuildDAG(m)
	if err != nil {
		t.Fatal(err)
	}
	if n := len(g.Predecessors["internal/repo/postgres.go"]); n != 0 {
		t.Errorf("self-pkg should produce 0 internal preds, got %d", n)
	}
}

// TestBuildDAG_PkgOutsideManifestIgnored: pkg pointing at a directory
// that no manifest entry produces. Silently allowed; the user may be
// referencing a stable hand-written package outside scaffold scope.
func TestBuildDAG_PkgOutsideManifestIgnored(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "internal/cache", pkg: "cache", files: []fileSpec{
			{output: "cache.go", sources: []scaffold.Source{
				{Pkg: "../some_external_package"}, // not in manifest
			}},
		}},
	)
	g, err := BuildDAG(m)
	if err != nil {
		t.Fatal(err)
	}
	if n := len(g.Predecessors["internal/cache/cache.go"]); n != 0 {
		t.Errorf("pkg outside manifest should produce 0 preds, got %d", n)
	}
}

// TestBuildDAG_LocalSymbolCreatesEdge: symbol with a local LOC behaves
// like pkg.
func TestBuildDAG_LocalSymbolCreatesEdge(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "domain", pkg: "domain", files: []fileSpec{
			{output: "user.go"},
		}},
		packageSpec{path: "service", pkg: "service", files: []fileSpec{
			{output: "svc.go", sources: []scaffold.Source{
				{Symbol: "User:../domain"},
			}},
		}},
	)
	g, err := BuildDAG(m)
	if err != nil {
		t.Fatal(err)
	}
	preds := g.Predecessors["service/svc.go"]
	if _, ok := preds["domain/user.go"]; !ok {
		t.Errorf("local symbol should create edge; preds = %v", preds)
	}
}

// TestBuildDAG_RemoteSymbolNoEdge: symbol with a remote LOC behaves
// like module.
func TestBuildDAG_RemoteSymbolNoEdge(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "domain", pkg: "domain", files: []fileSpec{
			{output: "user.go"},
		}},
		packageSpec{path: "service", pkg: "service", files: []fileSpec{
			{output: "svc.go", sources: []scaffold.Source{
				{Symbol: "Command:github.com/spf13/cobra"},
			}},
		}},
	)
	g, err := BuildDAG(m)
	if err != nil {
		t.Fatal(err)
	}
	if n := len(g.Predecessors["service/svc.go"]); n != 0 {
		t.Errorf("remote symbol should produce no internal edges, got %d", n)
	}
}

// TestBuildDAG_DuplicateNodeIDs: same package + same output declared
// twice -> error at DAG build time. (Validation in scaffold catches
// most of this, but weave double-checks because it indexes by ID.)
func TestBuildDAG_DuplicateNodeIDs(t *testing.T) {
	// Manually construct because manifest validator would reject this
	// before we got here. Bypass the validator to test weave's own check.
	m := &scaffold.Manifest{
		Packages: []scaffold.Package{
			{
				Path:    "a",
				Package: "a",
				Files: []scaffold.File{
					{Output: "x.go", Task: "t"},
					{Output: "x.go", Task: "t"},
				},
			},
		},
	}
	_, err := BuildDAG(m)
	if err == nil {
		t.Fatal("expected duplicate ID error, got nil")
	}
	if !strings.Contains(err.Error(), "duplicate") {
		t.Errorf("error should mention duplicate: %v", err)
	}
}

// TestIsLocalSymbol covers the heuristic on its own.
func TestIsLocalSymbol(t *testing.T) {
	cases := []struct {
		symbol string
		want   bool
	}{
		{"", false},
		{"User:./domain", true},
		{"User:../domain", true},
		{"User:domain", true}, // bare local name
		{"User:/abs/path", true},
		{"User:~/code", true},
		{"Command:github.com/spf13/cobra", false},
		{"Command:github.com/spf13/cobra@v1.10.1", false},
		{"Command:gopkg.in/yaml.v3", false},
		{"NoColon", false},
	}
	for _, c := range cases {
		got := isLocalSymbol(c.symbol)
		if got != c.want {
			t.Errorf("isLocalSymbol(%q) = %v, want %v", c.symbol, got, c.want)
		}
	}
}

// TestNode_IsGoOutput covers the suffix check.
func TestNode_IsGoOutput(t *testing.T) {
	cases := []struct {
		out  string
		want bool
	}{
		{"main.go", true},
		{"user_cache.go", true},
		{"../../main.go", true},
		{"Makefile", false},
		{"Dockerfile", false},
		{"docker-compose.yml", false},
		{"../../Makefile", false},
	}
	for _, c := range cases {
		n := Node{Output: c.out}
		if n.IsGoOutput() != c.want {
			t.Errorf("IsGoOutput(%q) = %v, want %v", c.out, n.IsGoOutput(), c.want)
		}
	}
}

// TestLevels_ParallelismWithinLevel: two independent paths produce
// nodes that should be in the same level (run in parallel).
func TestLevels_ParallelismWithinLevel(t *testing.T) {
	m := buildManifest(
		// Two independent seeds and one cross-cutting consumer.
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{{output: "b.go"}}},
		packageSpec{path: "c", pkg: "c", files: []fileSpec{
			{output: "c.go", sources: []scaffold.Source{
				{Pkg: "../a"},
				{Pkg: "../b"},
			}},
		}},
	)
	g, err := BuildDAG(m)
	if err != nil {
		t.Fatal(err)
	}
	levels := g.Levels()
	if len(levels) != 2 {
		t.Fatalf("want 2 levels (seeds||, consumer), got %d", len(levels))
	}
	if len(levels[0]) != 2 {
		t.Errorf("level 0 should have 2 parallel nodes, got %d", len(levels[0]))
	}
	if len(levels[1]) != 1 {
		t.Errorf("level 1 should have 1 node, got %d", len(levels[1]))
	}
}

// TestBuildDAG_Determinism: run twice, level sequences must match
// exactly. (Maps in Go iterate randomly; if our level construction
// doesn't normalize, this catches it.)
func TestBuildDAG_Determinism(t *testing.T) {
	m := buildManifest(
		packageSpec{path: "a", pkg: "a", files: []fileSpec{{output: "a.go"}}},
		packageSpec{path: "b", pkg: "b", files: []fileSpec{{output: "b.go"}}},
		packageSpec{path: "c", pkg: "c", files: []fileSpec{{output: "c.go"}}},
		packageSpec{path: "d", pkg: "d", files: []fileSpec{
			{output: "d.go", sources: []scaffold.Source{
				{Pkg: "../a"}, {Pkg: "../b"}, {Pkg: "../c"},
			}},
		}},
	)
	for i := 0; i < 25; i++ {
		g, err := BuildDAG(m)
		if err != nil {
			t.Fatal(err)
		}
		levels := g.Levels()
		ids := levelsAsIDs(levels)
		want := [][]string{
			{"a/a.go", "b/b.go", "c/c.go"},
			{"d/d.go"},
		}
		if !reflect.DeepEqual(ids, want) {
			t.Errorf("run %d: non-deterministic levels: got %v, want %v",
				i, ids, want)
		}
	}
}

func levelsAsIDs(levels [][]Node) [][]string {
	out := make([][]string, len(levels))
	for i, lvl := range levels {
		ids := make([]string, len(lvl))
		for j, n := range lvl {
			ids[j] = n.ID()
		}
		sort.Strings(ids)
		out[i] = ids
	}
	return out
}
