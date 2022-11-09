package types

// EventHarvesterFarmingInfo is the event data for `farming_info` from the harvester
type EventHarvesterFarmingInfo struct {
	ChallengeHash Bytes32 `json:"challenge_hash"`
	TotalPlots    uint64  `json:"total_plots"`
	FoundProofs   uint64  `json:"found_proofs"`
	EligiblePlots uint64  `json:"eligible_plots"`
	Time          float64 `json:"time"`
}

// PlotInfo contains information about a plot, as used in get_plots rpc
type PlotInfo struct {
	FileSize               uint64     `json:"file_size"`
	Filename               string     `json:"filename"`
	PlotID                 Bytes32    `json:"plot_id"`
	PlotPublicKey          G1Element  `json:"plot_public_key"`
	PoolContractPuzzleHash *Bytes32   `json:"pool_contract_puzzle_hash"`
	PoolPublicKey          *G1Element `json:"pool_public_key"`
	Size                   uint8      `json:"size"`
	TimeModified           int        `json:"time_modified"`
}
