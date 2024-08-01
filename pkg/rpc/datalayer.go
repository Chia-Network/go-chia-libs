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

// GetVersion returns the application version for the service
func (s *DataLayerService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
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

// DatalayerGetSubscriptionsOptions options for get_subscriptions
type DatalayerGetSubscriptionsOptions struct{}

// DatalayerGetSubscriptionsResponse response for get_subscriptions
type DatalayerGetSubscriptionsResponse struct {
	Response
	StoreIDs []string `json:"store_ids"`
}

// GetSubscriptions is just an alias for Subscriptions, since the CLI command is get_subscriptions
// Makes this easier to find
func (s *DataLayerService) GetSubscriptions(opts *DatalayerGetSubscriptionsOptions) (*DatalayerGetSubscriptionsResponse, *http.Response, error) {
	return s.Subscriptions(opts)
}

// Subscriptions calls the subscriptions endpoint to list all subscriptions
func (s *DataLayerService) Subscriptions(opts *DatalayerGetSubscriptionsOptions) (*DatalayerGetSubscriptionsResponse, *http.Response, error) {
	request, err := s.NewRequest("subscriptions", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &DatalayerGetSubscriptionsResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// DatalayerGetOwnedStoresOptions Options for get_owned_stores
type DatalayerGetOwnedStoresOptions struct {}

// DatalayerGetOwnedStoresResponse Response for get_owned_stores
type DatalayerGetOwnedStoresResponse struct {
	Response
	StoreIDs []string `json:"store_ids"`
}

// GetOwnedStores RPC endpoint get_owned_stores
func (s *DataLayerService) GetOwnedStores(opts *DatalayerGetOwnedStoresOptions) (*DatalayerGetOwnedStoresResponse, *http.Response, error) {
	request, err := s.NewRequest("get_owned_stores", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &DatalayerGetOwnedStoresResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
