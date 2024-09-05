package rpc

import (
	"github.com/samber/mo"
)

// GetNetworkInfoOptions options for the get_network_info rpc calls
type GetNetworkInfoOptions struct{}

// GetNetworkInfoResponse common get_network_info response from all RPC services
type GetNetworkInfoResponse struct {
	Response
	NetworkName   mo.Option[string] `json:"network_name"`
	NetworkPrefix mo.Option[string] `json:"network_prefix"`
}

// GetVersionOptions options for the get_version rpc calls
type GetVersionOptions struct{}

// GetVersionResponse is the response of get_version from all RPC services
type GetVersionResponse struct {
	Response
	Version string `json:"version"`
}

// ServiceFullName are the full names to services that things like the daemon will recognize
type ServiceFullName string

const (
	// ServiceFullNameDaemon name of the daemon service
	ServiceFullNameDaemon ServiceFullName = "daemon"

	// ServiceFullNameDataLayer name of the data layer service
	ServiceFullNameDataLayer ServiceFullName = "chia_data_layer"

	// ServiceFullNameDataLayerHTTP name of data layer http service
	ServiceFullNameDataLayerHTTP ServiceFullName = "chia_data_layer_http"

	// ServiceFullNameWallet name of the wallet service
	ServiceFullNameWallet ServiceFullName = "chia_wallet"

	// ServiceFullNameNode name of the full node service
	ServiceFullNameNode ServiceFullName = "chia_full_node"

	// ServiceFullNameHarvester name of the harvester service
	ServiceFullNameHarvester ServiceFullName = "chia_harvester"

	// ServiceFullNameFarmer name of the farmer service
	ServiceFullNameFarmer ServiceFullName = "chia_farmer"

	// ServiceFullNameIntroducer name of the introducer service
	ServiceFullNameIntroducer ServiceFullName = "chia_introducer"

	// ServiceFullNameTimelord name of the timelord service
	ServiceFullNameTimelord ServiceFullName = "chia_timelord"

	// ServiceFullNameTimelordLauncher name of the timelord launcher service
	ServiceFullNameTimelordLauncher ServiceFullName = "chia_timelord_launcher"

	// ServiceFullNameSimulator name of the simulator service
	ServiceFullNameSimulator ServiceFullName = "chia_full_node_simulator"

	// ServiceFullNameSeeder name of the seeder service
	ServiceFullNameSeeder ServiceFullName = "chia_seeder"

	// ServiceFullNameCrawler name of the crawler service
	ServiceFullNameCrawler ServiceFullName = "chia_crawler"
)
