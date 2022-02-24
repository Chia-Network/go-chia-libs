package httpclient

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/patrickmn/go-cache"
)

// CachedTransport is an http transport with cache on top
type CachedTransport struct {
	cache             *cache.Cache
	originalTransport http.RoundTripper
}

// NewCachedTransport returns a new transport wrapped in cache
func NewCachedTransport(expiration time.Duration, transport http.RoundTripper) *CachedTransport {
	return &CachedTransport{
		cache:             cache.New(expiration, expiration/2),
		originalTransport: transport,
	}
}

// key returns the cache key for the request
func (c *CachedTransport) key(r *http.Request) string {
	method := r.Method
	url := r.URL.String()
	body := ""
	if r.Body != nil || r.Body != http.NoBody {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		body = fmt.Sprintf("%x", sha256.Sum256(bodyBytes))
	}

	return fmt.Sprintf("%s%s%s", method, url, body)
}

// RoundTrip executes a single HTTP transaction, returning
// a Response for the provided Request.
func (c *CachedTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	// MUST get this now, or else the body will be read and no longer available at the end
	cacheKey := c.key(r)

	// If the response is cached, we can just respond with the cached version
	if cached, found := c.cache.Get(cacheKey); found != false {
		cachedBytes := cached.([]byte)
		return c.cachedResponse(cachedBytes, r)
	}

	resp, err := c.originalTransport.RoundTrip(r)

	if err != nil {
		return nil, err
	}

	// Grab the body and stick it in the cache
	buf, err := httputil.DumpResponse(resp, true)

	if err != nil {
		return nil, err
	}

	// Add the response bytes to cache
	c.cache.Set(cacheKey, buf, cache.DefaultExpiration)

	return resp, nil
}

// cachedResponse returns a response built from cached data
func (c *CachedTransport) cachedResponse(b []byte, r *http.Request) (*http.Response, error) {
	buf := bytes.NewBuffer(b)
	return http.ReadResponse(bufio.NewReader(buf), r)
}
