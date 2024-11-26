package rpc

import (
	"net/http"

	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/protocols"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
)

// HarvesterService encapsulates harvester RPC methods
type HarvesterService struct {
	client *Client
}

// NewRequest returns a new request specific to the wallet service
func (s *HarvesterService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceHarvester, rpcEndpoint, opt)
}

// GetClient returns the active client for the service
func (s *HarvesterService) GetClient() rpcinterface.Client {
	return s.client
}

// GetConnections returns connections
func (s *HarvesterService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
	return Do(s, "get_connections", opts, &GetConnectionsResponse{})
}

// GetNetworkInfo gets the network name and prefix from the harvester
func (s *HarvesterService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
	return Do(s, "get_network_info", opts, &GetNetworkInfoResponse{})
}

// GetVersion returns the application version for the service
func (s *HarvesterService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
	return Do(s, "get_version", opts, &GetVersionResponse{})
}

// HarvesterGetPlotsResponse get_plots response format
type HarvesterGetPlotsResponse struct {
	rpcinterface.Response
	Plots                 mo.Option[[]protocols.Plot] `json:"plots"`
	FailedToOpenFilenames mo.Option[[]string]         `json:"failed_to_open_filenames"`
	NotFoundFilenames     mo.Option[[]string]         `json:"not_found_filenames"`
}

// GetPlots returns connections
func (s *HarvesterService) GetPlots() (*HarvesterGetPlotsResponse, *http.Response, error) {
	return Do(s, "get_plots", nil, &HarvesterGetPlotsResponse{})
}
