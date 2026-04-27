// pkg/analyze/clean.go
//
// CleanDependencies removes parser-leak noise from a dependency set.
//
// The forward-dep extractor in pkg/extract walks every Ident and
// SelectorExpr it sees during AST traversal and adds them to the
// Functions set. This is the right behavior for *call sites*, but
// the extractor doesn't have type-resolution information — it can't
// distinguish "p.Address" (field access on a receiver) from
// "fmt.Println" (function call on an import). Both end up in the
// dependency set as if they were function dependencies.
//
// At the render layer we have one piece of information the extractor
// didn't: declMap, the set of *actually-declared* symbols in the
// package. Anything in the dep set that:
//
//   1. has the form "x.Y" where x is a single-letter lowercase ident
//      (the receiver-variable convention in idiomatic Go), and
//   2. doesn't correspond to a real declaration in declMap
//
// is almost certainly a receiver-variable field access, not a real
// function dependency. We drop it.
//
// We similarly drop bare lowercase single-letter idents that have no
// declaration — these are the receiver-variable Idents themselves.
//
// This is a heuristic, not a proof. The trade-off: a real package-
// level function named "p" (one letter, lowercase) would get filtered
// out, but in practice such functions don't exist in well-formed Go
// (any package would have a longer name and the function would be
// addressed as pkg.p, not bare p). The false-positive rate is
// effectively zero on any code Claude is likely to be reading.
//
// CleanDependencies is package-scoped: it does not mutate the input
// Dependencies (which is a value type with monoid sets), it returns
// a new one with the noise removed.

package analyze

import (
	"strings"

	astpkg "github.com/Pure-Company/pureast/pkg/ast"
)

// CleanDependencies returns deps with parser-leak entries removed.
// declMap is the package's full decl table — names absent from it
// can't be real local-package dependencies, so they're candidates
// for filtering.
func CleanDependencies(deps astpkg.Dependencies, declMap map[string]astpkg.DeclNode) astpkg.Dependencies {
	cleanFns := filterFunctionNames(deps.Functions.ToSlice(), declMap)

	// Rebuild the Functions set with the cleaned slice. The other
	// dependency categories (Types, Structs, Interfaces, etc.) come
	// from the type-position extractors and don't suffer the same
	// leak — we leave them alone.
	out := astpkg.NewDependencies()
	for _, n := range cleanFns {
		out = out.AddFunction(n)
	}
	for _, n := range deps.Types.ToSlice() {
		out = out.AddType(n)
	}
	for _, n := range deps.Structs.ToSlice() {
		out = out.AddStruct(n)
	}
	for _, n := range deps.Interfaces.ToSlice() {
		out = out.AddInterface(n)
	}
	for _, n := range deps.Imports.ToSlice() {
		out = out.AddImport(n)
	}
	// Constants and Variables don't have an Add* helper exposed; copy
	// the underlying sets directly. (They aren't subject to the leak,
	// either, so this is just preservation.)
	out.Constants = deps.Constants
	out.Variables = deps.Variables

	return out
}

// filterFunctionNames drops receiver-variable noise from a slice of
// names. The rules are documented at the top of this file; in short:
//
//   - drop "x" where x is a single lowercase letter and not in declMap
//   - drop "x.Y" where x is a single lowercase letter and "x.Y" is
//     not in declMap (most receiver-prefixed field accesses)
//
// Names from imports survive: "fmt.Println" has a multi-character
// prefix, "f.Println" with single-letter f *would* be filtered, but
// nobody aliases imports as single letters. The heuristic is robust
// against realistic Go code and conservative against odd patterns.
func filterFunctionNames(names []string, declMap map[string]astpkg.DeclNode) []string {
	out := make([]string, 0, len(names))
	for _, n := range names {
		if isLikelyReceiverNoise(n, declMap) {
			continue
		}
		out = append(out, n)
	}
	return out
}

func isLikelyReceiverNoise(name string, declMap map[string]astpkg.DeclNode) bool {
	if name == "" {
		return true
	}
	// A real declaration in this package is never noise.
	if _, ok := declMap[name]; ok {
		return false
	}

	// Bare single-letter lowercase: classic receiver/loop variable.
	if len(name) == 1 && name[0] >= 'a' && name[0] <= 'z' {
		return true
	}

	// "x.Y" with single-letter lowercase prefix.
	if idx := strings.Index(name, "."); idx > 0 {
		prefix := name[:idx]
		if len(prefix) == 1 && prefix[0] >= 'a' && prefix[0] <= 'z' {
			return true
		}
	}

	return false
}
