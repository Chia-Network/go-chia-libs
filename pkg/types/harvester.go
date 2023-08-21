package types

// EventHarvesterFarmingInfo is the event data for `farming_info` from the harvester
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/harvester/harvester_api.py#L232
type EventHarvesterFarmingInfo struct {
	ChallengeHash Bytes32 `json:"challenge_hash"`
	TotalPlots    uint64  `json:"total_plots"`
	FoundProofs   uint64  `json:"found_proofs"`
	EligiblePlots uint64  `json:"eligible_plots"`
	Time          float64 `json:"time"`
}

// HarvestingMode is the mode the harvester is using to harvest CPU or GPU
type HarvestingMode int

const (
	// HarvestingModeCPU Using CPU to harvest
	HarvestingModeCPU = HarvestingMode(1)

	// HarvestingModeGPU Using CPU to harvest
	HarvestingModeGPU = HarvestingMode(2)
)
