package types

// Addition an addition for a spend
type Addition struct {
	Amount     uint64 `json:"amount"`
	PuzzleHash string `json:"puzzle_hash"`
	Memos      []byte `json:"memos,omitempty"`
}
