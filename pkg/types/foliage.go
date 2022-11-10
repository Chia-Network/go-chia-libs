package types

import (
	"github.com/samber/mo"
)

// FoliageBlockData FoliageBlockData
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/foliage.py#L41
// @TODO Streamable
type FoliageBlockData struct {
	UnfinishedRewardBlockHash Bytes32              `json:"unfinished_reward_block_hash"`
	PoolTarget                PoolTarget           `json:"pool_target"`
	PoolSignature             mo.Option[G2Element] `json:"pool_signature"`
	FarmerRewardPuzzleHash    Bytes32              `json:"farmer_reward_puzzle_hash"`
	ExtensionData             Bytes32              `json:"extension_data"`
}

// Foliage Foliage
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/foliage.py#L52
// @TODO Streamable
type Foliage struct {
	PrevBlockHash                    Bytes32              `json:"prev_block_hash"`
	RewardBlockHash                  Bytes32              `json:"reward_block_hash"`
	FoliageBlockData                 FoliageBlockData     `json:"foliage_block_data"`
	FoliageBlockDataSignature        G2Element            `json:"foliage_block_data_signature"`
	FoliageTransactionBlockHash      mo.Option[Bytes32]   `json:"foliage_transaction_block_hash"`
	FoliageTransactionBlockSignature mo.Option[G2Element] `json:"foliage_transaction_block_signature"`
}

// FoliageTransactionBlock foliage transaction block
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/foliage.py#L29
// @TODO Streamable
type FoliageTransactionBlock struct {
	PrevTransactionBlockHash Bytes32   `json:"prev_transaction_block_hash"`
	Timestamp                Timestamp `json:"timestamp"`
	FilterHash               Bytes32   `json:"filter_hash"`
	AdditionsRoot            Bytes32   `json:"additions_root"`
	RemovalsRoot             Bytes32   `json:"removals_root"`
	TransactionsInfoHash     Bytes32   `json:"transactions_info_hash"`
}

// TransactionsInfo transactions info
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/foliage.py#L17
// @TODO Streamable
type TransactionsInfo struct {
	GeneratorRoot            Bytes32   `json:"generator_root"`
	GeneratorRefsRoot        Bytes32   `json:"generator_refs_root"`
	AggregatedSignature      G2Element `json:"aggregated_signature"`
	Fees                     uint64    `json:"fees"`
	Cost                     uint64    `json:"cost"`
	RewardClaimsIncorporated []Coin    `json:"reward_claims_incorporated"`
}
