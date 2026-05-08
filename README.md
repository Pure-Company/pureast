# 🧬 **PureAST**

[![Go Reference](https://pkg.go.dev/badge/github.com/Pure-Company/pureast.svg)](https://pkg.go.dev/github.com/Pure-Company/pureast)
[![Go Report Card](https://goreportcard.com/badge/github.com/Pure-Company/pureast)](https://goreportcard.com/report/github.com/Pure-Company/pureast)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Pure-Company/pureast)](https://github.com/Pure-Company/pureast/blob/main/go.mod)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/Pure-Company/pureast/pulls)

**Functorial AST Extraction and Code Generation Toolkit**

🧠 The MIT + Functorial Brain-Melt License (FBML)
*Read the [LICENSE](./LICENSE) file.*

---

## 🚀 Overview

**PureAST** is a **pure functional** Go library and CLI tool for
**AST extraction, dependency analysis, and LLM-assisted code generation**
— all built on top of **PureKernels** (monoids, folds, and functors).

Two halves:

* **Read your code.** Extract structs, interfaces, functions; analyze
  dependencies; compress dependencies into compact LLM-ready bundles
  (`--gomod`, `--module`); generate dependency graphs.

* **Write your code.** Generate source files, Makefiles, Dockerfiles, and
  compose files from a single YAML manifest (`scaffold` + `claude-edit`),
  with a content-addressed cache that preserves hand edits between
  regenerations.

Everything composes via **monoids and folds**, keeping it **pure,
deterministic, and composable**.

---

## 🧩 Core Concepts

| Concept                     | Description                                                                       |
| --------------------------- | --------------------------------------------------------------------------------- |
| **AST as Category**         | Each node is an object, and dependency edges are morphisms.                       |
| **Monoid Composition**      | Dependencies, files, and packages all compose using monoidal laws.                |
| **Functorial Extraction**   | Visitors and traversals are pure functions `(Deps → Deps)`.                       |
| **Concurrent Applicatives** | Directory extraction leverages applicative concurrency via `purekernels/functor`. |

---

## 📦 Installation

```bash
go install github.com/Pure-Company/pureast/cmd/pureast@latest
```

Or clone locally:

```bash
git clone https://github.com/Pure-Company/pureast.git
cd pureast
go build ./cmd/pureast
```

---

## ⚙️ Usage

PureAST uses subcommands. Each verb does one thing:

```bash
pureast dump ./examples/app                 # every symbol, signatures only — LLM context
pureast extract UserService ./examples/app  # one symbol with transitive deps
pureast deps Profile ./examples/app         # dependency analysis
pureast diff main ./examples/app            # symbols in files changed since main
pureast search "Handler" ./examples/app     # fuzzy symbol search
pureast list ./examples/app                 # enumerate all symbols
pureast types ./examples/app                # [deprecated] use 'dump --kind' instead
```

### Verbs

| Verb           | Purpose                                                        |
| -------------- | -------------------------------------------------------------- |
| `dump`         | Compact dump of every symbol — the LLM-context flagship.       |
| `extract`      | Extract one symbol plus its transitive dependencies.           |
| `deps`         | Forward or reverse dependency analysis (`--reverse`, `--locations`).|
| `diff`         | Symbols whose lines changed since a git ref (PR-review use).   |
| `search`       | Fuzzy search for symbols by name pattern.                      |
| `list`         | Enumerate all symbols in a package, optionally grouped.        |
| `claude-edit`  | LLM-driven codegen with a content-addressed cache. Designed for `//go:generate`. |
| `scaffold`     | Materialize a project skeleton from a YAML manifest.           |
| `types`        | **Deprecated**: use `dump --kind struct\|interface` instead.    |

Run `pureast <verb> --help` for the flags specific to each verb.

### Module modes

In addition to local paths, pureast can resolve **any public Go module**
on demand. These flags work with every verb (and with bare `pureast`):

| Flag        | Purpose                                                                |
| ----------- | ---------------------------------------------------------------------- |
| `--module`  | Dump a single Go module by path (uses `go mod download` under the hood).|
| `--gomod`   | Dump every direct dependency in a `go.mod` file. Indirect deps are skipped.|

```bash
# Single remote module — no clone, no manual setup
pureast --module github.com/spf13/cobra
pureast --module github.com/gin-gonic/gin@v1.10.0
pureast --module github.com/spf13/cobra/doc          # sub-package

# Every direct dependency in a project, in one shot
pureast --gomod ./go.mod
pureast --gomod ./go.mod --kind interface            # contracts only
pureast --gomod ./go.mod --skip-module github.com/aws/aws-sdk-go-v2
pureast --gomod ./go.mod --only-module github.com/redis/go-redis/v9
```

Module resolution piggybacks on `go mod download`, so `GOPROXY`,
`GOPRIVATE`, vendoring, and `replace` directives all work the way Go
already configures them. The first call hits the network; subsequent
calls for the same `(module, version)` are filesystem-cache hits.

---

## 🧠 Example Workflows

### 🧠 Compact context for an LLM

```bash
pureast dump ./examples/app -o context.txt           # signatures only
pureast dump ./examples/app --bodies -o full.txt     # include implementations
pureast dump ./examples/app --kind interface         # only interfaces
pureast dump ./examples/app --exported               # only exported symbols
pureast dump ./examples/app --max-tokens 4000        # fit a token budget
```

`--max-tokens` is symbol-aware: it drops trailing whole declarations
rather than slicing through one, so the output is always syntactically
complete Go (matters when the LLM re-parses it). Available on `dump`,
`extract`, and `diff`.

### 🪐 Context compression for an entire project (`--gomod`)

`--gomod` reads your project's `go.mod`, dumps every direct dependency
in one pass, and emits a single LLM-ready bundle. Indirect deps are
skipped automatically — what you get is your project's *actual API
surface*, not its transitive plumbing.

```bash
# Compress every direct dep in the project — typical result is 100×
# smaller than raw source while preserving the full callable surface.
pureast --gomod ./go.mod --kind interface > deps.md

# Pipe straight into Claude Code (or any LLM CLI that reads stdin)
pureast --gomod ./go.mod --exported | \
  claude -p "review these dependencies for security advisories or staleness"

# Generate a feature against the actual dep contracts — no hallucinated APIs
pureast --gomod ./go.mod --exported | \
  claude -p "implement a JWT-validated session cached in Redis with 15min TTL"

# Narrow the context — whitelist or blacklist specific deps
pureast --gomod ./go.mod --only-module github.com/redis/go-redis/v9
pureast --gomod ./go.mod --skip-module github.com/spf13/cobra
```

A typical Go service has 5–15 direct deps totaling millions of lines
of source. Through `pureast --gomod --kind interface`, that same
surface area lands at ~10k tokens — small enough for one prompt, large
enough that the LLM can generate code grounded in real signatures
rather than guesses from training data.

### 🧱 Extract a symbol with its dependencies

```bash
pureast extract UserService ./examples/app -o service.go
pureast extract Profile ./examples/app --minimal     # direct deps only
```

### 🧾 Dependency analysis

Forward (what does X depend on):

```bash
pureast deps Profile ./examples/app                  # text report (default)
pureast deps Profile ./examples/app --depth 1        # one-hop only
pureast deps Profile ./examples/app --locations      # show file:line on each entry
pureast deps Profile ./examples/app --format dot > graph.dot
dot -Tpng graph.dot -o graph.png
```

Reverse (who depends on X — impact analysis before a refactor):

```bash
pureast deps User ./examples/app --reverse           # every caller, with file:line
pureast deps User ./examples/app --reverse --depth 1 # direct callers only
pureast deps User ./examples/app --reverse --format json
```

`--reverse` always emits file:line locations (jump-to-source is the
point); for forward queries `--locations` is opt-in. JSON output is
deterministic — alphabetically sorted, fixed field order — so it
caches well in LLM prompts.

Errors go to stderr; successful output goes to stdout. Pipes stay
clean: `pureast deps Foo ./pkg --format json | jq .` either parses
or prints nothing on failure, never both.

### 🔀 PR review with `diff`

`diff` extracts only the symbols whose lines actually changed since
a git ref — useful for feeding an LLM the relevant context for a
code review without paying for the rest of the repo.

```bash
pureast diff main                                    # changes in current dir
pureast diff main ./examples/app                     # scoped to a path
pureast diff HEAD~5 --bodies                         # include implementations
pureast diff origin/main --format md -o pr.md        # markdown for context windows
pureast diff main --max-tokens 8000                  # fit a budget
```

Filtering is hunk-level by default (only symbols whose line range
intersects a changed hunk). Use `--whole-file` if you want every
symbol from any modified file — heavier output, useful when the
diff has wide-reaching refactors.

### 🔍 Find a symbol

```bash
pureast search "Handler" ./examples/app
pureast list ./examples/app                          # grouped by kind (default)
pureast list ./examples/app --grouped=false          # flat list
```

---

## 🛠 Project codegen — `scaffold` + `claude-edit`

PureAST also drives **LLM-assisted code generation** through two cooperating
verbs that fit Go's existing `//go:generate` workflow. The split:

* **`scaffold`** turns a YAML manifest into a tree of `gen.go` files. Pure
  structure, no LLM. Deterministic — same manifest always produces the same
  directives.
* **`claude-edit`** runs inside each `gen.go`'s `//go:generate` directive,
  feeds Claude the project's actual API surface (via `--pkg`, `--module`,
  `--gomod`, `--symbol`), and writes the resulting source file with a
  content-addressed cache header.

Two phases, two cache disciplines, both reviewable as ordinary diffs.

```bash
pureast scaffold                              # phase 1: structure (no LLM)
go generate -tags ignore ./...                # phase 2: content (LLM, cached)
```

The `-tags ignore` is required because each `gen.go` is built with
`//go:build ignore` so it never compiles into the final binary — only
`go generate` sees it, and only when the tag is requested.

### `claude-edit` — LLM codegen with a cache that respects hand edits

`claude-edit` is the verb you put after `//go:generate`:

```go
//go:generate pureast claude-edit --task "..." --pkg ../repo --module github.com/redis/go-redis/v9 --kind interface --output user_cache.go
```

It assembles a context bundle from one or more sources, computes
`sha256(version, task, context)`, and only invokes Claude when that key
differs from the one stored in the output file's header. Hand edits are
preserved by the cache: when nothing upstream has changed, regeneration
is a no-op and your edits stay.

| Flag                | Purpose                                                                |
| ------------------- | ---------------------------------------------------------------------- |
| `--task TEXT`       | What you want produced. Required.                                      |
| `--output PATH`     | Where to write. Comment syntax adapts: `//` for `.go`, `#` for Make/Docker/YAML. |
| `--pkg PATH`        | Local package as context (repeatable). Import path injected automatically from `go.mod`. |
| `--module SPEC`     | Remote module as context (repeatable). Resolved via `go mod download`. |
| `--gomod PATH`      | Every direct dep in a `go.mod`.                                        |
| `--symbol NAME:LOC` | A single symbol with its deps (repeatable).                            |
| `--kind`, `--exported`, `--max-tokens`, `--skip-module`, `--only-module` | Filters apply to all sources. |
| `--model NAME`      | Pass-through to `claude -p --model`. Use `opus` for the best results.  |
| `--dry-run`         | Print the assembled prompt; don't call Claude.                         |

Direct CLI use (no `gen.go` required):

```bash
pureast claude-edit \
  --task "Implement Cache wrapping repo.UserRepo: Redis-backed, 5min TTL." \
  --pkg ../repo \
  --module github.com/redis/go-redis/v9 \
  --kind interface \
  --output user_cache.go \
  --model opus
```

Run twice; the second run cache-hits and exits without calling Claude. Edit
the file by hand; the third run still cache-hits because nothing upstream
changed. Change the task or the underlying interface; the cache misses and
Claude regenerates.

The model occasionally produces code that doesn't compile — scope errors,
wrong import paths, missing braces. PureAST does not run `go build` on the
output. **You own the file. Review every regeneration as you would any PR
diff, and hand-edit fixes when needed.** The cache then preserves your
fixes until something upstream actually moves.

### `scaffold` — a YAML manifest is the source of truth

`scaffold` reads `pureast.yaml` (or another file via `--manifest`) and
produces one `gen.go` per package described:

```yaml
# pureast.yaml at the project root
module: github.com/example/shortener   # optional sanity-check vs go.mod

packages:
  - path: internal/domain
    package: domain
    doc: Core domain types — the seed of the project.
    files:
      - output: link.go
        task: "Define a Link struct with ID, LongURL, CreatedAt..."
        model: opus

  - path: internal/repo
    package: repo
    files:
      - output: link_repo.go
        task: "Define LinkRepo interface for persisting domain.Link..."
        sources:
          - pkg: ../domain
        kind: interface
        model: opus

      - output: postgres.go
        task: "PostgresLinkRepo implementing LinkRepo using pgxpool.Pool..."
        sources:
          - pkg: .
          - pkg: ../domain
          - module: github.com/jackc/pgx/v5/pgxpool
        model: opus

  # Build files — one tools/gen package emits Makefile, Dockerfile, compose
  # at the project root. --output walks up via "../../" and the cache
  # header uses '#' instead of '//'.
  - path: tools/gen
    package: gen
    files:
      - output: ../../Makefile
        task: "Makefile with build/test/lint/run/clean targets..."
        sources:
          - gomod: ../../go.mod
        kind: interface
        model: opus

      - output: ../../Dockerfile
        task: "Multi-stage Dockerfile, golang:1.22-alpine builder..."
        sources:
          - gomod: ../../go.mod
        kind: interface
        model: opus

      - output: ../../docker-compose.yml
        task: "docker-compose.yml for local dev with postgres, redis..."
        sources:
          - gomod: ../../go.mod
        kind: interface
        model: opus
```

Run scaffold:

```bash
pureast scaffold
#   created  cmd/shortener/gen.go
#   created  internal/domain/gen.go
#   created  internal/repo/gen.go
#   created  tools/gen/gen.go
#   ...
```

Each `gen.go` is build-ignored and contains a `//go:generate
pureast claude-edit ...` directive per output file. Validation catches
common mistakes before any code runs (missing tasks, paths escaping the
project root, duplicate outputs, `gen.go` reserved name, invalid Go
identifiers, unknown YAML fields like `pacakge:`).

### End-to-end workflow

```bash
# 0. one go.mod and one pureast.yaml at the project root
cat > go.mod <<EOF
module github.com/example/myproject
go 1.22
EOF
# (author pureast.yaml describing your packages and their generation tasks)

# 1. lay down the structure
pureast scaffold

# 2. fill in the source files
go generate -tags ignore ./...

# 3. build files often need a populated go.mod, so run the standard tidy
go mod tidy

# 4. regenerate the build files now that go.mod has direct deps
go generate -tags ignore ./tools/...

# 5. compile, hand-fix any LLM glitches, ship
go build ./...
```

The whole pipeline is reproducible. Every artifact (`pureast.yaml`,
each `gen.go`, every generated source file with its cache header)
lives in version control. The only non-deterministic step — the
Claude call — is gated by a cache so that re-running on a clean
checkout produces no API calls when inputs are unchanged.

### What this is *not*

* Not a substitute for code review. Generated files commit like any
  other source; reviewers see the diff and approve or request changes.
* Not a closed-loop "make it compile" tool. PureAST writes whatever
  Claude returns (after stripping markdown fences) and lets the user
  catch compile errors with `go build`. A multi-agent retry loop is
  out of scope for this project.
* Not Go-only on the output side. `claude-edit` adapts comment syntax
  for Makefile, Dockerfile, YAML, shell, TOML, and similar text
  formats — any text file becomes a generation target.

---

## 🤖 Use with Claude (MCP)

PureAST ships with an [MCP](https://modelcontextprotocol.io) server so Claude
(Desktop, Code, or any MCP-compatible client) can call it directly during a
conversation. Build the server, point Claude at it, and the LLM gets on-demand
access to your codebase as compact, deterministic context.

```bash
go build -o /usr/local/bin/pureast-mcp ./cmd/pureast-mcp
```

Add to `claude_desktop_config.json` (or your client's equivalent):

```json
{
  "mcpServers": {
    "pureast": {
      "command": "/usr/local/bin/pureast-mcp",
      "env": { "PUREAST_WORKERS": "8" }
    }
  }
}
```

The server exposes eight tools:

| Tool                | Use case                                                         |
| ------------------- | ---------------------------------------------------------------- |
| `dump_package`      | Compact orientation map for an unfamiliar package (LLM flagship) |
| `extract_symbol`    | Pull one symbol with its transitive dependencies                 |
| `reverse_deps`      | "Who uses X?" — impact analysis before refactors                 |
| `show_dependencies` | "What does X depend on?" — forward graph                         |
| `diff_since`        | PR-review context: only what changed since a git ref             |
| `search_symbols`    | Fuzzy symbol search                                              |
| `list_symbols`      | Enumerate everything                                             |
| `extract_types`     | Type declarations only                                           |

`dump_package`, `extract_symbol`, and `diff_since` accept `format: "md"` for
markdown-fenced output and `maxTokens: N` to fit a budget — useful when feeding
context-bounded models.

---

## 🧮 Architecture

```
pureast/
├── cmd/pureast/          # CLI entry point + subcommands
│   └── commands/         # one file per verb (dump, extract, deps, diff,
│                         #   search, list, claude-edit, scaffold, types)
├── cmd/pureast-mcp/      # MCP server entry point
├── pkg/
│   ├── ast/              # Core AST node structures & visitors
│   ├── extract/          # AST extraction (pure functional)
│   ├── analyze/          # Dependency graph analysis
│   ├── codegen/          # Code & report generation
│   ├── scaffold/         # Manifest schema + scaffolder for project skeletons
│   ├── cli/              # Type-safe command builder
│   ├── mcp/              # MCP server: tool handlers + JSON-RPC protocol
│   └── ...               # Composition via monoids and folds
└── examples/             # Sample Go packages for testing
```

---

## 🧰 Built On

* 🧱 [PureKernels](https://github.com/Pure-Company/purekernels):
  composable monoids, folds, functors, and concurrency primitives.
* 🧩 Go standard library (`go/ast`, `go/parser`, `token`, `format`).
* ⚙️ Functional composition patterns (`FoldLeft`, `Map`, `Reduce`).
* 🐍 [spf13/cobra](https://github.com/spf13/cobra) for the subcommand surface.

---

## 💡 Philosophy

PureAST treats Go's syntax tree as a **mathematical object**:

> *ASTs compose, dependencies fold, and code emerges through monoidal harmony.*

It's not a parser — it's a **categorical extractor**.

---

## 🧠 License

🧠 The MIT + Functorial Brain-Melt License (FBML)
See the [LICENSE](./LICENSE) file for details.
