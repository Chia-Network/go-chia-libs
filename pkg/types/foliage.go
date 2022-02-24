package types

// FoliageBlockData FoliageBlockData
type FoliageBlockData struct {
	UnfinishedRewardBlockHash string      `json:"unfinished_reward_block_hash"`
	PoolTarget                *PoolTarget `json:"pool_target"`
	PoolSignature             *G2Element  `json:"pool_signature"`
	FarmerRewardPuzzleHash    string      `json:"farmer_reward_puzzle_hash"`
	ExtensionData             string      `json:"extension_data"`
}

// Foliage Foliage
type Foliage struct {
	PrevBlockHash                    string            `json:"prev_block_hash"`
	RewardBlockHash                  string            `json:"reward_block_hash"`
	FoliageBlockData                 *FoliageBlockData `json:"foliage_block_data"`
	FoliageBlockDataSignature        *G2Element        `json:"foliage_block_data_signature"`
	FoliageTransactionBlockHash      string            `json:"foliage_transaction_block_hash"`
	FoliageTransactionBlockSignature *G2Element        `json:"foliage_transaction_block_signature"`
}

// FoliageTransactionBlock foliage transaction block
type FoliageTransactionBlock struct {
	PrevTransactionBlockHash string `json:"prev_transaction_block_hash"`
	Timestamp                uint64 `json:"timestamp"` // @TODO time.Time?
	FilterHash               string `json:"filter_hash"`
	AdditionsRoot            string `json:"additions_root"`
	RemovalsRoot             string `json:"removals_root"`
	TransactionsInfoHash     string `json:"transactions_info_hash"`
}

// TransactionsInfo transactions info
type TransactionsInfo struct {
	GeneratorRoot            string     `json:"generator_root"`
	GeneratorRefsRoot        string     `json:"generator_refs_root"`
	AggregatedSignature      *G2Element `json:"aggregated_signature"`
	Fees                     uint64     `json:"fees"`
	Cost                     uint64     `json:"cost"`
	RewardClaimsIncorporated []*Coin    `json:"reward_claims_incorporated"`
}
