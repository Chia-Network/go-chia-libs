package rpc

import (
	"net/http"

	"github.com/samber/mo"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

// FullNodeService encapsulates full node RPC methods
type FullNodeService struct {
	client *Client
}

// NewRequest returns a new request specific to the wallet service
func (s *FullNodeService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceFullNode, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *FullNodeService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetConnectionsOptions options to filter get_connections
type GetConnectionsOptions struct {
	NodeType types.NodeType `json:"node_type,omitempty"`
}

// GetConnectionsResponse get_connections response format
type GetConnectionsResponse struct {
	Response
	Connections mo.Option[[]types.Connection] `json:"connections"`
}

// GetConnections returns connections
func (s *FullNodeService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_connections", opts)
	if err != nil {
		return nil, nil, err
	}

	c := &GetConnectionsResponse{}
	resp, err := s.Do(request, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// GetBlockchainStateResponse is the blockchain state RPC response
type GetBlockchainStateResponse struct {
	Response
	BlockchainState mo.Option[types.BlockchainState] `json:"blockchain_state,omitempty"`
}

// GetBlockchainState returns blockchain state
func (s *FullNodeService) GetBlockchainState() (*GetBlockchainStateResponse, *http.Response, error) {
	request, err := s.NewRequest("get_blockchain_state", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetBlockchainStateResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlockOptions options for get_block rpc call
type GetBlockOptions struct {
	HeaderHash types.Bytes32 `json:"header_hash"`
}

// GetBlockResponse response for get_block rpc call
type GetBlockResponse struct {
	Response
	Block mo.Option[types.FullBlock] `json:"block"`
}

// GetBlock full_node->get_block RPC method
func (s *FullNodeService) GetBlock(opts *GetBlockOptions) (*GetBlockResponse, *http.Response, error) {
	request, err := s.NewRequest("get_block", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetBlockResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlocksOptions options for get_blocks rpc call
type GetBlocksOptions struct {
	Start int `json:"start"`
	End   int `json:"end"`
	// ExcludeHeaderHash if set to true, omits the `header_hash` key from the response
	ExcludeHeaderHash bool `json:"exclude_header_hash"`
	// ExcludeReorged if set to true, excludes reorged blocks from the response
	ExcludeReorged bool `json:"exclude_reorged"`
}

// GetBlocksResponse response for get_blocks rpc call
type GetBlocksResponse struct {
	Response
	Blocks mo.Option[[]types.FullBlock] `json:"blocks"`
}

// GetBlocks full_node->get_blocks RPC method
func (s *FullNodeService) GetBlocks(opts *GetBlocksOptions) (*GetBlocksResponse, *http.Response, error) {
	request, err := s.NewRequest("get_blocks", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetBlocksResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlockCountMetricsResponse response for get_block_count_metrics rpc call
type GetBlockCountMetricsResponse struct {
	Response
	Metrics mo.Option[types.BlockCountMetrics] `json:"metrics"`
}

// GetBlockCountMetrics gets metrics about blocks
func (s *FullNodeService) GetBlockCountMetrics() (*GetBlockCountMetricsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_block_count_metrics", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetBlockCountMetricsResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlockByHeightOptions options for get_block_record_by_height and get_block rpc call
type GetBlockByHeightOptions struct {
	BlockHeight int `json:"height"`
}

// GetBlockRecordResponse response from get_block_record_by_height
type GetBlockRecordResponse struct {
	Response
	BlockRecord mo.Option[types.BlockRecord] `json:"block_record"`
}

// GetBlockRecordByHeight full_node->get_block_record_by_height RPC method
func (s *FullNodeService) GetBlockRecordByHeight(opts *GetBlockByHeightOptions) (*GetBlockRecordResponse, *http.Response, error) {
	// Get Block Record
	request, err := s.NewRequest("get_block_record_by_height", opts)
	if err != nil {
		return nil, nil, err
	}

	record := &GetBlockRecordResponse{}
	resp, err := s.Do(request, record)
	if err != nil {
		return nil, resp, err
	}

	// I believe this happens when the node is not yet synced to this height
	if record == nil || record.BlockRecord.IsAbsent() {
		return nil, nil, nil
	}

	return record, resp, nil
}

// GetBlockByHeight helper function to get a full block by height, calls full_node->get_block_record_by_height RPC method then full_node->get_block RPC method
func (s *FullNodeService) GetBlockByHeight(opts *GetBlockByHeightOptions) (*GetBlockResponse, *http.Response, error) {
	// Get Block Record
	record, resp, err := s.GetBlockRecordByHeight(opts)
	if err != nil {
		return nil, resp, err
	}

	request, err := s.NewRequest("get_block", GetBlockOptions{
		HeaderHash: record.BlockRecord.OrEmpty().HeaderHash,
	})
	if err != nil {
		return nil, nil, err
	}

	// Get Full Block
	block := &GetBlockResponse{}
	resp, err = s.Do(request, block)
	if err != nil {
		return nil, resp, err
	}

	return block, resp, nil
}

// GetAdditionsAndRemovalsOptions options for get_additions_and_removals
type GetAdditionsAndRemovalsOptions struct {
	HeaderHash types.Bytes32 `json:"header_hash"`
}

// GetAdditionsAndRemovalsResponse response for get_additions_and_removals
type GetAdditionsAndRemovalsResponse struct {
	Response
	Additions []types.CoinRecord `json:"additions"`
	Removals  []types.CoinRecord `json:"removals"`
}

// GetAdditionsAndRemovals Gets additions and removals for a particular block hash
func (s *FullNodeService) GetAdditionsAndRemovals(opts *GetAdditionsAndRemovalsOptions) (*GetAdditionsAndRemovalsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_additions_and_removals", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetAdditionsAndRemovalsResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetCoinRecordsByPuzzleHashOptions request options for /get_coin_records_by_puzzle_hash
type GetCoinRecordsByPuzzleHashOptions struct {
	PuzzleHash        types.Bytes32 `json:"puzzle_hash"`
	IncludeSpentCoins bool          `json:"include_spent_coins"`
	StartHeight       uint32        `json:"start_height,omitempty"`
	EndHeight         uint32        `json:"end_height,omitempty"`
}

// GetCoinRecordsByPuzzleHashResponse Response for /get_coin_records_by_puzzle_hash
type GetCoinRecordsByPuzzleHashResponse struct {
	Response
	CoinRecords []types.CoinRecord `json:"coin_records"`
}

// GetCoinRecordsByPuzzleHash returns coin records for a specified puzzle hash
func (s *FullNodeService) GetCoinRecordsByPuzzleHash(opts *GetCoinRecordsByPuzzleHashOptions) (*GetCoinRecordsByPuzzleHashResponse, *http.Response, error) {
	request, err := s.NewRequest("get_coin_records_by_puzzle_hash", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetCoinRecordsByPuzzleHashResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetCoinRecordsByPuzzleHashesOptions request options for /get_coin_records_by_puzzle_hash
type GetCoinRecordsByPuzzleHashesOptions struct {
	PuzzleHash        []types.Bytes32 `json:"puzzle_hashes"`
	IncludeSpentCoins bool            `json:"include_spent_coins"`
	StartHeight       uint32          `json:"start_height,omitempty"`
	EndHeight         uint32          `json:"end_height,omitempty"`
}

// GetCoinRecordsByPuzzleHashesResponse Response for /get_coin_records_by_puzzle_hashes
type GetCoinRecordsByPuzzleHashesResponse struct {
	Response
	CoinRecords []types.CoinRecord `json:"coin_records"`
}

// GetCoinRecordsByPuzzleHashes returns coin records for a specified list of puzzle hashes
func (s *FullNodeService) GetCoinRecordsByPuzzleHashes(opts *GetCoinRecordsByPuzzleHashesOptions) (*GetCoinRecordsByPuzzleHashesResponse, *http.Response, error) {
	request, err := s.NewRequest("get_coin_records_by_puzzle_hashes", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetCoinRecordsByPuzzleHashesResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetCoinRecordByNameOptions request options for /get_coin_record_by_name
type GetCoinRecordByNameOptions struct {
	Name string `json:"name"`
}

// GetCoinRecordByNameResponse response from get_coin_record_by_name endpoint
type GetCoinRecordByNameResponse struct {
	Response
	CoinRecord mo.Option[types.CoinRecord] `json:"coin_record"`
}

// GetCoinRecordByName request to get_coin_record_by_name endpoint
func (s *FullNodeService) GetCoinRecordByName(opts *GetCoinRecordByNameOptions) (*GetCoinRecordByNameResponse, *http.Response, error) {
	request, err := s.NewRequest("get_coin_record_by_name", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetCoinRecordByNameResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// FullNodePushTXOptions options for pushing tx to full node mempool
type FullNodePushTXOptions struct {
	SpendBundle types.SpendBundle `json:"spend_bundle"`
}

// FullNodePushTXResponse Response from full node push_tx
type FullNodePushTXResponse struct {
	Response
	Status mo.Option[string] `json:"status"`
}

// PushTX pushes a transaction to the full node
func (s *FullNodeService) PushTX(opts *FullNodePushTXOptions) (*FullNodePushTXResponse, *http.Response, error) {
	request, err := s.NewRequest("push_tx", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &FullNodePushTXResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetPuzzleAndSolution full_node-> get_puzzle_and_solution RPC method
func (s *FullNodeService) GetPuzzleAndSolution(opts *GetPuzzleAndSolutionOptions) (*GetPuzzleAndSolutionResponse, *http.Response, error) {
	request, err := s.NewRequest("get_puzzle_and_solution", opts)
	if err != nil {
		return nil, nil, err
	}

	record := &GetPuzzleAndSolutionResponse{}
	resp, err := s.Do(request, record)
	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetFeeEstimateOptions inputs to get a fee estimate
// TargetTimes is a list of values corresponding to "seconds from now" to get a fee estimate for
// The estimated fee is the estimate of the fee required to complete the TX by the target time seconds
type GetFeeEstimateOptions struct {
	SpendBundle *types.SpendBundle `json:"spend_bundle,omitempty"`
	Cost        uint64             `json:"cost,omitempty"`
	TargetTimes []uint64           `json:"target_times"`
}

// GetFeeEstimateResponse response for get_fee_estimate
type GetFeeEstimateResponse struct {
	Response
	Estimates         mo.Option[[]uint64] `json:"estimates"`
	TargetTimes       mo.Option[[]uint64] `json:"target_times"`
	CurrentFeeRate    mo.Option[uint64]   `json:"current_fee_rate"`
	MempoolSize       mo.Option[uint64]   `json:"mempool_size"`
	MempoolMaxSize    mo.Option[uint64]   `json:"mempool_max_size"`
	FullNodeSynced    mo.Option[bool]     `json:"full_node_synced"`
	PeakHeight        mo.Option[uint32]   `json:"peak_height"`
	LastPeakTimestamp mo.Option[uint64]   `json:"last_peak_timestamp"`
	NodeTimeUTC       mo.Option[uint64]   `json:"node_time_utc"`
}

// GetFeeEstimate endpoint
func (s *FullNodeService) GetFeeEstimate(opts *GetFeeEstimateOptions) (*GetFeeEstimateResponse, *http.Response, error) {
	request, err := s.NewRequest("get_fee_estimate", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetFeeEstimateResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetPuzzleAndSolutionOptions options for get_puzzle_and_solution rpc call
type GetPuzzleAndSolutionOptions struct {
	CoinID types.Bytes32 `json:"coin_id"`
	Height uint32        `json:"height"`
}

// GetPuzzleAndSolutionResponse response from get_puzzle_and_solution
type GetPuzzleAndSolutionResponse struct {
	Response
	CoinSolution mo.Option[types.CoinSpend] `json:"coin_solution"`
}
