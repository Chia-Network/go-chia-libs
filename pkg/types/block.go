package types

// BlockRecord a single block record
type BlockRecord struct {
	HeaderHash                 string             `json:"header_hash"`
	PrevHash                   string             `json:"prev_hash"`
	Height                     uint32             `json:"height"`
	Weight                     Uint128            `json:"weight"`
	TotalIters                 Uint128            `json:"total_iters"`
	SignagePointIndex          uint8              `json:"signage_point_index"`
	ChallengeVDFOutput         *ClassgroupElement `json:"challenge_vdf_output"`
	InfusedChallengeVDFOutput  *ClassgroupElement `json:"infused_challenge_vdf_output"`
	RewardInfusionNewChallenge string             `json:"reward_infusion_new_challenge"`
	ChallengeBlockInfoHash     string             `json:"challenge_block_info_hash"`
	SubSlotIters               uint64             `json:"sub_slot_iters"`
	PoolPuzzleHash             *PuzzleHash        `json:"pool_puzzle_hash"`
	FarmerPuzzleHash           *PuzzleHash        `json:"farmer_puzzle_hash"`
	RequiredIters              uint64             `json:"required_iters"`
	Deficit                    uint8              `json:"deficit"`
	Overflow                   bool               `json:"overflow"`
	PrevTransactionBlockHeight uint32             `json:"prev_transaction_block_height"`

	// Transaction Block - Present if is_transaction_block
	Timestamp                uint64  `json:"timestamp"` // @TODO time.Time ?
	PrevTransactionBlockHash string  `json:"prev_transaction_block_hash"`
	Fees                     uint64  `json:"fees"` // @TODO proper unit (mojo/xch)?
	RewardClaimsIncorporated []*Coin `json:"reward_claims_incorporated"`

	// Slot - present if this is the first SB in sub slot
	FinishedChallengeSlotHashes        []string `json:"finished_challenge_slot_hashes"`
	FinishedInfusedChallengeSlotHashes []string `json:"finished_infused_challenge_slot_hashes"`
	FinishedRewardSlotHashes           []string `json:"finished_reward_slot_hashes"`

	// Sub-epoch - present if this is the first SB after sub-epoch
	SubEpochSummaryIncluded *SubEpochSummary `json:"sub_epoch_summary_included"`
}

// FullBlock a full block
type FullBlock struct {
	FinishedSubSlots             []*EndOfSubSlotBundle    `json:"finished_sub_slots"`
	RewardChainBlock             *RewardChainBlock        `json:"reward_chain_block"`
	ChallengeChainSPProof        *VDFProof                `json:"challenge_chain_sp_proof"`
	ChallengeChainIPProof        *VDFProof                `json:"challenge_chain_ip_proof"`
	RewardChainSPProof           *VDFProof                `json:"reward_chain_sp_proof"`
	RewardChainIPProof           *VDFProof                `json:"reward_chain_ip_proof"`
	InfusedChallengeChainIPProof *VDFProof                `json:"infused_challenge_chain_ip_proof"`
	Foliage                      *Foliage                 `json:"foliage"`
	FoliageTransactionBlock      *FoliageTransactionBlock `json:"foliage_transaction_block"`
	TransactionsInfo             *TransactionsInfo        `json:"transactions_info"`
	TransactionsGenerator        *SerializedProgram       `json:"transactions_generator"`          // @TODO Verify this is correct
	TransactionsGeneratorRefList []uint32                 `json:"transactions_generator_ref_list"` // @TODO Verify this is correct
}

// RewardChainBlock Reward Chain Block
type RewardChainBlock struct {
	Weight                     Uint128       `json:"weight"`
	Height                     uint32        `json:"height"`
	TotalIters                 Uint128       `json:"total_iters"`
	SignagePointIndex          uint8         `json:"signage_point_index"`
	POSSSCCChallengeHash       string        `json:"pos_ss_cc_challenge_hash"`
	ProofOfSpace               *ProofOfSpace `json:"proof_of_space"`
	ChallengeChainSPVDF        *VDFInfo      `json:"challenge_chain_sp_vdf"`
	ChallengeChainSPSignature  *G2Element    `json:"challenge_chain_sp_signature"`
	ChallengeChainIPVDF        *VDFInfo      `json:"challenge_chain_ip_vdf"`
	RewardChainSPVDF           *VDFInfo      `json:"reward_chain_sp_vdf"` // Not present for first sp in slot
	RewardChainSPSignature     *G2Element    `json:"reward_chain_sp_signature"`
	RewardChainIPVDF           *VDFInfo      `json:"reward_chain_ip_vdf"`
	InfusedChallengeChainIPVDF *VDFInfo      `json:"infused_challenge_chain_ip_vdf"` // Iff deficit < 16
	IsTransactionBlock         bool          `json:"is_transaction_block"`
}

// BlockCountMetrics metrics from get_block_count_metrics endpoint
type BlockCountMetrics struct {
	CompactBlocks   uint32 `json:"compact_blocks"`
	UncompactBlocks uint32 `json:"uncompact_blocks"`
	HintCount       uint64 `json:"hint_count"`
}

// ReceiveBlockResult When Blockchain.receive_block(b) is called, one of these results is returned,
// showing whether the block was added to the chain (extending the peak),
// and if not, why it was not added.
// These values match values in chia blockchain. Must not be arbitrarily changed
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
type BlockEvent struct {
	TransactionBlock              bool               `json:"transaction_block"`
	KSize                         uint8              `json:"k_size"`
	HeaderHash                    string             `json:"header_hash"`
	Height                        uint32             `json:"height"`
	BlockCost                     uint64             `json:"block_cost"`
	BlockFees                     uint64             `json:"block_fees"`
	TransactionGeneratorSizeBytes uint64             `json:"transaction_generator_size_bytes"`
	TransactionGeneratorRefList   []uint32           `json:"transaction_generator_ref_list"`
	ReceiveBlockResult            ReceiveBlockResult `json:"receive_block_result"`
	//Timestamp                     type                `json:"timestamp"`
}
