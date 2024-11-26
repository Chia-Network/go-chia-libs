package rpc

import (
	"github.com/samber/mo"
)

// iResponse is the interface that must be satisfied by return types so we can properly wrap RPC errors
type iResponse interface {
	isSuccessful() bool
	hasError() bool
	getError() string
}

// Response is the base response elements that may be in any response from an RPC server in Chia
type Response struct {
	Success bool              `json:"success"`
	Error   mo.Option[string] `json:"error,omitempty"`
}

// IsSuccessful returns whether the RPC request has success: true
func (r *Response) isSuccessful() bool {
	return r.Success
}

// HasError returns whether there is an error in the response
func (r *Response) hasError() bool {
	return r.Error.IsPresent()
}

// GetError returns the error if present or an empty string
func (r *Response) getError() string {
	return r.Error.OrEmpty()
}

// ChiaRPCError is the specific error returned when the RPC request succeeds, but returns success: false and an error
type ChiaRPCError struct {
	Message string
}

// Error satisfies the error interface
func (e *ChiaRPCError) Error() string {
	return e.Message
}
