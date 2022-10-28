package types

// NFT is an NFT
type NFT struct {
	ChainInfo          string   `json:"chain_info"`
	DataHash           string   `json:"data_hash"`
	DataUris           []string `json:"data_uris"`
	LauncherID         string   `json:"launcher_id"`
	LauncherPuzhash    string   `json:"launcher_puzhash"`
	LicenseHash        string   `json:"license_hash"`
	LicenseURIs        []string `json:"license_uris"`
	MetadataHash       string   `json:"metadata_hash"`
	MetadataURIs       []string `json:"metadata_uris"`
	MintHeight         uint32   `json:"mint_height"`
	MinterDid          string   `json:"minter_did"`
	NftCoinID          string   `json:"nft_coin_id"`
	OwnerDid           string   `json:"owner_did"`
	P2Address          string   `json:"p2_address"`
	PendingTransaction bool     `json:"pending_transaction"`
	RoyaltyPercentage  uint32   `json:"royalty_percentage"`
	RoyaltyPuzzleHash  string   `json:"royalty_puzzle_hash"`
	EditionNumber      uint32   `json:"edition_number"`
	EditionCount       uint32   `json:"edition_count"`
	SupportsDid        bool     `json:"supports_did"`
	UpdaterPuzhash     string   `json:"updater_puzhash"`
}
