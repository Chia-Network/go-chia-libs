package types

import (
	"fmt"
)

// Bytes32ToBytes returns []byte from [32]byte
func Bytes32ToBytes(bytes32 [32]byte) []byte {
	return bytes32[:]
}

// BytesToBytes32 Returns [32]byte from []byte
// If input is shorter than 32 bytes, the end will be padded
// If the input is longer than 32 bytes, an error will be returned
func BytesToBytes32(bytes []byte) ([32]byte, error) {
	var fixedLen [32]byte
	if len(bytes) > 32 {
		return fixedLen, fmt.Errorf("input bytes is longer than 32 bytes")
	}

	copy(fixedLen[:], bytes)
	return fixedLen, nil
}
