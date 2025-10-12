// Type definitions extracted by pureast

// Suitable for LLM context - contains only type structures

package analyze

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/functor"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type DependencyGraph struct {
	Decls map[ // DependencyGraph represents the full dependency structure
	string]astpkg.DeclNode
}

type DependencyStats struct {
	TotalTypes      int
	TotalFunctions  int
	TotalStructs    int
	TotalInterfaces int
	TotalImports    int
	MaxDepth        int
} // DependencyStats computes statistics about dependencies

type InterfaceChecker struct{ graph DependencyGraph } // InterfaceChecker checks interface implementations

type Dependencies struct {
	Types      monoid.SetMonoid[string]
	Functions  monoid.SetMonoid[string]
	Structs    monoid.SetMonoid[string]
	Interfaces monoid.SetMonoid[string]
	Imports    monoid.SetMonoid[string]
	Constants  monoid.SetMonoid[string]
	Variables  monoid.SetMonoid[string]
} // Dependencies accumulates AST dependencies using monoids
// This is our core composition structure - everything combines via monoid operations

type DependencyMonoid struct{} // DependencyMonoid - monoid instance for Dependencies

type DeclNode struct {
	Name string
	Decl ast.Decl
	Deps Dependencies
} // DeclNode represents a declaration with its dependencies

type FileNode struct {
	Name  string
	File  *ast.File
	Decls []DeclNode// FileNode represents a file with all its declarations

	Imports []string
	Deps    Dependencies
}

type PackageNode struct {
	Name  string
	Files []FileNode// PackageNode represents a package with multiple files

	Deps Dependencies
}

type MethodNode struct {
	ReceiverType string
	MethodName   string
	Func         *ast.FuncDecl
	Deps         Dependencies
} // MethodNode represents a method with receiver

type InterfaceImplementation struct {
	TypeName      string
	InterfaceName string
	Methods       []MethodNode// InterfaceImplementation tracks interface implementation

	MissingMethods []string
}

type Generator struct{ fset *token.FileSet } // Generator generates Go code from AST nodes

type Code struct {
	Lines []string// Code represents generated code as a monoid
}

type CodeMonoid struct{} // CodeMonoid - monoid for combining code

type bytesBuffer struct {
	data []byte// bytesBuffer implements io.Writer for printer
}

type FileDiscovery struct {
	Root      string
	Recursive bool
} // FileDiscovery finds all Go files in a directory tree

type FileNodeMonoid struct{} // FileNodeMonoid combines FileNodes

type SymbolInfo struct {
	Name     string
	Kind     string
	Package  string
	Receiver string
} // SymbolInfo represents a discovered symbol
// For methods: the receiver type

type TypeDeclaration struct {
	Name string
	Kind TypeKind
	Decl astpkg.DeclNode
} // TypeDeclaration represents a type with its kind

type TypeSummary struct {
	TotalTypes     int
	StructCount    int
	InterfaceCount int
	OtherCount     int
	Names          []string// TypeSummary provides a summary of type declarations

}

type FuzzyScore struct {
	Matched bool
	Score   int
} // FuzzyScore represents a fuzzy match score (this is a monoid!)

type FuzzyScoreMonoid struct{} // FuzzyScoreMonoid combines fuzzy scores

type ScoredSymbol struct {
	Entry SymbolEntry
	Score FuzzyScore
} // ScoredSymbol represents a symbol with its match score

type ScoredSymbolMonoid struct{} // ScoredSymbolMonoid combines scored symbols (for parallel search)

type SymbolEntry struct {
	Name        string
	Kind        string
	PackageName string
	File        string
	Line        int
} // SymbolEntry represents a symbol in the index

type Index struct {
	Symbols map[ // Index represents a searchable symbol index
	string]SymbolEntry
	PackageName string
	Files       []string
}

type SearchPattern struct {
	SymbolPattern string
	Kind          string
	PackageName   string
} // SearchPattern represents search criteria

type HandlerRegistry struct {
	handlers map[ // HandlerRegistry manages request handlers
	string]Handler
}

type ResponseListMonoid struct{} // ResponseListMonoid for combining response lists

type MCPRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
} // MCPRequest represents an MCP protocol request

type MCPResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
} // MCPResponse represents an MCP protocol response (this is a monoid!)

type MCPError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
} // MCPError represents an error response

type ResponseMonoid struct{} // ResponseMonoid combines responses (takes first non-error)

type ToolCall struct {
	Name               string `json:"name"`
	Arguments          map[   // ToolCall represents a tool invocation
	string]interface{}        `json:"arguments"`
}

type ToolResult struct {
	Content []ContentBlock `json:"content"`// ToolResult represents a tool result (monoid for combining results!)

	IsError bool `json:"isError,omitempty"`
}

type ContentBlock struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	Data string `json:"data,omitempty"`
} // ContentBlock represents a piece of content

type ToolResultMonoid struct{} // ToolResultMonoid combines tool results

type Server struct {
	registry *HandlerRegistry
	executor *ToolExecutor
} // Server is the MCP server

type ToolExecutor struct{ workers int } // ToolExecutor executes pureast tools using applicative kernels
