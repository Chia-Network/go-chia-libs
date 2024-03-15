package rpc

import (
	"net/http"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
)

// TimelordService encapsulates crawler RPC methods
type TimelordService struct {
	client *Client
}

// NewRequest returns a new request specific to the crawler service
func (s *TimelordService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceCrawler, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *TimelordService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetConnections returns connections
func (s *TimelordService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_connections", opts)
	if err != nil {
		return nil, nil, err
	}

	c := &GetConnectionsResponse{}
	resp, err := s.Do(request, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// GetNetworkInfo gets the network name and prefix from the full node
func (s *TimelordService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
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
