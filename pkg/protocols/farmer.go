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
