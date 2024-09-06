package vdf

/*
#cgo CXXFLAGS: -std=c++17
#cgo LDFLAGS: -lstdc++ -lchiavdfc -lgmp -lstdc++ -lm
#include "c_wrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"encoding/hex"
	"unsafe"
)

// CreateDiscriminant Creates discriminant
func CreateDiscriminant(seed []byte, length int) string {
	cSeed := C.CBytes(seed)
	defer C.free(cSeed)

	resultSize := (length + 7) / 8
	result := make([]byte, resultSize)
	C.create_discriminant_wrapper(
		(*C.uint8_t)(cSeed),
		C.size_t(len(seed)),
		C.size_t(length),
		(*C.uint8_t)(unsafe.Pointer(&result[0])),
	)

	return hex.EncodeToString(result)
}

// Prove generates a proof
func Prove(challengeHash []byte, initialEL []byte, discriminantSizeBits int, numIterations uint64) []byte {
	cChallengeHash := C.CBytes(challengeHash)
	defer C.free(cChallengeHash)

	cInitialEL := C.CBytes(initialEL)
	defer C.free(cInitialEL)

	cResult := C.prove_wrapper(
		(*C.uint8_t)(cChallengeHash),
		C.size_t(len(challengeHash)),
		(*C.uint8_t)(cInitialEL),
		C.size_t(len(initialEL)),
		C.size_t(discriminantSizeBits),
		C.uint64_t(numIterations),
	)
	defer C.free(unsafe.Pointer(cResult.data))

	// Convert C.ByteArray to Go []byte
	length := int(cResult.length)
	resultSlice := C.GoBytes(unsafe.Pointer(cResult.data), C.int(length))

	return resultSlice
}

// VerifyNWesolowski checks an N Wesolowski proof.
func VerifyNWesolowski(discriminant string, xS, proofBlob []byte, numIterations, discSizeBits, recursion uint64) bool {
	discriminantBytes, err := hex.DecodeString(discriminant)
	if err != nil {
		return false
	}

	cXS := C.CBytes(xS)
	defer C.free(cXS)

	cProofBlob := C.CBytes(proofBlob)
	defer C.free(cProofBlob)

	result := C.verify_n_wesolowski_wrapper(
		(*C.uint8_t)(unsafe.Pointer(&discriminantBytes[0])),
		C.size_t(len(discriminantBytes)),
		(*C.uchar)(cXS),
		(*C.uchar)(cProofBlob),
		C.size_t(len(proofBlob)),
		C.uint64_t(numIterations),
		C.uint64_t(recursion),
	)

	return bool(result)
}
