package types

// Addition an addition for a spend
// There isn't really good documentation for what is expected here, other than inspecting the transaction RPCs
// Amounts are known to be uint64
// PuzzleHash is known to be Bytes32
// Memos is a list of strings only in the RPCs, that is converted to bytes in chia. Can't send raw bytes
// Memos Ref: https://github.com/Chia-Network/chia-blockchain/blob/main/chia/rpc/wallet_rpc_api.py#L2358
type Addition struct {
	Amount     uint64   `json:"amount"`
	PuzzleHash Bytes32  `json:"puzzle_hash"`
	Memos      []string `json:"memos,omitempty"`
}
