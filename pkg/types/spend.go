package types

// Addition an addition for a spend
type Addition struct {
	Amount     uint64   `json:"amount"`
	PuzzleHash Bytes32  `json:"puzzle_hash"`
	Memos      []string `json:"memos,omitempty"`
}
