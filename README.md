# 🧬 **PureAST**

**Functorial AST Extraction and Code Generation Toolkit**

🧠 The MIT + Functorial Brain-Melt License (FBML)
*Read the [LICENSE](./LICENSE) file.*

---

## 🚀 Overview

**PureAST** is a **pure functional** Go library and CLI tool that performs **AST extraction, dependency analysis, and code generation** — all built on top of **PureKernels** (monoids, folds, and functors).

It lets you:

* Extract **structs**, **interfaces**, and **functions** from any Go package.
* Analyze **dependencies** between types, functions, and imports.
* Generate **minimal compilable subgraphs** for specific symbols.
* Output **compact, LLM-friendly code contexts**.
* Produce **DOT graphs** for visualization.

Everything composes via **monoids and folds**, keeping it **pure, deterministic, and composable**.

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
go install github.com/vinodhalaharvi/pureast/cmd/pureast@latest
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

| Verb      | Purpose                                                        |
| --------- | -------------------------------------------------------------- |
| `dump`    | Compact dump of every symbol — the LLM-context flagship.       |
| `extract` | Extract one symbol plus its transitive dependencies.           |
| `deps`    | Analyze what a symbol depends on (`--format text\|dot\|json`). |
| `diff`    | Symbols in `.go` files changed since a git ref. PR-review use. |
| `search`  | Fuzzy search for symbols by name pattern.                      |
| `list`    | Enumerate all symbols in a package, optionally grouped.        |
| `types`   | **Deprecated**: use `dump --kind struct\|interface` instead.    |

Run `pureast <verb> --help` for the flags specific to each verb.

---

## 🧠 Example Workflows

### 🧠 Compact context for an LLM

```bash
pureast dump ./examples/app -o context.txt           # signatures only
pureast dump ./examples/app --bodies -o full.txt     # include implementations
pureast dump ./examples/app --kind interface         # only interfaces
pureast dump ./examples/app --exported               # only exported symbols
```

### 🧱 Extract a symbol with its dependencies

```bash
pureast extract UserService ./examples/app -o service.go
pureast extract Profile ./examples/app --minimal     # direct deps only
```

### 🧾 Dependency analysis

```bash
pureast deps Profile ./examples/app                  # text report (default)
pureast deps Profile ./examples/app --format dot > graph.dot
dot -Tpng graph.dot -o graph.png
```

### 🔍 Find a symbol

```bash
pureast search "Handler" ./examples/app
pureast list ./examples/app --grouped
```

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
│   └── commands/         # one file per verb
├── cmd/pureast-mcp/      # MCP server entry point
├── pkg/
│   ├── ast/              # Core AST node structures & visitors
│   ├── extract/          # AST extraction (pure functional)
│   ├── analyze/          # Dependency graph analysis
│   ├── codegen/          # Code & report generation
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
