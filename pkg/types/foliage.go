package types

import (
	"github.com/samber/mo"
)

// FoliageBlockData FoliageBlockData
// https://github.com/Chia-Network/chia_rs/blob/main/crates/chia-protocol/src/foliage.rs#L31
type FoliageBlockData struct {
	UnfinishedRewardBlockHash Bytes32              `json:"unfinished_reward_block_hash" streamable:""`
	PoolTarget                PoolTarget           `json:"pool_target" streamable:""`
	PoolSignature             mo.Option[G2Element] `json:"pool_signature" streamable:""`
	FarmerRewardPuzzleHash    Bytes32              `json:"farmer_reward_puzzle_hash" streamable:""`
	ExtensionData             Bytes32              `json:"extension_data" streamable:""`
}

// Foliage Foliage
// https://github.com/Chia-Network/chia_rs/blob/main/crates/chia-protocol/src/foliage.rs#L41
type Foliage struct {
	PrevBlockHash                    Bytes32              `json:"prev_block_hash" streamable:""`
	RewardBlockHash                  Bytes32              `json:"reward_block_hash" streamable:""`
	FoliageBlockData                 FoliageBlockData     `json:"foliage_block_data" streamable:""`
	FoliageBlockDataSignature        G2Element            `json:"foliage_block_data_signature" streamable:""`
	FoliageTransactionBlockHash      mo.Option[Bytes32]   `json:"foliage_transaction_block_hash" streamable:""`
	FoliageTransactionBlockSignature mo.Option[G2Element] `json:"foliage_transaction_block_signature" streamable:""`
}

// FoliageTransactionBlock foliage transaction block
// https://github.com/Chia-Network/chia_rs/blob/main/crates/chia-protocol/src/foliage.rs#L20
type FoliageTransactionBlock struct {
	PrevTransactionBlockHash Bytes32   `json:"prev_transaction_block_hash" streamable:""`
	Timestamp                Timestamp `json:"timestamp" streamable:"Timestamp"`
	FilterHash               Bytes32   `json:"filter_hash" streamable:""`
	AdditionsRoot            Bytes32   `json:"additions_root" streamable:""`
	RemovalsRoot             Bytes32   `json:"removals_root" streamable:""`
	TransactionsInfoHash     Bytes32   `json:"transactions_info_hash" streamable:""`
}

// TransactionsInfo transactions info
// https://github.com/Chia-Network/chia_rs/blob/main/crates/chia-protocol/src/foliage.rs#L9
// @TODO Streamable
type TransactionsInfo struct {
	GeneratorRoot            Bytes32   `json:"generator_root" streamable:""`
	GeneratorRefsRoot        Bytes32   `json:"generator_refs_root" streamable:""`
	AggregatedSignature      G2Element `json:"aggregated_signature" streamable:""`
	Fees                     uint64    `json:"fees" streamable:""`
	Cost                     uint64    `json:"cost" streamable:""`
	RewardClaimsIncorporated []Coin    `json:"reward_claims_incorporated" streamable:""`
}
