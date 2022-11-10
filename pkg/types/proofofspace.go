package types

import (
	"github.com/samber/mo"
)

// ProofOfSpace Proof of Space
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/proof_of_space.py#L20
// @TODO Streamable
type ProofOfSpace struct {
	Challenge              Bytes32              `json:"challenge"`
	PoolPublicKey          mo.Option[G1Element] `json:"pool_public_key"` // Only one of these two should be present
	PoolContractPuzzleHash mo.Option[Bytes32]   `json:"pool_contract_puzzle_hash"`
	PlotPublicKey          G1Element            `json:"plot_public_key"`
	Size                   uint8                `json:"size"`
	Proof                  Bytes                `json:"proof"`
}
