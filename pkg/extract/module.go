// pkg/extract/module.go
//
// Module resolution via `go mod download`.
//
// The conceptual story: pureast extracts ASTs from on-disk Go source.
// `go mod download` extracts on-disk Go source from remote modules.
// Compose them and pureast works on any public Go module without
// teaching pureast anything about HTTP, git, auth, versioning, or
// caching — those are all `go`'s problem, and `go` already solved
// them. We're a thin shim: parse the spec, ask `go` to materialize
// it, return the directory.
//
// This file deliberately does NOT depend on cobra or the cmd/ tree.
// It's a library function so MCP, library users, and the CLI all
// share the same resolution logic.
package extract

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ModuleResolution is what ResolveModule returns: enough info to
// reach the source on disk plus enough to identify what we got
// (useful for headers, error messages, and caching keys).
type ModuleResolution struct {
	// Dir is the absolute, on-disk path that contains the requested
	// package's .go files. Read-only (it lives in $GOMODCACHE).
	// If the user requested a sub-package, Dir points at the
	// sub-package, NOT the module root.
	Dir string

	// ModulePath is the canonical module path go resolved to
	// (e.g. "github.com/spf13/cobra"). May be a prefix of the
	// originally-requested spec when the user pointed at a
	// sub-package inside a module.
	ModulePath string

	// Version is the resolved version (e.g. "v1.10.2"). When the
	// caller asked for "@latest" or omitted the version, this
	// holds the concrete version go picked.
	Version string

	// SubPath is the path inside the module that the user asked
	// for, relative to the module root. Empty when the user
	// pointed at the module root itself.
	SubPath string
}

// ResolveModule turns a module spec into an on-disk path.
//
// Accepted spec forms:
//
//	github.com/foo/bar               -> @latest
//	github.com/foo/bar@v1.2.3        -> exact version
//	github.com/foo/bar@latest        -> explicit latest
//	github.com/foo/bar/sub/pkg       -> @latest, points at sub/pkg
//	github.com/foo/bar/sub@v1.2.3    -> sub-package at exact version
//
// The function is tolerant about sub-packages: if the literal spec
// isn't a module (because it's a package path inside one), we walk
// path segments off the right until `go mod download` succeeds, then
// re-attach the trimmed segments as SubPath. This matches the way
// users naturally talk about Go code — they think in import paths
// (`github.com/gin-gonic/gin/render`), not module paths.
//
// Side effects: runs `go mod download` in a temp directory. The
// download itself lands in the user's $GOMODCACHE (global, content-
// addressed, persistent across runs). Repeat calls for the same
// (path, version) are filesystem-cache hits and effectively free.
func ResolveModule(spec string) (ModuleResolution, error) {
	if spec == "" {
		return ModuleResolution{}, fmt.Errorf("empty module spec")
	}
	path, version := splitSpec(spec)
	if version == "" {
		version = "latest"
	}

	if _, err := exec.LookPath("go"); err != nil {
		return ModuleResolution{}, fmt.Errorf(
			"--module requires the `go` toolchain on PATH: %w", err)
	}

	// `go mod download` needs to run from inside a module, otherwise
	// it errors with "go: go.mod file not found". The cheapest way to
	// satisfy that is a throwaway directory with a stub go.mod —
	// the actual download still lands in the global $GOMODCACHE.
	workDir, cleanup, err := mkTempModule()
	if err != nil {
		return ModuleResolution{}, err
	}
	defer cleanup()

	// Try the literal path first, then progressively trim trailing
	// segments. Most users will hit on the first try (their spec IS
	// a module path); the trim loop is for the package-inside-module
	// case. We cap the trim at the second segment to avoid resolving
	// nonsense like `github.com@latest`.
	parts := strings.Split(path, "/")
	for n := len(parts); n >= 2; n-- {
		modCandidate := strings.Join(parts[:n], "/")
		res, ok, err := tryDownload(workDir, modCandidate, version)
		if err != nil {
			// Hard error (go invocation failed, JSON malformed,
			// etc.) — surface immediately rather than masking it
			// behind the trim loop.
			return ModuleResolution{}, err
		}
		if !ok {
			continue
		}

		sub := strings.Join(parts[n:], "/")
		dir := res.Dir
		if sub != "" {
			dir = filepath.Join(res.Dir, sub)
			if st, err := os.Stat(dir); err != nil || !st.IsDir() {
				return ModuleResolution{}, fmt.Errorf(
					"resolved module %s@%s but sub-package %q does not exist on disk at %s",
					res.Path, res.Version, sub, dir)
			}
		}
		return ModuleResolution{
			Dir:        dir,
			ModulePath: res.Path,
			Version:    res.Version,
			SubPath:    sub,
		}, nil
	}

	return ModuleResolution{}, fmt.Errorf(
		"could not resolve %q as a Go module (tried path and progressively shorter prefixes)", spec)
}

// splitSpec separates "path@version" into ("path", "version").
// Missing version returns empty string for the version component;
// callers fill in the default. We split on the last '@' to be
// robust against paths that legitimately contain '@' (rare but
// not impossible in pseudo-versions if mis-typed).
func splitSpec(spec string) (path, version string) {
	if i := strings.LastIndex(spec, "@"); i >= 0 {
		return spec[:i], spec[i+1:]
	}
	return spec, ""
}

// downloadResult mirrors the `go mod download -json` output. We
// only decode the fields we use; go can add fields freely without
// breaking us.
type downloadResult struct {
	Path    string
	Version string
	Dir     string
	Error   string
}

// tryDownload runs `go mod download -json path@version` and returns:
//
//   - (result, true, nil)  on success,
//   - (_, false, nil)      on a "not a known module" style error
//     (caller should keep trying shorter prefixes),
//   - (_, false, err)      on a hard failure (toolchain broken,
//     network down with no cache, malformed JSON, etc.).
//
// The "soft" / "hard" distinction is what makes the trim loop in
// ResolveModule work: we only loop on errors that genuinely look
// like "wrong path", not on errors that mean "go itself is sad".
func tryDownload(workDir, path, version string) (downloadResult, bool, error) {
	spec := path + "@" + version
	cmd := exec.Command("go", "mod", "download", "-x", "-json", spec)
	cmd.Dir = workDir
	// Inherit env so GOPROXY, GOPRIVATE, GOSUMDB, netrc, etc. work
	// the way the user already configured them. This is the whole
	// point of piggybacking on `go mod`: zero config for us.
	out, err := cmd.Output()
	if err != nil {
		// `go mod download` exits non-zero when resolution fails,
		// but it still emits JSON on stdout describing the error.
		// Prefer the JSON message — it's much more informative
		// than the bare "exit status 1".
		if ee, ok := err.(*exec.ExitError); ok && len(out) > 0 {
			var r downloadResult
			if jerr := json.Unmarshal(out, &r); jerr == nil && r.Error != "" {
				if isSoftResolutionError(r.Error) {
					return downloadResult{}, false, nil
				}
				return downloadResult{}, false, fmt.Errorf(
					"go mod download %s: %s", spec, r.Error)
			}
			return downloadResult{}, false, fmt.Errorf(
				"go mod download %s: %s", spec, strings.TrimSpace(string(ee.Stderr)))
		}
		return downloadResult{}, false, fmt.Errorf("go mod download %s: %w", spec, err)
	}

	var r downloadResult
	if err := json.Unmarshal(out, &r); err != nil {
		return downloadResult{}, false, fmt.Errorf(
			"parse go mod download output: %w", err)
	}
	if r.Error != "" {
		if isSoftResolutionError(r.Error) {
			return downloadResult{}, false, nil
		}
		return downloadResult{}, false, fmt.Errorf(
			"go mod download %s: %s", spec, r.Error)
	}
	if r.Dir == "" {
		return downloadResult{}, false, fmt.Errorf(
			"go mod download %s: empty Dir in result", spec)
	}
	return r, true, nil
}

// isSoftResolutionError reports whether a go-mod-download error
// message is the kind we should swallow during the prefix-trim
// loop. The substrings here are stable across recent Go versions.
// Network errors, auth failures, etc. are NOT in this list — those
// should propagate so the user sees a real diagnostic.
func isSoftResolutionError(msg string) bool {
	soft := []string{
		"not a known dependency",
		"no matching versions",
		"unknown revision",
		"invalid version",
		"is not a valid module path",
		"no required module provides",
	}
	low := strings.ToLower(msg)
	for _, s := range soft {
		if strings.Contains(low, strings.ToLower(s)) {
			return true
		}
	}
	return false
}

// mkTempModule creates an empty directory with a minimal go.mod so
// `go mod download` is willing to run there. Returns the directory
// and a cleanup func. The cleanup is best-effort — leaving a stub
// go.mod in /tmp on a crash is harmless.
func mkTempModule() (string, func(), error) {
	dir, err := os.MkdirTemp("", "pureast-modresolve-")
	if err != nil {
		return "", nil, fmt.Errorf("mktemp: %w", err)
	}
	gomod := "module pureast.local/modresolve\n\ngo 1.22\n"
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(gomod), 0644); err != nil {
		_ = os.RemoveAll(dir)
		return "", nil, fmt.Errorf("write stub go.mod: %w", err)
	}
	cleanup := func() { _ = os.RemoveAll(dir) }
	return dir, cleanup, nil
}
