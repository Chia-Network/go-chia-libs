package peerprotocol

import (
	"github.com/chia-network/go-chia-libs/pkg/protocols"
)

// FullNodeProtocol is for interfacing with full nodes via the peer protocol
type FullNodeProtocol struct {
	*Connection
}

// NewFullNodeProtocol returns a new instance of the full node protocol
func NewFullNodeProtocol(connection *Connection) (*FullNodeProtocol, error) {
	return &FullNodeProtocol{connection}, nil
}

// Handshake performs the handshake with the peer
func (c *FullNodeProtocol) Handshake() error {
	return c.handshake(protocols.NodeTypeFullNode)
}

// RequestPeers asks the current peer to respond with their current peer list
func (c *FullNodeProtocol) RequestPeers() error {
	return c.Do(protocols.ProtocolMessageTypeRequestPeers, &protocols.RequestPeers{})
}

// RequestBlock asks the current peer to respond with a block
func (c *FullNodeProtocol) RequestBlock(data *protocols.RequestBlock) error {
	return c.Do(protocols.ProtocolMessageTypeRequestBlock, data)
}
