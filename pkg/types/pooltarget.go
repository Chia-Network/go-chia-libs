package types

// PoolTarget PoolTarget
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/pool_target.py#L12
// @TODO Streamable
type PoolTarget struct {
	PuzzleHash Bytes32 `json:"puzzle_hash"`
	MaxHeight  uint32  `json:"max_height"`
}
