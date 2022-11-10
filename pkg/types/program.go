package types

import (
	"encoding/json"
)

// SerializedProgram An opaque representation of a clvm program. It has a more limited interface than a full SExp
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/program.py#L232
type SerializedProgram Bytes

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
