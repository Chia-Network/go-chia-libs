package types

// PoolTarget PoolTarget
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/pool_target.py#L12
type PoolTarget struct {
	PuzzleHash Bytes32 `json:"puzzle_hash" streamable:""`
	MaxHeight  uint32  `json:"max_height" streamable:""`
}
