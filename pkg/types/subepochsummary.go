package types

import (
	"github.com/samber/mo"
)

// SubEpochSummary sub epoch summary
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/sub_epoch_summary.py#L13
// @TODO Streamable
type SubEpochSummary struct {
	PrevSubEpochSummaryHash Bytes32           `json:"prev_subepoch_summary_hash"`
	RewardChainHash         Bytes32           `json:"reward_chain_hash"`
	NumBlocksOverflow       uint8             `json:"num_blocks_overflow"`
	NewDifficulty           mo.Option[uint64] `json:"new_difficulty"`
	NewSubSlotIters         mo.Option[uint64] `json:"new_sub_slot_iters"`
}
