package types

// TransactionRecord Single Transaction
type TransactionRecord struct {
	ConfirmedAtHeight uint32       `json:"confirmed_at_height"`
	CreatedAtTime     uint64       `json:"created_at_time"` // @TODO time.Time?
	ToPuzzleHash      *PuzzleHash  `json:"to_puzzle_hash"`
	Amount            uint64       `json:"amount"`
	FeeAmount         uint64       `json:"fee_amount"`
	Confirmed         bool         `json:"confirmed"`
	Sent              uint32       `json:"sent"`
	SpendBundle       *SpendBundle `json:"spend_bundle"`
	Additions         []*Coin      `json:"additions"`
	Removals          []*Coin      `json:"removals"`
	WalletID          uint32       `json:"wallet_id"`
	//SentTo            SentTo          `json:"sent_to"` // @TODO need to properly unserialize this
	TradeID string           `json:"trade_id"`
	Type    *TransactionType `json:"type"`
	Name    string           `json:"name"` // @TODO bytes32 / hex
	// ToAddress is not on the official type, but some endpoints return it anyways
	ToAddress *Address `json:"to_address"`
}

// Address Own type for future methods to encode/decode
type Address string

// SentTo Represents the list of peers that we sent the transaction to, whether each one
// included it in the mempool, and what the error message (if any) was
// sent_to: List[Tuple[str, uint8, Optional[str]]]
// @TODO need to parse from the json
type SentTo struct {
	Peer                   string
	MempoolInclusionStatus *MempoolInclusionStatus
	Error                  string
}

// MempoolInclusionStatus status of being included in the mempool
type MempoolInclusionStatus uint8

const (
	// MempoolInclusionStatusSuccess Successfully added to mempool
	MempoolInclusionStatusSuccess = MempoolInclusionStatus(1)

	// MempoolInclusionStatusPending Pending being added to the mempool
	MempoolInclusionStatusPending = MempoolInclusionStatus(2)

	// MempoolInclusionStatusFailed Failed being added to the mempool
	MempoolInclusionStatusFailed = MempoolInclusionStatus(3)
)

// TransactionType type of transaction
type TransactionType uint32

const (
	// TransactionTypeIncomingTX incoming transaction
	TransactionTypeIncomingTX = TransactionType(0)

	// TransactionTypeOutgoingTX outgoing transaction
	TransactionTypeOutgoingTX = TransactionType(1)

	// TransactionTypeCoinbaseReward coinbase reward
	TransactionTypeCoinbaseReward = TransactionType(2)

	// TransactionTypeFeeReward fee reward
	TransactionTypeFeeReward = TransactionType(3)

	// TransactionTypeIncomingTrade incoming trade
	TransactionTypeIncomingTrade = TransactionType(4)

	// TransactionTypeOutgoingTrade outgoing trade
	TransactionTypeOutgoingTrade = TransactionType(5)
)

// SpendBundle Spend Bundle...
type SpendBundle struct {
	AggregatedSignature string          `json:"aggregated_signature"`
	CoinSolutions       []*CoinSolution `json:"coin_solutions"`
}
