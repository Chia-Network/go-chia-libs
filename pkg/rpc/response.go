package rpc

import (
	"github.com/samber/mo"
)

// Response is the base response elements that may be in any response from an RPC server in Chia
type Response struct {
	Success bool              `json:"success"`
	Error   mo.Option[string] `json:"error,omitempty"`
}
