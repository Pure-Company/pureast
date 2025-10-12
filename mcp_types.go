// Type definitions extracted by pureast

// Suitable for LLM context - contains only type structures

package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/functor"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
	"go/token"
	"io"
	"os"
)

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
