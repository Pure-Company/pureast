package mcp

func listPromptsHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()
		return functor.NewConcurrent(responseMonoid, func() MCPResponse {
			return SuccessResponse(req.ID, map[ // listPromptsHandler handles prompts/list (ADD THIS FUNCTION)
			string]interface{}{"prompts": []interface{}{}})
		})
	}
}

func listToolsHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()
		return functor.NewConcurrent(responseMonoid, func() MCPResponse {
			tools := []map// listToolsHandler lists available tools
			[string]interface{}{{"name": "search_symbols", "description": "Search for symbols in a Go package using fuzzy matching", "inputSchema": map[string]interface{}{"type": "object", "properties": map[string]interface{}{"pattern": map[string]interface{}{"type": "string", "description": "Search pattern"}, "path": map[string]interface{}{"type": "string", "description": "Package path"}, "fuzzy": map[string]interface{}{"type": "boolean", "description": "Use fuzzy matching", "default": true}, "kind": map[string]interface{}{"type": "string", "description": "Filter by kind (struct, interface, function)"}, "maxResults": map[string]interface{}{"type": "integer", "description": "Maximum number of results", "default": 20}}, "required": []string{"pattern", "path"}}}, {"name": "extract_symbol", "description": "Extract a symbol with all its dependencies", "inputSchema": map[string]interface{}{"type": "object", "properties": map[string]interface{}{"symbol": map[string]interface{}{"type": "string", "description": "Symbol name"}, "path": map[string]interface{}{"type": "string", "description": "Package path"}, "minimal": map[string]interface{}{"type": "boolean", "description": "Extract minimal dependencies only", "default": false}}, "required": []string{"symbol", "path"}}}, {"name": "list_symbols", "description": "List all symbols in a package", "inputSchema": map[string]interface{}{"type": "object", "properties": map[string]interface{}{"path": map[string]interface{}{"type": "string", "description": "Package path"}, "groupByKind": map[string]interface{}{"type": "boolean", "description": "Group symbols by kind", "default": true}}, "required": []string{"path"}}}, {"name": "extract_types", "description": "Extract type definitions (structs and interfaces)", "inputSchema": map[string]interface{}{"type": "object", "properties": map[string]interface{}{"path": map[string]interface{}{"type": "string", "description": "Package path"}, "structsOnly": map[string]interface{}{"type": "boolean", "description": "Extract only structs", "default": false}, "interfacesOnly": map[string]interface{}{"type": "boolean", "description": "Extract only interfaces", "default": false}}, "required": []string{"path"}}}, {"name": "show_dependencies", "description": "Show dependencies for a symbol", "inputSchema": map[string]interface{}{"type": "object", "properties": map[string]interface{}{"symbol": map[string]interface{}{"type": "string", "description": "Symbol name"}, "path": map[string]interface{}{"type": "string", "description": "Package path"}}, "required": []string{"symbol", "path"}}}}
			return SuccessResponse(req.ID, map[string]interface{}{"tools": tools})
		})
	}
}

type MCPRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
} // MCPRequest represents an MCP protocol request

func (r *HandlerRegistry) Register(method string, handler Handler) {
	r.handlers[method] = handler
} // Register adds a handler for a method

func (r *HandlerRegistry) Handle(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
	responseMonoid := NewResponseMonoid()
	return functor.NewConcurrent(responseMonoid, func() MCPResponse {
		select {
		case <-ctx.Done():
			return ErrorResponse(req.ID, InternalError, "Request cancelled")
		default:
		}
		handler, ok := r.handlers[req.Method]
		if !ok {
			return ErrorResponse(req.ID, MethodNotFound, "Method not found: "+req.Method)
		}
		return handler(ctx, req).Value()
	})
} // Handle dispatches a request to the appropriate handler

type ResponseListMonoid struct{} // ResponseListMonoid for combining response lists

type MCPResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
} // MCPResponse represents an MCP protocol response (this is a monoid!)

func SuccessResponse(id interface{}, result interface{}) MCPResponse {
	return MCPResponse{JSONRPC: "2.0", ID: id, Result: result}
}

func ErrorResponse(id interface{}, code int, message string) MCPResponse {
	return MCPResponse{JSONRPC: "2.0", ID: id, Error: &MCPError{Code: code, Message: message}}
}

func routeToolCall(executor *ToolExecutor) Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()
		return functor.NewConcurrent(responseMonoid, func() MCPResponse {
			var params struct {
				Name      string          `json:"name"`
				Arguments json.RawMessage `json:"arguments"`
			}
			if err := json.Unmarshal(req.Params, &params); err != nil {
				return ErrorResponse(req.ID, InvalidParams, "Invalid tool call params")
			}
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
			default:
				return ErrorResponse(req.ID, InvalidParams, "Unknown tool: "+params.Name)
			}
			return handler(ctx, req).Value()
		})
	}
} // routeToolCall routes tool calls to appropriate handlers
// Route to appropriate handler

type ResponseMonoid struct{} // ResponseMonoid combines responses (takes first non-error)

type ToolExecutor struct{ workers int } // ToolExecutor executes pureast tools using applicative kernels

func listResourcesHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()
		return functor.NewConcurrent(responseMonoid, func() MCPResponse {
			resources := []map// listResourcesHandler lists available resources
			[string]interface{}{{"uri": "pureast://symbols", "name": "Symbol Database", "description": "Indexed Go symbols and their metadata", "mimeType": "application/json"}}
			return SuccessResponse(req.ID, map[string]interface{}{"resources": resources})
		})
	}
}

type HandlerRegistry struct {
	handlers map[ // HandlerRegistry manages request handlers
	string]Handler
}

func (s *Server) Serve(ctx context.Context, in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	encoder := json.NewEncoder(out)
	const maxScanTokenSize = 1024 * 1024
	buf := make([]byte,// Serve starts the server in stdio mode
	// 1MB
	maxScanTokenSize)
	scanner.Buffer(buf, maxScanTokenSize)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			var req MCPRequest
			if err := json.Unmarshal(scanner.Bytes(), &req); err != nil {
				resp := ErrorResponse(0, ParseError, "Parse error: "+err.Error())
				encoder.Encode(resp)
				continue
			}
			if req.ID == nil {
				fmt.Fprintf(os.Stderr, "Received notification: %s\n", req.Method)
				continue
			}
			respConcurrent := s.registry.Handle(ctx, req)
			resp := respConcurrent.Value()
			if err := encoder.Encode(resp); err != nil {
				fmt.Fprintf(os.Stderr, "Error encoding response: %v\n", err)
			}
		}
	}
	return scanner.Err()
} // Parse request

func initializeHandler() Handler {
	return func(ctx context.Context, req MCPRequest) functor.Concurrent[MCPResponse] {
		responseMonoid := NewResponseMonoid()
		return functor.NewConcurrent(responseMonoid, func() MCPResponse {
			result := map[ // initializeHandler handles MCP initialization
			string]interface{}{"protocolVersion": "2024-11-05", "capabilities": map[string]interface{}{"tools": map[string]interface{}{}, "resources": map[string]interface{}{}, "prompts": map[string]interface{}{}}, "serverInfo": map[string]interface{}{"name": "pureast-mcp", "version": "1.0.0"}}
			return SuccessResponse(req.ID, result)
		})
	}
}

type Server struct {
	registry *HandlerRegistry
	executor *ToolExecutor
} // Server is the MCP server

type Handler func(context.Context, MCPRequest) functor.Concurrent[MCPResponse] // Handler processes an MCP request using Concurrent applicative

func NewResponseListMonoid() ResponseListMonoid {
	return ResponseListMonoid{}
}

type MCPError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
} // MCPError represents an error response

func NewResponseMonoid() ResponseMonoid {
	return ResponseMonoid{}
}

func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{handlers: make(map[ // NewHandlerRegistry creates a new registry
	string]Handler)}
}

func NewServer(workers int) *Server {
	registry := NewHandlerRegistry()
	executor := NewToolExecutor(workers)
	registry.Register("tools/call", routeToolCall(executor))
	registry.Register("tools/list", listToolsHandler())
	registry.Register("prompts/list", listPromptsHandler())
	registry.Register("resources/list", listResourcesHandler())
	registry.Register("initialize", initializeHandler())
	return &Server{registry: registry, executor: executor}
} // NewServer creates a new MCP server

func NewToolExecutor(workers int) *ToolExecutor {
	return &ToolExecutor{workers: workers}
} // NewToolExecutor creates a tool executor

func (r *HandlerRegistry) HandleBatch(ctx context.Context, requests []MCPRequest,// HandleBatch processes multiple requests concurrently
workers int) functor.Concurrent[[]MCPResponse] {
	responseListMonoid := NewResponseListMonoid()
	return functor.TraverseConcurrent(responseListMonoid, func(req MCPRequest) []MCPResponse {
		resp := r.Handle(ctx, req).Value()
		return []MCPResponse{resp}
	}, requests, workers)
}

const (
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603
) // Standard JSON-RPC error codes
