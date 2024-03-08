package vdf

/*
#cgo CXXFLAGS: -std=c++17
#cgo LDFLAGS: -lstdc++ -lchiavdfc
#include "c_wrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

// CreateDiscriminant Creates discriminant
func CreateDiscriminant(seed []byte, length int) string {
	cSeed := C.CBytes(seed)
	defer C.free(cSeed)

	cResultStr := C.create_discriminant_wrapper((*C.uint8_t)(cSeed), C.size_t(len(seed)), C.int(length))
	defer C.free(unsafe.Pointer(cResultStr))

	// Convert the C-string to a Go string
	resultStr := C.GoString(cResultStr)

	return resultStr
}

// Prove generates a proof
func Prove(challengeHash []byte, initialEL []byte, discriminantSizeBits int, numIterations uint64) []byte {
	cChallengeHash := C.CBytes(challengeHash)
	defer C.free(cChallengeHash)

	cInitialEL := C.CBytes(initialEL)
	defer C.free(cInitialEL)

	cResult := C.prove_wrapper((*C.uint8_t)(cChallengeHash), C.size_t(len(challengeHash)), (*C.uint8_t)(cInitialEL), C.size_t(len(initialEL)), C.int(discriminantSizeBits), C.uint64_t(numIterations))
	defer C.free(unsafe.Pointer(cResult.data))

	// Convert C.ByteArray to Go []byte
	length := int(cResult.length)
	resultSlice := C.GoBytes(unsafe.Pointer(cResult.data), C.int(length))

	return resultSlice
}

// VerifyNWesolowski checks an N Wesolowski proof.
func VerifyNWesolowski(discriminant string, xS, proofBlob []byte, numIterations, discSizeBits, recursion uint64) bool {
	cDiscriminant := C.CString(discriminant)
	defer C.free(unsafe.Pointer(cDiscriminant))

	cXS := C.CBytes(xS)
	defer C.free(cXS)

	cProofBlob := C.CBytes(proofBlob)
	defer C.free(cProofBlob)

	result := C.verify_n_wesolowski_wrapper((*C.char)(cDiscriminant), C.size_t(len(discriminant)), (*C.char)(cXS), C.size_t(len(xS)), (*C.char)(cProofBlob), C.size_t(len(proofBlob)), C.uint64_t(numIterations), C.uint64_t(discSizeBits), C.uint64_t(recursion))

	return result == 1
}
