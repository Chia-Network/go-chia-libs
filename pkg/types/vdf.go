package types

// VDFInfo VDF Info
// https://github.com/Chia-Network/chia_rs/blob/main/crates/chia-protocol/src/vdf.rs#L7
type VDFInfo struct {
	Challenge          Bytes32           `json:"challenge" streamable:""`
	NumberOfIterations uint64            `json:"number_of_iterations" streamable:""`
	Output             ClassgroupElement `json:"output" streamable:""`
}

// VDFProof VDF Proof
// https://github.com/Chia-Network/chia_rs/blob/main/crates/chia-protocol/src/vdf.rs#L14
type VDFProof struct {
	WitnessType          uint8 `json:"witness_type" streamable:""`
	Witness              Bytes `json:"witness" streamable:""`
	NormalizedToIdentity bool  `json:"normalized_to_identity" streamable:""`
}
