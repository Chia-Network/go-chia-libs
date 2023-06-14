package httpclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
)

// HTTPClient connects to Chia RPC via standard HTTP requests
type HTTPClient struct {
	config  *config.ChiaConfig
	baseURL *url.URL

	// If set > 0, will configure http requests with a cache
	cacheValidTime time.Duration

	// Request timeout
	Timeout time.Duration

	nodePort    uint16
	nodeKeyPair *tls.Certificate
	nodeClient  *http.Client

	farmerPort    uint16
	farmerKeyPair *tls.Certificate
	farmerClient  *http.Client

	harvesterPort    uint16
	harvesterKeyPair *tls.Certificate
	harvesterClient  *http.Client

	walletPort    uint16
	walletKeyPair *tls.Certificate
	walletClient  *http.Client

	crawlerPort    uint16
	crawlerKeyPair *tls.Certificate
	crawlerClient  *http.Client

	datalayerPort    uint16
	datalayerKeyPair *tls.Certificate
	datalayerClient  *http.Client
}

// NewHTTPClient returns a new HTTP client that satisfies the rpcinterface.Client interface
func NewHTTPClient(cfg *config.ChiaConfig, options ...rpcinterface.ClientOptionFunc) (*HTTPClient, error) {
	c := &HTTPClient{
		config: cfg,

		Timeout: 10 * time.Second, // Default, overridable with client option

		nodePort:      cfg.FullNode.RPCPort,
		farmerPort:    cfg.Farmer.RPCPort,
		harvesterPort: cfg.Harvester.RPCPort,
		walletPort:    cfg.Wallet.RPCPort,
		crawlerPort:   cfg.Seeder.CrawlerConfig.RPCPort,
		datalayerPort: cfg.DataLayer.RPCPort,
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

	u.Host = fmt.Sprintf("%s:%d", u.Host, c.portForService(service))

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

func (c *HTTPClient) generateHTTPClientForService(service rpcinterface.ServiceType) (*http.Client, error) {
	var (
		keyPair *tls.Certificate
		err     error
	)

	switch service {
	case rpcinterface.ServiceFullNode:
		if c.nodeKeyPair == nil {
			c.nodeKeyPair, err = c.config.FullNode.SSL.LoadPrivateKeyPair(c.config.ChiaRoot)
			if err != nil {
				return nil, err
			}
		}
		keyPair = c.nodeKeyPair
	case rpcinterface.ServiceFarmer:
		if c.farmerKeyPair == nil {
			c.farmerKeyPair, err = c.config.Farmer.SSL.LoadPrivateKeyPair(c.config.ChiaRoot)
			if err != nil {
				return nil, err
			}
		}
		keyPair = c.farmerKeyPair
	case rpcinterface.ServiceHarvester:
		if c.harvesterKeyPair == nil {
			c.harvesterKeyPair, err = c.config.Harvester.SSL.LoadPrivateKeyPair(c.config.ChiaRoot)
			if err != nil {
				return nil, err
			}
		}
		keyPair = c.harvesterKeyPair
	case rpcinterface.ServiceWallet:
		if c.walletKeyPair == nil {
			c.walletKeyPair, err = c.config.Wallet.SSL.LoadPrivateKeyPair(c.config.ChiaRoot)
			if err != nil {
				return nil, err
			}
		}
		keyPair = c.walletKeyPair
	case rpcinterface.ServiceCrawler:
		if c.crawlerKeyPair == nil {
			c.crawlerKeyPair, err = c.config.Seeder.CrawlerConfig.SSL.LoadPrivateKeyPair(c.config.ChiaRoot)
			if err != nil {
				// Fall back to just using the full node certs in this case
				// This should only happen on old installations that didn't have the crawler in the config initially
				c.crawlerKeyPair, err = c.config.FullNode.SSL.LoadPrivateKeyPair(c.config.ChiaRoot)
				if err != nil {
					return nil, err
				}
			}
		}
		keyPair = c.crawlerKeyPair
	case rpcinterface.ServiceDataLayer:
		if c.datalayerKeyPair == nil {
			c.datalayerKeyPair, err = c.config.DataLayer.SSL.LoadPrivateKeyPair(c.config.ChiaRoot)
			if err != nil {
				return nil, err
			}
		}
		keyPair = c.datalayerKeyPair
	default:
		return nil, fmt.Errorf("unknown service")
	}

	var transport http.RoundTripper

	transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{*keyPair},
			InsecureSkipVerify: true, // Cert is apparently for chia.net - can't validate until it matches hostname
		},
	}

	if c.cacheValidTime > 0 {
		transport = NewCachedTransport(c.cacheValidTime, transport)
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   c.Timeout,
	}

	return client, nil
}

// portForService returns the configured port for the service
func (c *HTTPClient) portForService(service rpcinterface.ServiceType) uint16 {
	var port uint16 = 0

	switch service {
	case rpcinterface.ServiceFullNode:
		port = c.nodePort
	case rpcinterface.ServiceFarmer:
		port = c.farmerPort
	case rpcinterface.ServiceHarvester:
		port = c.harvesterPort
	case rpcinterface.ServiceWallet:
		port = c.walletPort
	case rpcinterface.ServiceCrawler:
		port = c.crawlerPort
	case rpcinterface.ServiceDataLayer:
		port = c.datalayerPort
	}

	return port
}

// httpClientForService returns the proper http client to use with the service
func (c *HTTPClient) httpClientForService(service rpcinterface.ServiceType) (*http.Client, error) {
	var (
		client *http.Client
		err    error
	)

	switch service {
	case rpcinterface.ServiceFullNode:
		if c.nodeClient == nil {
			c.nodeClient, err = c.generateHTTPClientForService(rpcinterface.ServiceFullNode)
			if err != nil {
				return nil, err
			}
		}
		client = c.nodeClient
	case rpcinterface.ServiceFarmer:
		if c.farmerClient == nil {
			c.farmerClient, err = c.generateHTTPClientForService(rpcinterface.ServiceFarmer)
			if err != nil {
				return nil, err
			}
		}
		client = c.farmerClient
	case rpcinterface.ServiceHarvester:
		if c.harvesterClient == nil {
			c.harvesterClient, err = c.generateHTTPClientForService(rpcinterface.ServiceHarvester)
			if err != nil {
				return nil, err
			}
		}
		client = c.harvesterClient
	case rpcinterface.ServiceWallet:
		if c.walletClient == nil {
			c.walletClient, err = c.generateHTTPClientForService(rpcinterface.ServiceWallet)
			if err != nil {
				return nil, err
			}
		}
		client = c.walletClient
	case rpcinterface.ServiceCrawler:
		if c.crawlerClient == nil {
			c.crawlerClient, err = c.generateHTTPClientForService(rpcinterface.ServiceCrawler)
			if err != nil {
				return nil, err
			}
		}
		client = c.crawlerClient
	case rpcinterface.ServiceDataLayer:
		if c.datalayerClient == nil {
			c.datalayerClient, err = c.generateHTTPClientForService(rpcinterface.ServiceDataLayer)
			if err != nil {
				return nil, err
			}
		}
		client = c.datalayerClient
	}

	if client == nil {
		return nil, fmt.Errorf("unknown service")
	}

	return client, nil
}

// The following are here to satisfy the interface, but are not used by the HTTP client

// SubscribeSelf subscribes to events in response to requests from this service
// Not applicable on the HTTP connection
func (c *HTTPClient) SubscribeSelf() error {
	return nil
}

// Subscribe adds a subscription to events from a particular service
// Not applicable on the HTTP connection
func (c *HTTPClient) Subscribe(service string) error {
	return nil
}

// ListenSync Listens for async responses over the connection in a synchronous fashion, blocking anything else
// Not applicable on the HTTP connection
func (c *HTTPClient) ListenSync(handler rpcinterface.WebsocketResponseHandler) error {
	return nil
}

// AddDisconnectHandler Not applicable to the HTTP client
func (c *HTTPClient) AddDisconnectHandler(onDisconnect rpcinterface.DisconnectHandler) {}

// AddReconnectHandler Not applicable to the HTTP client
func (c *HTTPClient) AddReconnectHandler(onReconnect rpcinterface.ReconnectHandler) {}
