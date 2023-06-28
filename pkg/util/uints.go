package util

import (
	"encoding/binary"
)

// Uint8ToBytes Converts uint8 to []byte
// Kind of pointless, since byte is uint8, but here for consistency with the other methods
func Uint8ToBytes(num uint8) []byte {
	return []byte{num}
}

// BytesToUint8 returns uint8 from []byte
// if you have more than one byte in your []byte this wont work like you think
func BytesToUint8(bytes []byte) uint8 {
	return bytes[0]
}

// Uint16ToBytes Converts uint16 to []byte
func Uint16ToBytes(num uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, num)

	return b
}

// BytesToUint16 returns uint16 from []byte
// if you have more than two bytes in your []byte this wont work like you think
func BytesToUint16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}

// Uint32ToBytes Converts uint32 to []byte
func Uint32ToBytes(num uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, num)

	return b
}

// BytesToUint32 returns uint32 from []byte
// if you have more than four bytes in your []byte this wont work like you think
func BytesToUint32(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes)
}

// Uint64ToBytes Converts uint64 to []byte
func Uint64ToBytes(num uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, num)

	return b
}

// BytesToUint64 returns uint64 from []byte
// if you have more than eight bytes in your []byte this wont work like you think
func BytesToUint64(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}
