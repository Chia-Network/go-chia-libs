package types

// Coin is a coin
// https://github.com/Chia-Network/chia_rs/blob/69908769e7df0ff2c10569aea9992cfecf3eb23a/wheel/src/coin.rs#L16
type Coin struct {
	ParentCoinInfo Bytes32 `json:"parent_coin_info"`
	PuzzleHash     Bytes32 `json:"puzzle_hash"`
	Amount         uint64  `json:"amount"`
}

// CoinSpend spend to a coin
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/coin_spend.py#L11
// @TODO Streamable
type CoinSpend struct {
	Coin         Coin              `json:"coin"`
	PuzzleReveal SerializedProgram `json:"puzzle_reveal"`
	Solution     SerializedProgram `json:"solution"`
}

// CoinAddedEvent data from coin-added websocket event
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/wallet/wallet_node.py#L1250
type CoinAddedEvent struct {
	Success  bool   `json:"success"`
	State    string `json:"state"`
	WalletID uint32 `json:"wallet_id"`
}
