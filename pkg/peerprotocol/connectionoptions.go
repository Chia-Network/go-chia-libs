package peerprotocol

import (
	"crypto/tls"
	"time"
)

// ConnectionOptionFunc can be used to customize a new Connection
type ConnectionOptionFunc func(connection *Connection) error

// WithHandshakeTimeout sets the handshake timeout
func WithHandshakeTimeout(timeout time.Duration) ConnectionOptionFunc {
	return func(c *Connection) error {
		c.handshakeTimeout = timeout
		return nil
	}
}

// WithPeerPort sets the port for the peer
func WithPeerPort(port uint16) ConnectionOptionFunc {
	return func(c *Connection) error {
		c.peerPort = port
		return nil
	}
}

// WithPeerKeyPair sets the keypair for the peer
func WithPeerKeyPair(keypair tls.Certificate) ConnectionOptionFunc {
	return func(c *Connection) error {
		c.peerKeyPair = &keypair
		return nil
	}
}

// WithNetworkID sets the network ID for the peer
func WithNetworkID(networkID string) ConnectionOptionFunc {
	return func(c *Connection) error {
		c.networkID = networkID
		return nil
	}
}
