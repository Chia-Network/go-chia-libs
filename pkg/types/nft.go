package types

// NFT is an NFT
type NFT struct {
	ChainInfo          string   `json:"chain_info"`
	DataHash           Bytes    `json:"data_hash"`
	DataUris           []string `json:"data_uris"`
	LauncherID         Bytes32  `json:"launcher_id"`
	LauncherPuzhash    Bytes32  `json:"launcher_puzhash"`
	LicenseHash        Bytes    `json:"license_hash"`
	LicenseURIs        []string `json:"license_uris"`
	MetadataHash       Bytes    `json:"metadata_hash"`
	MetadataURIs       []string `json:"metadata_uris"`
	MintHeight         uint32   `json:"mint_height"`
	MinterDid          *Bytes32 `json:"minter_did"`
	NftCoinID          Bytes32  `json:"nft_coin_id"`
	OwnerDid           *Bytes32 `json:"owner_did"`
	P2Address          Bytes32  `json:"p2_address"`
	PendingTransaction bool     `json:"pending_transaction"`
	RoyaltyPercentage  uint32   `json:"royalty_percentage"`
	RoyaltyPuzzleHash  *Bytes32 `json:"royalty_puzzle_hash"`
	EditionNumber      uint32   `json:"edition_number"`
	EditionCount       uint32   `json:"edition_count"`
	SupportsDid        bool     `json:"supports_did"`
	UpdaterPuzhash     Bytes32  `json:"updater_puzhash"`
}
