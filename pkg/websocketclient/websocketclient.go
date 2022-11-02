package websocketclient

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"

	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

const origin string = "go-chia-rpc"

// WebsocketClient connects to Chia RPC via websockets
type WebsocketClient struct {
	config  *config.ChiaConfig
	baseURL *url.URL

	daemonPort    uint16
	daemonKeyPair *tls.Certificate
	daemonDialer  *websocket.Dialer

	conn *websocket.Conn

	listenSyncActive bool

	// subscriptions Keeps track of subscribed topics, so we can re-subscribe if we lose a connection and reconnect
	subscriptions []string

	disconnectHandlers []rpcinterface.DisconnectHandler
	reconnectHandlers  []rpcinterface.ReconnectHandler
}

// NewWebsocketClient returns a new websocket client that satisfies the rpcinterface.Client interface
func NewWebsocketClient(cfg *config.ChiaConfig, options ...rpcinterface.ClientOptionFunc) (*WebsocketClient, error) {
	c := &WebsocketClient{
		config: cfg,

		daemonPort: cfg.DaemonPort,
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
		if err := fn(c); err != nil {
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

// SetBaseURL sets the base URL for API requests to a custom endpoint.
func (c *WebsocketClient) SetBaseURL(url *url.URL) error {
	c.baseURL = url

	return nil
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
// *http.Response is always nil in this return, and exists to satisfy the interface that existed prior to
// websockets being supported in this library
func (c *WebsocketClient) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	err := c.ensureConnection()
	if err != nil {
		return nil, err
	}

	var destination string
	switch req.Service {
	case rpcinterface.ServiceDaemon:
		destination = "daemon"
	case rpcinterface.ServiceFullNode:
		destination = "chia_full_node"
	case rpcinterface.ServiceFarmer:
		destination = "chia_farmer" // @TODO validate the correct string for this
	case rpcinterface.ServiceHarvester:
		destination = "chia_harvester" // @TODO validate the correct string for this
	case rpcinterface.ServiceWallet:
		destination = "chia_wallet"
	case rpcinterface.ServiceCrawler:
		destination = "chia_crawler"
	default:
		return nil, fmt.Errorf("unknown service")
	}

	data := req.Data
	if data == nil {
		data = map[string]interface{}{}
	}
	request := &types.WebsocketRequest{
		Command:     string(req.Endpoint),
		Origin:      origin,
		Destination: destination,
		Data:        data,
	}

	return nil, c.conn.WriteJSON(request)
}

// SubscribeSelf calls subscribe for any requests that this client makes to the server
// Different from Subscribe with a custom service - that is more for subscribing to built in events emitted by Chia
// This call will subscribe `go-chia-rpc` origin for any requests we specifically make of the server
func (c *WebsocketClient) SubscribeSelf() error {
	return c.Subscribe(origin)
}

// Subscribe adds a subscription to a particular service
func (c *WebsocketClient) Subscribe(service string) error {
	c.subscriptions = append(c.subscriptions, service)

	return c.doSubscribe(service)
}

func (c *WebsocketClient) doSubscribe(service string) error {
	request, err := c.NewRequest(rpcinterface.ServiceDaemon, "register_service", types.WebsocketSubscription{Service: service})
	if err != nil {
		return err
	}

	_, err = c.Do(request, nil)
	return err
}

// ListenSync Listens for responses over the websocket connection in the foreground
// The error returned from this function would only correspond to an error setting up the listener
// Errors returned by ReadMessage, or some other part of the websocket request/response will be
// passed to the handler to deal with
func (c *WebsocketClient) ListenSync(handler rpcinterface.WebsocketResponseHandler) error {
	if !c.listenSyncActive {
		c.listenSyncActive = true

		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println(err.Error())
				if _, isCloseErr := err.(*websocket.CloseError); !isCloseErr {
					closeConnErr := c.conn.Close()
					if closeConnErr != nil {
						log.Printf("Error closing connection after error: %s\n", closeConnErr.Error())
					}
				}
				c.conn = nil
				c.reconnectLoop()
				continue
			}
			resp := &types.WebsocketResponse{}
			err = json.Unmarshal(message, resp)
			handler(resp, err)
		}
	}

	return nil
}

// AddDisconnectHandler the function to call when the client is disconnected
func (c *WebsocketClient) AddDisconnectHandler(onDisconnect rpcinterface.DisconnectHandler) {
	c.disconnectHandlers = append(c.disconnectHandlers, onDisconnect)
}

// AddReconnectHandler the function to call when the client is reconnected after a disconnect
func (c *WebsocketClient) AddReconnectHandler(onReconnect rpcinterface.ReconnectHandler) {
	c.reconnectHandlers = append(c.reconnectHandlers, onReconnect)
}

func (c *WebsocketClient) reconnectLoop() {
	for _, handler := range c.disconnectHandlers {
		handler()
	}
	for {
		log.Println("Trying to reconnect...")
		err := c.ensureConnection()
		if err == nil {
			log.Println("Reconnected!")
			for _, topic := range c.subscriptions {
				err = c.doSubscribe(topic)
				if err != nil {
					log.Printf("Error subscribing to topic %s: %s\n", topic, err.Error())
				}
			}
			for _, handler := range c.reconnectHandlers {
				handler()
			}
			return
		}

		log.Printf("Unable to reconnect: %s\n", err.Error())
		time.Sleep(5 * time.Second)
	}
}

// Sets the initial key pairs based on config
func (c *WebsocketClient) initialKeyPairs() error {
	var err error

	c.daemonKeyPair, err = c.config.DaemonSSL.LoadPrivateKeyPair()
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

// ensureConnection ensures there is an open websocket connection
func (c *WebsocketClient) ensureConnection() error {
	if c.conn == nil {
		u := url.URL{Scheme: "wss", Host: fmt.Sprintf("%s:%d", c.baseURL.Host, c.daemonPort), Path: "/"}
		var err error
		c.conn, _, err = c.daemonDialer.Dial(u.String(), nil)
		if err != nil {
			return err
		}
	}

	return nil
}
