package peerprotocol

import (
	"github.com/chia-network/go-chia-libs/pkg/protocols"
)

// FarmerProtocol is for interfacing with full nodes via the peer protocol
type FarmerProtocol struct {
	*Connection
}

// NewFarmerProtocol returns a new instance of the full node protocol
func NewFarmerProtocol(connection *Connection) (*FarmerProtocol, error) {
	return &FarmerProtocol{connection}, nil
}

// Handshake performs the handshake with the peer
func (c *FarmerProtocol) Handshake() error {
	return c.handshake(protocols.NodeTypeFarmer)
}

// DeclareProofOfSpace sends a DeclareProofOfSpace message to the peer
func (c *FarmerProtocol) DeclareProofOfSpace(data *protocols.DeclareProofOfSpace) error {
	return c.Do(protocols.ProtocolMessageTypeDeclareProofOfSpace, data)
}

func (c *FarmerProtocol) SignedValues(data *protocols.SignedValues) error {
	return c.Do(protocols.ProtocolMessageTypeSignedValues, data)

}
