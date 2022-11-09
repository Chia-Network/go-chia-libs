package types

import (
	"encoding/json"
	"net"
	"strings"
)

// NodeType is the type of peer (farmer, full node, etc)
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
)

// IPAddress wraps net.IP so we can fix unmarshalling issues
// due to invalid [] wrapping ipv6 addresses in Chia RPC
type IPAddress struct {
	net.IP
}

// Connection represents a single peer or internal connection
type Connection struct {
	BytesRead    uint64 `json:"bytes_read"`
	BytesWritten uint64 `json:"bytes_written"`
	//CreationTime // @TODO parse to time - is seconds as float
	//LastMessageTime // @TODO parse to time - is seconds as float
	LocalPort      uint16    `json:"local_port"`
	NodeID         string    `json:"node_id"`   // @TODO Should this be Bytes or Bytes32?
	PeakHash       string    `json:"peak_hash"` // @TODO Should this be bytes32?
	PeakHeight     uint32    `json:"peak_height"`
	PeakWeight     Uint128   `json:"peak_weight"`
	PeerHost       IPAddress `json:"peer_host"`
	PeerPort       uint16    `json:"peer_port"`
	PeerServerPort uint16    `json:"peer_server_port"`
	Type           NodeType  `json:"type"`
}

// UnmarshalJSON Unmarshals IP
func (ip *IPAddress) UnmarshalJSON(data []byte) error {
	strData := strings.ReplaceAll(string(data), "[", "")
	strData = strings.ReplaceAll(strData, "]", "")

	ipval := &net.IP{}
	err := json.Unmarshal([]byte(strData), ipval)
	if err != nil {
		return err
	}

	ip.IP = *ipval
	return nil
}
