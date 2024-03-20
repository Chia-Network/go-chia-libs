package types

import "github.com/samber/mo"

// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/slots.py

// ChallengeBlockInfo is a type of challenge_block_info
type ChallengeBlockInfo struct {
	ProofOfSpace              ProofOfSpace       `json:"proof_of_space" streamable:""`
	ChallengeChainSPVDF       mo.Option[VDFInfo] `json:"challenge_chain_sp_vdf" streamable:""`
	ChallengeChainSPSignature G2Element          `json:"challenge_chain_sp_signature" streamable:""`
	ChallengeChainIPVDF       VDFInfo            `json:"challenge_chain_ip_vdf" streamable:""`
}

// ChallengeChainSubSlot is a type of challenge_chain_sub_slot
type ChallengeChainSubSlot struct {
	ChallengeChainEndOfSlotVDF       VDFInfo            `json:"challenge_chain_end_of_slot_vdf" streamable:""`
	InfusedChallengeChainSubSlotHash mo.Option[Bytes32] `json:"infused_challenge_chain_sub_slot_hash" streamable:""`
	SubepochSummaryHash              mo.Option[Bytes32] `json:"subepoch_summary_hash" streamable:""`
	NewSubSlotIters                  mo.Option[uint64]  `json:"new_sub_slot_iters" streamable:""`
	NewDifficulty                    mo.Option[uint64]  `json:"new_difficulty" streamable:""`
}

// InfusedChallengeChainSubSlot is a type of infused_challenge_chain_sub_slot
type InfusedChallengeChainSubSlot struct {
	InfusedChallengeChainEndOfSlotVDF VDFInfo `json:"infused_challenge_chain_end_of_slot_vdf" streamable:""`
}

// RewardChainSubSlot is a type of reward_chain_sub_slot
type RewardChainSubSlot struct {
	EndOfSlotVDF                     VDFInfo            `json:"end_of_slot_vdf" streamable:""`
	ChallengeChainSubSlotHash        Bytes32            `json:"challenge_chain_sub_slot_hash" streamable:""`
	InfusedChallengeChainSubSlotHash mo.Option[Bytes32] `json:"infused_challenge_chain_sub_slot_hash" streamable:""`
	Deficit                          uint8              `json:"deficit" streamable:""`
}

// SubSlotProofs is a type of sub_slot_proofs
type SubSlotProofs struct {
	ChallengeChainSlotProof        VDFProof            `json:"challenge_chain_slot_proof" streamable:""`
	InfusedChallengeChainSlotProof mo.Option[VDFProof] `json:"infused_challenge_chain_slot_proof" streamable:""`
	RewardChainSlotProof           VDFProof            `json:"reward_chain_slot_proof" streamable:""`
}
