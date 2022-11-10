package tuple_test

import (
	"encoding/json"
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
