package bech32m_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/bech32m"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

func TestKnownAddressConversions(t *testing.T) {
	// Address: Hexstr
	combinations := map[string]map[string]string{
		"xch": map[string]string{
			"xch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqm6ks6e8mvy": "000000000000000000000000000000000000000000000000000000000000dead",
			"xch1arjpkq2a5kjd7t2st93wxqd0axcnfpq04xzyjespkr0xxakslcvq3wwwdh": "e8e41b015da5a4df2d505962e301afe9b134840fa984496601b0de6376d0fe18", // Random Keys
		},
		"txch": map[string]string{
			"txch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqm6ksh7qddh": "000000000000000000000000000000000000000000000000000000000000dead",
			"txch1arjpkq2a5kjd7t2st93wxqd0axcnfpq04xzyjespkr0xxakslcvquffcvy": "e8e41b015da5a4df2d505962e301afe9b134840fa984496601b0de6376d0fe18", // Random Keys
		},
	}

	for prefix, tests := range combinations {
		for address, hexstr := range tests {
			hexbytes, err := hex.DecodeString(hexstr)
			assert.NoError(t, err)
			hexbytes32, err := types.BytesToBytes32(hexbytes)
			assert.NoError(t, err)

			// Test encoding
			generatedAddress, err := bech32m.EncodePuzzleHash(hexbytes32, prefix)
			assert.NoError(t, err)
			assert.Equal(t, address, generatedAddress)

			// Test decoding
			generatedBytes, err := bech32m.DecodePuzzleHash(address)
			assert.NoError(t, err)
			assert.Equal(t, hexbytes32, generatedBytes)
		}
	}
}
