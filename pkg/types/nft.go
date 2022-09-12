package types

type NFT struct {
	ChainInfo          string   `json:"chain_info"`
	DataHash           string   `json:"data_hash"`
	DataUris           []string `json:"data_uris"`
	LauncherID         string   `json:"launcher_id"`
	LauncherPuzhash    string   `json:"launcher_puzhash"`
	LicenseHash        string   `json:"license_hash"`
	LicenseUris        []string `json:"license_uris"`
	MetadataHash       string   `json:"metadata_hash"`
	MetadataUris       []string `json:"metadata_uris"`
	MintHeight         int      `json:"mint_height"`
	NftCoinID          string   `json:"nft_coin_id"`
	OwnerDid           string   `json:"owner_did"`
	PendingTransaction bool     `json:"pending_transaction"`
	RoyaltyPercentage  int      `json:"royalty_percentage"`
	RoyaltyPuzzleHash  string   `json:"royalty_puzzle_hash"`
	EditionNumber      int      `json:"edition_number"`
	EditionCount       int      `json:"edition_count"`
	SupportsDid        bool     `json:"supports_did"`
	UpdaterPuzhash     string   `json:"updater_puzhash"`
}
