package rpc

import (
	"net/http"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
)

// TimelordService encapsulates timelord RPC methods
type TimelordService struct {
	client *Client
}

// NewRequest returns a new request specific to the crawler service
func (s *TimelordService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceTimelord, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *TimelordService) Do(req *rpcinterface.Request, v iResponse) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetConnections returns connections
func (s *TimelordService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_connections", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetConnectionsResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}

// GetNetworkInfo gets the network name and prefix from the full node
func (s *TimelordService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("get_network_info", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetNetworkInfoResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}

// GetVersion returns the application version for the service
func (s *TimelordService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
	request, err := s.NewRequest("get_version", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetVersionResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}
