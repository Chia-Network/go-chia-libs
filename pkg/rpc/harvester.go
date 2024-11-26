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

// Do is just a shortcut to the client's Do method
func (s *HarvesterService) Do(req *rpcinterface.Request, v iResponse) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetConnections returns connections
func (s *HarvesterService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_connections", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetConnectionsResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}

// GetNetworkInfo gets the network name and prefix from the harvester
func (s *HarvesterService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("get_network_info", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetNetworkInfoResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}

// GetVersion returns the application version for the service
func (s *HarvesterService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
	request, err := s.NewRequest("get_version", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetVersionResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}

// HarvesterGetPlotsResponse get_plots response format
type HarvesterGetPlotsResponse struct {
	Response
	Plots                 mo.Option[[]protocols.Plot] `json:"plots"`
	FailedToOpenFilenames mo.Option[[]string]         `json:"failed_to_open_filenames"`
	NotFoundFilenames     mo.Option[[]string]         `json:"not_found_filenames"`
}

// GetPlots returns connections
func (s *HarvesterService) GetPlots() (*HarvesterGetPlotsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_plots", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &HarvesterGetPlotsResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}
