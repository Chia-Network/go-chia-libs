package types

// ClassgroupElement Classgroup Element
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/classgroup.py#L12
type ClassgroupElement struct {
	Data Bytes100 `json:"data" streamable:""`
}
