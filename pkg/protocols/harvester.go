package protocols

import (
	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

// PoolDifficulty
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L23
type PoolDifficulty struct {
	Difficulty             uint64        `streamable:""`
	SubSlotIters           uint64        `streamable:""`
	PoolContractPuzzleHash types.Bytes32 `streamable:""`
}

// HarvesterHandshake
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L31
type HarvesterHandshake struct {
	FarmerPublicKeys []types.G1Element `streamable:""`
	PoolPublicKeys   []types.G1Element `streamable:""`
}

// NewSignagePointHarvester
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L38
type NewSignagePointHarvester struct {
	ChallengeHash     types.Bytes32    `streamable:""`
	Difficulty        uint64           `streamable:""`
	SubSlotIters      uint64           `streamable:""`
	SignagePointIndex uint8            `streamable:""`
	SPHash            types.Bytes32    `streamable:""`
	PoolDifficulties  []PoolDifficulty `streamable:""`
	FilterPrefixBits  uint8            `streamable:""`
}

// Plot is the plot definition in the harvester protocol
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L114
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
