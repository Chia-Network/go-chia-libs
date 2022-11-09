package types

// VDFInfo VDF Info
type VDFInfo struct {
	Challenge          Bytes32            `json:"challenge"`
	NumberOfIterations uint64             `json:"number_of_iterations"`
	Output             *ClassgroupElement `json:"output"`
}

// VDFProof VDF Proof
type VDFProof struct {
	WitnessType          uint8 `json:"witness_type"`
	Witness              Bytes `json:"witness"`
	NormalizedToIdentity bool  `json:"normalized_to_identity"`
}
