package rpc

import (
	"net/http"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

// DaemonService encapsulates direct daemon RPC methods
type DaemonService struct {
	client *Client
}

// NewRequest returns a new request specific to the crawler service
func (s *DaemonService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceDaemon, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *DaemonService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetNetworkInfo gets the network name and prefix from the full node
func (s *DaemonService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
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
func (s *DaemonService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
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

// GetKeysOptions configures how keys are returned in get_keys
type GetKeysOptions struct {
	IncludeSecrets bool `json:"include_secrets"`
}

// GetKeysResponse response from get_keys RPC call
type GetKeysResponse struct {
	Response
	Keys []types.KeyData `json:"keys"`
}

// GetKeys returns key information
func (s *DaemonService) GetKeys(opts *GetKeysOptions) (*GetKeysResponse, *http.Response, error) {
	request, err := s.NewRequest("get_keys", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetKeysResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
