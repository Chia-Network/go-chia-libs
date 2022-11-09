package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// Bytes100 Helper type with custom json handling for [100]byte
type Bytes100 [100]byte

// MarshalJSON marshals Bytes into hex for json
func (b Bytes100) MarshalJSON() ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b[:])
	final := []byte(`"0x`)
	final = append(final, dst...)
	final = append(final, []byte(`"`)...)
	return final, nil
}

// UnmarshalJSON unmarshals hex into []byte
func (b *Bytes100) UnmarshalJSON(data []byte) error {
	data = bytes.TrimLeft(data, `"`)
	data = bytes.TrimRight(data, `"`)
	data = bytes.TrimPrefix(data, []byte(`0x`))

	dest := make(Bytes, 100)
	_, err := hex.Decode(dest, data)
	if err != nil {
		return err
	}
	b100, err := BytesToBytes100(dest)
	if err != nil {
		return err
	}
	*b = b100
	return err
}

// Bytes100ToBytes returns []byte from [100]byte
func Bytes100ToBytes(bytes100 Bytes100) Bytes {
	return bytes100[:]
}

// BytesToBytes100 Returns Bytes100 from Bytes
// If input is shorter than 100 bytes, the end will be padded
// If the input is longer than 100 bytes, an error will be returned
func BytesToBytes100(bytes Bytes) (Bytes100, error) {
	var fixedLen Bytes100
	if len(bytes) > 100 {
		return fixedLen, fmt.Errorf("input bytes is longer than 100 bytes")
	}

	copy(fixedLen[:], bytes)
	return fixedLen, nil
}
