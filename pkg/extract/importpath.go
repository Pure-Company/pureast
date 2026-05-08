// pkg/extract/importpath.go
//
// Resolve an on-disk directory to its module-qualified Go import path.
//
// Why this exists: when pureast extracts a package's signatures and feeds
// them to an LLM for codegen, the LLM also needs to know the *import path*
// of that package. Without it, the model guesses — often wrongly, often
// with a relative path like "../repo" that isn't even legal Go. The fix
// is mechanical: walk up to find go.mod, read its module declaration,
// compute (module-path + relative-subdir).
//
// This is a small piece of logic but it's separated from the cmd/ tree
// so that any code path that produces signatures-for-LLM (claude-edit
// today, the MCP server tomorrow) can use the same resolution rule.
package extract

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

// ImportPathFor returns the Go import path that corresponds to a local
// directory, by locating the enclosing go.mod and combining its module
// path with the directory's location relative to the module root.
//
// Returns ("", nil) when the directory exists but isn't inside any
// Go module (no go.mod found by walking upward). Callers should treat
// this as "no canonical import path available" — usually a no-op.
//
// Returns an error only on I/O failures or malformed go.mod files.
// "Not in a module" is the expected condition for ad-hoc directories
// and isn't error-worthy.
func ImportPathFor(dir string) (string, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("resolve %s: %w", dir, err)
	}
	if st, err := os.Stat(abs); err != nil {
		return "", fmt.Errorf("stat %s: %w", abs, err)
	} else if !st.IsDir() {
		return "", fmt.Errorf("%s is not a directory", abs)
	}

	// Walk up looking for go.mod. We bound the walk so a totally
	// unrelated directory (e.g. /tmp) doesn't traverse forever.
	// In practice the loop terminates at the filesystem root long
	// before this bound — it's a paranoia guard, not a real limit.
	cur := abs
	for i := 0; i < 64; i++ {
		gomodPath := filepath.Join(cur, "go.mod")
		if data, err := os.ReadFile(gomodPath); err == nil {
			// Found it. Parse out the module path.
			mf, perr := modfile.Parse(gomodPath, data, nil)
			if perr != nil {
				return "", fmt.Errorf("parse %s: %w", gomodPath, perr)
			}
			if mf.Module == nil || mf.Module.Mod.Path == "" {
				return "", fmt.Errorf("%s has no module declaration", gomodPath)
			}
			modulePath := mf.Module.Mod.Path

			// Compute the import path = module path + (dir relative
			// to module root). When dir IS the module root, the
			// relative part is "." which we treat as empty.
			rel, err := filepath.Rel(cur, abs)
			if err != nil {
				return "", fmt.Errorf("relpath: %w", err)
			}
			rel = filepath.ToSlash(rel) // module paths use forward slashes regardless of OS
			if rel == "." {
				return modulePath, nil
			}
			return strings.TrimSuffix(modulePath, "/") + "/" + rel, nil
		}

		parent := filepath.Dir(cur)
		if parent == cur {
			// Hit filesystem root without finding go.mod.
			return "", nil
		}
		cur = parent
	}
	return "", nil
}
