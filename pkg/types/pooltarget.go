package types

// PoolTarget PoolTarget
// https://github.com/Chia-Network/chia_rs/blob/main/crates/chia-protocol/src/pool_target.rs#L6
type PoolTarget struct {
	PuzzleHash Bytes32 `json:"puzzle_hash" streamable:""`
	MaxHeight  uint32  `json:"max_height" streamable:""`
}
