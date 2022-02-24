package types

// PuzzleHash Own type for future methods to encode/decode
type PuzzleHash string

// SerializedProgram Just represent as a string for now
type SerializedProgram string

// ClassgroupElement Classgroup Element
type ClassgroupElement struct {
	Data string `json:"data"`
}

// EndOfSubSlotBundle end of subslot bundle
type EndOfSubSlotBundle struct {
	// @TODO
}

// G1Element String for now, can make better later if we need
type G1Element string

// G2Element String for now, can make better later if we need
type G2Element string
