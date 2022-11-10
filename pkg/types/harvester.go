package types

import (
	"github.com/samber/mo"
)

// EventHarvesterFarmingInfo is the event data for `farming_info` from the harvester
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/harvester/harvester_api.py#L232
type EventHarvesterFarmingInfo struct {
	ChallengeHash Bytes32 `json:"challenge_hash"`
	TotalPlots    uint64  `json:"total_plots"`
	FoundProofs   uint64  `json:"found_proofs"`
	EligiblePlots uint64  `json:"eligible_plots"`
	Time          float64 `json:"time"`
}

// PlotInfo contains information about a plot, as used in get_plots rpc
// There is also a PlotInfo type in chia, that is NOT used in the RPC, that has the first 5 fields as defined here
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/harvester/harvester.py#L139
type PlotInfo struct {
	PoolPublicKey          mo.Option[G1Element] `json:"pool_public_key"`
	PoolContractPuzzleHash mo.Option[Bytes32]   `json:"pool_contract_puzzle_hash"`
	PlotPublicKey          G1Element            `json:"plot_public_key"`
	FileSize               uint64               `json:"file_size"`
	TimeModified           Timestamp            `json:"time_modified"`
	Filename               string               `json:"filename"`
	PlotID                 Bytes32              `json:"plot_id"`
	Size                   uint8                `json:"size"` // https://github.com/Chia-Network/chiapos/blob/main/src/prover_disk.hpp#L181
}
