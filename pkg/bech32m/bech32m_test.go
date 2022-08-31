package bech32m_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/bech32m"
)

func TestKnownAddressConversions(t *testing.T) {
	// Address: Hexstr
	combinations := map[string]map[string]string{
		"xch": map[string]string{
			"xch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqm6ks6e8mvy": "000000000000000000000000000000000000000000000000000000000000dead",
		},
		"txch": map[string]string{
			"txch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqm6ksh7qddh": "000000000000000000000000000000000000000000000000000000000000dead",
		},
	}

	for prefix, tests := range combinations {
		for address, hexstr := range tests {
			hexbytes, err := hex.DecodeString(hexstr)
			assert.NoError(t, err)

			// Test encoding
			generatedAddress, err := bech32m.EncodePuzzleHash(hexbytes, prefix)
			assert.NoError(t, err)
			assert.Equal(t, address, generatedAddress)

			// Test decoding
			generatedBytes, err := bech32m.DecodePuzzleHash(address)
			assert.NoError(t, err)
			assert.Equal(t, hexbytes, generatedBytes)
		}
	}
}
