package types

import (
	"github.com/samber/mo"
)

// NFTInfo is an NFT
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/wallet/nft_wallet/nft_info.py#L21
// @TODO Streamable
type NFTInfo struct {
	LauncherID         Bytes32            `json:"launcher_id"`
	NftCoinID          Bytes32            `json:"nft_coin_id"`
	OwnerDid           mo.Option[Bytes32] `json:"owner_did"`
	RoyaltyPercentage  mo.Option[uint16]  `json:"royalty_percentage"`
	RoyaltyPuzzleHash  mo.Option[Bytes32] `json:"royalty_puzzle_hash"`
	DataUris           []string           `json:"data_uris"`
	DataHash           Bytes              `json:"data_hash"`
	MetadataURIs       []string           `json:"metadata_uris"`
	MetadataHash       Bytes              `json:"metadata_hash"`
	LicenseURIs        []string           `json:"license_uris"`
	LicenseHash        Bytes              `json:"license_hash"`
	EditionTotal       uint64             `json:"edition_total"`
	EditionNumber      uint64             `json:"edition_number"`
	UpdaterPuzhash     Bytes32            `json:"updater_puzhash"`
	ChainInfo          string             `json:"chain_info"`
	MintHeight         uint32             `json:"mint_height"`
	SupportsDid        bool               `json:"supports_did"`
	P2Address          Bytes32            `json:"p2_address"`
	PendingTransaction bool               `json:"pending_transaction"`
	MinterDid          mo.Option[Bytes32] `json:"minter_did"`
	LauncherPuzhash    Bytes32            `json:"launcher_puzhash"`
	OffChainMetadata   mo.Option[string]  `json:"off_chain_metadata"`
}
