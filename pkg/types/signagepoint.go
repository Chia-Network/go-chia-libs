package types

// SignagePointEvent is the data received for each signage point
type SignagePointEvent struct {
	Success         bool            `json:"success"`
	BroadcastFarmer NewSignagePoint `json:"broadcast_farmer"`
}

// NewSignagePoint is the event broadcast to farmers for a new signage point
type NewSignagePoint struct {
	ChallengeHash      string `json:"challenge_hash"`
	ChallengeChainHash string `json:"challenge_chain_hash"`
	RewardChainSP      string `json:"reward_chain_sp"`
	Difficulty         uint64 `json:"difficulty"`
	SubSlotIters       uint64 `json:"sub_slot_iters"`
	SignagePointIndex  uint8  `json:"signage_point_index"`
}
