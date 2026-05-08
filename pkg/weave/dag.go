// Package weave orchestrates parallel execution of the directives
// declared in a pureast manifest. The manifest already encodes the
// dependency graph (each File's sources say which packages it reads);
// weave reads that graph, topologically sorts it, and runs each level
// in parallel.
//
// Phase 1 — this package — is intentionally narrow: parallel execution
// of independent directives, level-by-level fan-out. No reconcile loop,
// no per-agent dialogue, no supervisor. Those are later phases.
//
// The level model:
//
//	Level 0:  every directive whose --pkg / local --symbol sources
//	          don't reference any other directive in the manifest.
//	          These run first, in parallel.
//
//	Level N:  every directive whose dependencies all completed in
//	          levels < N. Run in parallel within the level; the level
//	          only starts once level N-1 has fully finished.
//
// Source-type semantics for ordering:
//
//	pkg / local symbol  : creates internal edge to every directive in
//	                      the target package (must run after they finish).
//	module / remote sym : no internal edge. External; resolved via
//	                      go mod download independently of weave's run.
//	gomod               : creates internal edge to every directive
//	                      whose output ends in .go. Cannot run until
//	                      the project's Go files exist; weave also
//	                      runs `go mod tidy` before any level
//	                      containing a gomod-sourced directive.
//
// These rules make the manifest's source graph isomorphic to the
// project's package import graph: seeds at the top, build files at
// the bottom, everything else at a level determined by its dep chain.
// Cycles in the manifest correspond to import cycles, which wouldn't
// compile anyway — so we reject them at DAG construction time.
package weave

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Pure-Company/pureast/pkg/scaffold"
)

// Node is a single executable directive — one file produced by one
// claude-edit invocation. NodeID uniquely identifies it across the
// whole manifest.
type Node struct {
	// PackagePath is the manifest package's path (canonical, slash-
	// separated, e.g. "internal/cache").
	PackagePath string
	// Output is the directive's output filename or relative path
	// (e.g. "link_cache.go" or "../../Makefile").
	Output string
	// File is the underlying manifest entry, kept around so the
	// runner can build the claude-edit invocation later.
	File scaffold.File
	// Package is the manifest package, kept for context the runner
	// may need (model defaults, doc comments, etc.).
	Package scaffold.Package
}

// ID returns a stable, unique identifier. Format is "pkgpath/output",
// joined with forward slashes regardless of OS, so two main.go files
// in cmd/foo and cmd/bar don't collide.
func (n Node) ID() string {
	return filepath.ToSlash(filepath.Join(n.PackagePath, n.Output))
}

// IsGoOutput reports whether this directive produces a .go file.
// Used by gomod-source rule (which only depends on .go-producing
// directives, not on Makefile/Dockerfile/YAML peers).
func (n Node) IsGoOutput() bool {
	return strings.HasSuffix(n.Output, ".go")
}

// HasGomodSource reports whether this directive has at least one
// gomod source. Such directives need (a) every .go directive
// finished, and (b) `go mod tidy` to have populated go.mod.
func (n Node) HasGomodSource() bool {
	for _, s := range n.File.Sources {
		if s.Gomod != "" {
			return true
		}
	}
	return false
}

// DAG is the dependency graph of the manifest. Edges point from a
// node to its predecessors; a node only runs after every predecessor
// has succeeded.
type DAG struct {
	// Nodes indexed by ID for O(1) lookup.
	Nodes map[string]Node
	// Predecessors[id] is the set of node IDs that `id` depends on.
	// Empty set = ready to run (level 0).
	Predecessors map[string]map[string]struct{}
}

// BuildDAG constructs the DAG from a parsed manifest. Returns an
// error if the manifest declares a dependency cycle (which would
// also indicate a Go import cycle in the eventual project).
//
// Edge construction:
//
//	pkg: <local>     — edge to every directive in the resolved package
//	symbol: NAME:<local> — same as pkg
//	gomod: <path>    — edge to every directive whose output ends in .go
//	module / remote symbol — no edge (external, doesn't gate)
//
// A Source.Pkg that resolves to a package not in the manifest is
// silently ignored. The user may legitimately --pkg into a stable,
// hand-maintained package outside the manifest's scope; the runner
// just won't gate on it.
func BuildDAG(m *scaffold.Manifest) (*DAG, error) {
	if m == nil || len(m.Packages) == 0 {
		return nil, fmt.Errorf("empty manifest")
	}

	// Index of (canonical package path) -> all node IDs in that package.
	// Used to translate "pkg ../repo" -> "every directive in internal/repo".
	pkgPathToNodeIDs := make(map[string][]string, len(m.Packages))
	nodes := make(map[string]Node)
	var allGoNodeIDs []string

	for _, pkg := range m.Packages {
		canonicalPkgPath := filepath.ToSlash(filepath.Clean(pkg.Path))
		for _, f := range pkg.Files {
			n := Node{
				PackagePath: canonicalPkgPath,
				Output:      f.Output,
				File:        f,
				Package:     pkg,
			}
			id := n.ID()
			if _, exists := nodes[id]; exists {
				return nil, fmt.Errorf("duplicate node ID %q "+
					"(same package + output appears twice in manifest)", id)
			}
			nodes[id] = n
			pkgPathToNodeIDs[canonicalPkgPath] = append(pkgPathToNodeIDs[canonicalPkgPath], id)
			if n.IsGoOutput() {
				allGoNodeIDs = append(allGoNodeIDs, id)
			}
		}
	}

	// Sort the all-go list for stable iteration order in the edge
	// construction below; affects nothing semantically but produces
	// deterministic test output.
	sort.Strings(allGoNodeIDs)

	preds := make(map[string]map[string]struct{}, len(nodes))
	for id := range nodes {
		preds[id] = make(map[string]struct{})
	}

	for id, n := range nodes {
		for _, src := range n.File.Sources {
			switch {
			case src.Gomod != "":
				// gomod: depend on every .go directive in the manifest.
				// (Other gomod-sourced directives — e.g. another build
				// file — are not Go outputs, so they don't gate each
				// other. This is correct: build files can run in
				// parallel within their level.)
				for _, depID := range allGoNodeIDs {
					if depID == id {
						continue // shouldn't happen (gomod files aren't .go) but be safe
					}
					preds[id][depID] = struct{}{}
				}

			case src.Pkg != "" || isLocalSymbol(src.Symbol):
				localPath := src.Pkg
				if localPath == "" {
					_, localPath, _ = strings.Cut(src.Symbol, ":")
				}
				resolved := filepath.ToSlash(
					filepath.Clean(filepath.Join(n.PackagePath, localPath)))
				depIDs, found := pkgPathToNodeIDs[resolved]
				if !found {
					// Points at something not in the manifest. OK —
					// silently ignored.
					continue
				}
				for _, depID := range depIDs {
					if depID == id {
						// Self-edge (e.g. pkg: .). Ignore.
						continue
					}
					preds[id][depID] = struct{}{}
				}

				// case src.Module != "" or remote symbol: no edge.
			}
		}
	}

	g := &DAG{Nodes: nodes, Predecessors: preds}

	if cycle := detectCycle(g); cycle != nil {
		return nil, fmt.Errorf("manifest has a dependency cycle: %s",
			strings.Join(cycle, " -> "))
	}

	return g, nil
}

// isLocalSymbol reports whether a Source.Symbol's LOC component is a
// local path (vs. a remote module spec). Mirrors the heuristic in
// claude-edit: starts with ./, /, ~ -> local; otherwise check for
// dot in first segment -> remote if dotted, local if not.
func isLocalSymbol(symbol string) bool {
	if symbol == "" {
		return false
	}
	_, loc, ok := strings.Cut(symbol, ":")
	if !ok || loc == "" {
		return false
	}
	if strings.HasPrefix(loc, ".") || strings.HasPrefix(loc, "/") || strings.HasPrefix(loc, "~") {
		return true
	}
	first, _, _ := strings.Cut(loc, "/")
	return !strings.Contains(first, ".")
}

// detectCycle does a DFS-based cycle check. Returns the cycle as a
// slice of node IDs in traversal order if one exists, nil otherwise.
// Iteration order is alphabetical for deterministic test output.
func detectCycle(g *DAG) []string {
	const (
		white = 0 // unvisited
		gray  = 1 // currently in DFS stack
		black = 2 // fully explored
	)
	color := make(map[string]int, len(g.Nodes))
	var cycle []string

	ids := make([]string, 0, len(g.Nodes))
	for id := range g.Nodes {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	var stack []string
	var visit func(id string) bool
	visit = func(id string) bool {
		if color[id] == gray {
			// Found a cycle. The cycle is from the first occurrence
			// of `id` in stack to the end, plus `id` again.
			for i, sid := range stack {
				if sid == id {
					cycle = append([]string{}, stack[i:]...)
					cycle = append(cycle, id)
					return true
				}
			}
			cycle = append([]string{}, stack...)
			cycle = append(cycle, id)
			return true
		}
		if color[id] == black {
			return false
		}
		color[id] = gray
		stack = append(stack, id)

		// Iterate predecessors in deterministic order.
		predIDs := make([]string, 0, len(g.Predecessors[id]))
		for pid := range g.Predecessors[id] {
			predIDs = append(predIDs, pid)
		}
		sort.Strings(predIDs)
		for _, pid := range predIDs {
			if visit(pid) {
				return true
			}
		}

		stack = stack[:len(stack)-1]
		color[id] = black
		return false
	}

	for _, id := range ids {
		if color[id] == white {
			if visit(id) {
				return cycle
			}
		}
	}
	return nil
}

// Levels partitions the DAG into execution levels. Level i contains
// every node whose predecessors all reside in levels < i. Returns a
// slice of slices: levels[0] runs first, levels[1] runs after
// levels[0] completes, etc.
//
// Within a level, nodes are sorted by ID so weave's logs are
// deterministic across runs (helpful for debugging and tests).
func (g *DAG) Levels() [][]Node {
	// Standard Kahn's algorithm with deterministic ordering.
	remaining := make(map[string]int, len(g.Nodes))
	for id, p := range g.Predecessors {
		remaining[id] = len(p)
	}
	// successor index: who depends on each node
	succ := make(map[string][]string, len(g.Nodes))
	for id, preds := range g.Predecessors {
		for predID := range preds {
			succ[predID] = append(succ[predID], id)
		}
	}

	var levels [][]Node
	for len(remaining) > 0 {
		var ready []string
		for id, deg := range remaining {
			if deg == 0 {
				ready = append(ready, id)
			}
		}
		if len(ready) == 0 {
			// Shouldn't happen — we cycle-checked at construction.
			// Defensive: bail rather than infinite-loop.
			break
		}
		sort.Strings(ready)
		level := make([]Node, 0, len(ready))
		for _, id := range ready {
			level = append(level, g.Nodes[id])
			delete(remaining, id)
			// "Run" the node: drop in-degree on its successors.
			for _, sid := range succ[id] {
				remaining[sid]--
			}
		}
		levels = append(levels, level)
	}
	return levels
}

// LevelContainsGomod reports whether any node in the level has a
// gomod source. Used by the runner to decide when to run go mod tidy.
func LevelContainsGomod(level []Node) bool {
	for _, n := range level {
		if n.HasGomodSource() {
			return true
		}
	}
	return false
}
