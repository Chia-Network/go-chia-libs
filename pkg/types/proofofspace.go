package types

// ProofOfSpace Proof of Space
type ProofOfSpace struct {
	Challenge              string      `json:"challenge"`
	PoolPublicKey          *G1Element  `json:"pool_public_key"` // Only one of these two should be present
	PoolContractPuzzleHash *PuzzleHash `json:"pool_contract_puzzle_hash"`
	PlotPublicKey          *G1Element  `json:"plot_public_key"`
	Size                   uint8       `json:"size"`
	Proof                  string      `json:"proof"`
}
