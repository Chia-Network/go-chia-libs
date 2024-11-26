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

// GetClient returns the active client for the service
func (s *DaemonService) GetClient() rpcinterface.Client {
	return s.client
}

// GetNetworkInfo gets the network name and prefix from the full node
func (s *DaemonService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
	return Do(s, "get_network_info", opts, &GetNetworkInfoResponse{})
}

// GetVersion returns the application version for the service
func (s *DaemonService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
	return Do(s, "get_version", opts, &GetVersionResponse{})
}

// GetKeysOptions configures how keys are returned in get_keys
type GetKeysOptions struct {
	IncludeSecrets bool `json:"include_secrets"`
}

// GetKeysResponse response from get_keys RPC call
type GetKeysResponse struct {
	rpcinterface.Response
	Keys []types.KeyData `json:"keys"`
}

// GetKeys returns key information
func (s *DaemonService) GetKeys(opts *GetKeysOptions) (*GetKeysResponse, *http.Response, error) {
	return Do(s, "get_keys", opts, &GetKeysResponse{})
}

// StartServiceOptions start service options
type StartServiceOptions struct {
	Service ServiceFullName `json:"service"`
}

// StartServiceResponse start service response
type StartServiceResponse struct {
	rpcinterface.Response
	Service ServiceFullName `json:"service"`
}

// StartService starts the given service
func (s *DaemonService) StartService(opts *StartServiceOptions) (*StartServiceResponse, *http.Response, error) {
	return Do(s, "start_service", opts, &StartServiceResponse{})
}

// StopServiceOptions start service options
type StopServiceOptions struct {
	Service ServiceFullName `json:"service"`
}

// StopServiceResponse stop service response
type StopServiceResponse struct {
	rpcinterface.Response
	Service ServiceFullName `json:"service"`
}

// StopService stops the given service
func (s *DaemonService) StopService(opts *StopServiceOptions) (*StopServiceResponse, *http.Response, error) {
	return Do(s, "stop_service", opts, &StopServiceResponse{})
}

// IsRunningOptions is service running options
type IsRunningOptions struct {
	Service ServiceFullName `json:"service"`
}

// IsRunningResponse is service running response
type IsRunningResponse struct {
	rpcinterface.Response
	ServiceName ServiceFullName `json:"service_name"`
	IsRunning   bool            `json:"is_running"`
}

// IsRunning returns whether a service is running
func (s *DaemonService) IsRunning(opts *IsRunningOptions) (*IsRunningResponse, *http.Response, error) {
	return Do(s, "is_running", opts, &IsRunningResponse{})
}

// RunningServicesResponse is service running response
type RunningServicesResponse struct {
	rpcinterface.Response
	RunningServices []ServiceFullName `json:"running_services"`
}

// RunningServices returns all running services
func (s *DaemonService) RunningServices() (*RunningServicesResponse, *http.Response, error) {
	return Do(s, "running_services", nil, &RunningServicesResponse{})
}

// ExitResponse shows information about the services that were stopped
type ExitResponse struct {
	rpcinterface.Response
	ServicesStopped []ServiceFullName `json:"services_stopped"`
}

// Exit tells the daemon to exit
func (s *DaemonService) Exit() (*ExitResponse, *http.Response, error) {
	return Do(s, "exit", nil, &ExitResponse{})
}

// DaemonDeleteAllKeysOpts options for delete all keys request
type DaemonDeleteAllKeysOpts struct{}

// DaemonDeleteAllKeysResponse response when deleting all keys
type DaemonDeleteAllKeysResponse struct {
	rpcinterface.Response
}

// DeleteAllKeys deletes all keys from the keychain
func (s *DaemonService) DeleteAllKeys(opts *DaemonDeleteAllKeysOpts) (*DaemonDeleteAllKeysResponse, *http.Response, error) {
	return Do(s, "delete_all_keys", opts, &DaemonDeleteAllKeysResponse{})
}
