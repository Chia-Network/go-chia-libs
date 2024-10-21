package protocols

import (
	"github.com/chia-network/go-chia-libs/pkg/types"
	"github.com/samber/mo"
)

// SPSubSlotSourceData is the format for the sp_sub_slot_source_data response
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L26
type SPSubSlotSourceData struct {
	CCSubSlot types.ChallengeChainSubSlot `streamable:""`
	RCSubSlot types.RewardChainSubSlot    `streamable:""`
}

// SPVDFSourceData is the format for the sp_vdf_source_data response
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L33
type SPVDFSourceData struct {
	CCVDF types.ClassgroupElement `streamable:""`
	RCVDF types.ClassgroupElement `streamable:""`
}

// NewSignagePoint is the format for the new_signage_point response
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L47
type NewSignagePoint struct {
	ChallengeHash     types.Bytes32                     `streamable:""`
	ChallengeChainSP  types.Bytes32                     `streamable:""`
	RewardChainSP     types.Bytes32                     `streamable:""`
	Difficulty        uint64                            `streamable:""`
	SubSlotIters      uint64                            `streamable:""`
	SignagePointIndex uint8                             `streamable:""`
	PeakHeight        uint32                            `streamable:""`
	SPSourceData      mo.Option[SignagePointSourceData] `streamable:""`
}

// SignagePointSourceData is the format for the signage_point_source_data response
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L40
type SignagePointSourceData struct {
	SubSlotData mo.Option[SPSubSlotSourceData] `streamable:""`
	VDFData     mo.Option[SPVDFSourceData]     `streamable:""`
}

// DeclareProofOfSpace matches to the farmer protocol type
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L60
type DeclareProofOfSpace struct {
	ChallengeHash              types.Bytes32               `json:"challenge_hash" streamable:""`
	ChallengeChainSP           types.Bytes32               `json:"challenge_chain_sp" streamable:""`
	SignagePointIndex          uint8                       `json:"signage_point_index" streamable:""`
	RewardChainSP              types.Bytes32               `json:"reward_chain_sp" streamable:""`
	ProofOfSpace               types.ProofOfSpace          `json:"proof_of_space" streamable:""`
	ChallengeChainSPSignature  types.G2Element             `json:"challenge_chain_sp_signature" streamable:""`
	RewardChainSPSignature     types.G2Element             `json:"reward_chain_sp_signature" streamable:""`
	FarmerPuzzleHash           types.Bytes32               `json:"farmer_puzzle_hash" streamable:""`
	PoolTarget                 mo.Option[types.PoolTarget] `json:"pool_target,omitempty" streamable:""`
	PoolSignature              mo.Option[types.G2Element]  `json:"pool_signature,omitempty" streamable:""`
	IncludeSignatureSourceData bool                        `json:"include_signature_source_data" streamable:""`
}

// RequestSignedValues is the format for the request_signed_values response
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L76
type RequestSignedValues struct {
	QualityString               types.Bytes32                               `streamable:""`
	FoliageBlockDataHash        types.Bytes32                               `streamable:""`
	FoliageTransactionBlockHash types.Bytes32                               `streamable:""`
	FoliageBlockData            mo.Option[types.FoliageBlockData]           `streamable:""`
	FoliageTransactionBlockData mo.Option[types.FoliageTransactionBlock]    `streamable:""`
	RCBlockUnfinished           mo.Option[types.RewardChainBlockUnfinished] `streamable:""`
}

// FarmingInfo is the format for the farming_info response
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L87
type FarmingInfo struct {
	ChallengeHash types.Bytes32 `streamable:""`
	SPHash        types.Bytes32 `streamable:""`
	Timestamp     uint64        `streamable:""`
	Passed        uint32        `streamable:""`
	Proofs        uint32        `streamable:""`
	TotalPlots    uint32        `streamable:""`
	LookupTime    uint64        `streamable:""`
}

// SignedValues is the format for the signed_values response
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L99
type SignedValues struct {
	QualityString                    types.Bytes32   `streamable:""`
	FoliageBlockDataSignature        types.G2Element `streamable:""`
	FoliageTransactionBlockSignature types.G2Element `streamable:""`
}
