package rpc

import (
	"net/http"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
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
