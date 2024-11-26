package rpc

import (
	"net/http"

	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

// CrawlerService encapsulates crawler RPC methods
type CrawlerService struct {
	client *Client
}

// NewRequest returns a new request specific to the crawler service
func (s *CrawlerService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceCrawler, rpcEndpoint, opt)
}

// GetClient returns the active client for the service
func (s *CrawlerService) GetClient() rpcinterface.Client {
	return s.client
}

// GetNetworkInfo gets the network name and prefix from the full node
func (s *CrawlerService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
	return Do(s, "get_network_info", opts, &GetNetworkInfoResponse{})
}

// GetVersion returns the application version for the service
func (s *CrawlerService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
	return Do(s, "get_version", opts, &GetVersionResponse{})
}

// GetPeerCountsResponse Response for get_get_peer_counts on crawler
type GetPeerCountsResponse struct {
	rpcinterface.Response
	PeerCounts mo.Option[types.CrawlerPeerCounts] `json:"peer_counts"`
}

// GetPeerCounts crawler rpc -> get_peer_counts
func (s *CrawlerService) GetPeerCounts() (*GetPeerCountsResponse, *http.Response, error) {
	return Do(s, "get_peer_counts", nil, &GetPeerCountsResponse{})
}

// GetIPsAfterTimestampOptions Options for the get_ips_after_timestamp RPC call
type GetIPsAfterTimestampOptions struct {
	After  int64 `json:"after"`
	Offset uint  `json:"offset"`
	Limit  uint  `json:"limit"`
}

// GetIPsAfterTimestampResponse Response for get_ips_after_timestamp
type GetIPsAfterTimestampResponse struct {
	rpcinterface.Response
	IPs   mo.Option[[]string] `json:"ips"`
	Total mo.Option[int]      `json:"total"`
}

// GetIPsAfterTimestamp Returns IP addresses seen by the network after a particular timestamp
func (s *CrawlerService) GetIPsAfterTimestamp(opts *GetIPsAfterTimestampOptions) (*GetIPsAfterTimestampResponse, *http.Response, error) {
	return Do(s, "get_ips_after_timestamp", opts, &GetIPsAfterTimestampResponse{})
}
