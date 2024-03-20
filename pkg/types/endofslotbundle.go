package types

import "github.com/samber/mo"

// EndOfSubSlotBundle end of subslot bundle
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/end_of_slot_bundle.py#L17
type EndOfSubSlotBundle struct {
	ChallengeChain        ChallengeChainSubSlot                   `json:"challenge_chain" streamable:""`
	InfusedChallengeChain mo.Option[InfusedChallengeChainSubSlot] `json:"infused_challenge_chain" streamable:""`
	RewardChain           RewardChainSubSlot                      `json:"reward_chain" streamable:""`
	Proofs                SubSlotProofs                           `json:"proofs" streamable:""`
}
