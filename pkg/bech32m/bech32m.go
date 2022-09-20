// Package bech32m reference implementation for Bech32m and segwit addresses.
// Copyright (c) 2021 Takatoshi Nakagawa
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package bech32m

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

var charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

var bech32mConst = 0x2bc830a3

var generator = []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}

func polymod(values []byte) int {
	// Internal function that computes the Bech32 checksum.
	chk := 1
	for _, v := range values {
		top := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ int(v)
		for i := 0; i < 5; i++ {
			if (top>>uint(i))&1 == 1 {
				chk ^= generator[i]
			} else {
				chk ^= 0
			}
		}
	}
	return chk
}

func hrpExpand(hrp string) []byte {
	// Expand the HRP into values for checksum computation.
	ret := []byte{}
	for _, c := range hrp {
		ret = append(ret, byte(c>>5))
	}
	ret = append(ret, 0)
	for _, c := range hrp {
		ret = append(ret, byte(c&31))
	}
	return ret
}

func verifyChecksum(hrp string, data []byte) bool {
	// Verify a checksum given HRP and converted data characters.
	c := polymod(append(hrpExpand(hrp), data...))
	return c == bech32mConst
}

func createChecksum(hrp string, data []byte) []byte {
	// Compute the checksum values given HRP and data.
	values := append(append(hrpExpand(hrp), data...), []byte{0, 0, 0, 0, 0, 0}...)
	mod := polymod(values) ^ bech32mConst
	ret := make([]byte, 6)
	for i := 0; i < len(ret); i++ {
		ret[i] = byte(mod>>uint(5*(5-i))) & 31
	}
	return ret
}

// Encode compute a Bech32m string given HRP and data values.
func Encode(hrp string, data []byte) string {
	combined := append(data, createChecksum(hrp, data)...)
	var ret bytes.Buffer
	ret.WriteString(hrp)
	ret.WriteString("1")
	for _, p := range combined {
		ret.WriteByte(charset[p])
	}
	return ret.String()
}

// Decode validate a Bech32m string, and determine HRP and data.
func Decode(bechString string) (string, []byte, error) {
	if len(bechString) > 90 {
		return "", nil, fmt.Errorf("overall max length exceeded")
	}
	if strings.ToLower(bechString) != bechString && strings.ToUpper(bechString) != bechString {
		return "", nil, fmt.Errorf("mixed case")
	}
	bechString = strings.ToLower(bechString)
	pos := strings.LastIndex(bechString, "1")
	if pos < 0 {
		return "", nil, fmt.Errorf("no separator character")
	}
	if pos < 1 {
		return "", nil, fmt.Errorf("empty HRP")
	}
	if pos+7 > len(bechString) {
		return "", nil, fmt.Errorf("too short checksum")
	}
	hrp := bechString[0:pos]
	for _, c := range hrp {
		if c < 33 || c > 126 {
			return "", nil, fmt.Errorf("HRP character out of range")
		}
	}
	data := []byte{}
	for p := pos + 1; p < len(bechString); p++ {
		d := strings.Index(charset, fmt.Sprintf("%c", bechString[p]))
		if d == -1 {
			if p+6 > len(bechString) {
				return "", nil, fmt.Errorf("invalid character in checksum")
			}
			return "", nil, fmt.Errorf("invalid data character")
		}
		data = append(data, byte(d))
	}
	validChecksum := verifyChecksum(hrp, data)
	if !validChecksum {
		return "", nil, fmt.Errorf("invalid checksum")
	}
	return hrp, data[:len(data)-6], nil
}

func convertbits(data []byte, frombits, tobits uint, pad bool) ([]byte, error) {
	// General power-of-2 base conversion.
	acc := 0
	bits := uint(0)
	var ret []byte
	maxv := (1 << tobits) - 1
	maxAcc := (1 << (frombits + tobits - 1)) - 1
	for _, value := range data {
		acc = ((acc << frombits) | int(value)) & maxAcc
		bits += frombits
		for bits >= tobits {
			bits -= tobits
			ret = append(ret, byte((acc>>bits)&maxv))
		}
	}
	if pad {
		if bits > 0 {
			ret = append(ret, byte((acc<<(tobits-bits))&maxv))
		}
	} else if bits >= frombits {
		return nil, fmt.Errorf("more than 4 padding bits")
	} else if ((acc << (tobits - bits)) & maxv) != 0 {
		return nil, fmt.Errorf("non-zero padding in %d-to-%d conversion", tobits, frombits)
	}
	return ret, nil
}

// EncodePuzzleHash encode to an address
func EncodePuzzleHash(puzzleHash [32]byte, prefix string) (string, error) {
	data, err := convertbits(types.Bytes32ToBytes(puzzleHash), 8, 5, true)
	if err != nil {
		return "", err
	}
	ret := Encode(prefix, data)
	_, _, err = DecodePuzzleHash(ret)
	if err != nil {
		return "", err
	}
	return ret, nil
}

// DecodePuzzleHash Decodes an address to a puzzle hash
func DecodePuzzleHash(addr string) (string, [32]byte, error) {
	hrp, data, err := Decode(addr)
	if err != nil {
		return "", [32]byte{}, err
	}
	if len(data) < 1 {
		return "", [32]byte{}, fmt.Errorf("empty data section")
	}
	res, err := convertbits(data, 5, 8, false)
	if err != nil {
		return "", [32]byte{}, err
	}
	b, err := types.BytesToBytes32(res)
	if err != nil {
		return "", [32]byte{}, err
	}
	return hrp, b, nil
}
