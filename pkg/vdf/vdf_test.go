package vdf_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/vdf"
)

const (
	challengeHash = "885b99e5f88f762ac2af47697712050280a858ee68925cdd41d89e15e2775518"
	hexSeed       = "885b99e5f88f762ac2af47697712050280a858ee68925cdd41d89e15e2775518"
	proofHex      = "02004f5b791b1a749e07bedb221a10d4fded3c5e45586ab908753bbdd9f882f89633da35036b79956998e7755ab979eb35c7a397fa7acceb9d8f9731284dba9aa306c97b54733265e1da7d265d383f6957569fef4a11c2682de574bbc1e13a3a200901000000b1cd61a82f0f76f7a1884ecb0cb35cfc8e3213ce5e53adb67c4033bbd88159503b5da94cdeda331c63eebd94e047dde9efb76337f10a90e0f1a6ca708c40b07b65fa94d25a8a4b8070fc26c4934910bc8f109e2837d1c6f586833c2b700276410100"
)

func TestCreateDiscriminant(t *testing.T) {
	seedBytes, err := hex.DecodeString(hexSeed)
	assert.NoError(t, err)

	discriminant := vdf.CreateDiscriminant(seedBytes, 1024)

	assert.Equal(t,
		"fbceda79c65c7ab6a245aae6608d19bce75037784b209feb4f3715bc984384faacd1702f340a1e5fb0fad015f5232f2204e3a56196f218b53462970c23fbc1279df20a751ecba3cc4fc89d985a110809cf99b91be2852403e6d4baccfa9a805859c7729c2251c5b6ac303afdde45bbcb505dc27a8a06923809916aa7a2449c5f",
		discriminant,
	)
}

func TestVerifyNWesolowski(t *testing.T) {
	seedBytes, err := hex.DecodeString(hexSeed)
	assert.NoError(t, err)

	proofBytes, err := hex.DecodeString(proofHex)
	assert.NoError(t, err)

	discriminant := vdf.CreateDiscriminant(seedBytes, 1024)
	initialEl := append([]byte{0x08}, make([]byte, 99)...)
	isValid := vdf.VerifyNWesolowski(discriminant, initialEl, proofBytes, 1<<20, 1024, 0)
	assert.True(t, isValid)
}

func TestProve(t *testing.T) {
	challengeBytes, err := hex.DecodeString(challengeHash)
	assert.NoError(t, err)

	initialEl := append([]byte{0x08}, make([]byte, 99)...)
	proof := vdf.Prove(challengeBytes, initialEl, 1024, 1<<20)
	generatedProofHex := hex.EncodeToString(proof)

	assert.Equal(t, proofHex, generatedProofHex)
}
