package types

import (
	"encoding/json"
	"fmt"

	"github.com/samber/mo"
)

// TransactionRecord Single Transaction
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/wallet/transaction_record.py#L26
// @TODO Streamable
type TransactionRecord struct {
	ConfirmedAtHeight uint32                 `json:"confirmed_at_height"`
	CreatedAtTime     Timestamp              `json:"created_at_time"`
	ToPuzzleHash      Bytes32                `json:"to_puzzle_hash"`
	Amount            uint64                 `json:"amount"`
	FeeAmount         uint64                 `json:"fee_amount"`
	Confirmed         bool                   `json:"confirmed"`
	Sent              uint32                 `json:"sent"`
	SpendBundle       mo.Option[SpendBundle] `json:"spend_bundle"`
	Additions         []Coin                 `json:"additions"`
	Removals          []Coin                 `json:"removals"`
	WalletID          uint32                 `json:"wallet_id"`
	SentTo            []SentTo               `json:"sent_to"` // List[Tuple[str, uint8, Optional[str]]]
	TradeID           mo.Option[Bytes32]     `json:"trade_id"`
	Type              TransactionType        `json:"type"`
	Name              Bytes32                `json:"name"`
	//Memos             []MemoTuple            `json:"memos"` // List[Tuple[bytes32, List[bytes]]]
	// ToAddress is not on the official type, but some endpoints return it anyways. This part is not streamable
	ToAddress string `json:"to_address"`
}

// SentTo Represents the list of peers that we sent the transaction to, whether each one
// included it in the mempool, and what the error message (if any) was
// sent_to: List[Tuple[str, uint8, Optional[str]]]
type SentTo struct {
	Peer                   string
	MempoolInclusionStatus MempoolInclusionStatus
	Error                  mo.Option[string]
}

// UnmarshalJSON unmarshals the SentTo tuple into the struct
func (s *SentTo) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&s.Peer, &s.MempoolInclusionStatus, &s.Error}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in SentTo: %d != %d", g, e)
	}

	return nil
}

// MempoolInclusionStatus status of being included in the mempool
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/mempool_inclusion_status.py#L6
type MempoolInclusionStatus uint8

const (
	// MempoolInclusionStatusSuccess Successfully added to mempool
	MempoolInclusionStatusSuccess MempoolInclusionStatus = 1

	// MempoolInclusionStatusPending Pending being added to the mempool
	MempoolInclusionStatusPending MempoolInclusionStatus = 2

	// MempoolInclusionStatusFailed Failed being added to the mempool
	MempoolInclusionStatusFailed MempoolInclusionStatus = 3
)

// TransactionType type of transaction
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/wallet/util/transaction_type.py#L6
type TransactionType uint32

const (
	// TransactionTypeIncomingTX incoming transaction
	TransactionTypeIncomingTX TransactionType = 0

	// TransactionTypeOutgoingTX outgoing transaction
	TransactionTypeOutgoingTX TransactionType = 1

	// TransactionTypeCoinbaseReward coinbase reward
	TransactionTypeCoinbaseReward TransactionType = 2

	// TransactionTypeFeeReward fee reward
	TransactionTypeFeeReward TransactionType = 3

	// TransactionTypeIncomingTrade incoming trade
	TransactionTypeIncomingTrade TransactionType = 4

	// TransactionTypeOutgoingTrade outgoing trade
	TransactionTypeOutgoingTrade TransactionType = 5
)

// SpendBundle Spend Bundle
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/spend_bundle.py#L20
// @TODO Streamable
type SpendBundle struct {
	CoinSpends          []CoinSpend `json:"coin_spends"`
	AggregatedSignature G2Element   `json:"aggregated_signature"`
}
