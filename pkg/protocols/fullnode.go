package protocols

import (
	"github.com/chia-network/go-chia-libs/pkg/types"
)

// RequestPeers is an empty struct
type RequestPeers struct{}

// RespondPeers is the format for the request_peers response
type RespondPeers struct {
	PeerList []types.TimestampedPeerInfo `streamable:""`
}
