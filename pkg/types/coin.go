package types

// Coin is a coin
type Coin struct {
	Amount         Uint128 `json:"amount"`
	ParentCoinInfo string  `json:"parent_coin_info"`
	PuzzleHash     string  `json:"puzzle_hash"`
}

// CoinSolution solution to a coin
type CoinSolution struct {
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
