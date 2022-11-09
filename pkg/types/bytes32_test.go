package types_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

func TestBytesToBytes32(t *testing.T) {
	shortBytes := types.Bytes{0, 1, 2, 3}
	shortBytes32, err := types.BytesToBytes32(shortBytes)
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32{0, 1, 2, 3}, shortBytes32)

	longBytes := types.Bytes{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	longBytes32, err := types.BytesToBytes32(longBytes)
	assert.Error(t, err)
	assert.Equal(t, types.Bytes32{}, longBytes32)

	okBytes := types.Bytes{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	okBytes32, err := types.BytesToBytes32(okBytes)
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}, okBytes32)
}

func TestMarshalBytes32(t *testing.T) {
	b32 := types.Bytes32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	marshalled, err := json.Marshal(b32)
	assert.NoError(t, err)
	assert.Equal(t, `"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"`, string(marshalled))
}

func TestUnmarshalBytes32(t *testing.T) {
	hexstr := `"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"`
	expected := types.Bytes32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	actual := types.Bytes32{}
	err := json.Unmarshal([]byte(hexstr), &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUnmarshalBytes32_Nil(t *testing.T) {
	hexstr := `""`
	expected := types.Bytes32{}
	actual := types.Bytes32{}
	err := json.Unmarshal([]byte(hexstr), &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
