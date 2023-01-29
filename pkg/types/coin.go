package types

import (
	"crypto/sha256"
	"encoding/binary"
)

// Coin is a coin
// https://github.com/Chia-Network/chia_rs/blob/69908769e7df0ff2c10569aea9992cfecf3eb23a/wheel/src/coin.rs#L16
type Coin struct {
	ParentCoinInfo Bytes32 `json:"parent_coin_info"`
	PuzzleHash     Bytes32 `json:"puzzle_hash"`
	Amount         uint64  `json:"amount"`
}

// ID returns the coin ID of the coin
func (c *Coin) ID() Bytes32 {
	hasher := sha256.New()
	hasher.Write(Bytes32ToBytes(c.ParentCoinInfo))
	hasher.Write(Bytes32ToBytes(c.PuzzleHash))

	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, c.Amount)

	if c.Amount >= 0x8000000000000000 {
		hasher.Write([]byte{0})
		hasher.Write(amountBytes)
	} else {
		start := 0
		switch {
		case c.Amount >= 0x80000000000000:
			start = 0
		case c.Amount >= 0x800000000000:
			start = 1
		case c.Amount >= 0x8000000000:
			start = 2
		case c.Amount >= 0x80000000:
			start = 3
		case c.Amount >= 0x800000:
			start = 4
		case c.Amount >= 0x8000:
			start = 5
		case c.Amount >= 0x80:
			start = 6
		case c.Amount > 0:
			start = 7
		default:
			start = 8
		}
		hasher.Write(amountBytes[start:])
	}

	var hash Bytes32
	copy(hash[:], hasher.Sum(nil))
	return hash
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
