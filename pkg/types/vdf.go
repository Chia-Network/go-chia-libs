package types

// VDFInfo VDF Info
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/vdf.py#L49
// @TODO Streamable
type VDFInfo struct {
	Challenge          Bytes32           `json:"challenge"`
	NumberOfIterations uint64            `json:"number_of_iterations"`
	Output             ClassgroupElement `json:"output"`
}

// VDFProof VDF Proof
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/vdf.py#L57
// @TODO Streamable
type VDFProof struct {
	WitnessType          uint8 `json:"witness_type"`
	Witness              Bytes `json:"witness"`
	NormalizedToIdentity bool  `json:"normalized_to_identity"`
}
