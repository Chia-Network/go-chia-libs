package rpc

import (
	"net/http"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
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
func (s *DataLayerService) Do(req *rpcinterface.Request, v iResponse) (*http.Response, error) {
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
	return r, resp, err
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
	return r, resp, err
}

// DatalayerGetOwnedStoresOptions Options for get_owned_stores
type DatalayerGetOwnedStoresOptions struct{}

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
	return r, resp, err
}

// DatalayerGetMirrorsOptions Options for get_mirrors
type DatalayerGetMirrorsOptions struct {
	ID string `json:"id"` // Hex String
}

// DatalayerGetMirrorsResponse Response from the get_mirrors RPC
type DatalayerGetMirrorsResponse struct {
	Response
	Mirrors []types.DatalayerMirror `json:"mirrors"`
}

// GetMirrors lists the mirrors for the given datalayer store
func (s *DataLayerService) GetMirrors(opts *DatalayerGetMirrorsOptions) (*DatalayerGetMirrorsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_mirrors", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &DatalayerGetMirrorsResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}

// DatalayerDeleteMirrorOptions options for delete_mirror RPC call
type DatalayerDeleteMirrorOptions struct {
	CoinID string `json:"coin_id"` // hex string
	Fee    uint64 `json:"fee"`     // not required
}

// DatalayerDeleteMirrorResponse response data for delete_mirror
type DatalayerDeleteMirrorResponse struct {
	Response
}

// DeleteMirror deletes a datalayer mirror
func (s *DataLayerService) DeleteMirror(opts *DatalayerDeleteMirrorOptions) (*DatalayerDeleteMirrorResponse, *http.Response, error) {
	request, err := s.NewRequest("delete_mirror", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &DatalayerDeleteMirrorResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}

// DatalayerAddMirrorOptions options for delete_mirror RPC call
type DatalayerAddMirrorOptions struct {
	ID     string   `json:"id"` // hex string datastore ID
	URLs   []string `json:"urls"`
	Amount uint64   `json:"amount"`
	Fee    uint64   `json:"fee"`
}

// DatalayerAddMirrorResponse response data for add_mirror
type DatalayerAddMirrorResponse struct {
	Response
}

// AddMirror deletes a datalayer mirror
func (s *DataLayerService) AddMirror(opts *DatalayerAddMirrorOptions) (*DatalayerAddMirrorResponse, *http.Response, error) {
	request, err := s.NewRequest("add_mirror", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &DatalayerAddMirrorResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}

// DatalayerSubscribeOptions options for subscribe
type DatalayerSubscribeOptions struct {
	ID   string   `json:"id"` // hex string datastore id
	URLs []string `json:"urls,omitempty"`
}

// DatalayerSubscribeResponse Response from subscribe. Always empty aside from standard fields
type DatalayerSubscribeResponse struct {
	Response
}

// Subscribe deletes a datalayer mirror
func (s *DataLayerService) Subscribe(opts *DatalayerSubscribeOptions) (*DatalayerSubscribeResponse, *http.Response, error) {
	request, err := s.NewRequest("subscribe", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &DatalayerSubscribeResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}

// DatalayerUnsubscribeOptions options for unsubscribing to a datastore
type DatalayerUnsubscribeOptions struct {
	ID         string `json:"id"` // hex string datastore id
	RetainData bool   `json:"retain"`
}

// DatalayerUnsubscribeResponse response data for unsubscribe
type DatalayerUnsubscribeResponse struct {
	Response
}

// Unsubscribe deletes a datalayer mirror
func (s *DataLayerService) Unsubscribe(opts *DatalayerUnsubscribeOptions) (*DatalayerUnsubscribeResponse, *http.Response, error) {
	request, err := s.NewRequest("unsubscribe", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &DatalayerUnsubscribeResponse{}
	resp, err := s.Do(request, r)
	return r, resp, err
}
