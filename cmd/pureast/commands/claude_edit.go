// cmd/pureast/commands/claude_edit.go
//
// `pureast claude-edit` — collaborate with Claude on a Go file via go:generate.
//
// The mental model: it's the copy-paste loop you're already doing with an
// LLM, automated as a `go generate` directive, with a cache so it doesn't
// fire when nothing meaningful upstream has changed.
//
// Typical use in a Go file:
//
//	//go:generate pureast claude-edit \
//	//   --task "implement caching wrapper for UserRepo with 5min TTL" \
//	//   --pkg ../../internal/repo \
//	//   --module github.com/redis/go-redis/v9 \
//	//   --kind interface \
//	//   --output user_cache.go
//
// On `go generate ./...`, this command:
//
//  1. Assembles a context bundle from the requested sources (--gomod,
//     --module, --pkg, --symbol — all repeatable, all composable).
//  2. Computes a SHA-256 over (cache-version, task, context-bundle).
//  3. Reads the existing output file's header. If the cached hash
//     matches, exits without calling Claude — your hand edits are
//     preserved untouched.
//  4. On miss (or when the file doesn't exist), pipes the prompt to
//     `claude -p`, captures stdout, writes the result to the output
//     file with a header recording the new cache key.
//
// What this is NOT: a sandbox. Claude can write whatever it wants to
// the output file when invoked. The cache key only controls *when*
// invocation happens; once it does, you own the result. Review every
// regeneration the same way you'd review a PR — a clear disclaimer is
// embedded in the generated file's header to make this explicit.
package commands

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/Pure-Company/pureast/pkg/cli"
	"github.com/Pure-Company/pureast/pkg/extract"
)

// cacheVersion bumps whenever the cache-key derivation changes. Old
// headers carrying a different version are treated as cache misses,
// which forces a regeneration on the first run after an upgrade. This
// is the standard "schema migration via versioned key" trick.
const cacheVersion = "v1"

// commentPrefixFor returns the line-comment syntax to use for a given
// output file. Pureast's cache header sits at the top of every
// generated file as a small block of comments; the prefix character
// has to match the target file's language or the file is corrupted
// the moment it's written.
//
// We key on file extension and known basenames. Default is `#` because
// the largest set of likely outputs (Make, Docker, YAML, shell,
// Python, Ruby, Terraform, Nginx) all use `#`. Go is the notable
// outlier with `//`. Add other languages as their need arises.
func commentPrefixFor(path string) string {
	base := filepath.Base(path)
	switch {
	case strings.HasSuffix(path, ".go"):
		return "// "
	case base == "Makefile" || base == "makefile" || strings.HasSuffix(path, ".mk"):
		return "# "
	case base == "Dockerfile" || strings.HasSuffix(path, ".dockerfile"):
		return "# "
	case strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml"):
		return "# "
	case strings.HasSuffix(path, ".sh") || strings.HasSuffix(path, ".bash"):
		return "# "
	case strings.HasSuffix(path, ".toml"):
		return "# "
	default:
		// Conservative default: `#` works for most config-language
		// formats. If the user's outputting something exotic (CSS,
		// HTML, SQL with `--` comments), they'll need to add a case
		// or accept that the cache header may be syntactically
		// awkward in their target.
		return "# "
	}
}

// cacheKeyHeaderLine returns the line prefix that readCachedKey
// scans for. Combines the file's comment syntax with the fixed
// pureast-cache-key marker.
func cacheKeyHeaderLine(commentPrefix string) string {
	return commentPrefix + "pureast-cache-key: "
}

// ClaudeEditArgs is the typed argument bundle for the verb. Built from
// cobra flags by parseClaudeEditArgs and consumed by claudeEditAction.
type ClaudeEditArgs struct {
	Task          string
	OutputFile    string
	ClaudeBin     string // path to `claude` binary (defaults to "claude" in PATH)
	ModelOverride string // optional --model passthrough

	// Context sources. All optional, all composable. At least one is
	// required (asking Claude to write code with no context to ground
	// it would be a footgun, so we error early in that case).
	GomodPath string
	Modules   []string // module specs (path[@version]), repeatable
	Pkgs      []string // local package paths, repeatable
	Symbols   []string // NAME:LOC pairs where LOC is a path or module spec

	// Filter flags — apply to every source uniformly.
	Kind         string // all|interface|struct|func|method|const|var|type
	ExportedOnly bool
	MaxTokens    int
	SkipModule   []string
	OnlyModule   []string

	// DryRun prints the assembled prompt and the cache decision but
	// doesn't call Claude or write any file. Useful for debugging the
	// directive in a `go generate -n`-style "what would this do" mode.
	DryRun bool

	// NoCache skips the cache check and always invokes Claude. The
	// header still gets written with the normal (version, task,
	// context) cache key — the override is for THIS invocation only.
	// Designed for use by external orchestrators (weave) that want
	// to retry with extra context without polluting the steady-state
	// cache discipline.
	NoCache bool

	// AppendContext is extra text appended to the prompt's context
	// section. Does NOT influence the cache key (which is computed
	// from the original task + sources only). Typical use: weave's
	// reconcile loop passes the build output as appended context so
	// Claude can fix compile errors. Pureast itself doesn't know or
	// care what's in this string — it's just additional prompt
	// material on this one call.
	AppendContext string
}

func NewClaudeEditCommand() *cobra.Command {
	cmd := cli.NewCommand[ClaudeEditArgs]("claude-edit").
		Short("Collaborate with Claude on a Go file (for use in //go:generate)").
		Long(`claude-edit is the verb you put after //go:generate to let pureast
drive a Claude-assisted edit of a single file.

Context sources (any combination, all repeatable):
  --gomod PATH           every direct dep in a go.mod
  --module SPEC          a single Go module (path@version)
  --pkg PATH             a local package directory
  --symbol NAME:LOC      a single symbol with its transitive deps

Filter flags compose with the sources:
  --kind interface       only this kind of symbol
  --exported             only exported names (default: true)
  --skip-module / --only-module
                         narrow which modules from --gomod
  --max-tokens N         per-source budget

The cache key is sha256(version, task, assembled-context). The file's
header records this key on each generation. Subsequent runs read the
header and skip the Claude call when the key matches — so hand edits
are preserved between regenerations.

You own this file. Pureast handles when Claude is invoked; you handle
review and correctness.`).
		ParseArgs(parseClaudeEditArgs).
		Action(claudeEditAction).
		Build()

	cmd.Flags().String("task", "",
		"Description of what Claude should produce/update in the output file (required).")
	cmd.Flags().StringP("output", "o", "",
		"Output file path (required). The file's first lines record the cache key.")
	cmd.Flags().String("claude-bin", "claude",
		"Path to the claude binary. Default: 'claude' on PATH.")
	cmd.Flags().String("model", "",
		"Optional model override passed through to claude (e.g. claude-opus-4-7).")

	// Context-source flags. StringSlice is repeatable: --module a --module b.
	cmd.Flags().String("gomod", "",
		"Path to a go.mod file; every direct dep is included as context.")
	cmd.Flags().StringSlice("module", nil,
		"Go module spec (path[@version]); repeatable.")
	cmd.Flags().StringSlice("pkg", nil,
		"Local package directory; repeatable.")
	cmd.Flags().StringSlice("symbol", nil,
		"Single symbol as NAME:LOC where LOC is a path or module spec; repeatable.")

	// Filter flags — same names/semantics as the dump verb so muscle
	// memory transfers.
	cmd.Flags().String("kind", "all",
		"Filter: all|interface|struct|func|method|const|var|type.")
	cmd.Flags().Bool("exported", true,
		"Only include exported symbols (default true).")
	cmd.Flags().Int("max-tokens", 0,
		"Per-source max tokens (0 = unbounded).")
	cmd.Flags().StringSlice("skip-module", nil,
		"With --gomod: exclude these modules.")
	cmd.Flags().StringSlice("only-module", nil,
		"With --gomod: include ONLY these modules.")

	cmd.Flags().Bool("dry-run", false,
		"Print the prompt and cache decision; do not call Claude or write the file.")

	cmd.Flags().Bool("no-cache", false,
		"Skip the cache-hit short-circuit; always invoke Claude. The header is still "+
			"written with the normal (version, task, context) cache key — this flag "+
			"only forces re-invocation for THIS call. Designed for orchestrators (e.g. "+
			"weave's reconcile loop) that want to retry with --append-context.")

	cmd.Flags().String("append-context", "",
		"Extra text appended to the prompt's context section for this call. Does NOT "+
			"influence the cache key. Typical use: weave's reconcile loop passes build "+
			"errors so Claude can fix them. Pureast doesn't interpret this string — "+
			"it's just additional prompt material.")

	_ = cmd.MarkFlagRequired("task")
	_ = cmd.MarkFlagRequired("output")

	return cmd
}

func parseClaudeEditArgs(cmd *cobra.Command, args []string) (ClaudeEditArgs, error) {
	if len(args) > 0 {
		return ClaudeEditArgs{}, fmt.Errorf("claude-edit takes no positional arguments (use --task and source flags)")
	}

	task, _ := cmd.Flags().GetString("task")
	out, _ := cmd.Flags().GetString("output")
	claudeBin, _ := cmd.Flags().GetString("claude-bin")
	model, _ := cmd.Flags().GetString("model")

	gomodPath, _ := cmd.Flags().GetString("gomod")
	modules, _ := cmd.Flags().GetStringSlice("module")
	pkgs, _ := cmd.Flags().GetStringSlice("pkg")
	symbols, _ := cmd.Flags().GetStringSlice("symbol")

	kind, _ := cmd.Flags().GetString("kind")
	exported, _ := cmd.Flags().GetBool("exported")
	maxTokens, _ := cmd.Flags().GetInt("max-tokens")
	skip, _ := cmd.Flags().GetStringSlice("skip-module")
	only, _ := cmd.Flags().GetStringSlice("only-module")
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	noCache, _ := cmd.Flags().GetBool("no-cache")
	appendCtx, _ := cmd.Flags().GetString("append-context")

	if !validDumpKind(kind) {
		return ClaudeEditArgs{}, fmt.Errorf(
			"invalid --kind %q (want: all|type|struct|interface|func|method|const|var)", kind)
	}
	// Note: zero context sources is allowed. This is the "seed file"
	// case — defining a new type or interface that doesn't depend on
	// anything beyond stdlib, where the task description is the entire
	// specification. The cache key still works (just hashes task +
	// empty context), and the prompt sent to Claude is well-formed
	// (the Context section becomes empty but the rest of the prompt
	// remains intact).

	// Sort the slice-typed sources so identical directives with
	// different argument orders produce identical hashes. The hash
	// depends on the bytes of the assembled context, and the
	// assembled context iterates these slices in order.
	sort.Strings(modules)
	sort.Strings(pkgs)
	sort.Strings(symbols)

	return ClaudeEditArgs{
		Task:          task,
		OutputFile:    out,
		ClaudeBin:     claudeBin,
		ModelOverride: model,
		GomodPath:     gomodPath,
		Modules:       modules,
		Pkgs:          pkgs,
		Symbols:       symbols,
		Kind:          kind,
		ExportedOnly:  exported,
		MaxTokens:     maxTokens,
		SkipModule:    skip,
		OnlyModule:    only,
		DryRun:        dryRun,
		NoCache:       noCache,
		AppendContext: appendCtx,
	}, nil
}

// claudeEditAction is the orchestrator. The flow:
//
//  1. Build a context bundle from all sources.
//  2. Hash (cacheVersion, task, bundle) -> cache key.
//  3. Read existing output file (if any), parse cached key from header.
//  4. If keys match -> exit, no work, no Claude.
//  5. Otherwise build a prompt, call Claude, write the result with a
//     fresh header carrying the new key.
//
// On any error we leave the existing file untouched. A failed
// regeneration should never leave the file in a worse state than
// before; that's the contract that lets users trust the cache.
func claudeEditAction(ctx context.Context, args ClaudeEditArgs) (cli.Output, error) {
	bundle, err := assembleContext(ctx, args)
	if err != nil {
		return cli.Output{}, fmt.Errorf("assemble context: %w", err)
	}

	key := computeCacheKey(args.Task, bundle)

	// Cache check. If the output file exists and its header carries
	// the same key, the upstream inputs haven't moved — bail out
	// without touching anything. This is the property that makes
	// hand-edits-between-regenerations safe: when nothing's changed,
	// nothing happens.
	//
	// --no-cache bypasses this short-circuit and always invokes
	// Claude. The HEADER still gets written with the normal
	// (version, task, context) cache key, because that's the only
	// identity pureast knows. Any extra context the caller wants in
	// the prompt arrives via --append-context, which is NOT part of
	// the key. This separation lets external orchestrators (weave's
	// reconcile loop) retry with extra prompt material without
	// polluting the steady-state cache discipline.
	if !args.NoCache {
		if cached, ok := readCachedKey(args.OutputFile); ok && cached == key {
			fmt.Fprintf(os.Stderr,
				"pureast claude-edit: cache hit (key %s); %s unchanged\n",
				shortHash(key), args.OutputFile)
			return cli.Output{}, nil
		}
	}

	prompt := buildPrompt(args, bundle)

	if args.DryRun {
		fmt.Fprintln(os.Stderr, "pureast claude-edit: --dry-run; would call claude with:")
		return cli.Output{Text: prompt}, nil
	}

	if args.NoCache {
		fmt.Fprintf(os.Stderr,
			"pureast claude-edit: --no-cache; calling claude (key %s)...\n",
			shortHash(key))
	} else {
		fmt.Fprintf(os.Stderr,
			"pureast claude-edit: cache miss (key %s); calling claude...\n",
			shortHash(key))
	}

	body, err := callClaude(ctx, args, prompt)
	if err != nil {
		return cli.Output{}, fmt.Errorf("claude: %w", err)
	}

	if err := writeWithHeader(args.OutputFile, args.Task, key, body); err != nil {
		return cli.Output{}, fmt.Errorf("write %s: %w", args.OutputFile, err)
	}

	fmt.Fprintf(os.Stderr,
		"pureast claude-edit: wrote %s (%d bytes)\n",
		args.OutputFile, len(body))
	return cli.Output{}, nil
}

// assembleContext builds the deterministic context bundle that feeds
// both the cache key and the Claude prompt. The order is fixed so
// the same logical inputs produce the same byte stream regardless of
// argument order on the command line.
//
// We reuse the existing dump/extract machinery — every source
// produces output in the same format pureast emits today — so users
// who've been driving the workflow with shell pipelines see the
// exact same bytes flowing through a different orchestrator.
func assembleContext(ctx context.Context, args ClaudeEditArgs) (string, error) {
	var sb strings.Builder

	dumpArgsFor := func(filePath string) DumpArgs {
		return DumpArgs{
			FilePath:     filePath,
			Kind:         args.Kind,
			Format:       "go", // canonical format for hashing — render flag is for Claude, not the cache
			Bodies:       false,
			ExportedOnly: args.ExportedOnly,
			IncludeTests: false,
			IncludeDocs:  true,
			MaxTokens:    args.MaxTokens,
		}
	}

	// 1) --gomod (every direct dep). Reuses gomodAction so all the
	//    --skip-module / --only-module / replace-directive logic is
	//    handled in one place.
	if args.GomodPath != "" {
		out, err := gomodAction(ctx, args.GomodPath, dumpArgsFor("."), args.OnlyModule, args.SkipModule)
		if err != nil {
			return "", fmt.Errorf("--gomod %s: %w", args.GomodPath, err)
		}
		// gomodAction already injects per-module headers with module
		// paths inline (// MODULE: github.com/...) — they double as
		// import-path hints for Claude. So we just wrap with the
		// outer source header and don't try to compute a single
		// import path (there are N of them).
		writeSourceHeader(&sb, "--gomod "+args.GomodPath, "")
		sb.WriteString(out.Text)
		sb.WriteString("\n")
	}

	// 2) --pkg (local directories). Sorted upstream by parseClaudeEditArgs.
	for _, pkg := range args.Pkgs {
		abs, err := filepath.Abs(pkg)
		if err != nil {
			return "", fmt.Errorf("--pkg %s: %w", pkg, err)
		}
		out, err := dumpAction(ctx, dumpArgsFor(abs))
		if err != nil {
			return "", fmt.Errorf("--pkg %s: %w", pkg, err)
		}
		// Resolve the import path automatically from the enclosing
		// go.mod. This is the fix for "Claude wrote import \"../repo\""
		// — without this hint, the model has no way to know the
		// canonical module-qualified path. With it, the path is
		// right next to the signatures.
		impPath, _ := extract.ImportPathFor(abs)
		writeSourceHeader(&sb, "--pkg "+pkg, impPath)
		sb.WriteString(out.Text)
		sb.WriteString("\n")
	}

	// 3) --module (remote modules). Each is resolved through
	//    ResolveModule (the same path that --module uses). The
	//    module path doubles as the import path — modulo any
	//    sub-package the user specified, which we already resolved.
	for _, modSpec := range args.Modules {
		res, err := extract.ResolveModule(modSpec)
		if err != nil {
			return "", fmt.Errorf("--module %s: %w", modSpec, err)
		}
		out, err := dumpAction(ctx, dumpArgsFor(res.Dir))
		if err != nil {
			return "", fmt.Errorf("--module %s: %w", modSpec, err)
		}
		impPath := res.ModulePath
		if res.SubPath != "" {
			impPath = strings.TrimSuffix(impPath, "/") + "/" + res.SubPath
		}
		writeSourceHeader(&sb,
			fmt.Sprintf("--module %s (%s@%s)", modSpec, res.ModulePath, res.Version),
			impPath)
		sb.WriteString(out.Text)
		sb.WriteString("\n")
	}

	// 4) --symbol NAME:LOC. LOC may be a local path or a module spec.
	//    We rely on the existing extract command's logic by replicating
	//    just the symbol-resolution piece here — a thin shim.
	for _, symSpec := range args.Symbols {
		name, loc, ok := strings.Cut(symSpec, ":")
		if !ok || name == "" || loc == "" {
			return "", fmt.Errorf("--symbol %q: must be NAME:LOC", symSpec)
		}
		// We don't have a public ExtractSymbolText API today, so we
		// shell out conceptually by reusing dumpAction with the
		// resolved location and letting Claude see the package's
		// signature surface that includes NAME. A future refactor
		// could narrow this to just the symbol + its deps; for now
		// the broader package context is fine and preserves the
		// "use exact types" guarantee.
		var (
			path    string
			impPath string
		)
		if isLikelyModuleSpec(loc) {
			res, err := extract.ResolveModule(loc)
			if err != nil {
				return "", fmt.Errorf("--symbol %s: resolve %s: %w", symSpec, loc, err)
			}
			path = res.Dir
			impPath = res.ModulePath
			if res.SubPath != "" {
				impPath = strings.TrimSuffix(impPath, "/") + "/" + res.SubPath
			}
		} else {
			abs, err := filepath.Abs(loc)
			if err != nil {
				return "", fmt.Errorf("--symbol %s: %w", symSpec, err)
			}
			path = abs
			impPath, _ = extract.ImportPathFor(abs)
		}
		out, err := dumpAction(ctx, dumpArgsFor(path))
		if err != nil {
			return "", fmt.Errorf("--symbol %s: %w", symSpec, err)
		}
		writeSourceHeader(&sb, "--symbol "+symSpec, impPath)
		sb.WriteString(out.Text)
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

// writeSourceHeader emits the section delimiter that introduces each
// context source in the assembled bundle. When importPath is non-empty,
// it's surfaced prominently so Claude uses the right import statement
// instead of guessing (or, worse, emitting a relative-path import that
// isn't even legal Go). The phrasing is direct because models read it.
func writeSourceHeader(sb *strings.Builder, sourceLabel, importPath string) {
	sb.WriteString("// === SOURCE: ")
	sb.WriteString(sourceLabel)
	sb.WriteString(" ===\n")
	if importPath != "" {
		sb.WriteString("// IMPORT PATH: ")
		sb.WriteString(importPath)
		sb.WriteString("\n")
		sb.WriteString("// (when importing this package, use the path above — NOT a relative path)\n")
	}
}

// isLikelyModuleSpec mirrors the heuristic in claude-with-stdlib: if
// the first path segment contains a dot (github.com, gopkg.in,
// golang.org), treat it as a remote module; otherwise as a local
// path or stdlib name. Strips an optional @version suffix first.
//
// We reject specs that start with a path-like prefix (".", "/", "~")
// outright, because those are unambiguously local paths even when
// they contain dots — `./local/pkg` and `~/code/foo.bar` should not
// be misread as remote module specs.
func isLikelyModuleSpec(spec string) bool {
	if strings.HasPrefix(spec, ".") || strings.HasPrefix(spec, "/") || strings.HasPrefix(spec, "~") {
		return false
	}
	if i := strings.LastIndex(spec, "@"); i > 0 {
		spec = spec[:i]
	}
	first, _, _ := strings.Cut(spec, "/")
	return strings.Contains(first, ".")
}

// computeCacheKey hashes (cacheVersion, task, bundle) into a stable
// identifier. The version prefix means future cache-format changes
// invalidate old entries cleanly without colliding with old hashes.
//
// The hash does NOT include the current file contents. This is
// deliberate: hand edits to the output file should NOT trigger
// regeneration — only changes to the inputs (task or context) should.
// That's the property that makes the cache trustworthy for users who
// edit the file between regenerations.
func computeCacheKey(task, bundle string) string {
	h := sha256.New()
	h.Write([]byte(cacheVersion))
	h.Write([]byte{0})
	h.Write([]byte(task))
	h.Write([]byte{0})
	h.Write([]byte(bundle))
	return cacheVersion + ":" + hex.EncodeToString(h.Sum(nil))
}

func shortHash(key string) string {
	// "v1:abc123def456..." → "v1:abc123de" — enough to disambiguate
	// at a glance in CI logs without dumping 64 hex chars.
	if i := strings.Index(key, ":"); i >= 0 && len(key) > i+10 {
		return key[:i+10]
	}
	return key
}

// readCachedKey scans the first ~30 lines of `path` for the cache-key
// header line. The line prefix depends on the file's comment syntax
// (`// pureast-cache-key:` for Go, `# pureast-cache-key:` for
// Makefile/Dockerfile/YAML/etc.).
//
// Returns (key, true) on a hit and ("", false) on any error or
// absence — both mean "treat as cache miss." Reading more than a
// handful of lines guards against headers that might've grown in
// some future version while keeping the read bounded.
func readCachedKey(path string) (string, bool) {
	f, err := os.Open(path)
	if err != nil {
		return "", false
	}
	defer f.Close()
	prefix := cacheKeyHeaderLine(commentPrefixFor(path))
	scanner := bufio.NewScanner(f)
	for i := 0; i < 30 && scanner.Scan(); i++ {
		line := scanner.Text()
		if strings.HasPrefix(line, prefix) {
			return strings.TrimSpace(strings.TrimPrefix(line, prefix)), true
		}
	}
	return "", false
}

// buildPrompt assembles what gets piped to `claude -p`. The structure
// is intentionally simple — a heading, the context, the current file
// (if any), the task, and a strict "output only Go code" instruction.
//
// The "current file" inclusion is what makes hand edits survive a
// regeneration: when something upstream changes and Claude has to
// rewrite the file, it sees the user's current state and is asked to
// preserve it where the directive doesn't explicitly call for a
// change. Not a guarantee — Claude is an LLM — but the prompt nudges
// hard, and the user reviews the diff before merging.
// isGoTarget reports whether the output file is a Go source file.
// Used to switch prompt instructions and the post-call sanity check
// — Go targets demand `package <name>` at the top; non-Go targets
// (Makefile, Dockerfile, YAML) have no such constraint.
func isGoTarget(path string) bool {
	return strings.HasSuffix(path, ".go")
}

func buildPrompt(args ClaudeEditArgs, bundle string) string {
	var sb strings.Builder
	goTarget := isGoTarget(args.OutputFile)

	sb.WriteString("# pureast claude-edit\n\n")
	if goTarget {
		sb.WriteString("You are editing a single Go source file. Below are:\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("You are editing %s. Below are:\n\n",
			filepath.Base(args.OutputFile)))
	}
	sb.WriteString("1. Context: signatures from project / module sources, extracted by pureast.\n")
	sb.WriteString("2. The current contents of the file you're editing (may be empty on first run).\n")
	sb.WriteString("3. The task: what to produce or change.\n\n")

	if strings.TrimSpace(bundle) == "" {
		sb.WriteString("## Context\n\n")
		sb.WriteString("(no external context — produce the file from the task description alone.)\n\n")
	} else {
		sb.WriteString("## Context (signatures only — use these EXACT types where applicable)\n\n")
		sb.WriteString("```go\n")
		sb.WriteString(strings.TrimRight(bundle, "\n"))
		sb.WriteString("\n```\n\n")
	}

	current, err := os.ReadFile(args.OutputFile)
	if err == nil && len(current) > 0 {
		sb.WriteString("## Current file contents\n\n")
		sb.WriteString("```\n")
		sb.WriteString(strings.TrimRight(string(current), "\n"))
		sb.WriteString("\n```\n\n")
		sb.WriteString("PRESERVE existing hand edits where the task does not explicitly call for changes. ")
		sb.WriteString("Add to or refine the file rather than rewriting it from scratch.\n\n")
	} else {
		sb.WriteString("## Current file contents\n\n(file does not exist yet — produce a fresh one)\n\n")
	}

	// Optional caller-supplied context. Used by orchestrators (weave's
	// reconcile loop) to feed build errors back into the prompt. We
	// don't interpret the string — it just becomes another section
	// the model can read. The section is omitted when the flag is
	// empty, so normal claude-edit invocations are unaffected.
	if args.AppendContext != "" {
		sb.WriteString("## Additional context\n\n")
		sb.WriteString(strings.TrimRight(args.AppendContext, "\n"))
		sb.WriteString("\n\n")
		sb.WriteString("If the additional context above describes a problem (e.g. a build error, ")
		sb.WriteString("a failing test, a lint finding), assess whether it relates to THIS file. ")
		sb.WriteString("If yes, fix it while preserving the rest of the file. If no, return the ")
		sb.WriteString("file essentially unchanged.\n\n")
	}

	sb.WriteString("## Task\n\n")
	sb.WriteString(args.Task)
	sb.WriteString("\n\n")

	sb.WriteString("## Output format requirements (CRITICAL)\n\n")
	if goTarget {
		sb.WriteString("Output ONLY raw Go source code. The entire response must be valid compilable Go.\n\n")
		sb.WriteString("Required:\n")
		sb.WriteString("- First non-comment line: `package <name>`\n")
		sb.WriteString("- Imports in a single block at the top\n")
		sb.WriteString("- Use only the exact types and signatures shown in the Context section\n\n")
		sb.WriteString("Forbidden:\n")
		sb.WriteString("- NO markdown fences (no ```go or ```)\n")
		sb.WriteString("- NO prose, summary, commentary before/after/between code\n")
		sb.WriteString("- NO 'Here is...', 'I've created...', '## Summary' or similar\n\n")
		sb.WriteString("Match idiomatic Go: error wrapping, context propagation, zero values.\n")
	} else {
		sb.WriteString(fmt.Sprintf(
			"Output ONLY the raw contents of %s. The entire response must be valid for that file format.\n\n",
			filepath.Base(args.OutputFile)))
		sb.WriteString("Forbidden:\n")
		sb.WriteString("- NO markdown fences (no ``` of any kind)\n")
		sb.WriteString("- NO prose, summary, commentary before/after/between content\n")
		sb.WriteString("- NO 'Here is...', 'I've created...', '## Summary' or similar\n\n")
		sb.WriteString("Use the file format's native comment syntax (# for Make/Docker/YAML, // for Go).\n")
	}

	return sb.String()
}

// callClaude shells out to the configured `claude` binary in -p mode,
// pipes the prompt to its stdin, and captures stdout. Stderr from
// claude is forwarded to our stderr so users see authentication
// problems, rate limits, etc. without us having to interpret them.
//
// We don't try to use any Anthropic SDK directly. Shelling out keeps
// auth, model selection, and account quota all the user's
// responsibility — Claude Code already manages those, and pureast
// shouldn't try to second-guess them.
func callClaude(ctx context.Context, args ClaudeEditArgs, prompt string) (string, error) {
	goTarget := isGoTarget(args.OutputFile)
	var sysPrompt string
	if goTarget {
		sysPrompt = "Output ONLY raw Go source code starting with 'package'. " +
			"No markdown, no prose, no summary. The entire response is a valid .go file."
	} else {
		sysPrompt = fmt.Sprintf("Output ONLY the raw contents of %s. "+
			"No markdown fences, no prose, no summary. The entire response IS the file content.",
			filepath.Base(args.OutputFile))
	}

	claudeArgs := []string{"-p", sysPrompt}
	if args.ModelOverride != "" {
		claudeArgs = append(claudeArgs, "--model", args.ModelOverride)
	}

	cmd := exec.CommandContext(ctx, args.ClaudeBin, claudeArgs...)
	cmd.Stdin = strings.NewReader(prompt)
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("start %s: %w", args.ClaudeBin, err)
	}

	body, err := io.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("read claude stdout: %w", err)
	}
	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("%s: %w", args.ClaudeBin, err)
	}

	clean := stripCodeFences(string(body))
	if !looksLikeContent(clean, goTarget) {
		// Don't write garbage to disk. The user can re-run after
		// fixing the prompt; meanwhile their existing file is
		// untouched (we haven't written anything yet).
		preview := clean
		if len(preview) > 200 {
			preview = preview[:200] + "...[truncated]"
		}
		expectation := "code (likely a summary instead)"
		if goTarget {
			expectation = "Go code starting with 'package' (likely a summary instead)"
		}
		return "", fmt.Errorf("claude returned content that doesn't look like %s; preview:\n%s",
			expectation, preview)
	}
	return clean, nil
}

// stripCodeFences removes leading/trailing ``` lines if Claude
// wrapped its output in markdown despite the instructions. -p mode
// usually obeys but occasionally relapses; this is the cheap defense.
func stripCodeFences(s string) string {
	lines := strings.Split(s, "\n")
	var out []string
	for _, ln := range lines {
		t := strings.TrimSpace(ln)
		if strings.HasPrefix(t, "```") {
			continue
		}
		out = append(out, ln)
	}
	return strings.TrimSpace(strings.Join(out, "\n")) + "\n"
}

// looksLikeContent is a sanity check: did Claude actually return
// file content, or did it slip into summary mode and produce prose?
//
// For Go targets, we require `package <name>` near the top — that's
// the cheapest reliable signal that we got code.
//
// For non-Go targets, we can't be as strict (Makefile/Dockerfile/YAML
// have no single mandatory leading token), so we apply a weaker
// heuristic: reject obvious prose openers like "Here is", "I've
// created", or markdown headings ("## Summary"). Anything else passes.
func looksLikeContent(s string, goTarget bool) bool {
	if goTarget {
		return looksLikeGo(s)
	}
	// Non-Go: reject if the first non-empty line looks like prose
	// or a markdown heading.
	for _, line := range strings.SplitN(s, "\n", 5) {
		t := strings.TrimSpace(line)
		if t == "" {
			continue
		}
		lower := strings.ToLower(t)
		switch {
		case strings.HasPrefix(t, "## "), strings.HasPrefix(t, "# Summary"), strings.HasPrefix(t, "# Output"):
			// Markdown heading openers — likely a summary, not the file content.
			return false
		case strings.HasPrefix(lower, "here is"), strings.HasPrefix(lower, "here's "),
			strings.HasPrefix(lower, "i've created"), strings.HasPrefix(lower, "i have created"),
			strings.HasPrefix(lower, "this is the"):
			return false
		}
		break
	}
	return strings.TrimSpace(s) != ""
}

// looksLikeGo is the Go-specific check used inside looksLikeContent.
// "package " near the top is the cheapest reliable signal.
func looksLikeGo(s string) bool {
	for i, line := range strings.SplitN(s, "\n", 50) {
		_ = i
		t := strings.TrimSpace(line)
		if t == "" || strings.HasPrefix(t, "//") {
			continue
		}
		return strings.HasPrefix(t, "package ")
	}
	return false
}

// writeWithHeader prepends the cache header to body and writes the
// combined content atomically (write tmp, rename) so a Ctrl-C in the
// middle doesn't corrupt the existing file.
//
// The header's comment prefix is derived from the output path's
// extension (// for .go, # for Makefile/Dockerfile/yaml/etc.) so the
// generated file remains syntactically valid in its target language.
func writeWithHeader(path, task, key, body string) error {
	hdr := buildHeader(task, key, commentPrefixFor(path))

	full := hdr + body
	if !strings.HasSuffix(full, "\n") {
		full += "\n"
	}

	dir := filepath.Dir(path)
	if dir == "" {
		dir = "."
	}
	tmp, err := os.CreateTemp(dir, ".pureast-claude-edit-*")
	if err != nil {
		return fmt.Errorf("create temp: %w", err)
	}
	tmpPath := tmp.Name()
	if _, err := tmp.WriteString(full); err != nil {
		tmp.Close()
		os.Remove(tmpPath)
		return fmt.Errorf("write temp: %w", err)
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpPath)
		return fmt.Errorf("close temp: %w", err)
	}
	if err := os.Rename(tmpPath, path); err != nil {
		os.Remove(tmpPath)
		return fmt.Errorf("rename temp -> %s: %w", path, err)
	}
	return nil
}

func buildHeader(task, key, commentPrefix string) string {
	// Multi-line preamble. The cache-key line must be machine-findable
	// (readCachedKey looks for cacheKeyHeaderLine(commentPrefix)); the
	// rest is human-facing context.
	cp := commentPrefix
	var sb strings.Builder
	sb.WriteString(cp + "Code generated by `pureast claude-edit`. DO NOT REGENERATE BY HAND;\n")
	sb.WriteString(cp + "instead, edit the directive that produced this file.\n")
	sb.WriteString(cp + "pureast-cache-key: " + key + "\n")
	sb.WriteString(cp + "pureast-task: " + strings.ReplaceAll(task, "\n", " ") + "\n")
	sb.WriteString(cp + "pureast-generated: " + time.Now().UTC().Format(time.RFC3339) + "\n")
	sb.WriteString(strings.TrimRight(cp, " ") + "\n")
	sb.WriteString(cp + "Hand edits between regenerations are preserved: pureast only\n")
	sb.WriteString(cp + "re-invokes Claude when the cache key (task + context inputs)\n")
	sb.WriteString(cp + "changes. You own this file. Review every diff.\n")
	sb.WriteString("\n")
	return sb.String()
}
