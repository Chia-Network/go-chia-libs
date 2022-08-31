package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

func TestBytesToBytes32(t *testing.T) {
	shortBytes := []byte{0, 1, 2, 3}
	shortBytes32, err := types.BytesToBytes32(shortBytes)
	assert.NoError(t, err)
	assert.Equal(t, [32]byte{0, 1, 2, 3}, shortBytes32)

	longBytes := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	longBytes32, err := types.BytesToBytes32(longBytes)
	assert.Error(t, err)
	assert.Equal(t, [32]byte{}, longBytes32)

	okBytes := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	okBytes32, err := types.BytesToBytes32(okBytes)
	assert.NoError(t, err)
	assert.Equal(t, [32]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}, okBytes32)
}
