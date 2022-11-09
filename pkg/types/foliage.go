package types

// FoliageBlockData FoliageBlockData
type FoliageBlockData struct {
	UnfinishedRewardBlockHash Bytes32     `json:"unfinished_reward_block_hash"`
	PoolTarget                *PoolTarget `json:"pool_target"`
	PoolSignature             *G2Element  `json:"pool_signature"`
	FarmerRewardPuzzleHash    Bytes32     `json:"farmer_reward_puzzle_hash"`
	ExtensionData             Bytes32     `json:"extension_data"`
}

// Foliage Foliage
type Foliage struct {
	PrevBlockHash                    Bytes32           `json:"prev_block_hash"`
	RewardBlockHash                  Bytes32           `json:"reward_block_hash"`
	FoliageBlockData                 *FoliageBlockData `json:"foliage_block_data"`
	FoliageBlockDataSignature        *G2Element        `json:"foliage_block_data_signature"`
	FoliageTransactionBlockHash      *Bytes32          `json:"foliage_transaction_block_hash"`
	FoliageTransactionBlockSignature *G2Element        `json:"foliage_transaction_block_signature"`
}

// FoliageTransactionBlock foliage transaction block
type FoliageTransactionBlock struct {
	PrevTransactionBlockHash Bytes32 `json:"prev_transaction_block_hash"`
	Timestamp                uint64  `json:"timestamp"` // @TODO time.Time?
	FilterHash               Bytes32 `json:"filter_hash"`
	AdditionsRoot            Bytes32 `json:"additions_root"`
	RemovalsRoot             Bytes32 `json:"removals_root"`
	TransactionsInfoHash     Bytes32 `json:"transactions_info_hash"`
}

// TransactionsInfo transactions info
type TransactionsInfo struct {
	GeneratorRoot            Bytes32    `json:"generator_root"`
	GeneratorRefsRoot        Bytes32    `json:"generator_refs_root"`
	AggregatedSignature      *G2Element `json:"aggregated_signature"`
	Fees                     uint64     `json:"fees"`
	Cost                     uint64     `json:"cost"`
	RewardClaimsIncorporated []*Coin    `json:"reward_claims_incorporated"`
}
