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

// NewPeek is the format for the new_peak response
type NewPeek struct {
	HeaderHash                types.Bytes32 `streamable:""`
	Height                    uint32        `streamable:""`
	Weight                    types.Uint128 `streamable:""`
	ForkPointWithPreviousPeak uint32        `streamable:""`
	UnfinishedRewardBlockHash types.Bytes32 `streamable:""`
}

// RequestBlock is the format for the request_block request
type RequestBlock struct {
	Height                  uint32 `streamable:""`
	IncludeTransactionBlock bool   `streamable:""`
}

// RespondBlock is the format for the respond_block response
type RespondBlock struct {
	Block types.FullBlock `streamable:""`
}

