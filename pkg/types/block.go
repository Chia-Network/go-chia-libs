package types

import (
	"github.com/samber/mo"
)

// BlockRecord a single block record
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/consensus/block_record.py#L18
// @TODO Streamable
type BlockRecord struct {
	HeaderHash                 Bytes32                      `json:"header_hash"`
	PrevHash                   Bytes32                      `json:"prev_hash"`
	Height                     uint32                       `json:"height"`
	Weight                     Uint128                      `json:"weight"`
	TotalIters                 Uint128                      `json:"total_iters"`
	SignagePointIndex          uint8                        `json:"signage_point_index"`
	ChallengeVDFOutput         ClassgroupElement            `json:"challenge_vdf_output"`
	InfusedChallengeVDFOutput  mo.Option[ClassgroupElement] `json:"infused_challenge_vdf_output"`
	RewardInfusionNewChallenge Bytes32                      `json:"reward_infusion_new_challenge"`
	ChallengeBlockInfoHash     Bytes32                      `json:"challenge_block_info_hash"`
	SubSlotIters               uint64                       `json:"sub_slot_iters"`
	PoolPuzzleHash             Bytes32                      `json:"pool_puzzle_hash"`
	FarmerPuzzleHash           Bytes32                      `json:"farmer_puzzle_hash"`
	RequiredIters              uint64                       `json:"required_iters"`
	Deficit                    uint8                        `json:"deficit"`
	Overflow                   bool                         `json:"overflow"`
	PrevTransactionBlockHeight uint32                       `json:"prev_transaction_block_height"`

	// Transaction Block - Present if is_transaction_block
	Timestamp                mo.Option[Timestamp] `json:"timestamp"`
	PrevTransactionBlockHash mo.Option[Bytes32]   `json:"prev_transaction_block_hash"`
	Fees                     mo.Option[uint64]    `json:"fees"`
	RewardClaimsIncorporated mo.Option[[]Coin]    `json:"reward_claims_incorporated"`

	// Slot - present if this is the first SB in sub slot
	FinishedChallengeSlotHashes        mo.Option[[]Bytes32] `json:"finished_challenge_slot_hashes"`
	FinishedInfusedChallengeSlotHashes mo.Option[[]Bytes32] `json:"finished_infused_challenge_slot_hashes"`
	FinishedRewardSlotHashes           mo.Option[[]Bytes32] `json:"finished_reward_slot_hashes"`

	// Sub-epoch - present if this is the first SB after sub-epoch
	SubEpochSummaryIncluded mo.Option[SubEpochSummary] `json:"sub_epoch_summary_included"`
}

// FullBlock a full block
// https://github.com/Chia-Network/chia-blockchain/blob/0befdec071f49708e26c7638656874492c52600a/chia/types/full_block.py#L16
// @TODO Streamable
type FullBlock struct {
	FinishedSubSlots             []EndOfSubSlotBundle               `json:"finished_sub_slots"`
	RewardChainBlock             RewardChainBlock                   `json:"reward_chain_block"`
	ChallengeChainSPProof        mo.Option[VDFProof]                `json:"challenge_chain_sp_proof"`
	ChallengeChainIPProof        VDFProof                           `json:"challenge_chain_ip_proof"`
	RewardChainSPProof           mo.Option[VDFProof]                `json:"reward_chain_sp_proof"`
	RewardChainIPProof           VDFProof                           `json:"reward_chain_ip_proof"`
	InfusedChallengeChainIPProof mo.Option[VDFProof]                `json:"infused_challenge_chain_ip_proof"`
	Foliage                      Foliage                            `json:"foliage"`
	FoliageTransactionBlock      mo.Option[FoliageTransactionBlock] `json:"foliage_transaction_block"`
	TransactionsInfo             mo.Option[TransactionsInfo]        `json:"transactions_info"`
	TransactionsGenerator        mo.Option[SerializedProgram]       `json:"transactions_generator"`
	TransactionsGeneratorRefList []uint32                           `json:"transactions_generator_ref_list"`
}

// RewardChainBlock Reward Chain Block
// https://github.com/Chia-Network/chia-blockchain/blob/0befdec071f49708e26c7638656874492c52600a/chia/types/blockchain_format/reward_chain_block.py#L30
// @TODO Streamable
type RewardChainBlock struct {
	Weight                     Uint128            `json:"weight"`
	Height                     uint32             `json:"height"`
	TotalIters                 Uint128            `json:"total_iters"`
	SignagePointIndex          uint8              `json:"signage_point_index"`
	POSSSCCChallengeHash       Bytes32            `json:"pos_ss_cc_challenge_hash"`
	ProofOfSpace               ProofOfSpace       `json:"proof_of_space"`
	ChallengeChainSPVDF        mo.Option[VDFInfo] `json:"challenge_chain_sp_vdf"`
	ChallengeChainSPSignature  G2Element          `json:"challenge_chain_sp_signature"`
	ChallengeChainIPVDF        VDFInfo            `json:"challenge_chain_ip_vdf"`
	RewardChainSPVDF           mo.Option[VDFInfo] `json:"reward_chain_sp_vdf"` // Not present for first sp in slot
	RewardChainSPSignature     G2Element          `json:"reward_chain_sp_signature"`
	RewardChainIPVDF           VDFInfo            `json:"reward_chain_ip_vdf"`
	InfusedChallengeChainIPVDF mo.Option[VDFInfo] `json:"infused_challenge_chain_ip_vdf"` // Iff deficit < 16
	IsTransactionBlock         bool               `json:"is_transaction_block"`
}

// BlockCountMetrics metrics from get_block_count_metrics endpoint
// https://github.com/Chia-Network/chia-blockchain/blob/0befdec071f49708e26c7638656874492c52600a/chia/rpc/full_node_rpc_api.py#L382
// Types are `int` in python, which is apparently unlimited in python3. Using uint64 as the largest native int in go
type BlockCountMetrics struct {
	CompactBlocks   uint64 `json:"compact_blocks"`
	UncompactBlocks uint64 `json:"uncompact_blocks"`
	HintCount       uint64 `json:"hint_count"`
}

// ReceiveBlockResult When Blockchain.receive_block(b) is called, one of these results is returned,
// showing whether the block was added to the chain (extending the peak),
// and if not, why it was not added.
// These values match values in chia blockchain. Must not be arbitrarily changed
// https://github.com/Chia-Network/chia-blockchain/blob/0befdec071f49708e26c7638656874492c52600a/chia/consensus/blockchain.py#L57
type ReceiveBlockResult uint8

const (
	// ReceiveBlockResultNewPeak Added to the peak of the blockchain
	ReceiveBlockResultNewPeak ReceiveBlockResult = 1

	// ReceiveBlockResultOrphan Added as an orphan/stale block (not a new peak of the chain)
	ReceiveBlockResultOrphan ReceiveBlockResult = 2

	// ReceiveBlockResultInvalidBlock Block was not added because it was invalid
	ReceiveBlockResultInvalidBlock ReceiveBlockResult = 3

	// ReceiveBlockResultAlreadyHaveBlock Block is already present in this blockchain
	ReceiveBlockResultAlreadyHaveBlock ReceiveBlockResult = 4

	// ReceiveBlockResultDisconnectedBlock Block's parent (previous pointer) is not in this blockchain
	ReceiveBlockResultDisconnectedBlock ReceiveBlockResult = 5
)

// BlockEvent data from block websocket event
// https://github.com/Chia-Network/chia-blockchain/blob/0befdec071f49708e26c7638656874492c52600a/chia/full_node/full_node.py#L1784
type BlockEvent struct {
	TransactionBlock              bool                          `json:"transaction_block"`
	KSize                         uint8                         `json:"k_size"`
	HeaderHash                    Bytes32                       `json:"header_hash"`
	ForkHeight                    uint32                        `json:"fork_height"`
	Height                        uint32                        `json:"height"`
	ValidationTime                float64                       `json:"validation_time"`
	PreValidationTime             float64                       `json:"pre_validation_time"`
	BlockCost                     mo.Option[uint64]             `json:"block_cost,omitempty"`
	BlockFees                     mo.Option[uint64]             `json:"block_fees,omitempty"`
	Timestamp                     mo.Option[Timestamp]          `json:"timestamp"`
	TransactionGeneratorSizeBytes mo.Option[uint64]             `json:"transaction_generator_size_bytes,omitempty"`
	TransactionGeneratorRefList   []uint32                      `json:"transaction_generator_ref_list"`
	ReceiveBlockResult            mo.Option[ReceiveBlockResult] `json:"receive_block_result,omitempty"`
}
