// pkg/analyze/reverse.go
//
// Reverse dependency analysis and depth-bounded traversal.
//
// The forward graph in deps.go answers "what does X depend on?".
// Two queries that are essential for the LLM-context use case but can't
// be answered by the forward graph alone:
//
//   1. "Who depends on X?" — needed when the user wants to understand
//      the impact of changing a symbol, or wants to gather context for
//      "I'm about to refactor X, show me everything affected."
//
//   2. "Show me X's dependencies up to depth N" — needed to keep
//      context size manageable. The forward graph computes the full
//      transitive closure; depth-bounding lets the user say "I want
//      direct dependencies plus one hop" and stop there.
//
// Both are implemented as new methods on DependencyGraph rather than
// modifications to existing ones. The original API is untouched.

package analyze

import (
	astpkg "github.com/Pure-Company/pureast/pkg/ast"
	"github.com/Pure-Company/purekernels/pkg/monoid"
)

// reverseEdges returns the inverted adjacency relation: for each
// declared symbol, the set of symbols whose direct dependencies
// include it. This is the underlying data structure for `Users`.
//
// The result is computed lazily and not cached — for large packages
// where reverse queries are common, callers should compute it once
// and reuse. We keep it a pure function on the graph rather than a
// field on DependencyGraph so the type stays a thin wrapper.
func (g DependencyGraph) reverseEdges() map[string]monoid.SetMonoid[string] {
	rev := make(map[string]monoid.SetMonoid[string], len(g.Decls))

	// Initialize an empty set for every known declaration. This means
	// `Users("Foo")` returns an empty set rather than nil even when
	// nothing depends on Foo, which simplifies the caller.
	for name := range g.Decls {
		rev[name] = monoid.NewSetMonoid[string]()
	}

	// For each declaration, walk its direct dependencies and add an
	// inverted edge from each dependency back to this declaration.
	// We only invert names that are themselves in the graph — external
	// imports (stdlib, third-party) aren't queryable as targets here.
	for name, decl := range g.Decls {
		addReverse(rev, name, decl.Deps.Types.ToSlice())
		addReverse(rev, name, decl.Deps.Functions.ToSlice())
		addReverse(rev, name, decl.Deps.Structs.ToSlice())
		addReverse(rev, name, decl.Deps.Interfaces.ToSlice())
		// Constants and variables can also be dependents, though the
		// forward extractor populates them less consistently.
		addReverse(rev, name, decl.Deps.Constants.ToSlice())
		addReverse(rev, name, decl.Deps.Variables.ToSlice())
	}

	return rev
}

func addReverse(rev map[string]monoid.SetMonoid[string], dependent string, targets []string) {
	for _, target := range targets {
		set, ok := rev[target]
		if !ok {
			// Target isn't in our declaration map — likely an external
			// reference. Skip rather than create an entry, so callers
			// querying for a stdlib name get an empty result instead
			// of a partial one.
			continue
		}
		rev[target] = set.Insert(dependent)
	}
}

// Users returns the set of symbols that directly reference the target.
// This is the one-hop reverse query: it does not transitively walk who
// uses the users. For transitive reverse closure, see UsersTransitive.
//
// The return type matches forward queries (astpkg.Dependencies) so
// callers can format the result with the same code path as `deps`.
// We populate the field that matches the kind of the *target*, but
// in practice callers just want the union, so all dependent names go
// into Functions (the most general bucket).
func (g DependencyGraph) Users(targetName string) astpkg.Dependencies {
	rev := g.reverseEdges()
	users, ok := rev[targetName]
	if !ok {
		return astpkg.NewDependencies()
	}

	deps := astpkg.NewDependencies()
	for _, name := range users.ToSlice() {
		deps = deps.AddFunction(name)
	}
	return deps
}

// UsersTransitive returns every symbol that transitively depends on the
// target — direct users, plus their users, recursively. Useful for
// "show me everything affected by changing X."
//
// The traversal uses a visited set to terminate on cycles, same shape
// as ResolveTransitive but walking inverted edges.
func (g DependencyGraph) UsersTransitive(targetName string) astpkg.Dependencies {
	rev := g.reverseEdges()
	visited := monoid.NewSetMonoid[string]()
	collected := monoid.NewSetMonoid[string]()
	collected = collectReverse(rev, targetName, visited, collected)

	deps := astpkg.NewDependencies()
	for _, name := range collected.ToSlice() {
		deps = deps.AddFunction(name)
	}
	return deps
}

func collectReverse(
	rev map[string]monoid.SetMonoid[string],
	name string,
	visited monoid.SetMonoid[string],
	acc monoid.SetMonoid[string],
) monoid.SetMonoid[string] {
	if visited.Contains(name) {
		return acc
	}
	visited = visited.Insert(name)

	users, ok := rev[name]
	if !ok {
		return acc
	}
	for _, u := range users.ToSlice() {
		acc = acc.Insert(u)
		acc = collectReverse(rev, u, visited, acc)
	}
	return acc
}

// ResolveBounded computes dependencies up to a maximum traversal depth.
// Depth 0 returns just the direct dependencies of the target (one hop).
// Depth 1 returns direct dependencies plus their direct dependencies.
// Depth -1 (or any negative) is treated as unbounded and is equivalent
// to ResolveTransitive.
//
// This is the depth-bounded analogue of ResolveTransitive. The
// implementation tracks remaining depth instead of just visited names —
// a node at depth 2 is still useful to descend from even if visited at
// depth 0 would have stopped early, but for our use case (LLM context
// budgeting) the simpler "shortest-path" semantics is fine: once we've
// seen a name at any depth, we don't revisit. This means depth N
// returns the union of all nodes reachable in ≤ N hops, which is what
// the user wants.
func (g DependencyGraph) ResolveBounded(targetName string, depth int) astpkg.Dependencies {
	if depth < 0 {
		return g.ResolveTransitive(targetName)
	}
	visited := monoid.NewSetMonoid[string]()
	return g.resolveBoundedRec(targetName, depth, visited)
}

func (g DependencyGraph) resolveBoundedRec(
	name string,
	remaining int,
	visited monoid.SetMonoid[string],
) astpkg.Dependencies {
	if visited.Contains(name) {
		return astpkg.NewDependencies()
	}
	visited = visited.Insert(name)

	decl, ok := g.Decls[name]
	if !ok {
		return astpkg.NewDependencies()
	}

	// Always include the target's own immediate deps. The depth bound
	// controls *transitive* expansion: at depth 0 we include the
	// target's direct edges but don't follow them; at depth 1 we
	// follow one hop; etc.
	immediate := decl.Deps
	if remaining <= 0 {
		return immediate
	}

	depMonoid := astpkg.NewDependencyMonoid()
	combined := immediate

	for _, child := range immediate.Types.ToSlice() {
		combined = depMonoid.Combine(combined,
			g.resolveBoundedRec(child, remaining-1, visited))
	}
	for _, child := range immediate.Functions.ToSlice() {
		combined = depMonoid.Combine(combined,
			g.resolveBoundedRec(child, remaining-1, visited))
	}
	for _, child := range immediate.Structs.ToSlice() {
		combined = depMonoid.Combine(combined,
			g.resolveBoundedRec(child, remaining-1, visited))
	}
	for _, child := range immediate.Interfaces.ToSlice() {
		combined = depMonoid.Combine(combined,
			g.resolveBoundedRec(child, remaining-1, visited))
	}
	return combined
}
