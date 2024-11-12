package websocketclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
	"github.com/chia-network/go-chia-libs/pkg/util"
)

// WebsocketClient connects to Chia RPC via websockets
type WebsocketClient struct {
	config  *config.ChiaConfig
	baseURL *url.URL
	logger  *slog.Logger
	origin  string

	// Request timeout
	Timeout time.Duration

	daemonPort    uint16
	daemonKeyPair *tls.Certificate
	daemonDialer  *websocket.Dialer

	conn *websocket.Conn
	lock sync.Mutex

	// listenCancel is the cancel method of the listen context to stop listening
	listenCancel context.CancelFunc

	// listenSyncActive is tracking whether a client has opted to temporarily listen in sync mode for ALL requests
	listenSyncActive bool

	// syncMode is tracking whether or not the actual RPC calls should behave like sync calls
	// in this mode, we'll generate a request ID, and block until we get the request ID response back
	// (or hit a timeout)
	syncMode bool

	// subscriptions Keeps track of subscribed topics, so we can re-subscribe if we lose a connection and reconnect
	subscriptions map[string]bool

	// All registered functions that want data back from the websocket
	websocketHandlerMutex sync.Mutex
	websocketHandlers     map[uuid.UUID]rpcinterface.WebsocketResponseHandler

	disconnectHandlers []rpcinterface.DisconnectHandler
	reconnectHandlers  []rpcinterface.ReconnectHandler
}

// NewWebsocketClient returns a new websocket client that satisfies the rpcinterface.Client interface
func NewWebsocketClient(cfg *config.ChiaConfig, options ...rpcinterface.ClientOptionFunc) (*WebsocketClient, error) {
	c := &WebsocketClient{
		config: cfg,
		logger: slog.New(rpcinterface.SlogInfo()),
		origin: fmt.Sprintf("go-chia-rpc-%d", time.Now().UnixNano()),

		Timeout: 10 * time.Second, // Default, overridable with client option

		daemonPort: cfg.DaemonPort,

		// Init the maps
		subscriptions:     map[string]bool{},
		websocketHandlers: map[uuid.UUID]rpcinterface.WebsocketResponseHandler{},
	}

	// Sets the default host. Can be overridden by client options
	err := c.SetBaseURL(&url.URL{
		Scheme: "wss",
		Host:   "localhost",
	})
	if err != nil {
		return nil, err
	}

	err = c.initialKeyPairs()
	if err != nil {
		return nil, err
	}

	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err = fn(c); err != nil {
			return nil, err
		}
	}

	// Generate the http clients and transports after any client options are applied, in case custom keypairs were provided
	err = c.generateDialer()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// resetOrigin so we get a unique identifier if we have to establish a new connection
// Helps ensure that we don't end up getting duplicate messages
func (c *WebsocketClient) resetOrigin() {
	c.origin = fmt.Sprintf("go-chia-rpc-%d", time.Now().UnixNano())
}

// SetBaseURL sets the base URL for API requests to a custom endpoint.
func (c *WebsocketClient) SetBaseURL(url *url.URL) error {
	c.baseURL = url

	return nil
}

// SetLogHandler sets a slog compatible log handler
func (c *WebsocketClient) SetLogHandler(handler slog.Handler) {
	c.logger = slog.New(handler)
}

// NewRequest creates an RPC request for the specified service
func (c *WebsocketClient) NewRequest(service rpcinterface.ServiceType, rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	request := &rpcinterface.Request{
		Service:  service,
		Endpoint: rpcEndpoint,
		Data:     opt,
		Request:  nil,
	}

	return request, nil
}

// Do sends an RPC request via the websocket
// *http.Response is always nil in this return in async mode
// call SetSyncMode() to ensure the calls return the data in a synchronous fashion
func (c *WebsocketClient) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	err := c.ensureConnection()
	if err != nil {
		return nil, fmt.Errorf("error ensuring connection: %w", err)
	}

	var destination string
	switch req.Service {
	case rpcinterface.ServiceDaemon:
		destination = "daemon"
	case rpcinterface.ServiceFullNode:
		destination = "chia_full_node"
	case rpcinterface.ServiceFarmer:
		destination = "chia_farmer"
	case rpcinterface.ServiceHarvester:
		destination = "chia_harvester"
	case rpcinterface.ServiceWallet:
		destination = "chia_wallet"
	case rpcinterface.ServiceCrawler:
		destination = "chia_crawler"
	case rpcinterface.ServiceTimelord:
		destination = "chia_timelord"
	default:
		return nil, fmt.Errorf("unknown service")
	}

	data := req.Data
	if data == nil {
		data = map[string]interface{}{}
	}
	request := &types.WebsocketRequest{
		Command:     string(req.Endpoint),
		Origin:      c.origin,
		Destination: destination,
		Data:        data,
		RequestID:   util.GenerateRequestID(),
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	err = c.conn.WriteJSON(request)
	if err != nil {
		return nil, err
	}

	return c.responseHelper(request, v)
}

// Close closes the client/websocket
func (c *WebsocketClient) Close() error {
	if c.listenCancel != nil {
		c.listenCancel()
	}
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// responseHelper implements the logic to either immediately return in async mode
// or to wait for the expected response up to the defined timeout and returns the
// response in a synchronous fashion
func (c *WebsocketClient) responseHelper(request *types.WebsocketRequest, v interface{}) (*http.Response, error) {
	if !c.syncMode {
		return nil, nil
	}
	// We're in sync mode, so wait up to the timeout for the desired response, or else return an error

	errChan := make(chan error)
	doneChan := make(chan bool)
	ctx, cancelCtx := context.WithTimeout(context.Background(), c.Timeout)
	defer cancelCtx()

	// Set up a handler to process responses and keep an eye out for the right one
	handlerID, err := c.AddHandler(func(response *types.WebsocketResponse, err error) {
		if response.RequestID == request.RequestID {
			var err error
			if v != nil {
				reader := bytes.NewReader(response.Data)
				if w, ok := v.(io.Writer); ok {
					_, err = io.Copy(w, reader)
				} else {
					err = json.NewDecoder(reader).Decode(v)
				}
				if err != nil {
					errChan <- err
				}
			}
			doneChan <- true
		}
	})

	if err != nil {
		return nil, err
	}

	for {
		select {
		case err = <-errChan:
			c.RemoveHandler(handlerID)
			return nil, err
		case doneResult := <-doneChan:
			if doneResult {
				c.RemoveHandler(handlerID)
				return nil, nil
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout of %s reached before getting a response", c.Timeout.String())
		}
	}
}

// SubscribeSelf calls subscribe for any requests that this client makes to the server
// Different from Subscribe with a custom service - that is more for subscribing to built in events emitted by Chia
// This call will subscribe `go-chia-rpc` origin for any requests we specifically make of the server
func (c *WebsocketClient) SubscribeSelf() error {
	return c.Subscribe(c.origin)
}

// Subscribe adds a subscription to a particular service
func (c *WebsocketClient) Subscribe(service string) error {
	if _, alreadySet := c.subscriptions[service]; alreadySet {
		return nil
	}

	err := c.doSubscribe(service)
	if err != nil {
		return err
	}

	c.subscriptions[service] = true
	return nil
}

func (c *WebsocketClient) doSubscribe(service string) error {
	request, err := c.NewRequest(rpcinterface.ServiceDaemon, "register_service", types.WebsocketSubscription{Service: service})
	if err != nil {
		return err
	}

	_, err = c.Do(request, nil)
	return err
}

// AddHandler Adds a new handler function and returns its UUID for removing it later or an error
func (c *WebsocketClient) AddHandler(handler rpcinterface.WebsocketResponseHandler) (uuid.UUID, error) {
	c.websocketHandlerMutex.Lock()
	defer c.websocketHandlerMutex.Unlock()

	handlerID := uuid.New()
	c.websocketHandlers[handlerID] = handler

	return handlerID, nil
}

// RemoveHandler removes the handler from the list of active response handlers
func (c *WebsocketClient) RemoveHandler(handlerID uuid.UUID) {
	c.websocketHandlerMutex.Lock()
	defer c.websocketHandlerMutex.Unlock()
	delete(c.websocketHandlers, handlerID)
}

// handlerProxy matches the websocketRespHandler signature to send requests back to any registered handlers
func (c *WebsocketClient) handlerProxy(resp *types.WebsocketResponse, err error) {
	for _, handler := range c.websocketHandlers {
		handler(resp, err)
	}
}

// AddDisconnectHandler the function to call when the client is disconnected
func (c *WebsocketClient) AddDisconnectHandler(onDisconnect rpcinterface.DisconnectHandler) {
	c.disconnectHandlers = append(c.disconnectHandlers, onDisconnect)
}

// AddReconnectHandler the function to call when the client is reconnected after a disconnect
func (c *WebsocketClient) AddReconnectHandler(onReconnect rpcinterface.ReconnectHandler) {
	c.reconnectHandlers = append(c.reconnectHandlers, onReconnect)
}

// SetSyncMode enforces synchronous request/response behavior
// RPC method calls return the actual expected RPC response when this mode is enabled
func (c *WebsocketClient) SetSyncMode() {
	c.syncMode = true
}

// SetAsyncMode sets the client to async mode (default)
// RPC method calls return empty versions of the response objects, and you must have your own
// listeners to get the responses and handle them
func (c *WebsocketClient) SetAsyncMode() {
	c.syncMode = false
}

func (c *WebsocketClient) reconnectLoop() {
	for _, handler := range c.disconnectHandlers {
		handler()
	}
	for {
		c.logger.Info("Trying to reconnect...")
		err := c.ensureConnection()
		if err == nil {
			c.logger.Info("Reconnected!")
			for topic := range c.subscriptions {
				err = c.doSubscribe(topic)
				if err != nil {
					c.logger.Error("Error subscribing to topic", "topic", topic, "error", err.Error())
				}
			}
			for _, handler := range c.reconnectHandlers {
				// This must be a goroutine in case the handler relies on a blocking request over the websocket
				// Without, this blocks the listener from receiving the message and passing it back
				go handler()
			}
			return
		}

		c.logger.Error("Unable to reconnect", "error", err.Error())
		time.Sleep(5 * time.Second)
	}
}

// Sets the initial key pairs based on config
func (c *WebsocketClient) initialKeyPairs() error {
	var err error

	c.daemonKeyPair, err = c.config.DaemonSSL.LoadPrivateKeyPair(c.config.ChiaRoot)
	if err != nil {
		return err
	}

	return nil
}

func (c *WebsocketClient) generateDialer() error {
	if c.daemonDialer == nil {
		c.daemonDialer = &websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 45 * time.Second,
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{*c.daemonKeyPair},
				InsecureSkipVerify: true,
			},
		}
	}

	return nil
}

// ensureConnection ensures there is an open websocket connection and the listener is listening
func (c *WebsocketClient) ensureConnection() error {
	if c.conn == nil {
		c.resetOrigin()
		u := url.URL{Scheme: "wss", Host: fmt.Sprintf("%s:%d", c.baseURL.Host, c.daemonPort), Path: "/"}
		var err error
		c.conn, _, err = c.daemonDialer.Dial(u.String(), nil)
		if err != nil {
			return err
		}
	}

	go c.listen()

	return nil
}

// listen sets up a listener for all events and sends them back to handlerProxy
// The error returned from this function would only correspond to an error setting up the listener
// Errors returned by ReadMessage, or some other part of the websocket request/response will be
// passed to the handler to deal with
func (c *WebsocketClient) listen() {
	if !c.listenSyncActive {
		var ctx context.Context
		ctx, c.listenCancel = context.WithCancel(context.Background())
		c.listenSyncActive = true
		defer func() {
			c.listenSyncActive = false
		}()

		messageChan := make(chan []byte)

		// This reads messages from the websocket in the background allow us to either receive
		// a message OR cancel
		go func() {
			for {
				_, message, err := c.conn.ReadMessage()
				if err != nil {
					select {
					case <-ctx.Done():
						return
					default:
						c.logger.Error("Error reading message on chia websocket", "error", err.Error())
						var closeError *websocket.CloseError
						if !errors.As(err, &closeError) {
							c.logger.Debug("Chia websocket sent close message, attempting to close connection...")
							closeConnErr := c.conn.Close()
							if closeConnErr != nil {
								c.logger.Error("Error closing chia websocket connection", "error", closeConnErr.Error())
							}
						}
						c.conn = nil
						c.reconnectLoop()
						continue
					}
				}
				messageChan <- message
			}
		}()

		for {
			message := <-messageChan
			resp := &types.WebsocketResponse{}
			err := json.Unmarshal(message, resp)
			// Has to be called in goroutine so that the handler can potentially call cancel, which
			// this select needs to also read in order to properly cancel
			go c.handlerProxy(resp, err)
		}
	}
}
