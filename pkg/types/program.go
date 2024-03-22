package types

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
)

// SerializedProgram An opaque representation of a clvm program. It has a more limited interface than a full SExp
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/program.py#L232
type SerializedProgram Bytes

// MaxSingleByte Max single byte
const MaxSingleByte byte = 0x7f

// BackReference back referencee marker
const BackReference byte = 0xfe

// ConsBoxMarker cons box marker
const ConsBoxMarker byte = 0xff

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

// SerializedLengthFromBytesTrusted returns the length
func SerializedLengthFromBytesTrusted(b []byte) (uint64, error) {
	reader := bytes.NewReader(b)
	var opsCounter uint64 = 1

	for opsCounter > 0 {
		opsCounter--

		var currentByte byte
		err := binary.Read(reader, binary.BigEndian, &currentByte)
		if err != nil {
			if err == io.EOF {
				return 0, errors.New("unexpected end of input")
			}
			return 0, err
		}

		if currentByte == ConsBoxMarker {
			opsCounter += 2
		} else if currentByte == BackReference {
			var firstByte byte
			err = binary.Read(reader, binary.BigEndian, &firstByte)
			if err != nil {
				return 0, errors.New("unexpected end of input")
			}
			if firstByte > MaxSingleByte {
				pathSize, err := decodeSize(reader, firstByte)
				if err != nil {
					return 0, err
				}
				_, err = reader.Seek(int64(pathSize), io.SeekCurrent)
				if err != nil {
					return 0, errors.New("bad encoding")
				}
			}
		} else if currentByte == 0x80 || currentByte <= MaxSingleByte {
			// This one byte we just read was the whole atom or the special case of NIL.
		} else {
			blobSize, err := decodeSize(reader, currentByte)
			if err != nil {
				return 0, err
			}
			_, err = reader.Seek(int64(blobSize), io.SeekCurrent)
			if err != nil {
				return 0, errors.New("bad encoding")
			}
		}
	}

	position, err := reader.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}
	return uint64(position), nil
}

func decodeSize(reader *bytes.Reader, initialB byte) (uint64, error) {
	_, length, err := decodeSizeWithOffset(reader, initialB)
	return length, err
}

func decodeSizeWithOffset(reader *bytes.Reader, initialB byte) (uint64, uint64, error) {
	bitMask := byte(0x80)

	if (initialB & bitMask) == 0 {
		return 0, 0, errors.New(internalErr)
	}

	var atomStartOffset uint64 = 0
	b := initialB
	for (b & bitMask) != 0 {
		atomStartOffset++
		b &= 0xff ^ bitMask
		bitMask >>= 1
	}

	sizeBlob := make([]byte, atomStartOffset)
	sizeBlob[0] = b

	if atomStartOffset > 1 {
		// We need to read atomStartOffset-1 more bytes
		_, err := io.ReadFull(reader, sizeBlob[1:])
		if err != nil {
			return 0, 0, err
		}
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
