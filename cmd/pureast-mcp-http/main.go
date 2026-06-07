// cmd/pureast-mcp-http/main.go
//
// HTTP transport for the PureAST MCP server, backed by the official
// github.com/modelcontextprotocol/go-sdk library.
//
// Two transports are exposed on the same HTTP mux:
//
//   POST /mcp        — Streamable HTTP (spec 2025-03-26+, the modern standard)
//   GET  /sse        — Legacy SSE transport (spec 2024-11-05, for older clients)
//
// The underlying tool logic lives entirely in pkg/mcp (the custom package).
// This file is a thin adapter: it registers the same tools with the official
// SDK server so that both transports delegate to the same executors.
//
// Usage:
//
//	pureast-mcp-http                      # listens on :8080
//	pureast-mcp-http --addr :9090
//	PUREAST_HTTP_ADDR=:9090 pureast-mcp-http
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	sdkmcp "github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Pure-Company/pureast/pkg/mcp"
)

func main() {
	defaultAddr := ":8080"
	if v := os.Getenv("PUREAST_HTTP_ADDR"); v != "" {
		defaultAddr = v
	}

	addr := flag.String("addr", defaultAddr, "TCP address to listen on (e.g. :8080)")
	flag.Parse()

	workers := runtime.NumCPU()
	if workerEnv := os.Getenv("PUREAST_WORKERS"); workerEnv != "" {
		fmt.Sscanf(workerEnv, "%d", &workers)
	}

	// Build the official SDK server. All tools are registered here;
	// each handler delegates to the existing pkg/mcp executor so the
	// business logic is never duplicated.
	sdkServer := buildSDKServer(workers)

	mux := http.NewServeMux()

	// Streamable HTTP — current spec (2025-03-26+).
	// Single endpoint, Mcp-Session-Id header, optional SSE upgrade per request.
	mux.Handle("/mcp", sdkmcp.NewStreamableHTTPHandler(
		func(r *http.Request) *sdkmcp.Server { return sdkServer },
		nil,
	))

	// Legacy SSE transport — spec 2024-11-05.
	// The SSEHandler is self-contained: GET /sse opens the stream, and the
	// handler itself issues a ?sessionid= endpoint URL for the client to POST
	// messages back to — all on the same /sse path. No second handler needed.
	mux.Handle("/sse", sdkmcp.NewSSEHandler(
		func(r *http.Request) *sdkmcp.Server { return sdkServer },
		nil,
	))

	// Health check — useful for load-balancer probes.
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ok","transport":["streamable-http","sse"]}`)
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Fprintln(os.Stderr, "pureast-mcp-http: shutting down")
		cancel()
	}()

	srv := &http.Server{Addr: *addr, Handler: mux}

	go func() {
		<-ctx.Done()
		srv.Shutdown(context.Background()) //nolint:errcheck
	}()

	fmt.Fprintf(os.Stderr,
		"pureast-mcp-http: listening on %s  (workers: %d)\n"+
			"  Streamable HTTP → POST %s/mcp\n"+
			"  Legacy SSE      → GET  %s/sse\n",
		*addr, workers, *addr, *addr,
	)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "pureast-mcp-http: %v\n", err)
		os.Exit(1)
	}
}

// buildSDKServer creates an official SDK server with all pureast tools
// registered. Each tool handler is a thin shim that marshals arguments
// into the MCPRequest shape that pkg/mcp executors expect, then unmarshals
// the MCPResponse back into SDK content.
func buildSDKServer(workers int) *sdkmcp.Server {
	executor := mcp.NewToolExecutor(workers)

	server := sdkmcp.NewServer(&sdkmcp.Implementation{
		Name:    "pureast-mcp",
		Version: "1.0.0",
	}, nil)

	// ── search_symbols ────────────────────────────────────────────────────────

	type SearchSymbolsArgs struct {
		Pattern    string `json:"pattern"    jsonschema:"Search pattern"`
		Path       string `json:"path"       jsonschema:"Package path"`
		Fuzzy      bool   `json:"fuzzy"      jsonschema:"Use fuzzy matching"`
		Kind       string `json:"kind"       jsonschema:"Filter by kind (struct, interface, function)"`
		MaxResults int    `json:"maxResults" jsonschema:"Maximum number of results"`
	}
	sdkmcp.AddTool(server,
		&sdkmcp.Tool{
			Name:        "search_symbols",
			Description: "Search for symbols in a Go package using fuzzy matching",
		},
		func(ctx context.Context, req *sdkmcp.CallToolRequest, args SearchSymbolsArgs) (*sdkmcp.CallToolResult, any, error) {
			return callPkgMCP(ctx, executor.SearchSymbolsHandler(), req.Params.Name, args)
		},
	)

	// ── extract_symbol ────────────────────────────────────────────────────────

	type ExtractSymbolArgs struct {
		Symbol  string `json:"symbol"  jsonschema:"Symbol name"`
		Path    string `json:"path"    jsonschema:"Package path"`
		Minimal bool   `json:"minimal" jsonschema:"Extract minimal dependencies only"`
	}
	sdkmcp.AddTool(server,
		&sdkmcp.Tool{
			Name:        "extract_symbol",
			Description: "Extract a symbol with all its dependencies",
		},
		func(ctx context.Context, req *sdkmcp.CallToolRequest, args ExtractSymbolArgs) (*sdkmcp.CallToolResult, any, error) {
			return callPkgMCP(ctx, executor.ExtractSymbolHandler(), req.Params.Name, args)
		},
	)

	// ── list_symbols ──────────────────────────────────────────────────────────

	type ListSymbolsArgs struct {
		Path        string `json:"path"        jsonschema:"Package path"`
		GroupByKind bool   `json:"groupByKind" jsonschema:"Group symbols by kind"`
	}
	sdkmcp.AddTool(server,
		&sdkmcp.Tool{
			Name:        "list_symbols",
			Description: "List all symbols in a package",
		},
		func(ctx context.Context, req *sdkmcp.CallToolRequest, args ListSymbolsArgs) (*sdkmcp.CallToolResult, any, error) {
			return callPkgMCP(ctx, executor.ListSymbolsHandler(), req.Params.Name, args)
		},
	)

	// ── extract_types ─────────────────────────────────────────────────────────

	type ExtractTypesArgs struct {
		Path           string `json:"path"           jsonschema:"Package path"`
		StructsOnly    bool   `json:"structsOnly"    jsonschema:"Extract only structs"`
		InterfacesOnly bool   `json:"interfacesOnly" jsonschema:"Extract only interfaces"`
	}
	sdkmcp.AddTool(server,
		&sdkmcp.Tool{
			Name:        "extract_types",
			Description: "Extract type definitions (structs and interfaces)",
		},
		func(ctx context.Context, req *sdkmcp.CallToolRequest, args ExtractTypesArgs) (*sdkmcp.CallToolResult, any, error) {
			return callPkgMCP(ctx, executor.ExtractTypesHandler(), req.Params.Name, args)
		},
	)

	// ── show_dependencies ─────────────────────────────────────────────────────

	type ShowDependenciesArgs struct {
		Symbol string `json:"symbol" jsonschema:"Symbol name"`
		Path   string `json:"path"   jsonschema:"Package path"`
	}
	sdkmcp.AddTool(server,
		&sdkmcp.Tool{
			Name:        "show_dependencies",
			Description: "Show dependencies for a symbol",
		},
		func(ctx context.Context, req *sdkmcp.CallToolRequest, args ShowDependenciesArgs) (*sdkmcp.CallToolResult, any, error) {
			return callPkgMCP(ctx, executor.ShowDependenciesHandler(), req.Params.Name, args)
		},
	)

	// ── dump_package ──────────────────────────────────────────────────────────

	type DumpPackageArgs struct {
		Path         string `json:"path"         jsonschema:"Package path"`
		Kind         string `json:"kind"         jsonschema:"Filter by kind: all, struct, interface, function, method, const, var"`
		ExportedOnly bool   `json:"exportedOnly" jsonschema:"Only show exported symbols"`
		Format       string `json:"format"       jsonschema:"Output format: go or md"`
		MaxTokens    int    `json:"maxTokens"    jsonschema:"Token budget (0 = unbounded)"`
	}
	sdkmcp.AddTool(server,
		&sdkmcp.Tool{
			Name: "dump_package",
			Description: "Compact, signatures-mostly dump of every symbol in a Go package. " +
				"Use first to get oriented in an unfamiliar package before drilling into specific symbols.",
		},
		func(ctx context.Context, req *sdkmcp.CallToolRequest, args DumpPackageArgs) (*sdkmcp.CallToolResult, any, error) {
			return callPkgMCP(ctx, executor.DumpPackageHandler(), req.Params.Name, args)
		},
	)

	// ── reverse_deps ──────────────────────────────────────────────────────────

	type ReverseDepsArgs struct {
		Symbol     string `json:"symbol"     jsonschema:"Target symbol name (e.g. 'UserService' or 'User.Validate')"`
		Path       string `json:"path"       jsonschema:"Package path"`
		Transitive bool   `json:"transitive" jsonschema:"Include indirect users (callers of callers)"`
	}
	sdkmcp.AddTool(server,
		&sdkmcp.Tool{
			Name: "reverse_deps",
			Description: "Find which symbols use a given symbol. Use for impact analysis " +
				"before suggesting a refactor.",
		},
		func(ctx context.Context, req *sdkmcp.CallToolRequest, args ReverseDepsArgs) (*sdkmcp.CallToolResult, any, error) {
			return callPkgMCP(ctx, executor.ReverseDepsHandler(), req.Params.Name, args)
		},
	)

	// ── diff_since ────────────────────────────────────────────────────────────

	type DiffSinceArgs struct {
		Ref       string `json:"ref"       jsonschema:"Git ref to diff against (branch, tag, or commit)"`
		Path      string `json:"path"      jsonschema:"Repository path (defaults to '.')"`
		Format    string `json:"format"    jsonschema:"Output format: go or md"`
		MaxTokens int    `json:"maxTokens" jsonschema:"Token budget (0 = unbounded)"`
	}
	sdkmcp.AddTool(server,
		&sdkmcp.Tool{
			Name: "diff_since",
			Description: "Dump symbols from Go files that have changed since a git ref. " +
				"Intended for PR-review context.",
		},
		func(ctx context.Context, req *sdkmcp.CallToolRequest, args DiffSinceArgs) (*sdkmcp.CallToolResult, any, error) {
			return callPkgMCP(ctx, executor.DiffSinceHandler(), req.Params.Name, args)
		},
	)

	return server
}

// callPkgMCP is the bridge between the official SDK and pkg/mcp handlers.
//
// The pkg/mcp handlers expect an MCPRequest whose Params field contains:
//
//	{"name": "<tool>", "arguments": <args>}
//
// This matches what the custom stdio server sends, so we replicate it here.
// The response is always MCPResponse; on success we pull out the text content
// blocks and return them to the SDK.
func callPkgMCP(
	ctx context.Context,
	handler mcp.Handler,
	toolName string,
	args any,
) (*sdkmcp.CallToolResult, any, error) {
	// Marshal arguments into the params shape pkg/mcp handlers expect.
	argsJSON, err := json.Marshal(args)
	if err != nil {
		return nil, nil, fmt.Errorf("marshal args: %w", err)
	}

	paramsPayload := struct {
		Name      string          `json:"name"`
		Arguments json.RawMessage `json:"arguments"`
	}{
		Name:      toolName,
		Arguments: argsJSON,
	}
	paramsJSON, err := json.Marshal(paramsPayload)
	if err != nil {
		return nil, nil, fmt.Errorf("marshal params: %w", err)
	}

	req := mcp.MCPRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "tools/call",
		Params:  json.RawMessage(paramsJSON),
	}

	resp := handler(ctx, req).Value()

	if resp.Error != nil {
		return &sdkmcp.CallToolResult{
			Content: []sdkmcp.Content{
				&sdkmcp.TextContent{Text: resp.Error.Message},
			},
			IsError: true,
		}, nil, nil
	}

	// resp.Result is map[string]interface{}{"content": [...], "isError": bool}
	// Re-marshal and decode into ToolResult to extract text blocks.
	resultJSON, err := json.Marshal(resp.Result)
	if err != nil {
		return nil, nil, fmt.Errorf("marshal result: %w", err)
	}

	var toolResult mcp.ToolResult
	if err := json.Unmarshal(resultJSON, &toolResult); err != nil {
		// Fallback: return raw JSON as text.
		return &sdkmcp.CallToolResult{
			Content: []sdkmcp.Content{
				&sdkmcp.TextContent{Text: string(resultJSON)},
			},
		}, nil, nil
	}

	content := make([]sdkmcp.Content, 0, len(toolResult.Content))
	for _, block := range toolResult.Content {
		content = append(content, &sdkmcp.TextContent{Text: block.Text})
	}

	return &sdkmcp.CallToolResult{
		Content: content,
		IsError: toolResult.IsError,
	}, nil, nil
}
