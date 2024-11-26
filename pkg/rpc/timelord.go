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

// GetClient returns the active client for the service
func (s *TimelordService) GetClient() rpcinterface.Client {
	return s.client
}

// GetConnections returns connections
func (s *TimelordService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
	return Do(s, "get_connections", opts, &GetConnectionsResponse{})
}

// GetNetworkInfo gets the network name and prefix from the full node
func (s *TimelordService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
	return Do(s, "get_network_info", opts, &GetNetworkInfoResponse{})
}

// GetVersion returns the application version for the service
func (s *TimelordService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
	return Do(s, "get_version", opts, &GetVersionResponse{})
}
