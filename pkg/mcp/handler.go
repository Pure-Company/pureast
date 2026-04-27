package mcp

import (
	"context"

	"github.com/Pure-Company/purekernels/pkg/functor"
	"github.com/Pure-Company/purekernels/pkg/result"
)

// Handler processes an MCP request using Concurrent applicative
type Handler func(context.Context, MCPRequest) functor.Concurrent[MCPResponse]

// HandlerRegistry manages request handlers
type HandlerRegistry struct {
	handlers map[string]Handler
}

// NewHandlerRegistry creates a new registry
func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		handlers: make(map[string]Handler),
	}
}

// Register adds a handler for a method
func (r *HandlerRegistry) Register(method string, handler Handler) {
	r.handlers[method] = handler
}

// Handle dispatches a request to the appropriate handler
func (r *HandlerRegistry) Handle(
	ctx context.Context,
	req MCPRequest,
) functor.Concurrent[MCPResponse] {
	responseMonoid := NewResponseMonoid()

	return functor.NewConcurrent(
		responseMonoid,
		func() MCPResponse {
			// Check for cancellation
			select {
			case <-ctx.Done():
				return ErrorResponse(req.ID, InternalError, "Request cancelled")
			default:
			}

			// Find handler
			handler, ok := r.handlers[req.Method]
			if !ok {
				return ErrorResponse(req.ID, MethodNotFound, "Method not found: "+req.Method)
			}

			// Execute handler concurrently
			return handler(ctx, req).Value()
		},
	)
}

// HandleBatch processes multiple requests concurrently
func (r *HandlerRegistry) HandleBatch(
	ctx context.Context,
	requests []MCPRequest,
	workers int,
) functor.Concurrent[[]MCPResponse] {
	responseListMonoid := NewResponseListMonoid()

	// Use TraverseConcurrent for parallel batch processing!
	return functor.TraverseConcurrent(
		responseListMonoid,
		func(req MCPRequest) []MCPResponse {
			resp := r.Handle(ctx, req).Value()
			return []MCPResponse{resp}
		},
		requests,
		workers,
	)
}

// ResponseListMonoid for combining response lists
type ResponseListMonoid struct{}

func NewResponseListMonoid() ResponseListMonoid {
	return ResponseListMonoid{}
}

func (ResponseListMonoid) Empty() []MCPResponse {
	return []MCPResponse{}
}

func (ResponseListMonoid) Combine(a, b []MCPResponse) []MCPResponse {
	return append(a, b...)
}

// WrapResult wraps a Result in a Concurrent handler
func WrapResult[T any](
	reqID interface{},
	res result.Result[T],
	toResponse func(T) interface{},
) functor.Concurrent[MCPResponse] {
	responseMonoid := NewResponseMonoid()

	return functor.NewConcurrent(
		responseMonoid,
		func() MCPResponse {
			if !res.IsOk() {
				return ErrorResponse(reqID, InternalError, res.Error().Error())
			}
			return SuccessResponse(reqID, toResponse(res.Unwrap()))
		},
	)
}
