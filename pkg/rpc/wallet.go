package rpc

import (
	"net/http"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

// WalletService encapsulates wallet RPC methods
type WalletService struct {
	client *Client
}

// NewRequest returns a new request specific to the wallet service
func (s *WalletService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceWallet, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *WalletService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetWalletSyncStatusResponse Response for get_sync_status on wallet
type GetWalletSyncStatusResponse struct {
	Success            bool `json:"success"`
	GenesisInitialized bool `json:"genesis_initialized"`
	Synced             bool `json:"synced"`
	Syncing            bool `json:"syncing"`
}

// GetSyncStatus wallet rpc -> get_sync_status
func (s *WalletService) GetSyncStatus() (*GetWalletSyncStatusResponse, *http.Response, error) {
	request, err := s.NewRequest("get_sync_status", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletSyncStatusResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletHeightInfoResponse response for get_height_info on wallet
type GetWalletHeightInfoResponse struct {
	Success bool   `json:"success"`
	Height  uint32 `json:"height"`
}

// GetHeightInfo wallet rpc -> get_height_info
func (s *WalletService) GetHeightInfo() (*GetWalletHeightInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("get_height_info", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletHeightInfoResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletNetworkInfoResponse response for get_height_info on wallet
type GetWalletNetworkInfoResponse struct {
	Success       bool   `json:"success"`
	NetworkName   string `json:"network_name"`
	NetworkPrefix string `json:"network_prefix"`
}

// GetNetworkInfo wallet rpc -> get_network_info
func (s *WalletService) GetNetworkInfo() (*GetWalletNetworkInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("get_network_info", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletNetworkInfoResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletsResponse wallet rpc -> get_wallets
type GetWalletsResponse struct {
	Success bool                `json:"success"`
	Wallets []*types.WalletInfo `json:"wallets"`
}

// GetWallets wallet rpc -> get_wallets
func (s *WalletService) GetWallets() (*GetWalletsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_wallets", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletsResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletBalanceOptions request options for get_wallet_balance
type GetWalletBalanceOptions struct {
	WalletID uint32 `json:"wallet_id"`
}

// GetWalletBalanceResponse is the wallet balance RPC response
type GetWalletBalanceResponse struct {
	Success bool                 `json:"success"`
	Balance *types.WalletBalance `json:"wallet_balance"`
}

// GetWalletBalance returns wallet balance
func (s *WalletService) GetWalletBalance(opts *GetWalletBalanceOptions) (*GetWalletBalanceResponse, *http.Response, error) {
	request, err := s.NewRequest("get_wallet_balance", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletBalanceResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletTransactionCountOptions options for get transaction count
type GetWalletTransactionCountOptions struct {
	WalletID uint32 `json:"wallet_id"`
}

// GetWalletTransactionCountResponse response for get_transaction_count
type GetWalletTransactionCountResponse struct {
	Success  bool   `json:"success"`
	WalletID uint32 `json:"wallet_id"`
	Count    int    `json:"count"`
}

// GetTransactionCount returns the total count of transactions for the specific wallet ID
func (s *WalletService) GetTransactionCount(opts *GetWalletTransactionCountOptions) (*GetWalletTransactionCountResponse, *http.Response, error) {
	request, err := s.NewRequest("get_transaction_count", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletTransactionCountResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletTransactionsOptions options for get wallet transactions
type GetWalletTransactionsOptions struct {
	WalletID  uint32 `json:"wallet_id"`
	Start     *int   `json:"start,omitempty"`
	End       *int   `json:"end,omitempty"`
	ToAddress string `json:"to_address,omitempty"`
}

// GetWalletTransactionsResponse response for get_wallet_transactions
type GetWalletTransactionsResponse struct {
	Success      bool                       `json:"success"`
	WalletID     uint32                     `json:"wallet_id"`
	Transactions []*types.TransactionRecord `json:"transactions"`
}

// GetTransactions wallet rpc -> get_transactions
func (s *WalletService) GetTransactions(opts *GetWalletTransactionsOptions) (*GetWalletTransactionsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_transactions", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletTransactionsResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletTransactionOptions options for getting a single wallet transaction
type GetWalletTransactionOptions struct {
	WalletID      uint32 `json:"wallet_id"`
	TransactionID string `json:"transaction_id"`
}

// GetWalletTransactionResponse response for get_wallet_transactions
type GetWalletTransactionResponse struct {
	Success       bool                     `json:"success"`
	Transaction   *types.TransactionRecord `json:"transaction"`
	TransactionID string                   `json:"transaction_id"`
}

// GetTransaction returns a single transaction record
func (s *WalletService) GetTransaction(opts *GetWalletTransactionOptions) (*GetWalletTransactionResponse, *http.Response, error) {
	request, err := s.NewRequest("get_transaction", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletTransactionResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// SendTransactionOptions represents the options for send_transaction
type SendTransactionOptions struct {
	WalletID uint32 `json:"wallet_id"`
	Amount   uint64 `json:"amount"`
	Address  string `json:"address"`
	Fee      uint64 `json:"fee"`
}

// SendTransactionResponse represents the response from send_transaction
type SendTransactionResponse struct {
	Success       bool                    `json:"success"`
	TransactionID string                  `json:"transaction_id"`
	Transaction   types.TransactionRecord `json:"transaction"`
}

// SendTransaction sends a transaction
func (s *WalletService) SendTransaction(opts *SendTransactionOptions) (*SendTransactionResponse, *http.Response, error) {
	request, err := s.NewRequest("send_transaction", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &SendTransactionResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
