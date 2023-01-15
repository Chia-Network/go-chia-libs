package rpc

import (
	"net/http"

	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
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
func (s *HarvesterService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetConnections returns connections
func (s *HarvesterService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
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

// HarvesterGetPlotsResponse get_plots response format
type HarvesterGetPlotsResponse struct {
	Response
	Plots                 mo.Option[[]types.PlotInfo] `json:"plots"`
	FailedToOpenFilenames mo.Option[[]string]         `json:"failed_to_open_filenames"`
	NotFoundFilenames     mo.Option[[]string]         `json:"not_found_filenames"`
}

// GetPlots returns connections
func (s *HarvesterService) GetPlots() (*HarvesterGetPlotsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_plots", nil)
	if err != nil {
		return nil, nil, err
	}

	p := &HarvesterGetPlotsResponse{}
	resp, err := s.Do(request, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}
