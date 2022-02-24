package types

// SkippingPeakEvent data that is sent when a timelord skips a peak because it was fastest/already had the peak
type SkippingPeakEvent struct {
	Success bool   `json:"success"`
	Height  uint32 `json:"height"`
}

// NewPeakEvent data that is sent when a timelord skips a peak because it was fastest/already had the peak
type NewPeakEvent struct {
	Success bool   `json:"success"`
	Height  uint32 `json:"height"`
}

// TimelordChain references a particular chain within timelord code
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
type FinishedPoTEvent struct {
	Success          bool          `json:"success"`
	Chain            TimelordChain `json:"chain"`
	EstimatedIPS     float64       `json:"estimated_ips"`
	IterationsNeeded uint64        `json:"iterations_needed"`
	VDFInfo          VDFInfo       `json:"vdf_info"`
	VDFProof         VDFProof      `json:"vdf_proof"`
}

// CompressibleVDFField Stores, for a given VDF, the field that uses it.
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
type NewCompactProofEvent struct {
	Success    bool                 `json:"success"`
	HeaderHash string               `json:"header_hash"`
	Height     uint32               `json:"height"`
	FieldVdf   CompressibleVDFField `json:"field_vdf"`
}
