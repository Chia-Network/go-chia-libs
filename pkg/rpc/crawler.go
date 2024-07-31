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

// Do is just a shortcut to the client's Do method
func (s *CrawlerService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetNetworkInfo gets the network name and prefix from the full node
func (s *CrawlerService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("get_network_info", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetNetworkInfoResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetVersion returns the application version for the service
func (s *CrawlerService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
	request, err := s.NewRequest("get_version", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetVersionResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetPeerCountsResponse Response for get_get_peer_counts on crawler
type GetPeerCountsResponse struct {
	Response
	PeerCounts mo.Option[types.CrawlerPeerCounts] `json:"peer_counts"`
}

// GetPeerCounts crawler rpc -> get_peer_counts
func (s *CrawlerService) GetPeerCounts() (*GetPeerCountsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_peer_counts", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetPeerCountsResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetIPsAfterTimestampOptions Options for the get_ips_after_timestamp RPC call
type GetIPsAfterTimestampOptions struct {
	After  int64 `json:"after"`
	Offset uint  `json:"offset"`
	Limit  uint  `json:"limit"`
}

// GetIPsAfterTimestampResponse Response for get_ips_after_timestamp
type GetIPsAfterTimestampResponse struct {
	Response
	IPs   mo.Option[[]string] `json:"ips"`
	Total mo.Option[int]      `json:"total"`
}

// GetIPsAfterTimestamp Returns IP addresses seen by the network after a particular timestamp
func (s *CrawlerService) GetIPsAfterTimestamp(opts *GetIPsAfterTimestampOptions) (*GetIPsAfterTimestampResponse, *http.Response, error) {
	request, err := s.NewRequest("get_ips_after_timestamp", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetIPsAfterTimestampResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
