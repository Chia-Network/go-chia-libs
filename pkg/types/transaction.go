package types

import (
	"encoding/json"

	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/tuple"
)

// TransactionRecord Single Transaction
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/wallet/transaction_record.py#L26
// @TODO Streamable
type TransactionRecord struct {
	ConfirmedAtHeight uint32                   `json:"confirmed_at_height"`
	CreatedAtTime     Timestamp                `json:"created_at_time"`
	ToPuzzleHash      Bytes32                  `json:"to_puzzle_hash"`
	Amount            uint64                   `json:"amount"`
	FeeAmount         uint64                   `json:"fee_amount"`
	Confirmed         bool                     `json:"confirmed"`
	Sent              uint32                   `json:"sent"`
	SpendBundle       mo.Option[SpendBundle]   `json:"spend_bundle"`
	Additions         []Coin                   `json:"additions"`
	Removals          []Coin                   `json:"removals"`
	WalletID          uint32                   `json:"wallet_id"`
	SentTo            []tuple.Tuple[SentTo]    `json:"sent_to"` // List[Tuple[str, uint8, Optional[str]]]
	TradeID           mo.Option[Bytes32]       `json:"trade_id"`
	Type              TransactionType          `json:"type"`
	Name              Bytes32                  `json:"name"`
	Memos             []tuple.Tuple[MemoTuple] `json:"-"`     // List[Tuple[bytes32, List[bytes]]]
	MemosDict         map[string]string        `json:"memos"` // The tuple above is translated to a dict{ coin_id: memo, coin_id: memo } before going into the response
	// ToAddress is not on the official type, but some endpoints return it anyways. This part is not streamable
	ToAddress string `json:"to_address"`
}

// MarshalJSON Handles the weird juggling between the tuple and map[string]string that goes on with memos on RPC
func (t TransactionRecord) MarshalJSON() ([]byte, error) {
	type tr TransactionRecord

	t.MemosDict = map[string]string{}
	for _, memo := range t.Memos {
		t.MemosDict[memo.Value().CoinID.String()] = memo.Value().Memo[0].String()
	}

	return json.Marshal(tr(t))
}

// UnmarshalJSON Handles the weird juggling between the tuple and map[string]string that goes on with memos on RPC
func (t *TransactionRecord) UnmarshalJSON(data []byte) error {
	type tr TransactionRecord
	err := json.Unmarshal(data, (*tr)(t))
	if err != nil {
		return err
	}

	// Move memos back to the expected Tuple form
	for coinID, memo := range t.MemosDict {
		coinIDBytes, err := BytesFromHexString(coinID)
		if err != nil {
			return err
		}
		cidb32, err := BytesToBytes32(coinIDBytes)
		if err != nil {
			return err
		}
		memoBytes, err := BytesFromHexString(memo)
		if err != nil {
			return err
		}
		t.Memos = append(t.Memos, tuple.Some(MemoTuple{
			CoinID: cidb32,
			Memo:   []Bytes{memoBytes},
		}))
	}

	// Don't use the dict directly
	t.MemosDict = nil

	return nil
}

// SentTo Represents the list of peers that we sent the transaction to, whether each one
// included it in the mempool, and what the error message (if any) was
// sent_to: List[Tuple[str, uint8, Optional[str]]]
type SentTo struct {
	Peer                   string
	MempoolInclusionStatus MempoolInclusionStatus
	Error                  mo.Option[string]
}

// MemoTuple corresponds to the fields in the memo tuple for TransactionRecord
type MemoTuple struct {
	CoinID Bytes32
	Memo   []Bytes
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
