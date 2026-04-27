# PureAST CLI Reference

PureAST extracts Go symbols, analyzes dependencies, and emits compact
representations suitable for feeding to an LLM as context.

The CLI is organized as seven verbs. Each one does one thing.

```
pureast dump      # every symbol, signatures only — LLM context flagship
pureast extract   # one symbol with its transitive dependencies
pureast deps      # what does a symbol depend on (or who depends on it)
pureast diff      # symbols in files changed since a git ref
pureast search    # fuzzy symbol search
pureast list      # enumerate all symbols
pureast types     # [deprecated] use 'dump --kind' instead
```

Path is always positional and defaults to the current directory:

```bash
pureast dump          # equivalent to: pureast dump .
pureast dump ./pkg    # explicit path
pureast dump --file ./pkg   # deprecated, still works, warns to stderr
```

---

## Global conventions

| Convention                | Notes                                                                  |
| ------------------------- | ---------------------------------------------------------------------- |
| `[PATH]`                  | Always positional, optional, defaults to `.`                           |
| `-o FILE`, `--output FILE`| Write to a file instead of stdout. Available wherever output is bulky. |
| `--format go\|md`         | Where applicable. `go` = raw, `md` = fenced markdown for LLM contexts. |
| `--max-tokens N`          | Where applicable. Line-aware truncation to fit a budget.               |
| `--file PATH`             | **Deprecated** alias for the positional `[PATH]`. Warns to stderr.     |

`--max-tokens` uses a 3.5-chars-per-token approximation (well-known
OpenAI guidance). Pipe through your model's actual tokenizer for exact
counts. Truncation is line-aware with a `// ... truncated to fit token
budget ...` marker so output stays valid Go.

---

## `pureast dump` — every symbol

Walks the package, emits every top-level symbol (struct, interface,
type alias, func, method, const, var). Default output is signatures
only. This is the verb to reach for when feeding a package to an LLM.

```
pureast dump [PATH] [flags]
```

| Flag             | Default | Effect                                                             |
| ---------------- | ------- | ------------------------------------------------------------------ |
| `--kind K`       | `all`   | Filter: `all\|type\|struct\|interface\|func\|method\|const\|var` |
| `--bodies`       | off     | Include function bodies (uses `go/printer`)                        |
| `--exported`     | off     | Only exported symbols                                              |
| `--include-tests`| off     | Include `_test.go` files                                           |
| `--no-docs`      | off     | Strip doc comments                                                 |
| `--format`       | `go`    | `go` or `md` (markdown-fenced)                                     |
| `--max-tokens N` | 0       | Truncate to fit budget (0 = unbounded)                             |
| `-o FILE`        |         | Output file                                                        |

Examples:

```bash
pureast dump ./pkg                      # everything, signatures
pureast dump ./pkg --bodies             # include implementations
pureast dump ./pkg --kind interface     # only interfaces
pureast dump ./pkg --exported           # public API surface
pureast dump ./pkg --max-tokens 4000    # fit a budget
pureast dump ./pkg --format md -o ctx.md
```

---

## `pureast extract` — one symbol with deps

Extracts a single symbol along with its transitive dependencies as
compilable code. Default includes constructors and methods.

```
pureast extract SYMBOL [PATH] [flags]
```

| Flag             | Default | Effect                                       |
| ---------------- | ------- | -------------------------------------------- |
| `--minimal`      | off     | Direct (non-transitive) dependencies only    |
| `-w, --workers N`| 0       | Concurrent workers (0 = NumCPU)              |
| `--format`       | `go`    | `go` or `md`                                 |
| `--max-tokens N` | 0       | Truncate to fit budget                       |
| `-o FILE`        |         | Output file                                  |

Examples:

```bash
pureast extract User ./pkg
pureast extract Profile ./pkg --minimal
pureast extract UserService ./pkg -o service.go
pureast extract User ./pkg --format md
```

---

## `pureast deps` — dependency analysis

Shows what a symbol depends on (or who depends on it, with `--reverse`).

```
pureast deps SYMBOL [PATH] [flags]
```

| Flag         | Default | Effect                                                 |
| ------------ | ------- | ------------------------------------------------------ |
| `--format`   | `text`  | `text\|dot\|json`                                      |
| `--depth N`  | -1      | Max traversal depth (`-1` = unbounded, `0` = direct)   |
| `--reverse`  | off     | Show users instead of dependencies                     |
| `--minimal`  | off     | Equivalent to a small bounded depth (legacy)           |

Mutually exclusive: `--minimal` and `--depth N` (use one).

JSON output is stable: lists are alphabetically sorted, fields are
fixed. Suitable for LLM caching where byte-identical output matters.

When `--reverse` is set, each user is shown with its file:line location
(relative to the path you passed). In `--format json`, names expand
into `{name, file, line}` objects. This makes "who calls X" output
directly actionable — you can jump straight to the call site:

```
$ pureast deps User ./pkg --reverse
Reverse dependencies (users) of User:

Functions (3):
  - NewProfile  (profile.go:18)
  - User.Validate  (user.go:26)
  - UserService.AddUser  (service.go:16)
```

Examples:

```bash
pureast deps User ./pkg                         # text report
pureast deps User ./pkg --depth 1               # one hop forward
pureast deps User ./pkg --reverse               # who uses User
pureast deps User ./pkg --reverse --depth 0     # direct callers only
pureast deps User ./pkg --format json
pureast deps User ./pkg --format dot > graph.dot
dot -Tpng graph.dot -o graph.png
```

---

## `pureast diff` — symbols in changed files

Extracts symbols whose lines actually changed between a git ref and
HEAD. The intended use is PR-review and "what's new" LLM context: feed
only the code that changed in this branch, not the whole repo.

```
pureast diff REF [PATH] [flags]
```

| Flag             | Default | Effect                                                |
| ---------------- | ------- | ----------------------------------------------------- |
| `--whole-file`   | off     | Include every symbol in changed files (legacy mode)   |
| `--bodies`       | off     | Include function bodies                               |
| `--format`       | `go`    | `go` or `md`                                          |
| `--max-tokens N` | 0       | Truncate to fit budget                                |
| `-o FILE`        |         | Output file                                           |

REF is any git revision: branch, tag, commit, `HEAD~N`.

By default, `pureast diff` runs `git diff --unified=0` to find the exact
lines that changed in HEAD, parses the hunk headers, and intersects them
against AST line ranges. A symbol is included only if its declaration
overlaps a changed hunk. This is dramatically more focused for PR review
— a 1-line edit in a 2000-line file no longer dumps the whole file.

Pass `--whole-file` to revert to file-level granularity (any symbol in
any touched file). Useful when you want the surrounding context for
heavily-refactored PRs.

Examples:

```bash
pureast diff main                          # only changed-line symbols
pureast diff main --whole-file             # everything in touched files
pureast diff origin/main ./pkg --bodies
pureast diff HEAD~5 --format md -o pr.md
```

---

## `pureast search` — fuzzy symbol search

Find symbols by approximate name match.

```
pureast search PATTERN [PATH] [flags]
```

| Flag               | Default | Effect                                  |
| ------------------ | ------- | --------------------------------------- |
| `--kind K`         |         | Filter: `struct\|interface\|function`   |
| `-n, --max-results`| 20      | Maximum results to return               |

Examples:

```bash
pureast search "Handler" ./pkg
pureast search "User" ./pkg --kind struct
pureast search "Process" ./pkg -n 5
```

---

## `pureast list` — enumerate symbols

List every symbol with optional grouping.

```
pureast list [PATH] [flags]
```

| Flag             | Default | Effect                          |
| ---------------- | ------- | ------------------------------- |
| `--grouped`      | true    | Group by kind                   |

Examples:

```bash
pureast list ./pkg
pureast list ./pkg --grouped=false
```

---

## `pureast types` — type declarations only

> **DEPRECATED.** This verb is being phased out. Use `pureast dump --kind <kind>`
> instead — it covers all the same cases and is the canonical path going
> forward. The verb still works for now (a stderr warning fires on every
> invocation; suppress with `TYPES_NO_DEPRECATION_WARN=1`), but a future
> release will remove it. Migration:
>
>     pureast types ./pkg                   →  pureast dump ./pkg --kind type
>     pureast types ./pkg --kind struct     →  pureast dump ./pkg --kind struct
>     pureast types ./pkg --kind interface  →  pureast dump ./pkg --kind interface

Extract only type declarations (structs, interfaces, aliases) — no
function bodies. For functions and methods, prefer
`pureast dump --kind func` or `--kind method`.

```
pureast types [PATH] [flags]
```

| Flag                 | Default | Effect                                     |
| -------------------- | ------- | ------------------------------------------ |
| `--kind K`           | `all`   | `all\|struct\|interface`                   |
| `--exported`         | off     | Only exported symbols                      |
| `--format`           | `go`    | `go` or `md`                               |
| `--max-tokens N`     | 0       | Truncate to fit budget                     |
| `--structs-only`     | off     | **Deprecated**, use `--kind struct`        |
| `--interfaces-only`  | off     | **Deprecated**, use `--kind interface`     |
| `--functions`        | off     | **Deprecated**, use `dump --kind func`     |
| `--methods`          | off     | **Deprecated**, use `dump --kind method`   |

Examples:

```bash
pureast types ./pkg
pureast types ./pkg --kind struct
pureast types ./pkg --kind interface --exported
```

---

## Migration from the legacy flat-flag CLI

The pre-Cobra CLI took every option as a flag on a single `pureast`
invocation (`pureast -file ./pkg -symbol User -deps -report`). That
form no longer exists. Migration is mostly mechanical:

| Old                                          | New                                              |
| -------------------------------------------- | ------------------------------------------------ |
| `-file ./pkg`                                | positional `./pkg` (also `--file ./pkg`, warns)  |
| `-symbol X` (alone)                          | `extract X`                                      |
| `-symbol X -deps`                            | `deps X`                                         |
| `-symbol X -deps -report`                    | `deps X` (text is default)                       |
| `-symbol X -deps -dot`                       | `deps X --format dot`                            |
| `-symbol X -minimal`                         | `extract X --minimal`                            |
| `-search -pattern X`                         | `search X`                                       |
| `-list-symbols`                              | `list`                                           |
| `-all-types`                                 | `types` (or `dump` for everything-everything)    |
| `-structs`                                   | `types --kind struct`                            |
| `-interfaces`                                | `types --kind interface`                         |
| `-types-summary`                             | `list` (counts at top)                           |
