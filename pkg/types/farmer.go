package types

import (
	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/tuple"
)

// FarmerMissingSignagePoints is the struct representation of the missing signage points tuple
type FarmerMissingSignagePoints struct {
	Timestamp Timestamp // uint64 in the Python Tuple
	Count     uint32
}

// EventFarmerNewSignagePoint is the event data for `new_signage_point` in the farmer service
type EventFarmerNewSignagePoint struct {
	SPHash               Bytes32 `json:"sp_hash"`
	MissingSignagePoints mo.Option[tuple.Tuple[FarmerMissingSignagePoints]]
}

// EventFarmerNewFarmingInfo is the event data for `new_farming_info` from the farmer
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/farmer/farmer_api.py#L535
type EventFarmerNewFarmingInfo struct {
	FarmingInfo struct {
		ChallengeHash Bytes32   `json:"challenge_hash"`
		SignagePoint  Bytes32   `json:"signage_point"`
		PassedFilter  uint32    `json:"passed_filter"`
		Proofs        uint32    `json:"proofs"`
		TotalPlots    uint32    `json:"total_plots"`
		Timestamp     Timestamp `json:"timestamp"`
		NodeID        Bytes32   `json:"node_id"`
		LookupTime    uint64    `json:"lookup_time"`
	} `json:"farming_info"`
}

// EventFarmerSubmittedPartial is the event data for `submitted_partial` from the farmer
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/farmer/farmer_api.py#L270
type EventFarmerSubmittedPartial struct {
	LauncherID                   Bytes32 `json:"launcher_id"`
	PoolURL                      string  `json:"pool_url"`
	CurrentDifficulty            uint64  `json:"current_difficulty"` // https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/pool_protocol.py#L97
	PointsAcknowledgedSinceStart uint64  `json:"points_acknowledged_since_start"`
}

// EventFarmerProof is the farmer event `proof`
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/farmer/farmer_api.py#L367
type EventFarmerProof struct {
	Proof        DeclareProofOfSpace `json:"proof"`
	PassedFilter bool                `json:"passed_filter"`
}

// DeclareProofOfSpace matches to the farmer protocol type
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L33
// @TODO Streamable
type DeclareProofOfSpace struct {
	ChallengeHash             Bytes32               `json:"challenge_hash"`
	ChallengeChainSP          Bytes32               `json:"challenge_chain_sp"`
	SignagePointIndex         uint8                 `json:"signage_point_index"`
	RewardChainSP             Bytes32               `json:"reward_chain_sp"`
	ProofOfSpace              ProofOfSpace          `json:"proof_of_space"`
	ChallengeChainSPSignature G2Element             `json:"challenge_chain_sp_signature"`
	RewardChainSPSignature    G2Element             `json:"reward_chain_sp_signature"`
	FarmerPuzzleHash          Bytes32               `json:"farmer_puzzle_hash"`
	PoolTarget                mo.Option[PoolTarget] `json:"pool_target,omitempty"`
	PoolSignature             mo.Option[G2Element]  `json:"pool_signature,omitempty"`
}
