package config

import (
	"github.com/chia-network/go-chia-libs/pkg/types"
)

// ChiaConfig the chia config.yaml
type ChiaConfig struct {
	ChiaRoot                 string                 `yaml:"-"`
	MinMainnetKSize          uint8                  `yaml:"min_mainnet_k_size"`
	PingInterval             uint16                 `yaml:"ping_interval"`
	SelfHostname             string                 `yaml:"self_hostname"`
	PreferIPv6               bool                   `yaml:"prefer_ipv6"`
	RPCTimeout               uint16                 `yaml:"rpc_timeout"`
	DaemonPort               uint16                 `yaml:"daemon_port"`
	DaemonMaxMessageSize     uint32                 `yaml:"daemon_max_message_size"`
	DaemonHeartbeat          uint16                 `yaml:"daemon_heartbeat"`
	DaemonAllowTLS12         bool                   `yaml:"daemon_allow_tls_1_2"`
	InboundRateLimitPercent  uint8                  `yaml:"inbound_rate_limit_percent"`
	OutboundRateLimitPercent uint8                  `yaml:"outbound_rate_limit_percent"`
	NetworkOverrides         *NetworkOverrides      `yaml:"network_overrides"`
	SelectedNetwork          *string                `yaml:"selected_network"`
	AlertsURL                string                 `yaml:"ALERTS_URL"`
	ChiaAlertsPubkey         string                 `yaml:"CHIA_ALERTS_PUBKEY"`
	PrivateSSLCA             CAConfig               `yaml:"private_ssl_ca"`
	ChiaSSLCA                CAConfig               `yaml:"chia_ssl_ca"`
	DaemonSSL                SSLConfig              `yaml:"daemon_ssl"`
	Logging                  LoggingConfig          `yaml:"logging"` // @TODO this would usually be an anchor
	Seeder                   SeederConfig           `yaml:"seeder"`
	Harvester                HarvesterConfig        `yaml:"harvester"`
	Pool                     PoolConfig             `yaml:"pool"`
	Farmer                   FarmerConfig           `yaml:"farmer"`
	TimelordLauncher         TimelordLauncherConfig `yaml:"timelord_launcher"`
	Timelord                 TimelordConfig         `yaml:"timelord"`
	FullNode                 FullNodeConfig         `yaml:"full_node"`
	UI                       UIConfig               `yaml:"ui"`
	Introducer               IntroducerConfig       `yaml:"introducer"`
	Wallet                   WalletConfig           `yaml:"wallet"`
	DataLayer                DataLayerConfig        `yaml:"data_layer"`
	Simulator                SimulatorConfig        `yaml:"simulator"`
}

// PortConfig common port settings found in many sections of the config
type PortConfig struct {
	Port    uint16 `yaml:"port,omitempty"`
	RPCPort uint16 `yaml:"rpc_port"`
}

// CAConfig config keys for CA
type CAConfig struct {
	Crt string `yaml:"crt"`
	Key string `yaml:"key"`
}

// SSLConfig common ssl settings found in many sections of the config
type SSLConfig struct {
	PrivateCRT string `yaml:"private_crt"`
	PrivateKey string `yaml:"private_key"`
	PublicCRT  string `yaml:"public_crt"`
	PublicKey  string `yaml:"public_key"`
}

// Peer is a host/port for a peer
type Peer struct {
	Host                  string `yaml:"host"`
	Port                  uint16 `yaml:"port"`
	EnablePrivateNetworks bool   `yaml:"enable_private_networks,omitempty"`
}

// NetworkOverrides is all network settings
type NetworkOverrides struct {
	Constants map[string]NetworkConstants `yaml:"constants"`
	Config    map[string]NetworkConfig    `yaml:"config"`
}

// NetworkConstants the constants for each network
type NetworkConstants struct {
	AggSigMeAdditionalData         string        `yaml:"AGG_SIG_ME_ADDITIONAL_DATA,omitempty"`
	DifficultyConstantFactor       types.Uint128 `yaml:"DIFFICULTY_CONSTANT_FACTOR,omitempty"`
	DifficultyStarting             uint64        `yaml:"DIFFICULTY_STARTING,omitempty"`
	EpochBlocks                    uint32        `yaml:"EPOCH_BLOCKS,omitempty"`
	GenesisChallenge               string        `yaml:"GENESIS_CHALLENGE"`
	GenesisPreFarmPoolPuzzleHash   string        `yaml:"GENESIS_PRE_FARM_POOL_PUZZLE_HASH"`
	GenesisPreFarmFarmerPuzzleHash string        `yaml:"GENESIS_PRE_FARM_FARMER_PUZZLE_HASH"`
	MempoolBlockBuffer             uint8         `yaml:"MEMPOOL_BLOCK_BUFFER,omitempty"`
	MinPlotSize                    uint8         `yaml:"MIN_PLOT_SIZE,omitempty"`
	NetworkType                    uint8         `yaml:"NETWORK_TYPE,omitempty"`
	SubSlotItersStarting           uint64        `yaml:"SUB_SLOT_ITERS_STARTING,omitempty"`
	HardForkHeight                 uint32        `yaml:"HARD_FORK_HEIGHT,omitempty"`
	SoftFork4Height                uint32        `yaml:"SOFT_FORK4_HEIGHT,omitempty"`
	SoftFork5Height                uint32        `yaml:"SOFT_FORK5_HEIGHT,omitempty"`
	PlotFilter128Height            uint32        `yaml:"PLOT_FILTER_128_HEIGHT,omitempty"`
	PlotFilter64Height             uint32        `yaml:"PLOT_FILTER_64_HEIGHT,omitempty"`
	PlotFilter32Height             uint32        `yaml:"PLOT_FILTER_32_HEIGHT,omitempty"`
}

// NetworkConfig specific network configuration settings
type NetworkConfig struct {
	AddressPrefix       string `yaml:"address_prefix"`
	DefaultFullNodePort uint16 `yaml:"default_full_node_port"`
}

// LoggingConfig configuration settings for the logger
type LoggingConfig struct {
	LogStdout           bool   `yaml:"log_stdout"`
	LogFilename         string `yaml:"log_filename"`
	LogLevel            string `yaml:"log_level"`
	LogMaxFilesRotation uint8  `yaml:"log_maxfilesrotation"`
	LogMaxBytesRotation uint32 `yaml:"log_maxbytesrotation"`
	LogUseGzip          bool   `yaml:"log_use_gzip"`
	LogSyslog           bool   `yaml:"log_syslog"`
	LogSyslogHost       string `yaml:"log_syslog_host"`
	LogSyslogPort       uint16 `yaml:"log_syslog_port"`
}

// SeederConfig seeder configuration section
type SeederConfig struct {
	Port                uint16            `yaml:"port"`
	OtherPeersPort      uint16            `yaml:"other_peers_port"`
	DNSPort             uint16            `yaml:"dns_port"`
	PeerConnectTimeout  uint16            `yaml:"peer_connect_timeout"`
	CrawlerDBPath       string            `yaml:"crawler_db_path"`
	BootstrapPeers      []string          `yaml:"bootstrap_peers"`
	MinimumHeight       uint32            `yaml:"minimum_height"`
	MinimumVersionCount uint32            `yaml:"minimum_version_count"`
	DomainName          string            `yaml:"domain_name"`
	Nameserver          string            `yaml:"nameserver"`
	TTL                 uint16            `yaml:"ttl"`
	SOA                 SeederSOA         `yaml:"soa"`
	NetworkOverrides    *NetworkOverrides `yaml:"network_overrides"`
	SelectedNetwork     *string           `yaml:"selected_network"`
	Logging             LoggingConfig     `yaml:"logging"`
	CrawlerConfig       CrawlerConfig     `yaml:"crawler"`
}

// SeederSOA dns SOA for seeder
type SeederSOA struct {
	Rname        string `yaml:"rname"`
	SerialNumber string `yaml:"serial_number"`
	Refresh      uint32 `yaml:"refresh"`
	Retry        uint32 `yaml:"retry"`
	Expire       uint32 `yaml:"expire"`
	Minimum      uint32 `yaml:"minimum"`
}

// CrawlerConfig is the subsection of the seeder config specific to the crawler
type CrawlerConfig struct {
	StartRPCServer bool `yaml:"start_rpc_server"`
	PortConfig     `yaml:",inline"`
	SSL            SSLConfig `yaml:"ssl"`
}

// HarvesterConfig harvester configuration section
type HarvesterConfig struct {
	FarmerPeers                []Peer                `yaml:"farmer_peers"`
	StartRPCServer             bool                  `yaml:"start_rpc_server"`
	NumThreads                 uint8                 `yaml:"num_threads"`
	PlotsRefreshParameter      PlotsRefreshParameter `yaml:"plots_refresh_parameter"`
	ParallelRead               bool                  `yaml:"parallel_read"`
	Logging                    LoggingConfig         `yaml:"logging"`
	NetworkOverrides           *NetworkOverrides     `yaml:"network_overrides"`
	SelectedNetwork            *string               `yaml:"selected_network"`
	PlotDirectories            []string              `yaml:"plot_directories"`
	RecursivePlotScan          bool                  `yaml:"recursive_plot_scan"`
	PortConfig                 `yaml:",inline"`
	SSL                        SSLConfig `yaml:"ssl"`
	PrivateSSLCA               CAConfig  `yaml:"private_ssl_ca"`
	ChiaSSLCA                  CAConfig  `yaml:"chia_ssl_ca"`
	ParallelDecompressorCount  uint8     `yaml:"parallel_decompressor_count"`
	DecompressorThreadCount    uint8     `yaml:"decompressor_thread_count"`
	DisableCPUAffinity         bool      `yaml:"disable_cpu_affinity"`
	MaxCompressionLevelAllowed uint8     `yaml:"max_compression_level_allowed"`
	UseGPUHarvesting           bool      `yaml:"use_gpu_harvesting"`
	GPUIndex                   uint8     `yaml:"gpu_index"`
	EnforceGPUIndex            bool      `yaml:"enforce_gpu_index"`
	DecompressorTimeout        uint16    `yaml:"decompressor_timeout"`
}

// PlotsRefreshParameter refresh params for harvester
type PlotsRefreshParameter struct {
	IntervalSeconds        uint16 `yaml:"interval_seconds"`
	RetryInvalidSeconds    uint16 `yaml:"retry_invalid_seconds"`
	BatchSize              uint16 `yaml:"batch_size"`
	BatchSleepMilliseconds uint16 `yaml:"batch_sleep_milliseconds"`
}

// PoolConfig configures pool settings
type PoolConfig struct {
	XCHTargetAddress string            `yaml:"xch_target_address,omitempty"`
	Logging          LoggingConfig     `yaml:"logging"`
	NetworkOverrides *NetworkOverrides `yaml:"network_overrides"`
	SelectedNetwork  *string           `yaml:"selected_network"`
}

// FarmerConfig farmer configuration section
type FarmerConfig struct {
	FullNodePeers      []Peer            `yaml:"full_node_peers"`
	PoolPublicKeys     types.WonkySet    `yaml:"pool_public_keys"`
	XCHTargetAddress   string            `yaml:"xch_target_address,omitempty"`
	StartRPCServer     bool              `yaml:"start_rpc_server"`
	EnableProfiler     bool              `yaml:"enable_profiler"`
	PoolShareThreshold uint32            `yaml:"pool_share_threshold"`
	Logging            LoggingConfig     `yaml:"logging"`
	NetworkOverrides   *NetworkOverrides `yaml:"network_overrides"`
	SelectedNetwork    *string           `yaml:"selected_network"`
	PortConfig         `yaml:",inline"`
	SSL                SSLConfig `yaml:"ssl"`
}

// TimelordLauncherConfig settings for vdf_client launcher
type TimelordLauncherConfig struct {
	Host         string        `yaml:"host"`
	Port         uint16        `yaml:"port"`
	ProcessCount uint8         `yaml:"process_count"`
	Logging      LoggingConfig `yaml:"logging"`
}

// TimelordConfig timelord configuration section
type TimelordConfig struct {
	VDFClients                 VDFClients        `yaml:"vdf_clients"`
	FullNodePeers              []Peer            `yaml:"full_node_peers"`
	MaxConnectionTime          uint16            `yaml:"max_connection_time"`
	VDFServer                  Peer              `yaml:"vdf_server"`
	Logging                    LoggingConfig     `yaml:"logging"`
	NetworkOverrides           *NetworkOverrides `yaml:"network_overrides"`
	SelectedNetwork            *string           `yaml:"selected_network"`
	FastAlgorithm              bool              `yaml:"fast_algorithm"`
	BlueboxMode                bool              `yaml:"bluebox_mode"`
	SlowBluebox                bool              `yaml:"slow_bluebox"`
	SlowBlueboxProcessCount    uint8             `yaml:"slow_bluebox_process_count"`
	MultiprocessingStartMethod string            `yaml:"multiprocessing_start_method"`
	StartRPCServer             bool              `yaml:"start_rpc_server"`
	PortConfig                 `yaml:",inline"`
	SSL                        SSLConfig `yaml:"ssl"`
}

// VDFClients is a list of allowlisted IPs for vdf_client
type VDFClients struct {
	IP          []string `yaml:"ip"`
	IPSEstimate []uint32 `yaml:"ips_estimate"`
}

// FullNodeConfig full node configuration section
type FullNodeConfig struct {
	PortConfig                       `yaml:",inline"`
	FullNodePeers                    []Peer            `yaml:"full_node_peers"`
	DBSync                           string            `yaml:"db_sync"`
	DBReaders                        uint8             `yaml:"db_readers"`
	DatabasePath                     string            `yaml:"database_path"`
	PeerDBPath                       string            `yaml:"peer_db_path"`
	PeersFilePath                    string            `yaml:"peers_file_path"`
	MultiprocessingStartMethod       string            `yaml:"multiprocessing_start_method"`
	MaxDuplicateUnfinishedBlocks     uint8             `yaml:"max_duplicate_unfinished_blocks"`
	StartRPCServer                   bool              `yaml:"start_rpc_server"`
	EnableUPNP                       bool              `yaml:"enable_upnp"`
	SyncBlocksBehindThreshold        uint16            `yaml:"sync_blocks_behind_threshold"`
	ShortSyncBlocksBehindThreshold   uint16            `yaml:"short_sync_blocks_behind_threshold"`
	BadPeakCacheSize                 uint16            `yaml:"bad_peak_cache_size"`
	ReservedCores                    uint8             `yaml:"reserved_cores"`
	SingleThreaded                   bool              `yaml:"single_threaded"`
	PeerConnectInterval              uint8             `yaml:"peer_connect_interval"`
	PeerConnectTimeout               uint8             `yaml:"peer_connect_timeout"`
	TargetPeerCount                  uint16            `yaml:"target_peer_count"`
	TargetOutboundPeerCount          uint16            `yaml:"target_outbound_peer_count"`
	ExemptPeerNetworks               []string          `yaml:"exempt_peer_networks"`
	MaxInboundWallet                 uint8             `yaml:"max_inbound_wallet"`
	MaxInboundFarmer                 uint8             `yaml:"max_inbound_farmer"`
	MaxInboundTimelord               uint8             `yaml:"max_inbound_timelord"`
	RecentPeerThreshold              uint16            `yaml:"recent_peer_threshold"`
	SendUncompactInterval            uint16            `yaml:"send_uncompact_interval"`
	TargetUncompactProofs            uint16            `yaml:"target_uncompact_proofs"`
	SanitizeWeightProofOnly          bool              `yaml:"sanitize_weight_proof_only"`
	WeightProofTimeout               uint16            `yaml:"weight_proof_timeout"`
	MaxSyncWait                      uint16            `yaml:"max_sync_wait"`
	EnableProfiler                   bool              `yaml:"enable_profiler"`
	ProfileBlockValidation           bool              `yaml:"profile_block_validation"`
	EnableMemoryProfiler             bool              `yaml:"enable_memory_profiler"`
	LogSqliteCmds                    bool              `yaml:"log_sqlite_cmds"`
	MaxSubscribeItems                uint32            `yaml:"max_subscribe_items"`
	MaxSubscribeResponseItems        uint32            `yaml:"max_subscribe_response_items"`
	TrustedMaxSubscribeItems         uint32            `yaml:"trusted_max_subscribe_items"`
	TrustedMaxSubscribeResponseItems uint32            `yaml:"trusted_max_subscribe_response_items"`
	DNSServers                       []string          `yaml:"dns_servers"`
	IntroducerPeer                   Peer              `yaml:"introducer_peer"`
	Logging                          LoggingConfig     `yaml:"logging"`
	NetworkOverrides                 *NetworkOverrides `yaml:"network_overrides"`
	SelectedNetwork                  *string           `yaml:"selected_network"`
	TrustedPeers                     map[string]string `yaml:"trusted_peers"`
	SSL                              SSLConfig         `yaml:"ssl"`
	UseChiaLoopPolicy                bool              `yaml:"use_chia_loop_policy"`
}

// UIConfig settings for the UI
type UIConfig struct {
	PortConfig       `yaml:",inline"`
	SSHFilename      string            `yaml:"ssh_filename"`
	Logging          LoggingConfig     `yaml:"logging"`
	NetworkOverrides *NetworkOverrides `yaml:"network_overrides"`
	SelectedNetwork  *string           `yaml:"selected_network"`
	DaemonHost       string            `yaml:"daemon_host"`
	DaemonPort       uint16            `yaml:"daemon_port"`
	DaemonSSL        SSLConfig         `yaml:"daemon_ssl"`
}

// IntroducerConfig settings for introducers
type IntroducerConfig struct {
	Host                string `yaml:"host"`
	PortConfig          `yaml:",inline"`
	MaxPeersToSend      uint16            `yaml:"max_peers_to_send"`
	RecentPeerThreshold uint16            `yaml:"recent_peer_threshold"`
	Logging             LoggingConfig     `yaml:"logging"`
	NetworkOverrides    *NetworkOverrides `yaml:"network_overrides"`
	SelectedNetwork     *string           `yaml:"selected_network"`
	SSL                 SSLConfig         `yaml:"ssl"`
}

// WalletConfig wallet configuration section
type WalletConfig struct {
	PortConfig                     `yaml:",inline"`
	EnableProfiler                 bool              `yaml:"enable_profiler"`
	EnableMemoryProfiler           bool              `yaml:"enable_memory_profiler"`
	DBSync                         string            `yaml:"db_sync"`
	DBReaders                      uint8             `yaml:"db_readers"`
	ConnectToUnknownPeers          bool              `yaml:"connect_to_unknown_peers"`
	InitialNumPublicKeys           uint16            `yaml:"initial_num_public_keys"`
	ReusePublicKeyForChange        map[string]bool   `yaml:"reuse_public_key_for_change"`
	DNSServers                     []string          `yaml:"dns_servers"`
	FullNodePeers                  []Peer            `yaml:"full_node_peers"`
	NFTMetadataCachePath           string            `yaml:"nft_cache"`
	NFTMetadataCacheHashLength     uint8             `yaml:"nft_metadata_cache_hash_length"`
	MultiprocessingStartMethod     string            `yaml:"multiprocessing_start_method"`
	Testing                        bool              `yaml:"testing"`
	DatabasePath                   string            `yaml:"database_path"`
	WalletPeersPath                string            `yaml:"wallet_peers_path"`
	WalletPeersFilePath            string            `yaml:"wallet_peers_file_path"`
	LogSqliteCmds                  bool              `yaml:"log_sqlite_cmds"`
	Logging                        LoggingConfig     `yaml:"logging"`
	NetworkOverrides               *NetworkOverrides `yaml:"network_overrides"`
	SelectedNetwork                *string           `yaml:"selected_network"`
	TargetPeerCount                uint16            `yaml:"target_peer_count"`
	PeerConnectInterval            uint8             `yaml:"peer_connect_interval"`
	RecentPeerThreshold            uint16            `yaml:"recent_peer_threshold"`
	IntroducerPeer                 Peer              `yaml:"introducer_peer"`
	SSL                            SSLConfig         `yaml:"ssl"`
	TrustedPeers                   map[string]string `yaml:"trusted_peers"`
	ShortSyncBlocksBehindThreshold uint16            `yaml:"short_sync_blocks_behind_threshold"`
	InboundRateLimitPercent        uint8             `yaml:"inbound_rate_limit_percent"`
	OutboundRateLimitPercent       uint8             `yaml:"outbound_rate_limit_percent"`
	WeightProofTimeout             uint16            `yaml:"weight_proof_timeout"`
	AutomaticallyAddUnknownCats    bool              `yaml:"automatically_add_unknown_cats"`
	TxResendTimeoutSecs            uint16            `yaml:"tx_resend_timeout_secs"`
	ResetSyncForFingerprint        *int              `yaml:"reset_sync_for_fingerprint"`
	SpamFilterAfterNTxs            uint16            `yaml:"spam_filter_after_n_txs"`
	XCHSpamAmount                  uint64            `yaml:"xch_spam_amount"`
	EnableNotifications            bool              `yaml:"enable_notifications"`
	RequiredNotificationAmount     uint64            `yaml:"required_notification_amount"`
	UseDeltaSync                   bool              `yaml:"use_delta_sync"`
	// PuzzleDecorators
	AutoClaim   AutoClaim `yaml:"auto_claim"`
	AutoSignTxs bool      `yaml:"auto_sign_txs"`
}

// AutoClaim settings for auto claim in wallet
type AutoClaim struct {
	Enabled   bool   `yaml:"enabled"`
	TxFee     uint64 `yaml:"tx_fee"`
	MinAmount uint64 `yaml:"min_amount"`
	BatchSize uint16 `yaml:"batch_size"`
}

// DataLayerConfig datalayer configuration section
type DataLayerConfig struct {
	WalletPeer                  Peer          `yaml:"wallet_peer"`
	DatabasePath                string        `yaml:"database_path"`
	ServerFilesLocation         string        `yaml:"server_files_location"`
	ClientTimeout               uint16        `yaml:"client_timeout"`
	ProxyURL                    string        `yaml:"proxy_url,omitempty"`
	HostIP                      string        `yaml:"host_ip"`
	HostPort                    uint16        `yaml:"host_port"`
	ManageDataInterval          uint16        `yaml:"manage_data_interval"`
	SelectedNetwork             *string       `yaml:"selected_network"`
	StartRPCServer              bool          `yaml:"start_rpc_server"`
	RPCServerMaxRequestBodySize uint32        `yaml:"rpc_server_max_request_body_size"`
	LogSqliteCmds               bool          `yaml:"log_sqlite_cmds"`
	EnableBatchAutoinsert       bool          `yaml:"enable_batch_autoinsert"`
	Logging                     LoggingConfig `yaml:"logging"`
	PortConfig                  `yaml:",inline"`
	SSL                         SSLConfig        `yaml:"ssl"`
	Plugins                     DataLayerPlugins `yaml:"plugins"`
	MaximumFullFileCount        uint16           `yaml:"maximum_full_file_count"`
}

// DataLayerPlugins Settings for data layer plugins
type DataLayerPlugins struct {
	// @TODO
}

// SimulatorConfig settings for simulator
type SimulatorConfig struct {
	AutoFarm       bool   `yaml:"auto_farm"`
	KeyFingerprint int    `yaml:"key_fingerprint"`
	FarmingAddress string `yaml:"farming_address"`
	PlotDirectory  string `yaml:"plot_directory"`
	UseCurrentTime bool   `yaml:"use_current_time"`
}
