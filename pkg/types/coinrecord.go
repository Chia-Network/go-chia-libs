package types

// CoinRecord type
type CoinRecord struct {
	Coin                Coin   `json:"coin"`
	ConfirmedBlockIndex uint32 `json:"confirmed_block_index"`
	SpentBlockIndex     uint32 `json:"spent_block_index"`
	Coinbase            bool   `json:"coinbase"`
	Timestamp           uint64 `json:"timestamp"`
}
