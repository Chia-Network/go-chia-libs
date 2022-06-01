package types

// EventHarvesterFarmingInfo is the event data for `farming_info` from the harvester
type EventHarvesterFarmingInfo struct {
	ChallengeHash string  `json:"challenge_hash"`
	TotalPlots    uint64  `json:"total_plots"`
	FoundProofs   uint64  `json:"found_proofs"`
	EligiblePlots uint64  `json:"eligible_plots"`
	Time          float64 `json:"time"`
}
