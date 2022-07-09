package types_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

// TestIPv4 Ensures ipv4 addresses unmarshal correctly
func TestIPv4(t *testing.T) {
	data := []byte(`{"peer_host":"1.1.1.1"}`)
	connections := &types.Connection{}
	err := json.Unmarshal(data, connections)
	assert.NoError(t, err)
	assert.Equal(t, "1.1.1.1", connections.PeerHost.String())
}

// TestStandardIPV6 Tests for the case of a standard IPv6 Address
func TestStandardIPV6(t *testing.T) {
	data := []byte(`{"peer_host":"2606:4700:4700::1111"}`)
	connections := &types.Connection{}
	err := json.Unmarshal(data, connections)
	assert.NoError(t, err)
	assert.Equal(t, "2606:4700:4700::1111", connections.PeerHost.String())
}

// TestBracketWrappedIPV6 Tests for the case of an ipv6 address being wrapped in []
// Some methods respond with ipv6 addresses this way in Chia
func TestBracketWrappedIPV6(t *testing.T) {
	data := []byte(`{"peer_host":"[2606:4700:4700::1111]"}`)
	connections := &types.Connection{}
	err := json.Unmarshal(data, connections)
	assert.NoError(t, err)
	assert.Equal(t, "2606:4700:4700::1111", connections.PeerHost.String())
}
