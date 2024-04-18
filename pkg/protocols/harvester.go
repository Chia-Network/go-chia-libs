package protocols

import (
	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

// HarvesterMode is the harvester mode
type HarvesterMode uint8

// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/farmer/farmer_api.py#L97
const (
	HarvesterModeCPU HarvesterMode = 1
	HarvesterModeGPU HarvesterMode = 2
)

// PoolDifficulty is the pool difficulty in the harvester protocol
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L23
type PoolDifficulty struct {
	Difficulty             uint64        `streamable:""`
	SubSlotIters           uint64        `streamable:""`
	PoolContractPuzzleHash types.Bytes32 `streamable:""`
}

// HarvesterHandshake is the handshake message in the harvester protocol
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L31
type HarvesterHandshake struct {
	FarmerPublicKeys []types.G1Element `streamable:""`
	PoolPublicKeys   []types.G1Element `streamable:""`
}

// NewSignagePointHarvester is the new signage point message in the harvester protocol
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

// ProofOfSpaceFeeInfo is the fee info in the harvester protocol
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L50
type ProofOfSpaceFeeInfo struct {
	AppliedFeeThreshold uint32 `streamable:""`
}

// NewProofOfSpace is the new proof of space message in the harvester protocol
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L56
type NewProofOfSpace struct {
	ChallengeHash               types.Bytes32                  `streamable:""`
	SPHash                      types.Bytes32                  `streamable:""`
	PlotIdentifier              string                         `streamable:""`
	Proof                       types.ProofOfSpace             `streamable:""`
	SignagePointIndex           uint8                          `streamable:""`
	IncludeSourceSignatureData  bool                           `streamable:""`
	FarmerRewardAddressOverride mo.Option[types.Bytes32]       `streamable:""`
	FeeInfo                     mo.Option[ProofOfSpaceFeeInfo] `streamable:""`
}

// SigningDataKind is the kind of signing data
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L68
type SigningDataKind uint8

// SigningDataKind values
const (
	SigningDataKindFoliageBlockData SigningDataKind = iota + 1
	SigningDataKindFoliageTransactionBlock
	SigningDataKindChallengeChainVdf
	SigningDataKindRewardChainVdf
	SigningDataKindChallengeChainSubSlot
	SigningDataKindRewardChainSubSlot
	SigningDataKindPartial
)

// SignatureRequestSourceData is the source data for the signature request
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L80
type SignatureRequestSourceData struct {
	Kind SigningDataKind `streamable:""`
	Data []byte          `streamable:""`
}

// RequestSignatures is the request signatures message in the harvester protocol
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L89
type RequestSignatures struct {
	PlotIdentifier    string                                             `streamable:""`
	ChallengeHash     types.Bytes32                                      `streamable:""`
	SPHash            types.Bytes32                                      `streamable:""`
	Messages          []types.Bytes32                                    `streamable:""`
	MessageData       mo.Option[[]mo.Option[SignatureRequestSourceData]] `streamable:""`
	RCBlockUnfinished mo.Option[types.RewardChainBlockUnfinished]        `streamable:""`
}

// MessageSignature is the message signature in the harvester protocol
type MessageSignature struct {
	types.Bytes32   `streamable:""`
	types.G2Element `streamable:""`
}

// RespondSignatures is the respond signatures message in the harvester protocol
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L101
type RespondSignatures struct {
	PlotIdentifier              string                   `streamable:""`
	ChallengeHash               types.Bytes32            `streamable:""`
	SPHash                      types.Bytes32            `streamable:""`
	LocalPK                     types.G1Element          `streamable:""`
	FarmerPK                    types.G1Element          `streamable:""`
	MessageSignatures           []MessageSignature       `streamable:""`
	IncludeSourceSignatureData  bool                     `streamable:""`
	FarmerRewardAddressOverride mo.Option[types.Bytes32] `streamable:""`
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

// PlotSyncStart is the plot sync start message
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L150
type PlotSyncStart struct {
	Identifier     PlotSyncIdentifier `streamable:""`
	Initial        bool               `streamable:""`
	LastSyncID     uint64             `streamable:""`
	PlotFileCount  uint32             `streamable:""`
	HarvestingMode uint8              `streamable:""`
}

// PlotSyncIdentifier is the plot sync identifier
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/harvester_protocol.py#L142
type PlotSyncIdentifier struct {
	Timestamp uint64 `streamable:""`
	SyncID    uint64 `streamable:""`
	MessageID uint64 `streamable:""`
}
