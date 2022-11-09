package types

// EventFarmerSubmittedPartial is the event data for `submitted_partial` from the farmer
type EventFarmerSubmittedPartial struct {
	LauncherID                   string `json:"launcher_id"`
	PoolURL                      string `json:"pool_url"`
	CurrentDifficulty            uint64 `json:"current_difficulty"`
	PointsAcknowledgedSinceStart uint64 `json:"points_acknowledged_since_start"`
}

// EventFarmerProof is the farmer event `proof`
type EventFarmerProof struct {
	Proof        DeclareProofOfSpace `json:"proof"`
	PassedFilter bool                `json:"passed_filter"`
}

// DeclareProofOfSpace matches to the farmer protocol type
type DeclareProofOfSpace struct {
	ChallengeHash             Bytes32      `json:"challenge_hash"`
	ChallengeChainSP          Bytes32      `json:"challenge_chain_sp"`
	SignagePointIndex         uint8        `json:"signage_point_index"`
	RewardChainSP             Bytes32      `json:"reward_chain_sp"`
	ProofOfSpace              ProofOfSpace `json:"proof_of_space"`
	ChallengeChainSPSignature G2Element    `json:"challenge_chain_sp_signature"`
	RewardChainSPSignature    G2Element    `json:"reward_chain_sp_signature"`
	FarmerPuzzleHash          Bytes32      `json:"farmer_puzzle_hash"`
	PoolTarget                *PoolTarget  `json:"pool_target,omitempty"`
	PoolSignature             *G2Element   `json:"pool_signature,omitempty"`
}
