package rpc

// GetNetworkInfoOptions options for the get_network_info rpc calls
type GetNetworkInfoOptions struct {}

// GetNetworkInfoResponse common get_network_info response from all RPC services
type GetNetworkInfoResponse struct {
	Response
	NetworkName   string `json:"network_name"`
	NetworkPrefix string `json:"network_prefix"`
}
