package types

import (
	"bytes"
	"encoding/hex"
)

// Bytes is a wrapper around []byte that marshals down to hex and more closely matches types in chia-blockchain
type Bytes []byte

// MarshalJSON marshals Bytes into hex for json
func (b Bytes) MarshalJSON() ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	final := []byte(`"0x`)
	final = append(final, dst...)
	final = append(final, []byte(`"`)...)
	return final, nil
}

// UnmarshalText unmarshals hex into []byte
func (b *Bytes) UnmarshalText(data []byte) error {
	data = bytes.TrimLeft(data, `"`)
	data = bytes.TrimRight(data, `"`)
	data = bytes.TrimPrefix(data, []byte(`0x`))

	dest := make(Bytes, hex.DecodedLen(len(data)))
	_, err := hex.Decode(dest, data)
	*b = dest
	return err
}
