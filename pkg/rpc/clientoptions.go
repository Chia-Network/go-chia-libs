package rpc

import (
	"net/url"
	"time"

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
		c.SetCacheValidTime(validTime)

		return nil
	}
}
