package mcp

type MCPError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
} // MCPError represents an error response

func ErrorResponse(id interface{}, code int, message string) MCPResponse {
	return MCPResponse{JSONRPC: "2.0", ID: id, Error: &MCPError{Code: code, Message: message}}
}

type MCPResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
} // MCPResponse represents an MCP protocol response (this is a monoid!)

func SuccessResponse(id interface{}, result interface{}) MCPResponse {
	return MCPResponse{JSONRPC: "2.0", ID: id, Result: result}
}
