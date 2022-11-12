package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// Bytes32 Helper type with custom json handling for [32]byte
type Bytes32 [32]byte

// String Converts to hex string
func (b Bytes32) String() string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(b[:]))
}

// Bytes32FromHexString parses a hex string into Bytes32
func Bytes32FromHexString(hexstr string) (Bytes32, error) {
	b, err := BytesFromHexString(hexstr)
	if err != nil {
		return Bytes32{}, err
	}

	dest, err := BytesToBytes32(b)
	if err != nil {
		return Bytes32{}, err
	}
	return dest, nil
}

// MarshalJSON marshals Bytes into hex for json
func (b Bytes32) MarshalJSON() ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b[:])
	final := []byte(`"0x`)
	final = append(final, dst...)
	final = append(final, []byte(`"`)...)
	return final, nil
}

// UnmarshalJSON unmarshals hex into []byte
func (b *Bytes32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	data = bytes.TrimLeft(data, `"`)
	data = bytes.TrimRight(data, `"`)
	data = bytes.TrimPrefix(data, []byte(`0x`))

	dest := make(Bytes, 32)
	_, err := hex.Decode(dest, data)
	if err != nil {
		return err
	}
	b32, err := BytesToBytes32(dest)
	if err != nil {
		return err
	}
	*b = b32
	return err
}

// Bytes32ToBytes returns []byte from [32]byte
func Bytes32ToBytes(bytes32 Bytes32) Bytes {
	return bytes32[:]
}

// BytesToBytes32 Returns Bytes32 from Bytes
// If input is shorter than 32 bytes, the end will be padded
// If the input is longer than 32 bytes, an error will be returned
func BytesToBytes32(bytes Bytes) (Bytes32, error) {
	var fixedLen Bytes32
	if len(bytes) > 32 {
		return fixedLen, fmt.Errorf("input bytes is longer than 32 bytes")
	}

	copy(fixedLen[:], bytes)
	return fixedLen, nil
}
