package types

// PoolTarget PoolTarget
type PoolTarget struct {
	PuzzleHash *PuzzleHash `json:"puzzle_hash"`
	MaxHeight  uint32      `json:"max_height"`
}
