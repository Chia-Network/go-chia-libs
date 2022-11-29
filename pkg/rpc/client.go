package rpc

import (
	"log"
	"net/http"

	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/httpclient"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
	"github.com/chia-network/go-chia-libs/pkg/websocketclient"
)

// Client is the RPC client
type Client struct {
	config *config.ChiaConfig

	activeClient rpcinterface.Client

	// Services for the different chia services
	FullNodeService  *FullNodeService
	WalletService    *WalletService
	HarvesterService *HarvesterService
	CrawlerService   *CrawlerService

	websocketHandlers []rpcinterface.WebsocketResponseHandler
}

// ConnectionMode specifies the method used to connect to the server (HTTP or Websocket)
type ConnectionMode uint8

const (
	// ConnectionModeHTTP uses HTTP for requests to the RPC server
	ConnectionModeHTTP ConnectionMode = iota

	// ConnectionModeWebsocket uses websockets for requests to the RPC server
	ConnectionModeWebsocket
)

// NewClient returns a new RPC Client
func NewClient(connectionMode ConnectionMode, configOption rpcinterface.ConfigOptionFunc, options ...rpcinterface.ClientOptionFunc) (*Client, error) {
	cfg, err := configOption()
	if err != nil {
		return nil, err
	}

	c := &Client{
		config: cfg,
	}

	var activeClient rpcinterface.Client
	switch connectionMode {
	case ConnectionModeHTTP:
		activeClient, err = httpclient.NewHTTPClient(cfg, options...)
	case ConnectionModeWebsocket:
		activeClient, err = websocketclient.NewWebsocketClient(cfg, options...)
	}
	if err != nil {
		return nil, err
	}
	c.activeClient = activeClient

	// Init Services
	c.FullNodeService = &FullNodeService{client: c}
	c.WalletService = &WalletService{client: c}
	c.HarvesterService = &HarvesterService{client: c}
	c.CrawlerService = &CrawlerService{client: c}

	return c, nil
}

// NewRequest is a helper that wraps the activeClient's NewRequest method
func (c *Client) NewRequest(service rpcinterface.ServiceType, rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return c.activeClient.NewRequest(service, rpcEndpoint, opt)
}

// Do is a helper that wraps the activeClient's Do method
func (c *Client) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return c.activeClient.Do(req, v)
}

// The following has a bunch of methods that are currently only used for the websocket implementation

// SubscribeSelf subscribes to responses to requests from this service
// This is currently only useful for websocket mode
func (c *Client) SubscribeSelf() error {
	return c.activeClient.SubscribeSelf()
}

// Subscribe adds a subscription to events from a particular service
// This is currently only useful for websocket mode
func (c *Client) Subscribe(service string) error {
	return c.activeClient.Subscribe(service)
}

// AddHandler adds a handler function to call when a message is received over the websocket
// This is expected to NOT be used in conjunction with ListenSync
// This will run in the background, and allow other things to happen in the foreground
// while ListenSync will take over the foreground process
func (c *Client) AddHandler(handler rpcinterface.WebsocketResponseHandler) error {
	c.websocketHandlers = append(c.websocketHandlers, handler)

	go func() {
		err := c.ListenSync(c.handlerProxy)
		if err != nil {
			log.Printf("Error calling ListenSync: %s\n", err.Error())
		}
	}()
	return nil
}

// AddDisconnectHandler the function to call when the client is disconnected
func (c *Client) AddDisconnectHandler(onDisconnect rpcinterface.DisconnectHandler) {
	c.activeClient.AddDisconnectHandler(onDisconnect)
}

// AddReconnectHandler the function to call when the client is disconnected
func (c *Client) AddReconnectHandler(onReconnect rpcinterface.ReconnectHandler) {
	c.activeClient.AddReconnectHandler(onReconnect)
}

// handlerProxy matches the websocketRespHandler signature to send requests back to any registered handlers
// Here to support multiple handlers for a single event in the future
func (c *Client) handlerProxy(resp *types.WebsocketResponse, err error) {
	for _, handler := range c.websocketHandlers {
		handler(resp, err)
	}
}

// ListenSync Listens for async responses over the connection in a synchronous fashion, blocking anything else
func (c *Client) ListenSync(handler rpcinterface.WebsocketResponseHandler) error {
	return c.activeClient.ListenSync(handler)
}
