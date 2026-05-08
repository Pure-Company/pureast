// pkg/scaffold/manifest.go
//
// Manifest schema and parser for `pureast scaffold`.
//
// Conceptual model: the manifest is a deterministic specification of a
// project's package layout and per-file generation directives. The
// scaffolder reads this YAML, materializes a tree of `gen.go` files,
// and stops. No LLM involvement at this layer — the structural shape
// of the project is the user's decision, written down explicitly,
// reviewed in PRs, version-controlled like any other source artifact.
//
// LLM-driven content generation happens AFTER scaffolding, inside the
// per-file `claude-edit` directives that scaffold writes into each
// `gen.go`. Two layers, two cache disciplines, both reviewable.
package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// Manifest is the root document — a list of packages, each containing
// a list of files to generate. Authored by humans (or by an LLM as a
// proposed starting point that humans then edit), parsed by the
// scaffolder.
type Manifest struct {
	// Module is the Go module path, optional. When set, it's used as
	// a sanity check against the project's actual go.mod — a mismatch
	// suggests the manifest was copied from another project without
	// updating, which is worth catching early. Empty disables the check.
	Module string `yaml:"module,omitempty"`

	// Packages is the ordered list of packages this project will have.
	// Order in the file is preserved for human readability but doesn't
	// affect generation: the scaffolder processes each independently.
	Packages []Package `yaml:"packages"`
}

// Package describes one Go package directory and the files it should
// own. Path is relative to the project root (where the manifest lives
// or wherever scaffold was invoked from with --root).
type Package struct {
	// Path is the package directory, relative to the project root.
	// Examples: "cmd/api", "internal/repo", "internal/cache".
	Path string `yaml:"path"`

	// Package is the Go package name (the name in `package <name>`).
	// Most of the time this matches the last segment of Path, but
	// `cmd/api` typically holds `package main`, so we make this
	// explicit rather than infer.
	Package string `yaml:"package"`

	// Doc is an optional package-level doc comment that scaffold
	// writes above the `package` declaration in gen.go. Useful for
	// documenting *why* a package exists, since gen.go is the only
	// committed file in the package before generation runs.
	Doc string `yaml:"doc,omitempty"`

	// Files is the list of per-file generation directives. Each
	// becomes one //go:generate line in this package's gen.go.
	Files []File `yaml:"files"`
}

// File describes one output file's generation directive. Translates
// almost 1:1 into a `pureast claude-edit` invocation.
type File struct {
	// Output is the filename (no directory) within the package's Path.
	// Examples: "main.go", "user_repo.go", "user_cache_gen.go".
	Output string `yaml:"output"`

	// Task is the natural-language description of what claude-edit
	// should produce. Multi-line YAML strings work cleanly here.
	Task string `yaml:"task"`

	// Sources are the context inputs claude-edit will assemble.
	// Each source becomes a flag on the directive line. May be
	// empty for a "seed" file that has no upstream context — e.g.
	// the file that declares the project's first interface.
	Sources []Source `yaml:"sources,omitempty"`

	// Filters apply to all sources uniformly. Same flags as
	// claude-edit accepts directly.
	Kind         string `yaml:"kind,omitempty"`          // interface|struct|func|... (default: all)
	ExportedOnly *bool  `yaml:"exported_only,omitempty"` // pointer so we can distinguish unset from false
	MaxTokens    int    `yaml:"max_tokens,omitempty"`    // 0 = unbounded
	Model        string `yaml:"model,omitempty"`         // e.g. "opus"; passes through to --model
}

// Source is one context input. Exactly one of Pkg, Module, Symbol,
// or Gomod must be set — multiple is a manifest error.
type Source struct {
	// Pkg is a relative or absolute path to a local package directory.
	// When relative, it's resolved against the parent file's package
	// directory (matching how `claude-edit --pkg` behaves at runtime).
	Pkg string `yaml:"pkg,omitempty"`

	// Module is a Go module spec (path[@version]). Resolved via
	// `go mod download` when claude-edit runs; not at scaffold time.
	Module string `yaml:"module,omitempty"`

	// Symbol is a NAME:LOC pair extracting one symbol with its deps.
	// LOC is a path or module spec (same dual rule as claude-edit).
	Symbol string `yaml:"symbol,omitempty"`

	// Gomod is a path to a go.mod file. Includes every direct dep.
	Gomod string `yaml:"gomod,omitempty"`
}

// LoadManifest reads and parses a YAML manifest from disk.
func LoadManifest(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read manifest %s: %w", path, err)
	}
	var m Manifest
	dec := yaml.NewDecoder(strings.NewReader(string(data)))
	dec.KnownFields(true) // catches typos like `pacakge:` instead of `package:`
	if err := dec.Decode(&m); err != nil {
		return nil, fmt.Errorf("parse manifest %s: %w", path, err)
	}
	if err := m.Validate(); err != nil {
		return nil, fmt.Errorf("validate manifest %s: %w", path, err)
	}
	return &m, nil
}

// Validate checks the manifest for the kinds of mistakes that would
// otherwise produce confusing errors much later in the pipeline. The
// goal is to catch every problem here, with a clear message naming
// the offending package or file, rather than letting them surface as
// cryptic failures during scaffold or generation.
func (m *Manifest) Validate() error {
	if len(m.Packages) == 0 {
		return fmt.Errorf("manifest has no packages")
	}

	seenPath := make(map[string]int, len(m.Packages))
	for i, pkg := range m.Packages {
		if pkg.Path == "" {
			return fmt.Errorf("packages[%d]: path is required", i)
		}
		// Reject paths that escape the project root via .. — we
		// honor relative paths up to the root but not above it.
		clean := filepath.Clean(pkg.Path)
		if filepath.IsAbs(clean) || strings.HasPrefix(clean, "..") {
			return fmt.Errorf("packages[%d] (%s): path must be relative and stay within project root",
				i, pkg.Path)
		}
		if prev, ok := seenPath[clean]; ok {
			return fmt.Errorf("packages[%d] (%s): duplicate path; previously declared at packages[%d]",
				i, pkg.Path, prev)
		}
		seenPath[clean] = i

		if pkg.Package == "" {
			return fmt.Errorf("packages[%d] (%s): package is required (e.g. 'main', 'repo')",
				i, pkg.Path)
		}
		if !isValidGoIdentifier(pkg.Package) {
			return fmt.Errorf("packages[%d] (%s): package %q is not a valid Go identifier",
				i, pkg.Path, pkg.Package)
		}

		if len(pkg.Files) == 0 {
			return fmt.Errorf("packages[%d] (%s): at least one file is required",
				i, pkg.Path)
		}

		seenOutput := make(map[string]int, len(pkg.Files))
		for j, f := range pkg.Files {
			// output is a path relative to the package's directory.
			// Most files use a plain filename ("user.go"); meta-files
			// like Makefile/Dockerfile/compose.yml legitimately walk
			// out via "../../Makefile" to land at the project root.
			//
			// We require: not absolute, and the joined path stays
			// within the project root. The latter prevents
			// pureast.yaml from being a path-traversal vector.
			if f.Output == "" {
				return fmt.Errorf("packages[%d] (%s).files[%d]: output is required",
					i, pkg.Path, j)
			}
			if filepath.IsAbs(f.Output) {
				return fmt.Errorf("packages[%d] (%s).files[%d] (%s): output must be relative, not absolute",
					i, pkg.Path, j, f.Output)
			}
			// Resolve (package-path)/(output) and clean. If the result
			// starts with "..", the output escapes the project root.
			combined := filepath.Clean(filepath.Join(pkg.Path, f.Output))
			if strings.HasPrefix(combined, "..") || combined == ".." {
				return fmt.Errorf("packages[%d] (%s).files[%d] (%s): output escapes project root (resolves to %s)",
					i, pkg.Path, j, f.Output, combined)
			}
			// gen.go is reserved — that's the file scaffold itself produces.
			if filepath.Base(f.Output) == "gen.go" {
				return fmt.Errorf("packages[%d] (%s).files[%d]: 'gen.go' is reserved (it's the file scaffold itself produces)",
					i, pkg.Path, j)
			}
			if prev, ok := seenOutput[combined]; ok {
				return fmt.Errorf("packages[%d] (%s).files[%d] (%s): duplicate resolved output %s; previously declared at files[%d]",
					i, pkg.Path, j, f.Output, combined, prev)
			}
			seenOutput[combined] = j

			if f.Task == "" {
				return fmt.Errorf("packages[%d] (%s).files[%d] (%s): task is required",
					i, pkg.Path, j, f.Output)
			}
			if f.Kind != "" && !validKind(f.Kind) {
				return fmt.Errorf("packages[%d] (%s).files[%d] (%s): invalid kind %q",
					i, pkg.Path, j, f.Output, f.Kind)
			}

			for k, src := range f.Sources {
				if err := src.validate(); err != nil {
					return fmt.Errorf("packages[%d] (%s).files[%d] (%s).sources[%d]: %w",
						i, pkg.Path, j, f.Output, k, err)
				}
			}
		}
	}
	return nil
}

// validate ensures exactly one of Pkg/Module/Symbol/Gomod is set and
// the value is non-empty. Multiple sources per entry would silently
// drop all but one when translated to flags, so we error early.
func (s Source) validate() error {
	set := 0
	if s.Pkg != "" {
		set++
	}
	if s.Module != "" {
		set++
	}
	if s.Symbol != "" {
		set++
	}
	if s.Gomod != "" {
		set++
	}
	switch set {
	case 0:
		return fmt.Errorf("must set exactly one of pkg/module/symbol/gomod")
	case 1:
		// Validate the format of NAME:LOC for symbol entries early.
		if s.Symbol != "" {
			if !strings.Contains(s.Symbol, ":") {
				return fmt.Errorf("symbol must be NAME:LOC, got %q", s.Symbol)
			}
		}
		return nil
	default:
		return fmt.Errorf("set exactly one source kind, got %d", set)
	}
}

// validKind is the same set claude-edit accepts. Kept in sync with
// validDumpKind in cmd/pureast/commands; we don't import that package
// to avoid a cyclic dependency, so the list lives here too.
func validKind(s string) bool {
	switch s {
	case "all", "type", "struct", "interface", "func", "method", "const", "var":
		return true
	}
	return false
}

// isValidGoIdentifier is a permissive check — first char letter or
// underscore, rest letter/digit/underscore. Sufficient to catch
// "package: 1foo" or "package: my-package" type typos.
func isValidGoIdentifier(s string) bool {
	if s == "" {
		return false
	}
	for i, r := range s {
		isLetter := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
		isDigit := r >= '0' && r <= '9'
		if i == 0 {
			if !isLetter {
				return false
			}
		} else if !isLetter && !isDigit {
			return false
		}
	}
	return true
}

// SortedPackages returns the manifest's packages in deterministic order
// (by path). Useful for scaffold output that needs to be reproducible.
func (m *Manifest) SortedPackages() []Package {
	out := make([]Package, len(m.Packages))
	copy(out, m.Packages)
	sort.Slice(out, func(i, j int) bool {
		return out[i].Path < out[j].Path
	})
	return out
}
