package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Pure-Company/purekernels/pkg/functor"
)

// Server is the MCP server
type Server struct {
	registry *HandlerRegistry
	executor *ToolExecutor
}

// NewServer creates a new MCP server
func NewServer(workers int) *Server {
	registry := NewHandlerRegistry()
	executor := NewToolExecutor(workers)

	// Register tool handlers
	registry.Register("tools/call", routeToolCall(executor))
	registry.Register("tools/list", listToolsHandler())
	registry.Register("prompts/list", listPromptsHandler()) // ADD THIS
	registry.Register("resources/list", listResourcesHandler())
	registry.Register("initialize", initializeHandler())

	return &Server{
		registry: registry,
		executor: executor,
	}
}

// routeToolCall routes tool calls to appropriate handlers
func routeToolCall(executor *ToolExecutor) Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				// Parse tool name
				var params struct {
					Name      string          `json:"name"`
					Arguments json.RawMessage `json:"arguments"`
				}

				if err := json.Unmarshal(req.Params, &params); err != nil {
					return ErrorResponse(req.ID, InvalidParams, "Invalid tool call params")
				}

				// Route to appropriate handler
				var handler Handler
				switch params.Name {
				case "search_symbols":
					handler = executor.SearchSymbolsHandler()
				case "extract_symbol":
					handler = executor.ExtractSymbolHandler()
				case "list_symbols":
					handler = executor.ListSymbolsHandler()
				case "extract_types":
					handler = executor.ExtractTypesHandler()
				case "show_dependencies":
					handler = executor.ShowDependenciesHandler()
				case "dump_package":
					handler = executor.DumpPackageHandler()
				case "reverse_deps":
					handler = executor.ReverseDepsHandler()
				case "diff_since":
					handler = executor.DiffSinceHandler()
				default:
					return ErrorResponse(req.ID, InvalidParams, "Unknown tool: "+params.Name)
				}

				// Execute handler
				return handler(ctx, req).Value()
			},
		)
	}
}

// listPromptsHandler handles prompts/list (ADD THIS FUNCTION)
func listPromptsHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				// Return empty prompts list
				return SuccessResponse(req.ID, map[string]interface{}{
					"prompts": []interface{}{},
				})
			},
		)
	}
}

// listToolsHandler lists available tools
func listToolsHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				tools := []map[string]interface{}{
					{
						"name": "search_symbols",
						"description": "Search for Go symbols (types, functions, interfaces) by name using fuzzy matching. " +
							"Use this when the user mentions a symbol by name but you don't know which file or package it lives in. " +
							"Also use when exploring an unfamiliar codebase to locate entry points before reading code.",
						"inputSchema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"pattern": map[string]interface{}{
									"type":        "string",
									"description": "Search pattern",
								},
								"path": map[string]interface{}{
									"type":        "string",
									"description": "Package path",
								},
								"fuzzy": map[string]interface{}{
									"type":        "boolean",
									"description": "Use fuzzy matching",
									"default":     true,
								},
								"kind": map[string]interface{}{
									"type":        "string",
									"description": "Filter by kind (struct, interface, function)",
								},
								"maxResults": map[string]interface{}{
									"type":        "integer",
									"description": "Maximum number of results",
									"default":     20,
								},
							},
							"required": []string{"pattern", "path"},
						},
					},
					{
						"name": "extract_symbol",
						"description": "Extract a Go symbol together with every type, function, and import it depends on — " +
							"producing a minimal compilable snippet. Use this when you need to deeply understand or modify a " +
							"specific symbol: before editing a struct, implementing an interface, or fixing a function. " +
							"Prefer dump_package first for orientation, then extract_symbol to drill into a specific target.",
						"inputSchema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"symbol": map[string]interface{}{
									"type":        "string",
									"description": "Symbol name",
								},
								"path": map[string]interface{}{
									"type":        "string",
									"description": "Package path",
								},
								"minimal": map[string]interface{}{
									"type":        "boolean",
									"description": "Extract minimal dependencies only",
									"default":     false,
								},
							},
							"required": []string{"symbol", "path"},
						},
					},
					{
						"name": "list_symbols",
						"description": "List all symbol names in a Go package grouped by kind (struct, interface, function, etc). " +
							"Use this for a quick inventory of what exists in a package — cheaper than dump_package when you " +
							"only need names, not signatures. Good first step when a user asks 'what's in this package?'.",
						"inputSchema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"path": map[string]interface{}{
									"type":        "string",
									"description": "Package path",
								},
								"groupByKind": map[string]interface{}{
									"type":        "boolean",
									"description": "Group symbols by kind",
									"default":     true,
								},
							},
							"required": []string{"path"},
						},
					},
					{
						"name": "extract_types",
						"description": "Extract all struct and interface definitions from a Go package. " +
							"Use this when planning a refactor that touches the type system, when checking interface compliance, " +
							"or when generating code that needs to match existing types. Use structsOnly or interfacesOnly " +
							"to narrow scope when you know what you're looking for.",
						"inputSchema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"path": map[string]interface{}{
									"type":        "string",
									"description": "Package path",
								},
								"structsOnly": map[string]interface{}{
									"type":        "boolean",
									"description": "Extract only structs",
									"default":     false,
								},
								"interfacesOnly": map[string]interface{}{
									"type":        "boolean",
									"description": "Extract only interfaces",
									"default":     false,
								},
							},
							"required": []string{"path"},
						},
					},
					{
						"name": "show_dependencies",
						"description": "Show the full dependency tree for a Go symbol — what types, functions, and imports it pulls in. " +
							"Call this before editing or moving a symbol to understand its blast radius. " +
							"Use format='json' with locations=true for precise file:line info when you need to navigate to dependents. " +
							"For the reverse direction (who depends ON this symbol), use reverse_deps instead.",
						"inputSchema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"symbol": map[string]interface{}{
									"type":        "string",
									"description": "Symbol name",
								},
								"path": map[string]interface{}{
									"type":        "string",
									"description": "Package path",
								},
							},
							"required": []string{"symbol", "path"},
						},
					},
					{
						"name": "dump_package",
						"description": "Compact, signatures-mostly dump of every symbol in a Go package. " +
							"This is the LLM-context flagship tool: call it first whenever starting work on an unfamiliar package, " +
							"before any refactor, feature addition, or code review. Gives you the full shape of the package " +
							"(all types, functions, methods, consts, vars) without requiring you to read individual files. " +
							"Use maxTokens to stay within context limits on large packages. " +
							"Follow up with extract_symbol to drill into specific symbols.",
						"inputSchema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"path": map[string]interface{}{
									"type":        "string",
									"description": "Package path",
								},
								"kind": map[string]interface{}{
									"type":        "string",
									"description": "Filter by kind: all, struct, interface, function, method, const, var",
									"default":     "all",
								},
								"exportedOnly": map[string]interface{}{
									"type":        "boolean",
									"description": "Only show exported symbols (capitalized names)",
									"default":     false,
								},
								"format": map[string]interface{}{
									"type":        "string",
									"description": "Output format: 'go' (raw) or 'md' (markdown-fenced)",
									"default":     "go",
								},
								"maxTokens": map[string]interface{}{
									"type":        "integer",
									"description": "Truncate output to fit a token budget (0 = unbounded). Line-aware truncation.",
									"default":     0,
								},
							},
							"required": []string{"path"},
						},
					},
					{
						"name": "reverse_deps",
						"description": "Find all symbols that depend on (use) a given symbol — the reverse of show_dependencies. " +
							"Use this for impact analysis before any rename, signature change, deletion, or refactor: " +
							"call reverse_deps first to see everything that will break. " +
							"Set transitive=true to find indirect callers too (callers of callers). " +
							"Complements show_dependencies: use both together for a full picture of a symbol's position in the codebase.",
						"inputSchema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"symbol": map[string]interface{}{
									"type":        "string",
									"description": "Target symbol name (e.g. 'UserService' or 'User.Validate')",
								},
								"path": map[string]interface{}{
									"type":        "string",
									"description": "Package path",
								},
								"transitive": map[string]interface{}{
									"type":        "boolean",
									"description": "Include indirect users (callers of callers). Default: direct only.",
									"default":     false,
								},
							},
							"required": []string{"symbol", "path"},
						},
					},
					{
						"name": "diff_since",
						"description": "Dump symbols only from Go files that changed since a git ref (branch, tag, or commit). " +
							"Use this automatically when reviewing a PR or branch: instead of loading the entire codebase, " +
							"call diff_since with ref='main' to see only what's new or changed. " +
							"Also useful before a merge to summarize what a branch introduces. " +
							"Set bodies=true when you need full function implementations, not just signatures.",
						"inputSchema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"ref": map[string]interface{}{
									"type":        "string",
									"description": "Git ref to diff against (branch, tag, or commit, e.g. 'main' or 'HEAD~5')",
								},
								"path": map[string]interface{}{
									"type":        "string",
									"description": "Repository path (defaults to '.')",
								},
								"bodies": map[string]interface{}{
									"type":        "boolean",
									"description": "Include function bodies (default: signatures only)",
									"default":     false,
								},
								"format": map[string]interface{}{
									"type":        "string",
									"description": "Output format: 'go' or 'md'",
									"default":     "go",
								},
								"maxTokens": map[string]interface{}{
									"type":        "integer",
									"description": "Token budget (0 = unbounded)",
									"default":     0,
								},
							},
							"required": []string{"ref"},
						},
					},
				}

				return SuccessResponse(req.ID, map[string]interface{}{"tools": tools})
			},
		)
	}
}

// listResourcesHandler lists available resources
func listResourcesHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				resources := []map[string]interface{}{
					{
						"uri":         "pureast://symbols",
						"name":        "Symbol Database",
						"description": "Indexed Go symbols and their metadata",
						"mimeType":    "application/json",
					},
				}

				return SuccessResponse(req.ID, map[string]interface{}{"resources": resources})
			},
		)
	}
}

// initializeHandler handles MCP initialization
func initializeHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()

		return functor.NewConcurrent(
			responseMonoid,
			func() MCPResponse {
				result := map[string]interface{}{
					"protocolVersion": "2024-11-05",
					"capabilities": map[string]interface{}{
						"tools":     map[string]interface{}{},
						"resources": map[string]interface{}{},
						"prompts":   map[string]interface{}{},
					},
					"serverInfo": map[string]interface{}{
						"name":    "pureast-mcp",
						"version": "1.0.0",
					},
				}

				return SuccessResponse(req.ID, result)
			},
		)
	}
}

// Serve starts the server in stdio mode
func (s *Server) Serve(ctx context.Context, in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	encoder := json.NewEncoder(out)

	// Set max buffer size for large responses
	const maxScanTokenSize = 1024 * 1024 // 1MB
	buf := make([]byte, maxScanTokenSize)
	scanner.Buffer(buf, maxScanTokenSize)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Parse request
			var req MCPRequest
			if err := json.Unmarshal(scanner.Bytes(), &req); err != nil {
				// Send parse error with id: 0 (not null)
				resp := ErrorResponse(0, ParseError, "Parse error: "+err.Error())
				encoder.Encode(resp)
				continue
			}

			// Check if this is a notification (no id field or id is nil/null)
			// Notifications don't get responses
			if req.ID == nil {
				// This is a notification - just log it and don't respond
				fmt.Fprintf(os.Stderr, "Received notification: %s\n", req.Method)
				continue
			}

			// Handle request using concurrent applicative
			respConcurrent := s.registry.Handle(ctx, req)
			resp := respConcurrent.Value()

			// Write response
			if err := encoder.Encode(resp); err != nil {
				fmt.Fprintf(os.Stderr, "Error encoding response: %v\n", err)
			}
		}
	}

	return scanner.Err()
}
