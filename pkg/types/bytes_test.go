package types_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

func TestMarshalBytes(t *testing.T) {
	noPtr := types.Bytes{0, 1, 2, 3}
	ptr := &types.Bytes{0, 1, 2, 3}

	encodedNoPtr, err := json.Marshal(noPtr)
	assert.NoError(t, err)
	assert.Equal(t, `"0x00010203"`, string(encodedNoPtr))

	encodedPtr, err := json.Marshal(ptr)
	assert.NoError(t, err)
	assert.Equal(t, `"0x00010203"`, string(encodedPtr))

	type testStruct struct {
		Data types.Bytes `json:"data"`
	}

	data := &testStruct{Data: types.Bytes{0, 1, 2, 3}}
	encodedData, err := json.Marshal(data)
	assert.NoError(t, err)
	assert.Equal(t, `{"data":"0x00010203"}`, string(encodedData))
}

func TestMarshalBytes_Empty(t *testing.T) {
	noPtr := types.Bytes{}
	ptr := &types.Bytes{}

	encodedNoPtr, err := json.Marshal(noPtr)
	assert.NoError(t, err)
	assert.Equal(t, `null`, string(encodedNoPtr))

	encodedPtr, err := json.Marshal(ptr)
	assert.NoError(t, err)
	assert.Equal(t, `null`, string(encodedPtr))

	type testStruct struct {
		Data types.Bytes `json:"data"`
	}

	data := &testStruct{Data: types.Bytes{}}
	encodedData, err := json.Marshal(data)
	assert.NoError(t, err)
	assert.Equal(t, `{"data":null}`, string(encodedData))

	var nilType types.Bytes
	encodedNil, err := json.Marshal(nilType)
	assert.NoError(t, err)
	assert.Equal(t, `null`, string(encodedNil))
}

func TestUnmarshalBytes(t *testing.T) {
	test := []byte(`"0x00010203"`)
	dest := &types.Bytes{}
	err := json.Unmarshal(test, dest)
	assert.NoError(t, err)
	assert.Equal(t, &types.Bytes{0, 1, 2, 3}, dest)

	type testStruct struct {
		Data types.Bytes `json:"data"`
	}
	jsonInput := []byte(`{"data":"0x00010203"}`)
	expected := &testStruct{Data: types.Bytes{0, 1, 2, 3}}
	destStruct := &testStruct{}
	err = json.Unmarshal(jsonInput, destStruct)
	assert.NoError(t, err)
	assert.Equal(t, expected, destStruct)
}
