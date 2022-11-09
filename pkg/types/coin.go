package types

// Coin is a coin
type Coin struct {
	ParentCoinInfo Bytes32 `json:"parent_coin_info"`
	PuzzleHash     Bytes32 `json:"puzzle_hash"`
	Amount         Uint128 `json:"amount"`
}

// CoinSpend spend to a coin
type CoinSpend struct {
	Coin         *Coin              `json:"coin"`
	PuzzleReveal *SerializedProgram `json:"puzzle_reveal"`
	Solution     *SerializedProgram `json:"solution"`
}

// CoinAddedEvent data from coin-added websocket event
type CoinAddedEvent struct {
	Success  bool   `json:"success"`
	State    string `json:"state"`
	WalletID uint32 `json:"wallet_id"`
}
