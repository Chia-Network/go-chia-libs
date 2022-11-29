package rpc

import (
	"net/url"
	"time"

	"github.com/chia-network/go-chia-libs/pkg/httpclient"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
)

// WithBaseURL sets the host for RPC requests
func WithBaseURL(url *url.URL) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		return c.SetBaseURL(url)
	}
}

// WithCache specify a duration http requests should be cached for
// If unset, cache will not be used
func WithCache(validTime time.Duration) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		typed, ok := c.(*httpclient.HTTPClient)
		if ok {
			typed.SetCacheValidTime(validTime)
		}

		return nil
	}
}

// WithTimeout sets the timeout for the requests
func WithTimeout(timeout time.Duration) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		typed, ok := c.(*httpclient.HTTPClient)
		if ok {
			typed.Timeout = timeout
		}

		return nil
	}
}

// WithDaemonPort sets the port for RPC requests
func WithDaemonPort(port uint16) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		return c.SetDaemonPort(port)
	}
}

// WithNodePort sets the port for RPC requests
func WithNodePort(port uint16) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		return c.SetNodePort(port)
	}
}

// WithFarmerPort sets the port for RPC requests
func WithFarmerPort(port uint16) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		return c.SetFarmerPort(port)
	}
}

// WithHarvesterPort sets the port for RPC requests
func WithHarvesterPort(port uint16) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		return c.SetHarvesterPort(port)
	}
}

// WithWalletPort sets the port for RPC requests
func WithWalletPort(port uint16) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		return c.SetWalletPort(port)
	}
}

// WithCrawlerPort sets the port for RPC requests
func WithCrawlerPort(port uint16) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		return c.SetCrawlerPort(port)
	}
}
