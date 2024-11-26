package rpc

import (
	"net/http"

	"github.com/samber/mo"

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

// GetClient returns the active client for the service
func (s *WalletService) GetClient() rpcinterface.Client {
	return s.client
}

// GetConnections returns connections
func (s *WalletService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
	return Do(s, "get_connections", opts, &GetConnectionsResponse{})
}

// GetNetworkInfo wallet rpc -> get_network_info
func (s *WalletService) GetNetworkInfo(opts *GetNetworkInfoOptions) (*GetNetworkInfoResponse, *http.Response, error) {
	return Do(s, "get_network_info", opts, &GetNetworkInfoResponse{})
}

// GetVersion returns the application version for the service
func (s *WalletService) GetVersion(opts *GetVersionOptions) (*GetVersionResponse, *http.Response, error) {
	return Do(s, "get_version", opts, &GetVersionResponse{})
}

// GetPublicKeysResponse response from get_public_keys
type GetPublicKeysResponse struct {
	rpcinterface.Response
	PublicKeyFingerprints mo.Option[[]int] `json:"public_key_fingerprints"`
}

// GetPublicKeys endpoint
func (s *WalletService) GetPublicKeys() (*GetPublicKeysResponse, *http.Response, error) {
	return Do(s, "get_public_keys", nil, &GetPublicKeysResponse{})
}

// GenerateMnemonicResponse Random new 24 words response
type GenerateMnemonicResponse struct {
	rpcinterface.Response
	Mnemonic mo.Option[[]string] `json:"mnemonic"`
}

// GenerateMnemonic Endpoint for generating a new random 24 words
func (s *WalletService) GenerateMnemonic() (*GenerateMnemonicResponse, *http.Response, error) {
	return Do(s, "generate_mnemonic", nil, &GenerateMnemonicResponse{})
}

// AddKeyOptions options for the add_key endpoint
type AddKeyOptions struct {
	Mnemonic []string `json:"mnemonic"`
}

// AddKeyResponse response from the add_key endpoint
type AddKeyResponse struct {
	rpcinterface.Response
	Word        mo.Option[string] `json:"word,omitempty"` // This is part of a unique error response
	Fingerprint mo.Option[int]    `json:"fingerprint,omitempty"`
}

// AddKey Adds a new key from 24 words to the keychain
func (s *WalletService) AddKey(opts *AddKeyOptions) (*AddKeyResponse, *http.Response, error) {
	return Do(s, "add_key", opts, &AddKeyResponse{})
}

// DeleteAllKeysResponse Delete keys response
type DeleteAllKeysResponse struct {
	rpcinterface.Response
}

// DeleteAllKeys deletes all keys from the keychain
func (s *WalletService) DeleteAllKeys() (*DeleteAllKeysResponse, *http.Response, error) {
	return Do(s, "delete_all_keys", nil, &DeleteAllKeysResponse{})
}

// GetNextAddressOptions options for get_next_address endpoint
type GetNextAddressOptions struct {
	NewAddress bool   `json:"new_address"`
	WalletID   uint32 `json:"wallet_id"`
}

// GetNextAddressResponse response from get next address
type GetNextAddressResponse struct {
	rpcinterface.Response
	WalletID mo.Option[uint32] `json:"wallet_id"`
	Address  mo.Option[string] `json:"address"`
}

// GetNextAddress returns the current address for the wallet. If NewAddress is true, it moves to the next address before responding
func (s *WalletService) GetNextAddress(opts *GetNextAddressOptions) (*GetNextAddressResponse, *http.Response, error) {
	return Do(s, "get_next_address", opts, &GetNextAddressResponse{})
}

// GetWalletSyncStatusResponse Response for get_sync_status on wallet
type GetWalletSyncStatusResponse struct {
	rpcinterface.Response
	GenesisInitialized mo.Option[bool] `json:"genesis_initialized"`
	Synced             mo.Option[bool] `json:"synced"`
	Syncing            mo.Option[bool] `json:"syncing"`
}

// GetSyncStatus wallet rpc -> get_sync_status
func (s *WalletService) GetSyncStatus() (*GetWalletSyncStatusResponse, *http.Response, error) {
	return Do(s, "get_sync_status", nil, &GetWalletSyncStatusResponse{})
}

// GetWalletHeightInfoResponse response for get_height_info on wallet
type GetWalletHeightInfoResponse struct {
	rpcinterface.Response
	Height mo.Option[uint32] `json:"height"`
}

// GetHeightInfo wallet rpc -> get_height_info
func (s *WalletService) GetHeightInfo() (*GetWalletHeightInfoResponse, *http.Response, error) {
	return Do(s, "get_height_info", nil, &GetWalletHeightInfoResponse{})
}

// GetWalletsOptions wallet rpc -> get_wallets
type GetWalletsOptions struct {
	Type types.WalletType `json:"type"`
}

// GetWalletsResponse wallet rpc -> get_wallets
type GetWalletsResponse struct {
	rpcinterface.Response
	Fingerprint mo.Option[int]                `json:"fingerprint"`
	Wallets     mo.Option[[]types.WalletInfo] `json:"wallets"`
}

// GetWallets wallet rpc -> get_wallets
func (s *WalletService) GetWallets(opts *GetWalletsOptions) (*GetWalletsResponse, *http.Response, error) {
	return Do(s, "get_wallets", opts, &GetWalletsResponse{})
}

// GetWalletBalanceOptions request options for get_wallet_balance
type GetWalletBalanceOptions struct {
	WalletID uint32 `json:"wallet_id"`
}

// GetWalletBalanceResponse is the wallet balance RPC response
type GetWalletBalanceResponse struct {
	rpcinterface.Response
	Balance mo.Option[types.WalletBalance] `json:"wallet_balance"`
}

// GetWalletBalance returns wallet balance
func (s *WalletService) GetWalletBalance(opts *GetWalletBalanceOptions) (*GetWalletBalanceResponse, *http.Response, error) {
	return Do(s, "get_wallet_balance", opts, &GetWalletBalanceResponse{})
}

// GetWalletTransactionCountOptions options for get transaction count
type GetWalletTransactionCountOptions struct {
	WalletID uint32 `json:"wallet_id"`
}

// GetWalletTransactionCountResponse response for get_transaction_count
type GetWalletTransactionCountResponse struct {
	rpcinterface.Response
	WalletID mo.Option[uint32] `json:"wallet_id"`
	Count    mo.Option[int]    `json:"count"`
}

// GetTransactionCount returns the total count of transactions for the specific wallet ID
func (s *WalletService) GetTransactionCount(opts *GetWalletTransactionCountOptions) (*GetWalletTransactionCountResponse, *http.Response, error) {
	return Do(s, "get_wallet_transaction_count", opts, &GetWalletTransactionCountResponse{})
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
	rpcinterface.Response
	WalletID     mo.Option[uint32]                    `json:"wallet_id"`
	Transactions mo.Option[[]types.TransactionRecord] `json:"transactions"`
}

// GetTransactions wallet rpc -> get_transactions
func (s *WalletService) GetTransactions(opts *GetWalletTransactionsOptions) (*GetWalletTransactionsResponse, *http.Response, error) {
	return Do(s, "get_transactions", opts, &GetWalletTransactionsResponse{})
}

// GetWalletTransactionOptions options for getting a single wallet transaction
type GetWalletTransactionOptions struct {
	WalletID      uint32 `json:"wallet_id"`
	TransactionID string `json:"transaction_id"`
}

// GetWalletTransactionResponse response for get_wallet_transactions
type GetWalletTransactionResponse struct {
	rpcinterface.Response
	Transaction   mo.Option[types.TransactionRecord] `json:"transaction"`
	TransactionID mo.Option[string]                  `json:"transaction_id"`
}

// GetTransaction returns a single transaction record
func (s *WalletService) GetTransaction(opts *GetWalletTransactionOptions) (*GetWalletTransactionResponse, *http.Response, error) {
	return Do(s, "get_transaction", opts, &GetWalletTransactionResponse{})
}

// SendTransactionOptions represents the options for send_transaction
type SendTransactionOptions struct {
	WalletID uint32        `json:"wallet_id"`
	Amount   uint64        `json:"amount"`
	Address  string        `json:"address"`
	Memos    []types.Bytes `json:"memos,omitempty"`
	Fee      uint64        `json:"fee"`
	Coins    []types.Coin  `json:"coins,omitempty"`
}

// SendTransactionResponse represents the response from send_transaction
type SendTransactionResponse struct {
	rpcinterface.Response
	TransactionID mo.Option[string]                  `json:"transaction_id"`
	Transaction   mo.Option[types.TransactionRecord] `json:"transaction"`
}

// SendTransaction sends a transaction
func (s *WalletService) SendTransaction(opts *SendTransactionOptions) (*SendTransactionResponse, *http.Response, error) {
	return Do(s, "send_transaction", opts, &SendTransactionResponse{})
}

// CatSpendOptions represents the options for cat_spend
type CatSpendOptions struct {
	WalletID uint32 `json:"wallet_id"`
	Amount   uint64 `json:"amount"`
	Address  string `json:"inner_address"`
	Fee      uint64 `json:"fee"`
}

// CatSpendResponse represents the response from cat_spend
type CatSpendResponse struct {
	rpcinterface.Response
	TransactionID mo.Option[string]                  `json:"transaction_id"`
	Transaction   mo.Option[types.TransactionRecord] `json:"transaction"`
}

// CatSpend sends a transaction
func (s *WalletService) CatSpend(opts *CatSpendOptions) (*CatSpendResponse, *http.Response, error) {
	return Do(s, "cat_spend", opts, &CatSpendResponse{})
}

// MintNFTOptions represents the options for nft_get_info
type MintNFTOptions struct {
	DidID             string   `json:"did_id"`             // not required
	EditionNumber     uint32   `json:"edition_number"`     // not required
	EditionCount      uint32   `json:"edition_count"`      // not required
	Fee               uint64   `json:"fee"`                // not required
	LicenseHash       string   `json:"license_hash"`       //not required
	LicenseURIs       []string `json:"license_uris"`       // not required
	MetaHash          string   `json:"meta_hash"`          // not required
	MetaURIs          []string `json:"meta_uris"`          // not required
	RoyaltyAddress    string   `json:"royalty_address"`    // not required
	RoyaltyPercentage uint32   `json:"royalty_percentage"` // not required
	TargetAddress     string   `json:"target_address"`     // not required
	Hash              string   `json:"hash"`
	URIs              []string `json:"uris"`
	WalletID          uint32   `json:"wallet_id"`
}

// MintNFTResponse represents the response from nft_get_info
type MintNFTResponse struct {
	rpcinterface.Response
	SpendBundle mo.Option[types.SpendBundle] `json:"spend_bundle"`
	WalletID    mo.Option[uint32]            `json:"wallet_id"`
}

// MintNFT Mint a new NFT
func (s *WalletService) MintNFT(opts *MintNFTOptions) (*MintNFTResponse, *http.Response, error) {
	return Do(s, "nft_mint_nft", opts, &MintNFTResponse{})
}

// GetNFTsOptions represents the options for nft_get_nfts
type GetNFTsOptions struct {
	WalletID   uint32         `json:"wallet_id"`
	StartIndex mo.Option[int] `json:"start_index"`
	Num        mo.Option[int] `json:"num"`
}

// GetNFTsResponse represents the response from nft_get_nfts
type GetNFTsResponse struct {
	rpcinterface.Response
	WalletID mo.Option[uint32]          `json:"wallet_id"`
	NFTList  mo.Option[[]types.NFTInfo] `json:"nft_list"`
}

// GetNFTs Show all NFTs in a given wallet
func (s *WalletService) GetNFTs(opts *GetNFTsOptions) (*GetNFTsResponse, *http.Response, error) {
	return Do(s, "nft_get_nfts", opts, &GetNFTsResponse{})
}

// TransferNFTOptions represents the options for nft_get_info
type TransferNFTOptions struct {
	Fee           uint64 `json:"fee"` // not required
	NFTCoinID     string `json:"nft_coin_id"`
	TargetAddress string `json:"target_address"`
	WalletID      uint32 `json:"wallet_id"`
}

// TransferNFTResponse represents the response from nft_get_info
type TransferNFTResponse struct {
	rpcinterface.Response
	SpendBundle mo.Option[types.SpendBundle] `json:"spend_bundle"`
	WalletID    mo.Option[uint32]            `json:"wallet_id"`
}

// TransferNFT Get info about an NFT
func (s *WalletService) TransferNFT(opts *TransferNFTOptions) (*TransferNFTResponse, *http.Response, error) {
	return Do(s, "nft_transfer_nft", opts, &TransferNFTResponse{})
}

// GetNFTInfoOptions represents the options for nft_get_info
type GetNFTInfoOptions struct {
	CoinID   string `json:"coin_id"`
	WalletID uint32 `json:"wallet_id"`
}

// GetNFTInfoResponse represents the response from nft_get_info
type GetNFTInfoResponse struct {
	rpcinterface.Response
	NFTInfo mo.Option[types.NFTInfo] `json:"nft_info"`
}

// GetNFTInfo Get info about an NFT
func (s *WalletService) GetNFTInfo(opts *GetNFTInfoOptions) (*GetNFTInfoResponse, *http.Response, error) {
	return Do(s, "nft_get_info", opts, &GetNFTInfoResponse{})
}

// NFTAddURIOptions represents the options for nft_add_uri
type NFTAddURIOptions struct {
	Fee       uint64 `json:"fee"` // not required
	Key       string `json:"key"`
	NFTCoinID string `json:"nft_coin_id"`
	URI       string `json:"uri"`
	WalletID  uint32 `json:"wallet_id"`
}

// NFTAddURIResponse represents the response from nft_add_uri
type NFTAddURIResponse struct {
	rpcinterface.Response
	SpendBundle mo.Option[types.SpendBundle] `json:"spend_bundle"`
	WalletID    mo.Option[uint32]            `json:"wallet_id"`
}

// NFTAddURI Get info about an NFT
func (s *WalletService) NFTAddURI(opts *NFTAddURIOptions) (*NFTAddURIResponse, *http.Response, error) {
	return Do(s, "nft_add_uri", opts, &NFTAddURIResponse{})
}

// NFTGetByDidOptions represents the options for nft_get_by_did
type NFTGetByDidOptions struct {
	DidID types.Bytes32 `json:"did_id,omitempty"`
}

// NFTGetByDidResponse represents the response from nft_get_by_did
type NFTGetByDidResponse struct {
	rpcinterface.Response
	WalletID mo.Option[uint32] `json:"wallet_id"`
}

// NFTGetByDid Get wallet ID by DID
func (s *WalletService) NFTGetByDid(opts *NFTGetByDidOptions) (*NFTGetByDidResponse, *http.Response, error) {
	return Do(s, "nft_get_by_did", opts, &NFTGetByDidResponse{})
}

// GetSpendableCoinsOptions Options for get_spendable_coins
type GetSpendableCoinsOptions struct {
	WalletID            uint32   `json:"wallet_id"`
	MinCoinAmount       *uint64  `json:"min_coin_amount,omitempty"`
	MaxCoinAmount       *uint64  `json:"max_coin_amount,omitempty"`
	ExcludedCoinAmounts []uint64 `json:"excluded_coin_amounts,omitempty"`
}

// GetSpendableCoinsResponse response from get_spendable_coins
type GetSpendableCoinsResponse struct {
	rpcinterface.Response
	ConfirmedRecords     mo.Option[[]types.CoinRecord] `json:"confirmed_records"`
	UnconfirmedRemovals  mo.Option[[]types.CoinRecord] `json:"unconfirmed_removals"`
	UnconfirmedAdditions mo.Option[[]types.CoinRecord] `json:"unconfirmed_additions"`
}

// GetSpendableCoins returns information about the coins in the wallet
func (s *WalletService) GetSpendableCoins(opts *GetSpendableCoinsOptions) (*GetSpendableCoinsResponse, *http.Response, error) {
	return Do(s, "get_spendable_coins", opts, &GetSpendableCoinsResponse{})
}

// CreateSignedTransactionOptions Options for create_signed_transaction endpoint
type CreateSignedTransactionOptions struct {
	WalletID           *uint32          `json:"wallet_id,omitempty"`
	Additions          []types.Addition `json:"additions"`
	Fee                *uint64          `json:"fee,omitempty"`
	MinCoinAmount      *uint64          `json:"min_coin_amount,omitempty"`
	MaxCoinAmount      *uint64          `json:"max_coin_amount,omitempty"`
	ExcludeCoinAmounts []*uint64        `json:"exclude_coin_amounts,omitempty"`
	Coins              []types.Coin     `json:"Coins,omitempty"`
	ExcludeCoins       []types.Coin     `json:"exclude_coins,omitempty"`
}

// CreateSignedTransactionResponse Response from create_signed_transaction
type CreateSignedTransactionResponse struct {
	rpcinterface.Response
	SignedTXs mo.Option[[]types.TransactionRecord] `json:"signed_txs"`
	SignedTX  mo.Option[types.TransactionRecord]   `json:"signed_tx"`
}

// CreateSignedTransaction generates a signed transaction based on the specified options
func (s *WalletService) CreateSignedTransaction(opts *CreateSignedTransactionOptions) (*CreateSignedTransactionResponse, *http.Response, error) {
	return Do(s, "create_signed_transaction", opts, &CreateSignedTransactionResponse{})
}

// SendTransactionMultiResponse Response from send_transaction_multi
type SendTransactionMultiResponse struct {
	rpcinterface.Response
	Transaction   mo.Option[types.TransactionRecord] `json:"transaction"`
	TransactionID mo.Option[string]                  `json:"transaction_id"`
}

// SendTransactionMulti allows sending a more detailed transaction with multiple inputs/outputs.
// Options are the same as create signed transaction since this is ultimately just a wrapper around that in Chia
func (s *WalletService) SendTransactionMulti(opts *CreateSignedTransactionOptions) (*SendTransactionMultiResponse, *http.Response, error) {
	return Do(s, "send_transaction_multi", opts, &SendTransactionMultiResponse{})
}
