package types

import (
	"github.com/samber/mo"
)

// DatalayerMirror represents a single mirror for a data store
type DatalayerMirror struct {
	CoinID            Bytes32           `json:"coin_id"`
	LauncherID        Bytes32           `json:"launcher_id"`
	Amount            uint64            `json:"amount"`
	URLs              []string          `json:"urls"`
	Ours              bool              `json:"ours"`
	ConfirmedAtHeight mo.Option[uint32] `json:"confirmed_at_height"`
}

// DatalayerKeyValue a key and value from a store ID
type DatalayerKeyValue struct {
	Atom  mo.Option[string] `json:"atom"`
	Hash  Bytes32           `json:"hash"`
	Key   Bytes             `json:"key"`
	Value Bytes             `json:"value"`
}
