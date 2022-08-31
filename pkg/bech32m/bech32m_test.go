package bech32m_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/bech32m"
)

func TestKnownAddressConversions(t *testing.T) {
	// Address: Hexstr
	mainnetCombinations := map[string]string{
		"xch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqm6ks6e8mvy": "000000000000000000000000000000000000000000000000000000000000dead",
	}

	testnetCombinations := map[string]string{
		"txch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqm6ksh7qddh": "000000000000000000000000000000000000000000000000000000000000dead",
	}

	for address, hexstr := range mainnetCombinations {
		hexbytes, err := hex.DecodeString(hexstr)
		assert.NoError(t, err)
		generatedAddress, err := bech32m.EncodePuzzleHash(hexbytes, "xch")
		assert.NoError(t, err)
		assert.Equal(t, address, generatedAddress)
	}

	for address, hexstr := range testnetCombinations {
		hexbytes, err := hex.DecodeString(hexstr)
		assert.NoError(t, err)
		generatedAddress, err := bech32m.EncodePuzzleHash(hexbytes, "txch")
		assert.NoError(t, err)
		assert.Equal(t, address, generatedAddress)
	}
}
