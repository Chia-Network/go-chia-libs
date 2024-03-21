package peerprotocol

import (
	"github.com/chia-network/go-chia-libs/pkg/protocols"
)

// FullNodeProtocol is for interfacing with full nodes via the peer protocol
type FullNodeProtocol struct {
	connection *Connection
}

// NewFullNodeProtocol returns a new instance of the full node protocol
func NewFullNodeProtocol(connection *Connection) (*FullNodeProtocol, error) {
	fnp := &FullNodeProtocol{connection: connection}

	return fnp, nil
}

// RequestPeers asks the current peer to respond with their current peer list
func (c *FullNodeProtocol) RequestPeers() error {
	return c.connection.Do(protocols.ProtocolMessageTypeRequestPeers, &protocols.RequestPeers{})
}

// RequestBlock asks the current peer to respond with a block
func (c *FullNodeProtocol) RequestBlock(data *protocols.RequestBlock) error {
	return c.connection.Do(protocols.ProtocolMessageTypeRequestBlock, data)
}
