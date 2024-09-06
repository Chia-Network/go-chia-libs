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

// StartServiceOptions start service options
type StartServiceOptions struct {
	Service ServiceFullName `json:"service"`
}

// StartServiceResponse start service response
type StartServiceResponse struct {
	Response
	Service ServiceFullName `json:"service"`
}

// StartService starts the given service
func (s *DaemonService) StartService(opts *StartServiceOptions) (*StartServiceResponse, *http.Response, error) {
	request, err := s.NewRequest("start_service", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &StartServiceResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// StopServiceOptions start service options
type StopServiceOptions struct {
	Service ServiceFullName `json:"service"`
}

// StopServiceResponse stop service response
type StopServiceResponse struct {
	Response
	Service ServiceFullName `json:"service"`
}

// StopService stops the given service
func (s *DaemonService) StopService(opts *StopServiceOptions) (*StopServiceResponse, *http.Response, error) {
	request, err := s.NewRequest("stop_service", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &StopServiceResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// IsRunningOptions is service running options
type IsRunningOptions struct {
	Service ServiceFullName `json:"service"`
}

// IsRunningResponse is service running response
type IsRunningResponse struct {
	Response
	ServiceName ServiceFullName `json:"service_name"`
	IsRunning   bool            `json:"is_running"`
}

// IsRunning returns whether a service is running
func (s *DaemonService) IsRunning(opts *IsRunningOptions) (*IsRunningResponse, *http.Response, error) {
	request, err := s.NewRequest("is_running", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &IsRunningResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// DaemonDeleteAllKeysOpts options for delete all keys request
type DaemonDeleteAllKeysOpts struct {}

// DaemonDeleteAllKeysResponse response when deleting all keys
type DaemonDeleteAllKeysResponse struct {
	Response
}

// DeleteAllKeys deletes all keys from the keychain
func (s *DaemonService) DeleteAllKeys(opts *DaemonDeleteAllKeysOpts) (*DaemonDeleteAllKeysResponse, *http.Response, error) {
	request, err := s.NewRequest("delete_all_keys", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &DaemonDeleteAllKeysResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
