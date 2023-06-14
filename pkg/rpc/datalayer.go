package rpc

import (
	"net/http"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
)

// DataLayerService encapsulates data layer RPC methods
type DataLayerService struct {
	client *Client
}

// NewRequest returns a new request specific to the wallet service
func (s *DataLayerService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceDataLayer, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *DataLayerService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}
