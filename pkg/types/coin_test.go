package types_test

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

var (
	parentCoin  types.Bytes32
	puzzleHash  types.Bytes32
	puzzleHash2 types.Bytes32
)

func init() {
	var err error

	parentCoin, err = types.BytesToBytes32([]byte("---foo---                       "))
	if err != nil { panic(err) }

	puzzleHash, err = types.BytesToBytes32([]byte("---bar---                       "))
	if err != nil { panic(err) }

	puzzleHash2, err = types.BytesToBytes32([]byte("---bar--- 2                     "))
	if err != nil { panic(err) }
}

// TestCoinID coin id tests adapted from https://github.com/Chia-Network/chia_rs/blob/main/tests/test_coin.py
// With additional static/expected hex values of the coin ID, just for my own sanity
// Static values computed using python like: Coin(parent_coin, puzzle_hash, 0xFF).name().hex()
func TestCoinID(t *testing.T) {
	assert.NotEqual(t, puzzleHash, puzzleHash2)

	c := types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0,
	}
	expected := sha256.Sum256(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...))
	expectedFromHex, err := types.BytesFromHexString("9dbfc291966738e809a8b55999027d0b4f049ca1e5f18601e60de3ff9bc087a4")
	result := types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         1,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{1}...))
	expectedFromHex, err = types.BytesFromHexString("3966fc259b7cb0b1a3e97da4dabc4ddbf77b4b37dcbbe72457cbcae95c6453ae")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	// 0xFF prefix
	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0xff,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0xff}...))
	expectedFromHex, err = types.BytesFromHexString("c7bdcad6f4354f346acc076cd90fa63458ca9f2f9408eca99fa6f077c26af1e9")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0xffff,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0xff, 0xff}...))
	expectedFromHex, err = types.BytesFromHexString("95afe6f9e7031614a99d404041cbf2b61df4cf5bb70b1dea63874f9fa14b4d95")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0xffffff,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0xff, 0xff, 0xff}...))
	expectedFromHex, err = types.BytesFromHexString("05a3787b942aca42ec73b39413063dc278ad99d7a24b245046ceea75c1a815f2")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0xffffffff,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0xff, 0xff, 0xff, 0xff}...))
	expectedFromHex, err = types.BytesFromHexString("f2c66f9433203d8f01744e84ddf9d5681e37c492c8bc5a237079d6b2ee6b563c")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0xffffffffff,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0xff, 0xff, 0xff, 0xff, 0xff}...))
	expectedFromHex, err = types.BytesFromHexString("a457d428fcc7d094dcb2d586ca2783bf088beaab8ef65bed2e1dc3f9c899dc8d")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0xffffffffffff,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}...))
	expectedFromHex, err = types.BytesFromHexString("62df1b5a9897314213c07084f040a12990d701d3a6b293d239fca8fd114d54ae")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0xffffffffffffff,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}...))
	expectedFromHex, err = types.BytesFromHexString("e841c3b82faeee77336c6e3a409719df050484b39097efd40c6d127a3bbaf932")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0xffffffffffffffff,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}...))
	expectedFromHex, err = types.BytesFromHexString("132b0db8d339a4f41c33d6b44c5699ca39d37ad8e2e80f8965bbc4c3b69e1f07")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	// 0x7F prefix
	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x7F,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0x7F}...))
	expectedFromHex, err = types.BytesFromHexString("b629a6f158eda42767ea353b577e31b3967c954d8aea68d1bf442db8f5576d4c")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x7FFF,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0x7F, 0xFF}...))
	expectedFromHex, err = types.BytesFromHexString("ed885979a51b59be836367cd56fd44e1b27ebdb1bbc0a7f11dfc1640be2b8708")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x7FFFFF,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0x7F, 0xFF, 0xFF}...))
	expectedFromHex, err = types.BytesFromHexString("d4a793f2f5631d312d42a51a29c8dad898342d581197534909f10c4daac8aaf1")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x7FFFFFFF,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0x7F, 0xFF, 0xFF, 0xFF}...))
	expectedFromHex, err = types.BytesFromHexString("95b683f86abcb62c8244b3397352e68c2fabf21d0ca0021bd6b636c403467235")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x7FFFFFFFFF,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF}...))
	expectedFromHex, err = types.BytesFromHexString("07928ee745f1dc9284e9c7df68f419df567ada0b93b0c32c819019df162bdb5b")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x7FFFFFFFFFFF,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}...))
	expectedFromHex, err = types.BytesFromHexString("82cb850d70d8bca77335058c742220a2d8d0d5bc11f32c4aad4554b52b19644c")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x7FFFFFFFFFFFFF,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}...))
	expectedFromHex, err = types.BytesFromHexString("aaa0fcbea9f2302db2a46041ce1857f2cfb0b990ea06abcc7646cc6b138531bd")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x7FFFFFFFFFFFFFFF,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}...))
	expectedFromHex, err = types.BytesFromHexString("85c2735fc17d163422ac85a63f54a2ea08736e4dc2e0c1d3b0bd33a6bed7f584")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	// 0x80 prefix
	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x80,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0x80}...))
	expectedFromHex, err = types.BytesFromHexString("19fbcae7129e4ad616327fbbcb52aecf165d1a09d7bc966251ca98fec3f07930")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x8000,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0x80, 0}...))
	expectedFromHex, err = types.BytesFromHexString("9abe734158bdfd0a1169b3ced41ffacd2c2d66a2596473ebdef1ad743c90d804")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x800000,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0x80, 0, 0}...))
	expectedFromHex, err = types.BytesFromHexString("f4126e16bbad038b04b054fd1dafaca75613777db13a50706fd13385f28ebe1d")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x80000000,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0x80, 0, 0, 0}...))
	expectedFromHex, err = types.BytesFromHexString("8fb916ebb70c5133e3eff2f33097ef8cc566258d63c8b70202776949cb14b101")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x8000000000,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0x80, 0, 0, 0, 0}...))
	expectedFromHex, err = types.BytesFromHexString("b1915ae092fd313f3ccb34048b5b84539004d46c87722e622dbf43b6be012af0")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x800000000000,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0x80, 0, 0, 0, 0, 0}...))
	expectedFromHex, err = types.BytesFromHexString("be0cca36039a26b3e5e993fe26a30f3b2e6620aa80a2c8eb24e29971e532aad1")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x80000000000000,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0x80, 0, 0, 0, 0, 0, 0}...))
	expectedFromHex, err = types.BytesFromHexString("0183aaec6ed4c0a07fce637443e4439c5258da10501b6c84e69f894e16d11d41")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)

	c = types.Coin{
		ParentCoinInfo: parentCoin,
		PuzzleHash:     puzzleHash,
		Amount:         0x8000000000000000,
	}
	expected = sha256.Sum256(append(append(types.Bytes32ToBytes(parentCoin), types.Bytes32ToBytes(puzzleHash)...), []byte{0, 0x80, 0, 0, 0, 0, 0, 0, 0}...))
	expectedFromHex, err = types.BytesFromHexString("dfd2839427f975d1ed55bbad67a9389d0d852730bc203926dc478c7a7953fd0d")
	result = types.Bytes32ToBytes(c.ID())
	assert.NoError(t, err)
	assert.Equal(t, types.Bytes32ToBytes(expected), result)
	assert.Equal(t, expectedFromHex, result)
}
