package mcp

import (
	"encoding/json"
)

// MCPRequest represents an MCP protocol request
type MCPRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

// MCPResponse represents an MCP protocol response (this is a monoid!)
type MCPResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
}

// MCPError represents an error response
type MCPError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Standard JSON-RPC error codes
const (
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603
)

// ResponseMonoid combines responses (takes first non-error)
type ResponseMonoid struct{}

func NewResponseMonoid() ResponseMonoid {
	return ResponseMonoid{}
}

func (ResponseMonoid) Empty() MCPResponse {
	return MCPResponse{
		JSONRPC: "2.0",
		Error: &MCPError{
			Code:    InternalError,
			Message: "No response",
		},
	}
}

func (ResponseMonoid) Combine(a, b MCPResponse) MCPResponse {
	// Prefer non-error responses
	if a.Error == nil {
		return a
	}
	return b
}

// ToolCall represents a tool invocation
type ToolCall struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// ToolResult represents a tool result (monoid for combining results!)
type ToolResult struct {
	Content []ContentBlock `json:"content"`
	IsError bool           `json:"isError,omitempty"`
}

// ContentBlock represents a piece of content
type ContentBlock struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	Data string `json:"data,omitempty"`
}

// ToolResultMonoid combines tool results
type ToolResultMonoid struct{}

func NewToolResultMonoid() ToolResultMonoid {
	return ToolResultMonoid{}
}

func (ToolResultMonoid) Empty() ToolResult {
	return ToolResult{Content: []ContentBlock{}}
}

func (ToolResultMonoid) Combine(a, b ToolResult) ToolResult {
	// Prefer non-error results
	if a.IsError && !b.IsError {
		return b
	}
	if !a.IsError && b.IsError {
		return a
	}
	// Both same status - concatenate content
	return ToolResult{
		Content: append(a.Content, b.Content...),
		IsError: a.IsError || b.IsError,
	}
}

// Helper constructors

func SuccessResponse(id interface{}, result interface{}) MCPResponse {
	return MCPResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}
}

func ErrorResponse(id interface{}, code int, message string) MCPResponse {
	return MCPResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &MCPError{
			Code:    code,
			Message: message,
		},
	}
}

func TextContent(text string) ContentBlock {
	return ContentBlock{Type: "text", Text: text}
}

func TextResult(text string) ToolResult {
	return ToolResult{
		Content: []ContentBlock{TextContent(text)},
	}
}

func ErrorResult(message string) ToolResult {
	return ToolResult{
		Content: []ContentBlock{TextContent(message)},
		IsError: true,
	}
}
