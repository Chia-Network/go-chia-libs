package types

import (
	"encoding/json"
)

// PublicKeyMPL is a public key
type PublicKeyMPL Bytes48

// G1Element is a public key
type G1Element PublicKeyMPL

// MarshalJSON custom hex marshaller
func (g G1Element) MarshalJSON() ([]byte, error) {
	return json.Marshal(Bytes48(g))
}

// UnmarshalJSON custom hex unmarshaller
func (g *G1Element) UnmarshalJSON(data []byte) error {
	b48 := Bytes48{}
	err := json.Unmarshal(data, &b48)
	if err != nil {
		return err
	}

	*g = G1Element(b48)

	return nil
}

// SignatureMPL is a signature
type SignatureMPL Bytes96

// G2Element is a signature
type G2Element SignatureMPL

// MarshalJSON custom hex marshaller
func (g G2Element) MarshalJSON() ([]byte, error) {
	return json.Marshal(Bytes96(g))
}

// UnmarshalJSON custom hex unmarshaller
func (g *G2Element) UnmarshalJSON(data []byte) error {
	b96 := Bytes96{}
	err := json.Unmarshal(data, &b96)
	if err != nil {
		return err
	}

	*g = G2Element(b96)

	return nil
}
