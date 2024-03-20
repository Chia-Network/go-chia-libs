package types

// VDFInfo VDF Info
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/vdf.py#L49
type VDFInfo struct {
	Challenge          Bytes32           `json:"challenge" streamable:""`
	NumberOfIterations uint64            `json:"number_of_iterations" streamable:""`
	Output             ClassgroupElement `json:"output" streamable:""`
}

// VDFProof VDF Proof
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/vdf.py#L57
type VDFProof struct {
	WitnessType          uint8 `json:"witness_type" streamable:""`
	Witness              Bytes `json:"witness" streamable:""`
	NormalizedToIdentity bool  `json:"normalized_to_identity" streamable:""`
}
