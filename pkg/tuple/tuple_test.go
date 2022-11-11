package tuple_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/tuple"
)

// basically the same as sent_to in transactions
type testTypeForTuples struct {
	Peer                   string
	MempoolInclusionStatus uint8
	Error                  mo.Option[string]
}

var (
	expectedStruct = []tuple.Tuple[testTypeForTuples]{
		tuple.Some(testTypeForTuples{
			Peer:                   "9a6aac3959e9192c16156cd73f954791934d79804eed131f20666f228c6cf192",
			MempoolInclusionStatus: 3,
			Error:                  mo.Some("INVALID_FEE_TOO_CLOSE_TO_ZERO"),
		}),
		tuple.Some(testTypeForTuples{
			Peer:                   "9a6aac3959e9192c16156cd73f954791934d79804eed131f20666f228c6cf192",
			MempoolInclusionStatus: 1,
		}),
	}
	expectedJSON = `[["9a6aac3959e9192c16156cd73f954791934d79804eed131f20666f228c6cf192",3,"INVALID_FEE_TOO_CLOSE_TO_ZERO"],["9a6aac3959e9192c16156cd73f954791934d79804eed131f20666f228c6cf192",1,null]]`
)

func TestTuple_MarshalJSON(t *testing.T) {
	marshalled, err := json.Marshal(expectedStruct)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(marshalled))
}

func TestTuple_UnmarshalJSON(t *testing.T) {
	var unmarshalled []tuple.Tuple[testTypeForTuples]
	err := json.Unmarshal([]byte(expectedJSON), &unmarshalled)
	assert.NoError(t, err)
	assert.Equal(t, expectedStruct, unmarshalled)
}

func BenchmarkTuple_UnmarshalJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var unmarshalled []tuple.Tuple[testTypeForTuples]
		err := json.Unmarshal([]byte(expectedJSON), &unmarshalled)
		if err != nil {
			b.Error(err)
		}
	}
}

// Test struct, same as SentTo
// sent_to: List[Tuple[str, uint8, Optional[str]]]
type BenchStruct struct {
	Peer                   string
	MempoolInclusionStatus uint8
	Error                  mo.Option[string]
}

// UnmarshalJSON unmarshals the BenchStruct tuple into the struct
func (s *BenchStruct) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&s.Peer, &s.MempoolInclusionStatus, &s.Error}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in SentTo: %d != %d", g, e)
	}

	return nil
}

// This is a silly test, becuse its a test that only exists to ensure the benchmark test is actually doing the same work
// as the real tuple test. But without it, the benchmark is not known to be valid
func TestTuple_UnmarshalJSON_NoTuple(t *testing.T) {
	var unmarshalled []BenchStruct
	expectedStruct := []BenchStruct{
		BenchStruct{
			Peer:                   "9a6aac3959e9192c16156cd73f954791934d79804eed131f20666f228c6cf192",
			MempoolInclusionStatus: 3,
			Error:                  mo.Some("INVALID_FEE_TOO_CLOSE_TO_ZERO"),
		},
		BenchStruct{
			Peer:                   "9a6aac3959e9192c16156cd73f954791934d79804eed131f20666f228c6cf192",
			MempoolInclusionStatus: 1,
		},
	}
	err := json.Unmarshal([]byte(expectedJSON), &unmarshalled)
	assert.NoError(t, err)
	assert.Equal(t, expectedStruct, unmarshalled)
}

// BenchmarkTuple_UnmarshalJSON_NoTuple tests the old way (custom marshal/unmarshal) on a struct
// Vs the Tuple wrapper with all the reflection
func BenchmarkTuple_UnmarshalJSON_NoTuple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var unmarshalled []BenchStruct
		err := json.Unmarshal([]byte(expectedJSON), &unmarshalled)
		if err != nil {
			b.Error(err)
		}
	}
}
