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
* Output **type summaries** or **LLM-friendly code contexts**.
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
git clone https://github.com/PureCompany/pureast.git
cd pureast
go build ./cmd/pureast
```

---

## ⚙️ Usage

Run `pureast` from the command line:

```bash
pureast -file ./examples/app -symbol UserService -report
```

### Available Flags

| Flag             | Description                                  |
| ---------------- | -------------------------------------------- |
| `-file <path>`   | Input Go file or directory (required).       |
| `-symbol <name>` | Target symbol to extract.                    |
| `-output <file>` | Write output to file (default: stdout).      |
| `-deps`          | Print dependencies for the given symbol.     |
| `-methods`       | List all methods for a type.                 |
| `-report`        | Generate human-readable dependency report.   |
| `-dot`           | Generate a Graphviz DOT dependency graph.    |
| `-minimal`       | Generate minimal compilable code for symbol. |
| `-recursive`     | Recurse into subdirectories.                 |
| `-workers <n>`   | Number of concurrent workers (0 = NumCPU).   |
| `-all-types`     | Extract all structs and interfaces.          |
| `-structs`       | Extract only struct declarations.            |
| `-interfaces`    | Extract only interface declarations.         |
| `-types-summary` | Show a summarized report of all types.       |

---

## 🧠 Example Workflows

### 🔍 Extract dependencies for a type

```bash
pureast -file ./examples/app -symbol Profile -deps
```

### 🧱 Generate minimal code for a symbol

```bash
pureast -file ./examples/app -symbol UserService -minimal -output extracted.go
```

### 🧾 Show dependency report

```bash
pureast -file ./examples/app -symbol Profile -report
```

### 🧩 Generate DOT graph

```bash
pureast -file ./examples/app -symbol Profile -dot > graph.dot
dot -Tpng graph.dot -o graph.png
```

### 🧠 Extract all types for LLM context

```bash
pureast -file ./examples/app -all-types > types.go
```

---

## 🧮 Architecture

```
pureast/
├── cmd/pureast/          # CLI entry point
├── pkg/
│   ├── ast/              # Core AST node structures & visitors
│   ├── extract/          # AST extraction (pure functional)
│   ├── analyze/          # Dependency graph analysis
│   ├── codegen/          # Code & report generation
│   └── ...               # Composition via monoids and folds
└── examples/             # Sample Go packages for testing
```

---

## 🧰 Built On

* 🧱 [PureKernels](https://github.com/PureCompany/purekernels):
  composable monoids, folds, functors, and concurrency primitives.
* 🧩 Go standard library (`go/ast`, `go/parser`, `token`, `format`).
* ⚙️ Functional composition patterns (`FoldLeft`, `Map`, `Reduce`).

---

## 🧪 Example Output

```bash
$ pureast -file ./examples/app -symbol Profile -report
```

**Output:**

```
# Dependency Report for: Profile

## Summary
- Types: 3
- Functions: 2
- Structs: 2
- Interfaces: 0
- Imports: 1

## Types
- User
- Address
- Profile
```

---

## 💡 Philosophy

PureAST treats Go’s syntax tree as a **mathematical object**:

> *ASTs compose, dependencies fold, and code emerges through monoidal harmony.*

It’s not a parser — it’s a **categorical extractor**.

---

## 🧠 License

🧠 The MIT + Functorial Brain-Melt License (FBML)
See the [LICENSE](./LICENSE) file for details.

---
