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

// MarshalYAML marshals the bytes32 to the appropriate format
func (g G1Element) MarshalYAML() (interface{}, error) {
	b48, err := BytesToBytes48(g[:])
	if err != nil {
		return nil, err
	}
	return b48.String(), nil
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

// UnmarshalYAML custom unmarshaller for G1element that can deal with the 0x prefixes and hex
func (g *G1Element) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}
	b48, err := Bytes48FromHexString(s)
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
