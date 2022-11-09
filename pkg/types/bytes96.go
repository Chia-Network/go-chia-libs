package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// Bytes96 Helper type with custom json handling for [96]byte
type Bytes96 [96]byte

// String Converts to hex string
func (b Bytes96) String() string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(b[:]))
}

// MarshalJSON marshals Bytes into hex for json
func (b Bytes96) MarshalJSON() ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b[:])
	final := []byte(`"0x`)
	final = append(final, dst...)
	final = append(final, []byte(`"`)...)
	return final, nil
}

// UnmarshalJSON unmarshals hex into []byte
func (b *Bytes96) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	data = bytes.TrimLeft(data, `"`)
	data = bytes.TrimRight(data, `"`)
	data = bytes.TrimPrefix(data, []byte(`0x`))

	dest := make(Bytes, 96)
	_, err := hex.Decode(dest, data)
	if err != nil {
		return err
	}
	b96, err := BytesToBytes96(dest)
	if err != nil {
		return err
	}
	*b = b96
	return err
}

// Bytes96ToBytes returns []byte from [96]byte
func Bytes96ToBytes(bytes96 Bytes96) Bytes {
	return bytes96[:]
}

// BytesToBytes96 Returns Bytes96 from Bytes
// If input is shorter than 96 bytes, the end will be padded
// If the input is longer than 96 bytes, an error will be returned
func BytesToBytes96(bytes Bytes) (Bytes96, error) {
	var fixedLen Bytes96
	if len(bytes) > 96 {
		return fixedLen, fmt.Errorf("input bytes is longer than 96 bytes")
	}

	copy(fixedLen[:], bytes)
	return fixedLen, nil
}
