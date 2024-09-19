package rpc

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/httpclient"
	"github.com/chia-network/go-chia-libs/pkg/publichttpclient"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/websocketclient"
)

// Client is the RPC client
type Client struct {
	config *config.ChiaConfig

	activeClient rpcinterface.Client

	// Services for the different chia services
	DaemonService    *DaemonService
	FullNodeService  *FullNodeService
	WalletService    *WalletService
	FarmerService    *FarmerService
	HarvesterService *HarvesterService
	CrawlerService   *CrawlerService
	DataLayerService *DataLayerService
	TimelordService  *TimelordService
}

// ConnectionMode specifies the method used to connect to the server (HTTP or Websocket)
type ConnectionMode uint8

const (
	// ConnectionModeHTTP uses HTTP for requests to the RPC server
	ConnectionModeHTTP ConnectionMode = iota

	// ConnectionModeWebsocket uses websockets for requests to the RPC server
	ConnectionModeWebsocket

	// ConnectionModePublicHTTP is for use with public http(s) servers that don't require cert auth but otherwise mirror the RPCs
	ConnectionModePublicHTTP
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
	case ConnectionModePublicHTTP:
		activeClient, err = publichttpclient.NewHTTPClient(options...)
	}
	if err != nil {
		return nil, err
	}
	c.activeClient = activeClient

	// Init Services
	c.DaemonService = &DaemonService{client: c}
	c.FullNodeService = &FullNodeService{client: c}
	c.WalletService = &WalletService{client: c}
	c.FarmerService = &FarmerService{client: c}
	c.HarvesterService = &HarvesterService{client: c}
	c.CrawlerService = &CrawlerService{client: c}
	c.DataLayerService = &DataLayerService{client: c}
	c.TimelordService = &TimelordService{client: c}

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
func (c *Client) AddHandler(handler rpcinterface.WebsocketResponseHandler) (uuid.UUID, error) {
	return c.activeClient.AddHandler(handler)
}

// RemoveHandler removes the handler from the list of active response handlers
func (c *Client) RemoveHandler(handlerID uuid.UUID) {
	c.activeClient.RemoveHandler(handlerID)
}

// AddDisconnectHandler the function to call when the client is disconnected
func (c *Client) AddDisconnectHandler(onDisconnect rpcinterface.DisconnectHandler) {
	c.activeClient.AddDisconnectHandler(onDisconnect)
}

// AddReconnectHandler the function to call when the client is disconnected
func (c *Client) AddReconnectHandler(onReconnect rpcinterface.ReconnectHandler) {
	c.activeClient.AddReconnectHandler(onReconnect)
}

// SetSyncMode sets the client to wait for responses before returning
// This is default (and only option) for HTTP client
// Websocket client defaults to async mode
func (c *Client) SetSyncMode() {
	c.activeClient.SetSyncMode()
}

// SetAsyncMode sets the client to async mode
// This does not apply to the HTTP client
// For the websocket client, this is the default mode and means that RPC function calls return immediate with empty
// versions of the structs that would otherwise contain the response, and you should have an async handler defined
// to receive the response
func (c *Client) SetAsyncMode() {
	c.activeClient.SetAsyncMode()
}
