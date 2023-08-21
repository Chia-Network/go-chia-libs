package protocols

import (
	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

// Plot is the plot definition in the harvester protocol
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L78
type Plot struct {
	Filename               string                     `json:"filename"`
	Size                   uint8                      `json:"size"`
	PlotID                 types.Bytes32              `json:"plot_id"`
	PoolPublicKey          mo.Option[types.G1Element] `json:"pool_public_key"`
	PoolContractPuzzleHash mo.Option[types.Bytes32]   `json:"pool_contract_puzzle_hash"`
	PlotPublicKey          types.G1Element            `json:"plot_public_key"`
	FileSize               uint64                     `json:"file_size"`
	TimeModified           types.Timestamp            `json:"time_modified"`
	CompressionLevel       mo.Option[uint8]           `json:"compression_level"`
}
