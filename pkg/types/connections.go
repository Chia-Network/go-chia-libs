package types

import (
	"github.com/samber/mo"
)

// NodeType is the type of peer (farmer, full node, etc)
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/server/outbound_message.py#L12
type NodeType uint8

const (
	// NodeTypeFullNode NodeType for full node
	NodeTypeFullNode NodeType = 1

	// NodeTypeHarvester NodeType for Harvester
	NodeTypeHarvester NodeType = 2

	// NodeTypeFarmer NodeType for Farmer
	NodeTypeFarmer NodeType = 3

	// NodeTypeTimelord NodeType for Timelord
	NodeTypeTimelord NodeType = 4

	// NodeTypeIntroducer NodeType for Introducer
	NodeTypeIntroducer NodeType = 5

	// NodeTypeWallet NodeType for Wallet
	NodeTypeWallet NodeType = 6

	// NodeTypeDataLayer Data Layer Node
	NodeTypeDataLayer NodeType = 7
)

// Connection represents a single peer or internal connection
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/rpc/rpc_server.py#L119
type Connection struct {
	Type            NodeType  `json:"type"`
	LocalPort       uint16    `json:"local_port"`
	PeerHost        string    `json:"peer_host"` // Can be hostname as well as IP
	PeerPort        uint16    `json:"peer_port"`
	PeerServerPort  uint16    `json:"peer_server_port"`
	NodeID          Bytes32   `json:"node_id"`
	CreationTime    Timestamp `json:"creation_time"`
	BytesRead       uint64    `json:"bytes_read"`
	BytesWritten    uint64    `json:"bytes_written"`
	LastMessageTime Timestamp `json:"last_message_time"`

	// Full Node
	PeakHash   mo.Option[Bytes32] `json:"peak_hash"`
	PeakHeight mo.Option[uint32]  `json:"peak_height"`
	PeakWeight mo.Option[Uint128] `json:"peak_weight"`
}
