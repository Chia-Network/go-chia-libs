package peerprotocol

import (
	"errors"
	"github.com/chia-network/go-chia-libs/pkg/protocols"
	"time"
)

// ErrInvalidNodeType is returned when the node type is invalid
var ErrInvalidNodeType = errors.New("invalid node type")

type IHarvesterProtocol interface {
	Handshake(nodeType protocols.NodeType) error
	HarvesterHandshake(data *protocols.HarvesterHandshake) error
	NewSignagePointHarvester(data *protocols.NewSignagePointHarvester) error
	RequestSignatures(data *protocols.RequestSignatures) error
	ReadOne(timeout time.Duration) (*protocols.Message, error)
}

// HarvesterProtocol is for interfacing with full nodes via the peer protocol
type HarvesterProtocol struct {
	conn *Connection
}

// NewHarvesterProtocol returns a new instance of the full node protocol
func NewHarvesterProtocol(connection *Connection) (IHarvesterProtocol, error) {
	return &HarvesterProtocol{connection}, nil
}

// Handshake performs the handshake with the peer
func (c *HarvesterProtocol) Handshake(nodeType protocols.NodeType) error {
	if nodeType != protocols.NodeTypeHarvester && nodeType != protocols.NodeTypeFarmer {
		return ErrInvalidNodeType
	}
	return c.conn.handshake(nodeType)
}

// HarvesterHandshake performs the handshake with the peer
func (c *HarvesterProtocol) HarvesterHandshake(data *protocols.HarvesterHandshake) error {
	return c.conn.Do(protocols.ProtocolMessageTypeHarvesterHandshake, data)
}

// NewSignagePointHarvester sends a new signage point to the harvester
func (c *HarvesterProtocol) NewSignagePointHarvester(data *protocols.NewSignagePointHarvester) error {
	return c.conn.Do(protocols.ProtocolMessageTypeNewSignagePointHarvester, data)
}

// RequestSignatures sends a request for signatures to the harvester
func (c *HarvesterProtocol) RequestSignatures(data *protocols.RequestSignatures) error {
	return c.conn.Do(protocols.ProtocolMessageTypeRequestSignatures, data)
}

// ReadOne reads a single message from the connection
func (c *HarvesterProtocol) ReadOne(timeout time.Duration) (*protocols.Message, error) {
	return c.conn.ReadOne(timeout)
}
