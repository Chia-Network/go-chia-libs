package types

// SerializedProgram An opaque representation of a clvm program. It has a more limited interface than a full SExp
type SerializedProgram Bytes

// ClassgroupElement Classgroup Element
type ClassgroupElement struct {
	Data Bytes100 `json:"data"`
}

// EndOfSubSlotBundle end of subslot bundle
type EndOfSubSlotBundle struct {
	// @TODO
}

// PublicKeyMPL is a public key
type PublicKeyMPL Bytes48

// G1Element is a public key
type G1Element PublicKeyMPL

// SignatureMPL is a signature
type SignatureMPL Bytes96

// G2Element is a signature
type G2Element SignatureMPL
