package types

// CrawlerPeerCounts peer_count data from the crawler
type CrawlerPeerCounts struct {
	TotalLast5Days uint           `json:"total_last_5_days"`
	ReliableNodes  uint           `json:"reliable_nodes"`
	IPV4Last5Days  uint           `json:"ipv4_last_5_days"`
	IPV6Last5Days  uint           `json:"ipv6_last_5_days"`
	Versions       map[string]int `json:"versions"`
}
