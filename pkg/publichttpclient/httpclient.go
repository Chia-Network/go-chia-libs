package publichttpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/google/uuid"

	"github.com/chia-network/go-chia-libs/pkg/httpclient"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
)

// HTTPClient connects to Chia RPC via standard HTTP requests
type HTTPClient struct {
	baseURL *url.URL
	logger  *slog.Logger

	// If set > 0, will configure http requests with a cache
	cacheValidTime time.Duration

	// Request timeout
	Timeout time.Duration

	nodeClient *http.Client
}

// NewHTTPClient returns a new HTTP client that satisfies the rpcinterface.Client interface for use with public RPC services
func NewHTTPClient(options ...rpcinterface.ClientOptionFunc) (*HTTPClient, error) {
	c := &HTTPClient{
		logger: slog.New(rpcinterface.SlogInfo()),

		Timeout: 10 * time.Second, // Default, overridable with client option
	}

	// Sets the default host. Can be overridden by client options
	err := c.SetBaseURL(&url.URL{
		Scheme: "https",
		Host:   "localhost",
	})
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

	return c, nil
}

// SetBaseURL sets the base URL for API requests to a custom endpoint.
func (c *HTTPClient) SetBaseURL(url *url.URL) error {
	c.baseURL = url

	return nil
}

// SetLogHandler sets a slog compatible log handler
func (c *HTTPClient) SetLogHandler(handler slog.Handler) {
	c.logger = slog.New(handler)
}

// SetCacheValidTime sets how long cache should be valid for
func (c *HTTPClient) SetCacheValidTime(validTime time.Duration) {
	c.cacheValidTime = validTime
}

// NewRequest creates an RPC request for the specified service
func (c *HTTPClient) NewRequest(service rpcinterface.ServiceType, rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	// Always POST
	// Supporting it as a variable in case that changes in the future, it can be passed in instead
	method := http.MethodPost

	u := *c.baseURL

	u.RawPath = fmt.Sprintf("/%s", rpcEndpoint)
	u.Path = fmt.Sprintf("/%s", rpcEndpoint)

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	var body []byte
	var err error
	switch {
	case method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		// Always need at least an empty json object in the body
		if opt == nil {
			body = []byte(`{}`)
		} else {
			body, err = json.Marshal(opt)
			if err != nil {
				return nil, err
			}
		}
	case opt != nil:
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return &rpcinterface.Request{
		Service: service,
		Request: req,
	}, nil
}

// Do sends an RPC request and returns the RPC response.
func (c *HTTPClient) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	client, err := c.httpClientForService(req.Service)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req.Request)
	if err != nil {
		return nil, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

// Close closes the connection. Not applicable for the HTTPClient, so just returns no error
func (c *HTTPClient) Close() error {
	return nil
}

func (c *HTTPClient) generateHTTPClientForService(service rpcinterface.ServiceType) (*http.Client, error) {
	var transport http.RoundTripper

	transport = &http.Transport{}

	if c.cacheValidTime > 0 {
		transport = httpclient.NewCachedTransport(c.cacheValidTime, transport)
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   c.Timeout,
	}

	return client, nil
}

// httpClientForService returns the proper http client to use with the service
func (c *HTTPClient) httpClientForService(service rpcinterface.ServiceType) (*http.Client, error) {
	var (
		client *http.Client
		err    error
	)

	switch service {
	case rpcinterface.ServiceDaemon:
		return nil, fmt.Errorf("daemon RPC calls must be made with the websocket client")
	case rpcinterface.ServiceFullNode:
		if c.nodeClient == nil {
			c.nodeClient, err = c.generateHTTPClientForService(rpcinterface.ServiceFullNode)
			if err != nil {
				return nil, err
			}
		}
		client = c.nodeClient
	}

	if client == nil {
		return nil, fmt.Errorf("unknown service")
	}

	return client, nil
}

// The following are here to satisfy the interface, but are not used by the HTTP client

// SubscribeSelf does not apply to the HTTP Client
func (c *HTTPClient) SubscribeSelf() error {
	return fmt.Errorf("subscriptions are not supported on the HTTP client - websockets are required for subscriptions")
}

// Subscribe does not apply to the HTTP Client
// Not applicable on the HTTP connection
func (c *HTTPClient) Subscribe(service string) error {
	return fmt.Errorf("subscriptions are not supported on the HTTP client - websockets are required for subscriptions")
}

// AddHandler does not apply to HTTP Client
func (c *HTTPClient) AddHandler(handler rpcinterface.WebsocketResponseHandler) (uuid.UUID, error) {
	return uuid.Nil, fmt.Errorf("handlers are not supported on the HTTP client - reponses are returned directly from the calling functions")
}

// RemoveHandler does not apply to HTTP Client
func (c *HTTPClient) RemoveHandler(handlerID uuid.UUID) {}

// AddDisconnectHandler does not apply to the HTTP Client
func (c *HTTPClient) AddDisconnectHandler(onDisconnect rpcinterface.DisconnectHandler) {}

// AddReconnectHandler does not apply to the HTTP Client
func (c *HTTPClient) AddReconnectHandler(onReconnect rpcinterface.ReconnectHandler) {}

// SetSyncMode does not apply to the HTTP Client
func (c *HTTPClient) SetSyncMode() {
	c.logger.Debug("Sync mode is default for HTTP client. SetSyncMode call is ignored")
}

// SetAsyncMode does not apply to the HTTP Client
func (c *HTTPClient) SetAsyncMode() {
	c.logger.Debug("Async mode is not applicable to the HTTP client. SetAsyncMode call is ignored")
}
