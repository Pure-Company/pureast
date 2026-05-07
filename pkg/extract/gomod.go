// pkg/extract/gomod.go
//
// Parse a go.mod file and extract its direct dependencies.
//
// "Direct" here means: listed in the `require` block(s) without an
// `// indirect` comment marker. These are the dependencies the project
// imports explicitly — the ones whose APIs are part of the project's
// own surface, not transitive plumbing.
//
// Why a separate file: the rest of pkg/extract operates on Go source
// trees. go.mod parsing is a different concern (it's about the module
// graph, not the AST), and the standard library doesn't expose a parser
// for go.mod's grammar. golang.org/x/mod/modfile does, and is the
// canonical choice — it's what `go` itself uses internally.
//
// We honor `replace` directives by substituting the right-hand-side
// (path + version, or local filesystem path) when a require entry has
// a matching replacement. This matches what `go build` actually
// compiles, which is what the user cares about when feeding their
// project's deps to an LLM.
package extract

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"golang.org/x/mod/modfile"
)

// ModuleRef identifies one module-and-version pair extracted from a
// go.mod's direct dependencies, with replace directives applied.
type ModuleRef struct {
	// Path is the module path as resolved (after `replace` has been
	// applied if there was a matching directive).
	Path string

	// Version is the module version. For local-path replacements
	// (replace foo => ../local), this is empty and LocalPath is set.
	Version string

	// LocalPath is set ONLY for `replace` directives that point at a
	// local filesystem path (the "==>" target was not a module path
	// but a relative or absolute directory). When non-empty, callers
	// should treat this as a directory to dump directly, NOT as a
	// module to resolve through `go mod download`.
	LocalPath string

	// OriginalPath is what was written in the require block, before
	// `replace` substitution. Useful for diagnostics — users wrote
	// "github.com/foo/bar" and we want to tell them what we did with it.
	OriginalPath string
}

// ParseGoMod reads a go.mod file and returns its direct dependencies,
// with replace directives applied. Indirect deps (marked `// indirect`)
// are skipped — the caller asked for the project's *own* API surface,
// not its transitive closure.
//
// If the path doesn't exist, doesn't parse, or has no direct deps, a
// descriptive error is returned. Empty results are NOT returned silently;
// "no direct deps" almost always means the user pointed at the wrong
// file and would rather hear about it.
//
// Output is sorted by module path for determinism — same input, same
// output, run after run, which makes the dumps cacheable and diffable.
func ParseGoMod(path string) ([]ModuleRef, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	// modfile.Parse does the heavy lifting: tokenizes, validates the
	// require/replace/exclude/retract blocks, attaches comments to the
	// statements they belong to. We pass `nil` as the version-fixer
	// because we don't want to upgrade pseudo-versions on the fly —
	// we want exactly what the user wrote.
	mf, err := modfile.Parse(path, data, nil)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	// Build a replace map keyed by the LHS module path. modfile already
	// represents replaces structurally; we just need O(1) lookup during
	// the require iteration. A nil/missing entry means "no replacement,
	// use the original."
	//
	// Note: replaces can also be keyed by (path, version) pair for
	// version-specific replacements. We treat those as more specific
	// matches than path-only replacements when both exist.
	type replaceKey struct{ Path, Version string }
	replaces := make(map[replaceKey]*modfile.Replace, len(mf.Replace))
	for _, r := range mf.Replace {
		replaces[replaceKey{r.Old.Path, r.Old.Version}] = r
	}

	var out []ModuleRef
	for _, req := range mf.Require {
		if req.Indirect {
			continue
		}
		ref := ModuleRef{
			Path:         req.Mod.Path,
			Version:      req.Mod.Version,
			OriginalPath: req.Mod.Path,
		}

		// Try version-specific replace first, then path-only replace.
		// The Go spec says version-specific wins.
		var rep *modfile.Replace
		if r, ok := replaces[replaceKey{req.Mod.Path, req.Mod.Version}]; ok {
			rep = r
		} else if r, ok := replaces[replaceKey{req.Mod.Path, ""}]; ok {
			rep = r
		}

		if rep != nil {
			// A replace's New.Version is empty when the target is a
			// filesystem path. modfile encodes this by leaving Version
			// empty; the Path then holds a relative or absolute dir.
			if rep.New.Version == "" {
				// Local path replacement. Resolve relative to the
				// go.mod's directory so users can write
				// `replace foo => ../local-foo` and have it work
				// regardless of where pureast was invoked from.
				localPath := rep.New.Path
				if !filepath.IsAbs(localPath) {
					localPath = filepath.Join(filepath.Dir(path), localPath)
				}
				ref.Path = rep.New.Path // Keep the original-as-written for headers
				ref.Version = ""
				ref.LocalPath = localPath
			} else {
				ref.Path = rep.New.Path
				ref.Version = rep.New.Version
			}
		}

		out = append(out, ref)
	}

	if len(out) == 0 {
		return nil, fmt.Errorf("%s: no direct dependencies found (only indirect deps, if any)", path)
	}

	sort.Slice(out, func(i, j int) bool {
		// Sort by the resolved path, not OriginalPath, so replaced
		// modules end up where they actually point. This keeps the
		// dump alphabetically stable on the *effective* deps.
		return out[i].Path < out[j].Path
	})

	return out, nil
}

// FilterModules applies optional skip/only filters to a module list.
//
// Both lists match against ModuleRef.OriginalPath (what the user wrote
// in go.mod) AND against the resolved Path (after replace), so users
// can refer to deps by either name. Matching is exact — if you want
// to skip everything under github.com/aws/, you pass each module name.
// (Glob support is a future-feature; YAGNI for now.)
//
// If `only` is non-empty, only modules whose path appears in `only` are
// kept; everything else is dropped. `skip` then runs on the survivors.
// In other words: only is a whitelist, skip is a blacklist, and they
// compose (you can use both — only narrows, then skip excludes from
// the narrowed set).
func FilterModules(refs []ModuleRef, only, skip []string) []ModuleRef {
	if len(only) == 0 && len(skip) == 0 {
		return refs
	}

	onlySet := make(map[string]struct{}, len(only))
	for _, p := range only {
		onlySet[p] = struct{}{}
	}
	skipSet := make(map[string]struct{}, len(skip))
	for _, p := range skip {
		skipSet[p] = struct{}{}
	}

	matches := func(ref ModuleRef, set map[string]struct{}) bool {
		if _, ok := set[ref.Path]; ok {
			return true
		}
		if _, ok := set[ref.OriginalPath]; ok {
			return true
		}
		return false
	}

	var out []ModuleRef
	for _, ref := range refs {
		if len(only) > 0 && !matches(ref, onlySet) {
			continue
		}
		if len(skip) > 0 && matches(ref, skipSet) {
			continue
		}
		out = append(out, ref)
	}
	return out
}

// Spec returns the canonical "path@version" string for a ModuleRef
// suitable for handing to ResolveModule. For local-path replacements,
// returns the local path (callers are expected to check LocalPath
// before calling Spec, but we make this a no-throw convenience).
func (r ModuleRef) Spec() string {
	if r.LocalPath != "" {
		return r.LocalPath
	}
	if r.Version == "" {
		return r.Path
	}
	return r.Path + "@" + r.Version
}
