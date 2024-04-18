package protocols_test

import (
	"encoding/hex"
	"github.com/chia-network/go-chia-libs/pkg/protocols"
	"github.com/chia-network/go-chia-libs/pkg/streamable"
	"github.com/chia-network/go-chia-libs/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHarvesterHandshakeMarshal(t *testing.T) {
	// Hex to bytes
	fpk1, err := hex.DecodeString("c31c4d1b4cc006eb8ea294656aba445f43e4e3d78910392c6880a1e9bd806d0419acd4e60c4abd533f11dafafbf23e7f")
	assert.NoError(t, err)

	fpk2, err := hex.DecodeString("af397582efbcdc276bbe5ee3b73981c0c9ff5b3fafeae5c34a0f9ea27a67023966a53d6a3f5c37ee9272cdbe3f3cfc9c")
	assert.NoError(t, err)

	ppk1, err := hex.DecodeString("0da874d9ede1dc629731c05f853c7d20eb987b00465f1acdd84d1256b17762a2b22b3bb2cbd25ac3178515b4d75311fd")
	assert.NoError(t, err)

	ppk2, err := hex.DecodeString("d74f2a6f8e9960417df6665fefd7598dbdc5a5fb75b9e0b0f2f3ffc9b270de54dc368af2fe63aa122e7b558f792bc891")
	assert.NoError(t, err)

	rp := protocols.HarvesterHandshake{
		FarmerPublicKeys: []types.G1Element{types.G1Element(fpk1), types.G1Element(fpk2)},
		PoolPublicKeys:   []types.G1Element{types.G1Element(ppk1), types.G1Element(ppk2)},
	}

	b, err := streamable.Marshal(rp)
	assert.NoError(t, err)

	assert.Len(t, b, 200)

	// Unmarshal
	rp2 := &protocols.HarvesterHandshake{}
	err = streamable.Unmarshal(b, rp2)
	assert.NoError(t, err)

	assert.Len(t, rp2.FarmerPublicKeys, 2)
	assert.Len(t, rp2.PoolPublicKeys, 2)

	assert.Equal(t, fpk1, rp2.FarmerPublicKeys[0][:])
	assert.Equal(t, fpk2, rp2.FarmerPublicKeys[1][:])
	assert.Equal(t, ppk1, rp2.PoolPublicKeys[0][:])
	assert.Equal(t, ppk2, rp2.PoolPublicKeys[1][:])
}
