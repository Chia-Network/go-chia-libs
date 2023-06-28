package protocol

import (
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
