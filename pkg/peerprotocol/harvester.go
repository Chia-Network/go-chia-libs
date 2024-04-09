package peerprotocol

import "github.com/chia-network/go-chia-libs/pkg/protocols"

// HarvesterProtocol is for interfacing with full nodes via the peer protocol
type HarvesterProtocol struct {
	*Connection
}

// NewHarvesterProtocol returns a new instance of the full node protocol
func NewHarvesterProtocol(connection *Connection) (*HarvesterProtocol, error) {
	return &HarvesterProtocol{connection}, nil
}

// Handshake performs the handshake with the peer
func (c *HarvesterProtocol) Handshake(data *protocols.HarvesterHandshake) error {
	return c.Do(protocols.ProtocolMessageTypeHarvesterHandshake, data)
}

// NewSignagePointHarvester sends a new signage point to the harvester
func (c *HarvesterProtocol) NewSignagePointHarvester(data *protocols.NewSignagePointHarvester) error {
	return c.Do(protocols.ProtocolMessageTypeNewSignagePointHarvester, data)
}
