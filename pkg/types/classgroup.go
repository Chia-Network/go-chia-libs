package types

// ClassgroupElement Classgroup Element
// https://github.com/Chia-Network/chia_rs/blob/main/crates/chia-protocol/src/classgroup.rs#L8
type ClassgroupElement struct {
	Data Bytes100 `json:"data" streamable:""`
}
