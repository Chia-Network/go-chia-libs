package rpc

import (
	"net/http"

	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/protocols"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

// FarmerService encapsulates farmer RPC methods
type FarmerService struct {
	client *Client
}

// NewRequest returns a new request specific to the wallet service
func (s *FarmerService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceFarmer, rpcEndpoint, opt)
}

// GetClient returns the active client for the service
func (s *FarmerService) GetClient() rpcinterface.Client {
	return s.client
}

// GetConnections returns connections
func (s *FarmerService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
	return Do(s, "get_connections", opts, &GetConnectionsResponse{})
}

// GetNetworkInfo gets the network name and prefix from the farmer
func (s *FarmerService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
	return Do(s, "get_network_info", opts, &GetNetworkInfoResponse{})
}

// GetVersion returns the application version for the service
func (s *FarmerService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
	return Do(s, "get_version", opts, &GetVersionResponse{})
}

// FarmerGetHarvestersOptions optoins for get_harvesters endpoint. Currently, accepts no options
type FarmerGetHarvestersOptions struct{}

// FarmerHarvester is a single harvester record returned by the farmer's get_harvesters endpoint
type FarmerHarvester struct {
	Connection struct {
		NodeID types.Bytes32 `json:"node_id"`
		Host   string        `json:"host"`
		Port   uint16        `json:"port"`
	} `json:"connection"`
	Plots                  []protocols.Plot `json:"plots"`
	FailedToOpenFilenames  []string         `json:"failed_to_open_filenames"`
	NoKeyFilenames         []string         `json:"no_key_filenames"`
	Duplicates             []string         `json:"duplicates"`
	TotalPlotSize          int              `json:"total_plot_size"`
	TotalEffectivePlotSize int              `json:"total_effective_plot_size"`
	Syncing                mo.Option[struct {
		Initial            bool   `json:"initial"`
		PlotFilesProcessed uint32 `json:"plot_files_processed"`
		PlotFilesTotal     uint32 `json:"plot_files_total"`
	}] `json:"syncing"`
	LastSyncTime   types.Timestamp                 `json:"last_sync_time"`
	HarvestingMode mo.Option[types.HarvestingMode] `json:"harvesting_mode"`
}

// FarmerGetHarvestersResponse get_harvesters response format
type FarmerGetHarvestersResponse struct {
	rpcinterface.Response
	Harvesters []FarmerHarvester `json:"harvesters"`
}

// GetHarvesters returns all harvester details for the farmer
func (s *FarmerService) GetHarvesters(opts *FarmerGetHarvestersOptions) (*FarmerGetHarvestersResponse, *http.Response, error) {
	return Do(s, "get_harvesters", opts, &FarmerGetHarvestersResponse{})
}
