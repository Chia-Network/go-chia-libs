package types

import "github.com/samber/mo"

// EndOfSubSlotBundle end of subslot bundle
// https://github.com/Chia-Network/chia_rs/blob/main/crates/chia-protocol/src/end_of_sub_slot_bundle.rs#L9
type EndOfSubSlotBundle struct {
	ChallengeChain        ChallengeChainSubSlot                   `json:"challenge_chain" streamable:""`
	InfusedChallengeChain mo.Option[InfusedChallengeChainSubSlot] `json:"infused_challenge_chain" streamable:""`
	RewardChain           RewardChainSubSlot                      `json:"reward_chain" streamable:""`
	Proofs                SubSlotProofs                           `json:"proofs" streamable:""`
}
