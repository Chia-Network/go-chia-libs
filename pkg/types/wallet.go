package types

// WalletType types of wallets
// This matches constants on the chia-blockchain end as well. Don't change to arbitrary values.
type WalletType uint8

const (
	// WalletTypeStandard Standard Wallet
	WalletTypeStandard = WalletType(0)

	// WalletTypeRateLimited Rate Limited Wallet
	WalletTypeRateLimited = WalletType(1)

	// WalletTypeAtomicSwap Atomic Swap
	WalletTypeAtomicSwap = WalletType(2)

	// WalletTypeAuthorizedPayee Authorized Payee
	WalletTypeAuthorizedPayee = WalletType(3)

	// WalletTypeMultiSig Multi Sig
	WalletTypeMultiSig = WalletType(4)

	// WalletTypeCustody Custody
	WalletTypeCustody = WalletType(5)

	// WalletTypeColouredCoin Coloured Coin
	WalletTypeColouredCoin = WalletType(6)

	// WalletTypeRecoverable Recoverable Wallet
	WalletTypeRecoverable = WalletType(7)

	// WalletTypeDistributedID DID Wallet
	WalletTypeDistributedID = WalletType(8)
)

// WalletInfo single wallet record
type WalletInfo struct {
	Data string      `json:"data"`
	ID   uint32      `json:"id"`
	Name string      `json:"name"`
	Type *WalletType `json:"type"`
}

// WalletBalance specific wallet balance information
type WalletBalance struct {
	WalletID                 int32   `json:"wallet_id"`
	Fingerprint              int     `json:"fingerprint"`
	ConfirmedWalletBalance   Uint128 `json:"confirmed_wallet_balance"`
	UnconfirmedWalletBalance Uint128 `json:"unconfirmed_wallet_balance"`
	SpendableBalance         Uint128 `json:"spendable_balance"`
	PendingChange            int64   `json:"pending_change"`
	MaxSendAmount            int64   `json:"max_send_amount"`
	UnspentCoinCount         int64   `json:"unspent_coin_count"`
	PendingCoinRemovalCount  int64   `json:"pending_coin_removal_count"`
}
