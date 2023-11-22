package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
)

// Bytes is a wrapper around []byte that marshals down to hex and more closely matches types in chia-blockchain
type Bytes []byte

// String Converts to hex string
func (b Bytes) String() string {
	if len(b) == 0 {
		return ""
	}

	return fmt.Sprintf("0x%s", hex.EncodeToString(b))
}

// BytesFromHexString parses a hex string into Bytes
func BytesFromHexString(hexstr string) (Bytes, error) {
	hexstr = strings.TrimLeft(hexstr, `"`)
	hexstr = strings.TrimRight(hexstr, `"`)
	hexstr = strings.TrimPrefix(hexstr, `0x`)

	hexStrBytes := []byte(hexstr)
	dest := make(Bytes, hex.DecodedLen(len(hexStrBytes)))
	_, err := hex.Decode(dest, hexStrBytes)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

// MarshalJSON marshals Bytes into hex for json
func (b Bytes) MarshalJSON() ([]byte, error) {
	if len(b) == 0 {
		return []byte("null"), nil
	}
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	final := []byte(`"0x`)
	final = append(final, dst...)
	final = append(final, []byte(`"`)...)
	return final, nil
}

// UnmarshalJSON unmarshals hex into []byte
func (b *Bytes) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	data = bytes.TrimLeft(data, `"`)
	data = bytes.TrimRight(data, `"`)
	data = bytes.TrimPrefix(data, []byte(`0x`))

	dest := make(Bytes, hex.DecodedLen(len(data)))
	_, err := hex.Decode(dest, data)
	*b = dest
	return err
}
