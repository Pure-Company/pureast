# 🎯 PureAST CLI Reference Guide

## **Complete Command Reference**

### **1. Search & Discovery**

#### **Fuzzy Symbol Search**
```bash
# Search for symbols with fuzzy matching
./pureast -file ./pkg/mcp -search -pattern "Handler"

# Search in Go stdlib
./pureast -file $GOROOT/src/net/http -search -pattern "Server"

# Search with more workers for speed
./pureast -file ./pkg -search -pattern "Concurrent" -workers 16
```

#### **List All Symbols**
```bash
# List all symbols in a package
./pureast -file ./pkg/mcp -list-symbols

# List symbols grouped by kind (default)
./pureast -file ./pkg/mcp -list-symbols -group

# List symbols without grouping
./pureast -file ./pkg/mcp -list-symbols -group=false

# List symbols in Go stdlib
./pureast -file $GOROOT/src/net/http -list-symbols
```

#### **Build Search Index**
```bash
# Build and save index for fast searches
./pureast -file ./pkg -index

# Build index with custom path
./pureast -file ./pkg -index -index-path ./my-index.json

# Use existing index for searches (loads automatically)
./pureast -file ./pkg -search -pattern "Server"
```

---

### **2. Extract Code**

#### **Extract Single Symbol**
```bash
# Extract a specific type with all dependencies
./pureast -file ./pkg/mcp -symbol MCPResponse

# Extract to file
./pureast -file ./pkg/mcp -symbol MCPResponse -output response.go

# Extract with minimal dependencies (fewer transitive deps)
./pureast -file ./pkg/mcp -symbol MCPResponse -minimal

# Extract from Go stdlib
./pureast -file $GOROOT/src/net/http -symbol Server -output server.go
```

#### **Extract Multiple Symbols**
```bash
# Extract multiple symbols (comma-separated)
./pureast -file ./pkg/mcp -symbol "Server,Handler,MCPResponse" -output mcp_core.go

# Extract with regex pattern
./pureast -file ./pkg/mcp -symbol "MCP.*" -output mcp_types.go
```

---

### **3. Extract Types**

#### **Extract All Types (Best for LLM Context)**
```bash
# Extract ALL structs and interfaces (perfect for AI analysis)
./pureast -file ./pkg/mcp -all-types -output mcp_types.go

# Extract types from Go stdlib for LLM context
./pureast -file $GOROOT/src/net/http -all-types -output http_api.go

# Extract types with workers
./pureast -file ./large-pkg -all-types -output types.go -workers 16
```

#### **Extract Only Structs**
```bash
# Extract struct definitions only
./pureast -file ./pkg/mcp -structs -output structs.go

# Structs from Go stdlib
./pureast -file $GOROOT/src/net/http -structs -output http_structs.go
```

#### **Extract Only Interfaces**
```bash
# Extract interface definitions only
./pureast -file ./pkg/mcp -interfaces -output interfaces.go

# Interfaces from Go stdlib
./pureast -file $GOROOT/src/net/http -interfaces -output http_interfaces.go
```

#### **Types Summary**
```bash
# Show summary of all types (counts, stats)
./pureast -file ./pkg/mcp -types-summary

# Summary for large codebase
./pureast -file ./entire-project -types-summary -workers 16
```

---

### **4. Dependency Analysis**

#### **Show Dependencies**
```bash
# Show dependencies for a symbol
./pureast -file ./pkg/mcp -symbol Server -deps

# Show dependencies only (no code generation)
./pureast -file ./pkg/mcp -symbol MCPResponse -deps

# Dependencies for stdlib type
./pureast -file $GOROOT/src/net/http -symbol Server -deps
```

#### **Show Methods**
```bash
# Show methods for a type
./pureast -file ./pkg/mcp -symbol Server -methods

# Show methods from stdlib
./pureast -file $GOROOT/src/net/http -symbol Server -methods
```

#### **Generate Dependency Report**
```bash
# Generate full dependency report
./pureast -file ./pkg/mcp -symbol Server -report

# Report with graph
./pureast -file ./pkg/mcp -symbol Server -report -output report.txt

# Report for multiple symbols
./pureast -file ./pkg -symbol "Server,Handler" -report
```

---

### **5. Graph Generation**

#### **DOT Graph**
```bash
# Generate DOT graph of dependencies
./pureast -file ./pkg/mcp -symbol Server -dot -output server.dot

# Convert to PNG
./pureast -file ./pkg/mcp -symbol Server -dot -output server.dot
dot -Tpng server.dot -o server.png

# Graph for entire package
./pureast -file ./pkg/mcp -dot -output mcp.dot
```

---

### **6. Performance Tuning**

#### **Worker Configuration**
```bash
# Auto-detect CPU count (default)
./pureast -file ./pkg -workers 0

# Specific worker count
./pureast -file ./pkg -workers 8

# Maximum parallelism
./pureast -file ./pkg -workers 32

# Single-threaded (debugging)
./pureast -file ./pkg -workers 1
```

#### **Batch Processing**
```bash
# Custom batch size for concurrent processing
./pureast -file ./large-pkg -batch 20

# Smaller batches (more memory efficient)
./pureast -file ./huge-pkg -batch 5

# Larger batches (more throughput)
./pureast -file ./pkg -batch 50
```

---

### **7. Recursive Processing**

```bash
# Process directories recursively (default)
./pureast -file ./pkg -recursive

# Non-recursive (single directory only)
./pureast -file ./pkg -recursive=false

# Recursive with all types
./pureast -file ./entire-project -all-types -recursive -output api.go
```

---

## **Common Use Cases**

### **🔍 Understanding a New Codebase**
```bash
# 1. List all symbols to get overview
./pureast -file ./new-project -list-symbols

# 2. Search for main types
./pureast -file ./new-project -search -pattern "Server"

# 3. Extract API surface for LLM analysis
./pureast -file ./new-project -all-types -output api.go
```

### **📚 Creating LLM Context**
```bash
# Extract complete API surface from Go stdlib
./pureast -file $GOROOT/src/net/http -all-types -output http_context.go

# Extract from your own package
./pureast -file ./pkg/myapp -all-types -output myapp_context.go

# Extract multiple packages
./pureast -file ./pkg/api -all-types -output api_context.go
./pureast -file ./pkg/db -all-types -output db_context.go
```

### **🔧 Extracting Specific Components**
```bash
# Extract a feature with all dependencies
./pureast -file ./pkg/auth -symbol "AuthService" -output auth.go

# Extract multiple related types
./pureast -file ./pkg/api -symbol "Handler,Middleware,Router" -output api_core.go

# Extract with minimal deps for testing
./pureast -file ./pkg -symbol "Calculator" -minimal -output calc.go
```

### **📊 Analyzing Dependencies**
```bash
# Show what a type depends on
./pureast -file ./pkg -symbol "Server" -deps

# Generate visual dependency graph
./pureast -file ./pkg -symbol "Server" -dot -output deps.dot
dot -Tpng deps.dot -o deps.png

# Full dependency report
./pureast -file ./pkg -symbol "Server" -report -output report.txt
```

### **🔎 Advanced Search**
```bash
# Fuzzy search with index for speed
./pureast -file ./large-project -index
./pureast -file ./large-project -search -pattern "Handler"

# Search and extract
./pureast -file ./pkg -search -pattern "Server"
./pureast -file ./pkg -symbol "Server" -output server.go
```

---

## **Performance Tips**

1. **Use workers** for large codebases: `-workers 16`
2. **Build index** for repeated searches: `-index`
3. **Use minimal** for faster extraction: `-minimal`
4. **Adjust batch size** for memory: `-batch 10`
5. **Non-recursive** for single dir: `-recursive=false`

---

## **Examples by Package Size**

### **Small Package (< 50 files)**
```bash
./pureast -file ./pkg -all-types -output api.go
```

### **Medium Package (50-500 files)**
```bash
./pureast -file ./pkg -all-types -workers 8 -output api.go
```

### **Large Package (> 500 files)**
```bash
./pureast -file ./pkg -index
./pureast -file ./pkg -all-types -workers 16 -batch 20 -output api.go
```

### **Go Standard Library**
```bash
./pureast -file $GOROOT/src/net/http -all-types -workers 12 -output http.go
```

---

## **Quick Reference Table**

| Task | Command |
|------|---------|
| Search symbols | `-search -pattern "Foo"` |
| List symbols | `-list-symbols` |
| Extract type | `-symbol "Foo"` |
| Extract multiple | `-symbol "Foo,Bar,Baz"` |
| All types (LLM) | `-all-types` |
| Only structs | `-structs` |
| Only interfaces | `-interfaces` |
| Show dependencies | `-deps` |
| Show methods | `-methods` |
| Generate graph | `-dot` |
| Build index | `-index` |
| Parallel processing | `-workers 16` |
| Output to file | `-output file.go` |

---

