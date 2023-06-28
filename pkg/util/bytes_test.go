package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/util"
)

func TestShiftNBytes(t *testing.T) {
	origBytes := []byte{
		uint8(0),
		uint8(1),
		uint8(2),
		uint8(3),
		uint8(4),
		uint8(5),
		uint8(6),
		uint8(7),
	}

	// Ensure we're in a good starting place before changing things around
	assert.Equal(t, 8, len(origBytes))
	assert.Equal(t, uint8(0), origBytes[0])
	assert.Equal(t, uint8(7), origBytes[7])

	shift2, origBytes, err := util.ShiftNBytes(2, origBytes)

	assert.NoError(t, err)

	// Check expected lengths
	assert.Len(t, shift2, 2)
	assert.Len(t, origBytes, 6)

	// Check actual expected values
	assert.Equal(t, uint8(0), shift2[0])
	assert.Equal(t, uint8(1), shift2[1])
	assert.Equal(t, uint8(2), origBytes[0])
	assert.Equal(t, uint8(7), origBytes[5])

	// Test pulling off too many bytes
	shiftTooMany, origBytes, err := util.ShiftNBytes(7, origBytes)

	assert.Error(t, err)
	assert.Nil(t, shiftTooMany)
	assert.Len(t, origBytes, 6)
}
