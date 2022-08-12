package rpcinterface

import (
	"net/http"
	"net/url"
	"time"
)

// Client defines the interface for a client
// HTTP (standard RPC) and websockets are the two supported now
type Client interface {
	NewRequest(service ServiceType, rpcEndpoint Endpoint, opt interface{}) (*Request, error)
	Do(req *Request, v interface{}) (*http.Response, error)
	SetBaseURL(url *url.URL) error
	SetCacheValidTime(validTime time.Duration)

	// The following are added for websocket compatibility
	// Any implementation that these don't make sense for should just do nothing / return nil as applicable

	// SubscribeSelf subscribes to responses to requests from this service
	SubscribeSelf() error
	// Subscribe adds a subscription to events from a particular service
	Subscribe(service string) error
	// ListenSync Listens for async responses over the connection in a synchronous fashion, blocking anything else
	ListenSync(handler WebsocketResponseHandler) error

	// AddDisconnectHandler adds a function to call if the connection is disconnected
	// Applies to websocket connections
	AddDisconnectHandler(onDisconnect DisconnectHandler)

	// AddReconnectHandler adds a function to call if the connection is reconnected
	// Applies to websocket connections
	AddReconnectHandler(onReconnect ReconnectHandler)
}
