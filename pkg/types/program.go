package types

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SerializedProgram An opaque representation of a clvm program. It has a more limited interface than a full SExp
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/program.py#L232
type SerializedProgram Bytes

const MAX_SINGLE_BYTE byte = 0x7f
const BACK_REFERENCE byte = 0xfe
const CONS_BOX_MARKER byte = 0xff

const (
	badEncErr   = "bad encoding"
	internalErr = "internal error"
)

// MarshalJSON custom hex marshaller
func (g SerializedProgram) MarshalJSON() ([]byte, error) {
	return json.Marshal(Bytes(g))
}

// UnmarshalJSON custom hex unmarshaller
func (g *SerializedProgram) UnmarshalJSON(data []byte) error {
	b := Bytes{}
	err := json.Unmarshal(data, &b)
	if err != nil {
		return err
	}

	*g = SerializedProgram(b)

	return nil
}

func SerializedLengthFromBytesTrusted(b []byte) (uint64, error) {
	var opsCounter uint64 = 1
	var position uint64 = 0
	start := len(b)

	for opsCounter > 0 {
		opsCounter--
		if len(b) == 0 {
			return 0, errors.New("unexpected end of input")
		}
		currentByte := b[0]
		b = b[1:]
		position++

		if currentByte == CONS_BOX_MARKER {
			opsCounter += 2
		} else if currentByte == BACK_REFERENCE {
			if len(b) == 0 {
				return 0, errors.New("unexpected end of input")
			}
			firstByte := b[0]
			b = b[1:]
			position++
			if firstByte > MAX_SINGLE_BYTE {
				_, length, err := decodeSize(b, firstByte)
				if err != nil {
					return 0, err
				}
				b = b[length:]
				position += length
			}
		} else if currentByte == 0x80 || currentByte <= MAX_SINGLE_BYTE {
			// This one byte we just read was the whole atom.
			// or the special case of NIL
		} else {
			_, length, err := decodeSize(b, currentByte)
			if err != nil {
				return 0, err
			}
			b = b[length:]
			position += length
		}

	}

	fmt.Println("read bytes", start, start-len(b), position)

	return position, nil
}

func decodeSize(input []byte, initialB byte) (byte, uint64, error) {

	bitMask := byte(0x80)

	if (initialB & bitMask) == 0 {
		return 0, 0, errors.New(internalErr)
	}

	var atomStartOffset byte

	b := initialB

	for (b & bitMask) != 0 {
		atomStartOffset++
		b &= 0xff ^ bitMask
		bitMask >>= 1
	}

	sizeBlob := make([]byte, atomStartOffset)
	sizeBlob[0] = b

	if atomStartOffset > 1 {
		copy(sizeBlob[1:], input)
	}

	var atomSize uint64 = 0

	if len(sizeBlob) > 6 {
		return 0, 0, errors.New(badEncErr)
	}

	for _, b := range sizeBlob {
		atomSize <<= 8
		atomSize += uint64(b)
	}

	if atomSize >= 0x400000000 {
		return 0, 0, errors.New(badEncErr)
	}

	return atomStartOffset, atomSize, nil
}
