package types

// ProofOfSpace Proof of Space
type ProofOfSpace struct {
	Challenge              Bytes32    `json:"challenge"`
	PoolPublicKey          *G1Element `json:"pool_public_key"` // Only one of these two should be present
	PoolContractPuzzleHash *Bytes32   `json:"pool_contract_puzzle_hash"`
	PlotPublicKey          *G1Element `json:"plot_public_key"`
	Size                   uint8      `json:"size"`
	Proof                  Bytes      `json:"proof"`
}
