package types

// SignagePointEvent is the data received for each signage point
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/full_node/full_node.py#L1413
type SignagePointEvent struct {
	Success         bool            `json:"success"`
	BroadcastFarmer NewSignagePoint `json:"broadcast_farmer"`
}

// NewSignagePoint is the event broadcast to farmers for a new signage point
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py#L22
// @TODO Streamable
type NewSignagePoint struct {
	ChallengeHash      Bytes32 `json:"challenge_hash"`
	ChallengeChainHash Bytes32 `json:"challenge_chain_hash"`
	RewardChainSP      Bytes32 `json:"reward_chain_sp"`
	Difficulty         uint64  `json:"difficulty"`
	SubSlotIters       uint64  `json:"sub_slot_iters"`
	SignagePointIndex  uint8   `json:"signage_point_index"`
}
