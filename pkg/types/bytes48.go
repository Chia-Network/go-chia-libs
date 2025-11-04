package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// Bytes48 Helper type with custom json handling for [48]byte
type Bytes48 [48]byte

// String Converts to hex string
func (b Bytes48) String() string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(b[:]))
}

// Bytes48FromHexString parses a hex string into Bytes48
func Bytes48FromHexString(hexstr string) (Bytes48, error) {
	b, err := BytesFromHexString(hexstr)
	if err != nil {
		return Bytes48{}, err
	}

	dest, err := BytesToBytes48(b)
	if err != nil {
		return Bytes48{}, err
	}
	return dest, nil
}

// MarshalJSON marshals Bytes into hex for json
func (b Bytes48) MarshalJSON() ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b[:])
	final := []byte(`"0x`)
	final = append(final, dst...)
	final = append(final, []byte(`"`)...)
	return final, nil
}

// UnmarshalJSON unmarshals hex into []byte
func (b *Bytes48) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	data = bytes.TrimLeft(data, `"`)
	data = bytes.TrimRight(data, `"`)
	data = bytes.TrimPrefix(data, []byte(`0x`))

	dest := make(Bytes, 48)
	_, err := hex.Decode(dest, data)
	if err != nil {
		return err
	}
	b48, err := BytesToBytes48(dest)
	if err != nil {
		return err
	}
	*b = b48
	return err
}

// Bytes48ToBytes returns []byte from [48]byte
func Bytes48ToBytes(bytes48 Bytes48) Bytes {
	return bytes48[:]
}

// BytesToBytes48 Returns Bytes48 from Bytes
// If input is shorter than 48 bytes, the end will be padded
// If the input is longer than 48 bytes, an error will be returned
func BytesToBytes48(bytes Bytes) (Bytes48, error) {
	var fixedLen Bytes48
	if len(bytes) > 48 {
		return fixedLen, fmt.Errorf("input bytes is longer than 48 bytes")
	}

	copy(fixedLen[:], bytes)
	return fixedLen, nil
}
