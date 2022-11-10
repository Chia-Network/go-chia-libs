package types

// SkippingPeakEvent data that is sent when a timelord skips a peak because it was fastest/already had the peak
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/timelord/timelord_api.py#L44
type SkippingPeakEvent struct {
	Success bool   `json:"success"`
	Height  uint32 `json:"height"`
}

// NewPeakEvent data that is sent when a timelord skips a peak because it was fastest/already had the peak
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/timelord/timelord_api.py#L49
type NewPeakEvent struct {
	Success bool   `json:"success"`
	Height  uint32 `json:"height"`
}

// TimelordChain references a particular chain within timelord code
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/timelord/types.py#L6
type TimelordChain uint8

const (
	// TimelordChainChallenge Challenge Chain
	TimelordChainChallenge TimelordChain = 1

	// TimelordChainReward Reward Chain
	TimelordChainReward TimelordChain = 2

	// TimelordChainInfusedChallenge Infused Challenge Chain
	TimelordChainInfusedChallenge TimelordChain = 3

	// TimelordChainBluebox Bluebox Chain
	TimelordChainBluebox TimelordChain = 4
)

// FinishedPoTEvent data every time a PoT Challenge is completed
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/timelord/timelord.py#L1050
type FinishedPoTEvent struct {
	Success          bool          `json:"success"`
	EstimatedIPS     float64       `json:"estimated_ips"`
	IterationsNeeded uint64        `json:"iterations_needed"`
	Chain            TimelordChain `json:"chain"`
	VDFInfo          VDFInfo       `json:"vdf_info"`
	VDFProof         VDFProof      `json:"vdf_proof"`
}

// CompressibleVDFField Stores, for a given VDF, the field that uses it.
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/vdf.py#L94
type CompressibleVDFField uint8

const (
	// CompressibleVDFFieldCCEOSVDF CC_EOS_VDF
	CompressibleVDFFieldCCEOSVDF CompressibleVDFField = 1

	// CompressibleVDFFieldICCEOSVDF ICC_EOS_VDF
	CompressibleVDFFieldICCEOSVDF CompressibleVDFField = 2

	// CompressibleVDFFieldCCSPVDF CC_SP_VDF
	CompressibleVDFFieldCCSPVDF CompressibleVDFField = 3

	// CompressibleVDFFieldCCIPVDF CC_IP_VDF
	CompressibleVDFFieldCCIPVDF CompressibleVDFField = 4
)

// NewCompactProofEvent is an event from the timelord every time a new compact proof is generated
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/timelord/timelord.py#L1074
type NewCompactProofEvent struct {
	Success    bool                 `json:"success"`
	HeaderHash Bytes32              `json:"header_hash"`
	Height     uint32               `json:"height"`
	FieldVdf   CompressibleVDFField `json:"field_vdf"`
}
