package types

// TimestampedPeerInfo contains information about peers with timestamps
type TimestampedPeerInfo struct {
	Host      string `streamable:""`
	Port      uint16 `streamable:""`
	Timestamp uint64 `streamable:""`
}
