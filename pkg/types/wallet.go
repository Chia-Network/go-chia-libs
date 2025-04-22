package types

// WalletType types of wallets
// This matches constants on the chia-blockchain end as well. Don't change to arbitrary values.
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/wallet/util/wallet_types.py#L12
type WalletType uint8

const (
	// WalletTypeStandard Standard Wallet
	WalletTypeStandard = WalletType(0)

	// WalletTypeRateLimited Deprecated: Rate Limited Wallet
	WalletTypeRateLimited = WalletType(1)

	// WalletTypeAtomicSwap Atomic Swap
	WalletTypeAtomicSwap = WalletType(2)

	// WalletTypeAuthorizedPayee Authorized Payee
	WalletTypeAuthorizedPayee = WalletType(3)

	// WalletTypeMultiSig Multi Sig
	WalletTypeMultiSig = WalletType(4)

	// WalletTypeCustody Custody
	WalletTypeCustody = WalletType(5)

	// WalletTypeCAT CAT Wallet
	WalletTypeCAT = WalletType(6)

	// WalletTypeRecoverable Recoverable Wallet
	WalletTypeRecoverable = WalletType(7)

	// WalletTypeDID DECENTRALIZED_ID Wallet
	WalletTypeDID = WalletType(8)

	// WalletTypePooling Pooling Wallet
	WalletTypePooling = WalletType(9)

	// WalletTypeNFT NFT Wallet
	WalletTypeNFT = WalletType(10)

	// WalletTypeDataLayer Data Layer Wallet
	WalletTypeDataLayer = WalletType(11)

	// WalletTypeDataLayerOffer Data Layer Offer wallet
	WalletTypeDataLayerOffer = WalletType(12)
)

// WalletInfo single wallet record
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/wallet/wallet_info.py#L12
// @TODO Streamable
type WalletInfo struct {
	ID   uint32     `json:"id"`
	Name string     `json:"name"`
	Type WalletType `json:"type"`
	Data string     `json:"data"`
}

// WalletBalance specific wallet balance information
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/rpc/wallet_rpc_api.py#L721
type WalletBalance struct {
	WalletID                 int32      `json:"wallet_id"`
	Fingerprint              int        `json:"fingerprint"`
	ConfirmedWalletBalance   Uint128    `json:"confirmed_wallet_balance"`
	UnconfirmedWalletBalance Uint128    `json:"unconfirmed_wallet_balance"`
	SpendableBalance         Uint128    `json:"spendable_balance"`
	PendingChange            uint64     `json:"pending_change"`
	MaxSendAmount            Uint128    `json:"max_send_amount"`
	UnspentCoinCount         uint32     `json:"unspent_coin_count"`
	PendingCoinRemovalCount  uint32     `json:"pending_coin_removal_count"`
	WalletType               WalletType `json:"wallet_type"`
	AssetID                  string     `json:"asset_id"`
}
