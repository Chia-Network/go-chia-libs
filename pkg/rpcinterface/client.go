package rpcinterface

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

// Client defines the interface for a client
// HTTP (standard RPC) and websockets are the two supported now
type Client interface {
	NewRequest(service ServiceType, rpcEndpoint Endpoint, opt interface{}) (*Request, error)
	Do(req *Request, v IResponse) (*http.Response, error)
	Close() error
	SetBaseURL(url *url.URL) error

	// SetLogHandler sets a slog compatible log handler
	SetLogHandler(handler slog.Handler)

	// The following are added for websocket compatibility
	// Any implementation that these don't make sense for should just do nothing / return nil as applicable

	// SubscribeSelf subscribes to responses to requests from this service
	SubscribeSelf() error
	// Subscribe adds a subscription to events from a particular service
	Subscribe(service string) error

	// AddHandler adds a handler function that will be called when a message is received over the websocket
	// Does not apply to HTTP client
	AddHandler(handler WebsocketResponseHandler) (uuid.UUID, error)

	// RemoveHandler removes the handler from the active websocket handlers
	RemoveHandler(handlerID uuid.UUID)

	// AddDisconnectHandler adds a function to call if the connection is disconnected
	// Applies to websocket connections
	AddDisconnectHandler(onDisconnect DisconnectHandler)

	// AddReconnectHandler adds a function to call if the connection is reconnected
	// Applies to websocket connections
	AddReconnectHandler(onReconnect ReconnectHandler)

	// SetSyncMode enforces synchronous request/response behavior
	// This is default for HTTP client, but websocket default is async, so this forces a different mode
	// Note that anything received by the websocket in sync mode that is not the current expected response
	// will be ignored
	SetSyncMode()

	// SetAsyncMode sets the client to async mode
	// This is not supported for the HTTP client, but will set the websocket client back to async mode
	// if it was set to sync mode temporarily
	SetAsyncMode()
}
