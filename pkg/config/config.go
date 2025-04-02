package config

import (
	"gopkg.in/yaml.v3"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

// ChiaConfig the chia config.yaml
type ChiaConfig struct {
	// Tracks where the config was loaded from so we can call Save()
	configPath               string
	ChiaRoot                 string                 `yaml:"-" json:"-"`
	UnknownFields            map[string]any         `yaml:",inline" json:",inline"`
	MinMainnetKSize          uint8                  `yaml:"min_mainnet_k_size" json:"min_mainnet_k_size"`
	PingInterval             uint16                 `yaml:"ping_interval" json:"ping_interval"`
	SelfHostname             string                 `yaml:"self_hostname" json:"self_hostname"`
	PreferIPv6               bool                   `yaml:"prefer_ipv6" json:"prefer_ipv6"`
	RPCTimeout               uint16                 `yaml:"rpc_timeout" json:"rpc_timeout"`
	DaemonPort               uint16                 `yaml:"daemon_port" json:"daemon_port"`
	DaemonMaxMessageSize     uint32                 `yaml:"daemon_max_message_size" json:"daemon_max_message_size"`
	DaemonHeartbeat          uint16                 `yaml:"daemon_heartbeat" json:"daemon_heartbeat"`
	DaemonAllowTLS12         bool                   `yaml:"daemon_allow_tls_1_2" json:"daemon_allow_tls_1_2"`
	InboundRateLimitPercent  uint8                  `yaml:"inbound_rate_limit_percent" json:"inbound_rate_limit_percent"`
	OutboundRateLimitPercent uint8                  `yaml:"outbound_rate_limit_percent" json:"outbound_rate_limit_percent"`
	NetworkOverrides         *NetworkOverrides      `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork          *string                `yaml:"selected_network" json:"selected_network"`
	AlertsURL                string                 `yaml:"ALERTS_URL,omitempty" json:"ALERTS_URL,omitempty"`
	ChiaAlertsPubkey         string                 `yaml:"CHIA_ALERTS_PUBKEY,omitempty" json:"CHIA_ALERTS_PUBKEY,omitempty"`
	PrivateSSLCA             CAConfig               `yaml:"private_ssl_ca" json:"private_ssl_ca"`
	ChiaSSLCA                CAConfig               `yaml:"chia_ssl_ca" json:"chia_ssl_ca"`
	DaemonSSL                SSLConfig              `yaml:"daemon_ssl" json:"daemon_ssl"`
	Logging                  *LoggingConfig         `yaml:"logging" json:"logging"`
	Seeder                   SeederConfig           `yaml:"seeder" json:"seeder"`
	Harvester                HarvesterConfig        `yaml:"harvester" json:"harvester"`
	Pool                     PoolConfig             `yaml:"pool" json:"pool"`
	Farmer                   FarmerConfig           `yaml:"farmer" json:"farmer"`
	TimelordLauncher         TimelordLauncherConfig `yaml:"timelord_launcher" json:"timelord_launcher"`
	Timelord                 TimelordConfig         `yaml:"timelord" json:"timelord"`
	FullNode                 FullNodeConfig         `yaml:"full_node" json:"full_node"`
	UI                       UIConfig               `yaml:"ui" json:"ui"`
	Introducer               IntroducerConfig       `yaml:"introducer" json:"introducer"`
	Wallet                   WalletConfig           `yaml:"wallet" json:"wallet"`
	DataLayer                DataLayerConfig        `yaml:"data_layer" json:"data_layer"`
	Simulator                SimulatorConfig        `yaml:"simulator" json:"simulator"`
	// Simulator Fork Settings
	HardForkHeight  uint32 `yaml:"HARD_FORK_HEIGHT" json:"HARD_FORK_HEIGHT"`
	SoftFork4Height uint32 `yaml:"SOFT_FORK4_HEIGHT" json:"SOFT_FORK4_HEIGHT"`
	SoftFork5Height uint32 `yaml:"SOFT_FORK5_HEIGHT" json:"SOFT_FORK5_HEIGHT"`
	SoftFork6Height uint32 `yaml:"SOFT_FORK6_HEIGHT" json:"SOFT_FORK6_HEIGHT"`
}

// PortConfig common port settings found in many sections of the config
type PortConfig struct {
	UnknownFields map[string]any `yaml:",inline" json:",inline"`
	Port          uint16         `yaml:"port,omitempty" json:"port,omitempty"`
	RPCPort       uint16         `yaml:"rpc_port,omitempty" json:"rpc_port,omitempty"`
}

// CAConfig config keys for CA
type CAConfig struct {
	UnknownFields map[string]any `yaml:",inline" json:",inline"`
	Crt           string         `yaml:"crt" json:"crt"`
	Key           string         `yaml:"key" json:"key"`
}

// SSLConfig common ssl settings found in many sections of the config
type SSLConfig struct {
	UnknownFields map[string]any `yaml:",inline" json:",inline"`
	PrivateCRT    string         `yaml:"private_crt,omitempty" json:"private_crt,omitempty"`
	PrivateKey    string         `yaml:"private_key,omitempty" json:"private_key,omitempty"`
	PublicCRT     string         `yaml:"public_crt,omitempty" json:"public_crt,omitempty"`
	PublicKey     string         `yaml:"public_key,omitempty" json:"public_key,omitempty"`
}

// Peer is a host/port for a peer
type Peer struct {
	UnknownFields         map[string]any `yaml:",inline" json:",inline"`
	Host                  string         `yaml:"host" json:"host"`
	Port                  uint16         `yaml:"port" json:"port"`
	EnablePrivateNetworks bool           `yaml:"enable_private_networks,omitempty" json:"enable_private_networks,omitempty"`
}

// NetworkOverrides is all network settings
type NetworkOverrides struct {
	yamlAnchor    *yaml.Node                  `yaml:"-" json:"-"` // Helps with serializing the anchors to yaml
	Constants     map[string]NetworkConstants `yaml:"constants" json:"constants"`
	Config        map[string]NetworkConfig    `yaml:"config" json:"config"`
}

// AnchorNode returns the node to be used in yaml anchors
func (nc *NetworkOverrides) AnchorNode() *yaml.Node {
	return nc.yamlAnchor
}

// SetAnchorNode sets the yaml.Node reference when marshaling
func (nc *NetworkOverrides) SetAnchorNode(node *yaml.Node) {
	nc.yamlAnchor = node
}

// NetworkConstants the constants for each network
type NetworkConstants struct {
	AggSigMeAdditionalData         string         `yaml:"AGG_SIG_ME_ADDITIONAL_DATA,omitempty" json:"AGG_SIG_ME_ADDITIONAL_DATA,omitempty"`
	DifficultyConstantFactor       types.Uint128  `yaml:"DIFFICULTY_CONSTANT_FACTOR,omitempty" json:"DIFFICULTY_CONSTANT_FACTOR,omitempty"`
	DifficultyStarting             uint64         `yaml:"DIFFICULTY_STARTING,omitempty" json:"DIFFICULTY_STARTING,omitempty"`
	EpochBlocks                    uint32         `yaml:"EPOCH_BLOCKS,omitempty" json:"EPOCH_BLOCKS,omitempty"`
	GenesisChallenge               string         `yaml:"GENESIS_CHALLENGE" json:"GENESIS_CHALLENGE"`
	GenesisPreFarmPoolPuzzleHash   string         `yaml:"GENESIS_PRE_FARM_POOL_PUZZLE_HASH" json:"GENESIS_PRE_FARM_POOL_PUZZLE_HASH"`
	GenesisPreFarmFarmerPuzzleHash string         `yaml:"GENESIS_PRE_FARM_FARMER_PUZZLE_HASH" json:"GENESIS_PRE_FARM_FARMER_PUZZLE_HASH"`
	MempoolBlockBuffer             uint8          `yaml:"MEMPOOL_BLOCK_BUFFER,omitempty" json:"MEMPOOL_BLOCK_BUFFER,omitempty"`
	MinPlotSize                    uint8          `yaml:"MIN_PLOT_SIZE,omitempty" json:"MIN_PLOT_SIZE,omitempty"`
	NetworkType                    uint8          `yaml:"NETWORK_TYPE,omitempty" json:"NETWORK_TYPE,omitempty"`
	SubSlotItersStarting           uint64         `yaml:"SUB_SLOT_ITERS_STARTING,omitempty" json:"SUB_SLOT_ITERS_STARTING,omitempty"`
	// All pointers that that 0 is an allowed value when marshaling with omitempty, but they will still be omitted from configs that dont have them defined
	HardForkHeight      *uint32 `yaml:"HARD_FORK_HEIGHT,omitempty" json:"HARD_FORK_HEIGHT,omitempty"`
	SoftFork4Height     *uint32 `yaml:"SOFT_FORK4_HEIGHT,omitempty" json:"SOFT_FORK4_HEIGHT,omitempty"`
	SoftFork5Height     *uint32 `yaml:"SOFT_FORK5_HEIGHT,omitempty" json:"SOFT_FORK5_HEIGHT,omitempty"`
	SoftFork6Height     *uint32 `yaml:"SOFT_FORK6_HEIGHT,omitempty" json:"SOFT_FORK6_HEIGHT,omitempty"`
	PlotFilter128Height *uint32 `yaml:"PLOT_FILTER_128_HEIGHT,omitempty" json:"PLOT_FILTER_128_HEIGHT,omitempty"`
	PlotFilter64Height  *uint32 `yaml:"PLOT_FILTER_64_HEIGHT,omitempty" json:"PLOT_FILTER_64_HEIGHT,omitempty"`
	PlotFilter32Height  *uint32 `yaml:"PLOT_FILTER_32_HEIGHT,omitempty" json:"PLOT_FILTER_32_HEIGHT,omitempty"`
}

// NetworkConfig specific network configuration settings
type NetworkConfig struct {
	AddressPrefix       string         `yaml:"address_prefix" json:"address_prefix"`
	DefaultFullNodePort uint16         `yaml:"default_full_node_port,omitempty" json:"default_full_node_port,omitempty"`
}

// LoggingConfig configuration settings for the logger
type LoggingConfig struct {
	yamlAnchor          *yaml.Node     `yaml:"-" json:"-"` // Helps with serializing the anchors to yaml
	UnknownFields       map[string]any `yaml:",inline" json:",inline"`
	LogStdout           bool           `yaml:"log_stdout" json:"log_stdout"`
	LogBackcompat       bool           `yaml:"log_backcompat" json:"log_backcompat"`
	LogFilename         string         `yaml:"log_filename" json:"log_filename"`
	LogLevel            string         `yaml:"log_level" json:"log_level"`
	LogMaxFilesRotation uint32         `yaml:"log_maxfilesrotation" json:"log_maxfilesrotation"`
	LogMaxBytesRotation uint32         `yaml:"log_maxbytesrotation" json:"log_maxbytesrotation"`
	LogUseGzip          bool           `yaml:"log_use_gzip" json:"log_use_gzip"`
	LogSyslog           bool           `yaml:"log_syslog" json:"log_syslog"`
	LogSyslogHost       string         `yaml:"log_syslog_host" json:"log_syslog_host"`
	LogSyslogPort       uint16         `yaml:"log_syslog_port" json:"log_syslog_port"`
}

// AnchorNode returns the node to be used in yaml anchors
func (lc *LoggingConfig) AnchorNode() *yaml.Node {
	return lc.yamlAnchor
}

// SetAnchorNode sets the yaml.Node reference when marshaling
func (lc *LoggingConfig) SetAnchorNode(node *yaml.Node) {
	lc.yamlAnchor = node
}

// SeederConfig seeder configuration section
type SeederConfig struct {
	UnknownFields       map[string]any    `yaml:",inline" json:",inline"`
	Port                uint16            `yaml:"port" json:"port"`
	OtherPeersPort      uint16            `yaml:"other_peers_port" json:"other_peers_port"`
	DNSPort             uint16            `yaml:"dns_port" json:"dns_port"`
	PeerConnectTimeout  uint16            `yaml:"peer_connect_timeout" json:"peer_connect_timeout"`
	CrawlerDBPath       string            `yaml:"crawler_db_path" json:"crawler_db_path"`
	BootstrapPeers      []string          `yaml:"bootstrap_peers" json:"bootstrap_peers"`
	StaticPeers         []string          `yaml:"static_peers" json:"static_peers"`
	MinimumHeight       uint32            `yaml:"minimum_height" json:"minimum_height"`
	MinimumVersionCount uint32            `yaml:"minimum_version_count" json:"minimum_version_count"`
	DomainName          string            `yaml:"domain_name" json:"domain_name"`
	Nameserver          string            `yaml:"nameserver" json:"nameserver"`
	TTL                 uint16            `yaml:"ttl" json:"ttl"`
	SOA                 SeederSOA         `yaml:"soa" json:"soa"`
	NetworkOverrides    *NetworkOverrides `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork     *string           `yaml:"selected_network" json:"selected_network"`
	Logging             *LoggingConfig    `yaml:"logging" json:"logging"`
	CrawlerConfig       CrawlerConfig     `yaml:"crawler" json:"crawler"`
}

// SeederSOA dns SOA for seeder
type SeederSOA struct {
	UnknownFields map[string]any `yaml:",inline" json:",inline"`
	Rname         string         `yaml:"rname" json:"rname"`
	SerialNumber  uint32         `yaml:"serial_number" json:"serial_number"`
	Refresh       uint32         `yaml:"refresh" json:"refresh"`
	Retry         uint32         `yaml:"retry" json:"retry"`
	Expire        uint32         `yaml:"expire" json:"expire"`
	Minimum       uint32         `yaml:"minimum" json:"minimum"`
}

// CrawlerConfig is the subsection of the seeder config specific to the crawler
type CrawlerConfig struct {
	UnknownFields  map[string]any `yaml:",inline" json:",inline"`
	StartRPCServer bool           `yaml:"start_rpc_server" json:"start_rpc_server"`
	PortConfig     `yaml:",inline" json:",inline"`
	PrunePeerDays  uint32    `yaml:"prune_peer_days" json:"prune_peer_days"`
	SSL            SSLConfig `yaml:"ssl" json:"ssl"`
}

// HarvesterConfig harvester configuration section
type HarvesterConfig struct {
	UnknownFields              map[string]any        `yaml:",inline" json:",inline"`
	FarmerPeers                []Peer                `yaml:"farmer_peers" json:"farmer_peers"`
	StartRPCServer             bool                  `yaml:"start_rpc_server" json:"start_rpc_server"`
	NumThreads                 uint8                 `yaml:"num_threads" json:"num_threads"`
	PlotsRefreshParameter      PlotsRefreshParameter `yaml:"plots_refresh_parameter" json:"plots_refresh_parameter"`
	ParallelRead               bool                  `yaml:"parallel_read" json:"parallel_read"`
	Logging                    *LoggingConfig        `yaml:"logging" json:"logging"`
	NetworkOverrides           *NetworkOverrides     `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork            *string               `yaml:"selected_network" json:"selected_network"`
	PlotDirectories            []string              `yaml:"plot_directories" json:"plot_directories"`
	RecursivePlotScan          bool                  `yaml:"recursive_plot_scan" json:"recursive_plot_scan"`
	RecursiveFollowLinks       bool                  `yaml:"recursive_follow_links" json:"recursive_follow_links"`
	PortConfig                 `yaml:",inline" json:",inline"`
	SSL                        SSLConfig `yaml:"ssl" json:"ssl"`
	PrivateSSLCA               CAConfig  `yaml:"private_ssl_ca" json:"private_ssl_ca"`
	ChiaSSLCA                  CAConfig  `yaml:"chia_ssl_ca" json:"chia_ssl_ca"`
	ParallelDecompressorCount  uint8     `yaml:"parallel_decompressor_count" json:"parallel_decompressor_count"`
	DecompressorThreadCount    uint8     `yaml:"decompressor_thread_count" json:"decompressor_thread_count"`
	DisableCPUAffinity         bool      `yaml:"disable_cpu_affinity" json:"disable_cpu_affinity"`
	MaxCompressionLevelAllowed uint8     `yaml:"max_compression_level_allowed" json:"max_compression_level_allowed"`
	UseGPUHarvesting           bool      `yaml:"use_gpu_harvesting" json:"use_gpu_harvesting"`
	GPUIndex                   uint8     `yaml:"gpu_index" json:"gpu_index"`
	EnforceGPUIndex            bool      `yaml:"enforce_gpu_index" json:"enforce_gpu_index"`
	DecompressorTimeout        uint16    `yaml:"decompressor_timeout" json:"decompressor_timeout"`
}

// PlotsRefreshParameter refresh params for harvester
type PlotsRefreshParameter struct {
	UnknownFields          map[string]any `yaml:",inline" json:",inline"`
	IntervalSeconds        uint16         `yaml:"interval_seconds" json:"interval_seconds"`
	RetryInvalidSeconds    uint16         `yaml:"retry_invalid_seconds" json:"retry_invalid_seconds"`
	BatchSize              uint16         `yaml:"batch_size" json:"batch_size"`
	BatchSleepMilliseconds uint16         `yaml:"batch_sleep_milliseconds" json:"batch_sleep_milliseconds"`
}

// PoolConfig configures pool settings
type PoolConfig struct {
	UnknownFields    map[string]any    `yaml:",inline" json:",inline"`
	XCHTargetAddress string            `yaml:"xch_target_address,omitempty" json:"xch_target_address,omitempty"`
	Logging          *LoggingConfig    `yaml:"logging" json:"logging"`
	NetworkOverrides *NetworkOverrides `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork  *string           `yaml:"selected_network" json:"selected_network"`
}

// FarmerConfig farmer configuration section
type FarmerConfig struct {
	UnknownFields      map[string]any    `yaml:",inline" json:",inline"`
	FullNodePeers      []Peer            `yaml:"full_node_peers" json:"full_node_peers"`
	PoolPublicKeys     types.WonkySet    `yaml:"pool_public_keys" json:"pool_public_keys"`
	XCHTargetAddress   string            `yaml:"xch_target_address,omitempty" json:"xch_target_address,omitempty"`
	StartRPCServer     bool              `yaml:"start_rpc_server" json:"start_rpc_server"`
	EnableProfiler     bool              `yaml:"enable_profiler" json:"enable_profiler"`
	PoolShareThreshold uint32            `yaml:"pool_share_threshold" json:"pool_share_threshold"`
	Logging            *LoggingConfig    `yaml:"logging" json:"logging"`
	NetworkOverrides   *NetworkOverrides `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork    *string           `yaml:"selected_network" json:"selected_network"`
	PortConfig         `yaml:",inline" json:",inline"`
	SSL                SSLConfig `yaml:"ssl" json:"ssl"`
}

// TimelordLauncherConfig settings for vdf_client launcher
type TimelordLauncherConfig struct {
	UnknownFields map[string]any `yaml:",inline" json:",inline"`
	Host          string         `yaml:"host" json:"host"`
	Port          uint16         `yaml:"port" json:"port"`
	ProcessCount  uint8          `yaml:"process_count" json:"process_count"`
	Logging       *LoggingConfig `yaml:"logging" json:"logging"`
}

// TimelordConfig timelord configuration section
type TimelordConfig struct {
	UnknownFields              map[string]any    `yaml:",inline" json:",inline"`
	VDFClients                 VDFClients        `yaml:"vdf_clients" json:"vdf_clients"`
	FullNodePeers              []Peer            `yaml:"full_node_peers" json:"full_node_peers"`
	MaxConnectionTime          uint16            `yaml:"max_connection_time" json:"max_connection_time"`
	VDFServer                  Peer              `yaml:"vdf_server" json:"vdf_server"`
	Logging                    *LoggingConfig    `yaml:"logging" json:"logging"`
	NetworkOverrides           *NetworkOverrides `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork            *string           `yaml:"selected_network" json:"selected_network"`
	BlueboxMode                bool              `yaml:"bluebox_mode" json:"bluebox_mode"`
	SlowBluebox                bool              `yaml:"slow_bluebox" json:"slow_bluebox"`
	SlowBlueboxProcessCount    uint8             `yaml:"slow_bluebox_process_count" json:"slow_bluebox_process_count"`
	MultiprocessingStartMethod string            `yaml:"multiprocessing_start_method" json:"multiprocessing_start_method"`
	StartRPCServer             bool              `yaml:"start_rpc_server" json:"start_rpc_server"`
	PortConfig                 `yaml:",inline" json:",inline"`
	SSL                        SSLConfig `yaml:"ssl" json:"ssl"`
}

// VDFClients is a list of allowlisted IPs for vdf_client
type VDFClients struct {
	UnknownFields map[string]any `yaml:",inline" json:",inline"`
	IP            []string       `yaml:"ip" json:"ip"`
	IPSEstimate   []uint32       `yaml:"ips_estimate" json:"ips_estimate"`
}

// FullNodeConfig full node configuration section
type FullNodeConfig struct {
	UnknownFields                    map[string]any `yaml:",inline" json:",inline"`
	PortConfig                       `yaml:",inline" json:",inline"`
	FullNodePeers                    []Peer            `yaml:"full_node_peers" json:"full_node_peers"`
	DBSync                           string            `yaml:"db_sync" json:"db_sync"`
	DBReaders                        uint8             `yaml:"db_readers" json:"db_readers"`
	DatabasePath                     string            `yaml:"database_path" json:"database_path"`
	PeerDBPath                       string            `yaml:"peer_db_path" json:"peer_db_path"`
	PeersFilePath                    string            `yaml:"peers_file_path" json:"peers_file_path"`
	MultiprocessingStartMethod       string            `yaml:"multiprocessing_start_method" json:"multiprocessing_start_method"`
	MaxDuplicateUnfinishedBlocks     uint8             `yaml:"max_duplicate_unfinished_blocks" json:"max_duplicate_unfinished_blocks"`
	StartRPCServer                   bool              `yaml:"start_rpc_server" json:"start_rpc_server"`
	EnableUPNP                       bool              `yaml:"enable_upnp" json:"enable_upnp"`
	SyncBlocksBehindThreshold        uint16            `yaml:"sync_blocks_behind_threshold" json:"sync_blocks_behind_threshold"`
	ShortSyncBlocksBehindThreshold   uint16            `yaml:"short_sync_blocks_behind_threshold" json:"short_sync_blocks_behind_threshold"`
	BadPeakCacheSize                 uint16            `yaml:"bad_peak_cache_size" json:"bad_peak_cache_size"`
	ReservedCores                    uint8             `yaml:"reserved_cores" json:"reserved_cores"`
	SingleThreaded                   bool              `yaml:"single_threaded" json:"single_threaded"`
	LogCoins                         bool              `yaml:"log_coins" json:"log_coins"`
	PeerConnectInterval              uint8             `yaml:"peer_connect_interval" json:"peer_connect_interval"`
	PeerConnectTimeout               uint8             `yaml:"peer_connect_timeout" json:"peer_connect_timeout"`
	TargetPeerCount                  uint16            `yaml:"target_peer_count" json:"target_peer_count"`
	TargetOutboundPeerCount          uint16            `yaml:"target_outbound_peer_count" json:"target_outbound_peer_count"`
	ExemptPeerNetworks               []string          `yaml:"exempt_peer_networks" json:"exempt_peer_networks"`
	MaxInboundWallet                 uint8             `yaml:"max_inbound_wallet" json:"max_inbound_wallet"`
	MaxInboundFarmer                 uint8             `yaml:"max_inbound_farmer" json:"max_inbound_farmer"`
	MaxInboundTimelord               uint8             `yaml:"max_inbound_timelord" json:"max_inbound_timelord"`
	RecentPeerThreshold              uint16            `yaml:"recent_peer_threshold" json:"recent_peer_threshold"`
	SendUncompactInterval            uint16            `yaml:"send_uncompact_interval" json:"send_uncompact_interval"`
	TargetUncompactProofs            uint16            `yaml:"target_uncompact_proofs" json:"target_uncompact_proofs"`
	SanitizeWeightProofOnly          bool              `yaml:"sanitize_weight_proof_only" json:"sanitize_weight_proof_only"`
	WeightProofTimeout               uint16            `yaml:"weight_proof_timeout" json:"weight_proof_timeout"`
	MaxSyncWait                      uint16            `yaml:"max_sync_wait" json:"max_sync_wait"`
	EnableProfiler                   bool              `yaml:"enable_profiler" json:"enable_profiler"`
	ProfileBlockValidation           bool              `yaml:"profile_block_validation" json:"profile_block_validation"`
	EnableMemoryProfiler             bool              `yaml:"enable_memory_profiler" json:"enable_memory_profiler"`
	LogMempool                       bool              `yaml:"log_mempool" json:"log_mempool"`
	LogSqliteCmds                    bool              `yaml:"log_sqlite_cmds" json:"log_sqlite_cmds"`
	MaxSubscribeItems                uint32            `yaml:"max_subscribe_items" json:"max_subscribe_items"`
	MaxSubscribeResponseItems        uint32            `yaml:"max_subscribe_response_items" json:"max_subscribe_response_items"`
	TrustedMaxSubscribeItems         uint32            `yaml:"trusted_max_subscribe_items" json:"trusted_max_subscribe_items"`
	TrustedMaxSubscribeResponseItems uint32            `yaml:"trusted_max_subscribe_response_items" json:"trusted_max_subscribe_response_items"`
	DNSServers                       []string          `yaml:"dns_servers" json:"dns_servers"`
	IntroducerPeer                   Peer              `yaml:"introducer_peer" json:"introducer_peer"`
	Logging                          *LoggingConfig    `yaml:"logging" json:"logging"`
	NetworkOverrides                 *NetworkOverrides `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork                  *string           `yaml:"selected_network" json:"selected_network"`
	TrustedPeers                     map[string]string `yaml:"trusted_peers" json:"trusted_peers"`
	SSL                              SSLConfig         `yaml:"ssl" json:"ssl"`
	UseChiaLoopPolicy                bool              `yaml:"use_chia_loop_policy" json:"use_chia_loop_policy"`
	// trusted_cidrs allows marking certain nodes as "trusted" in the full node and wallet
	// Not in the initial config anywhere, since it's a more advanced option
	TrustedCIDRs []string `yaml:"trusted_cidrs,omitempty" json:"trusted_cidrs,omitempty"`
}

// UIConfig settings for the UI
type UIConfig struct {
	UnknownFields    map[string]any `yaml:",inline" json:",inline"`
	PortConfig       `yaml:",inline" json:",inline"`
	SSHFilename      string            `yaml:"ssh_filename" json:"ssh_filename"`
	Logging          *LoggingConfig    `yaml:"logging" json:"logging"`
	NetworkOverrides *NetworkOverrides `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork  *string           `yaml:"selected_network" json:"selected_network"`
	DaemonHost       string            `yaml:"daemon_host" json:"daemon_host"`
	DaemonPort       uint16            `yaml:"daemon_port" json:"daemon_port"`
	DaemonSSL        SSLConfig         `yaml:"daemon_ssl" json:"daemon_ssl"`
}

// IntroducerConfig settings for introducers
type IntroducerConfig struct {
	UnknownFields       map[string]any `yaml:",inline" json:",inline"`
	Host                string         `yaml:"host" json:"host"`
	PortConfig          `yaml:",inline" json:",inline"`
	MaxPeersToSend      uint16            `yaml:"max_peers_to_send" json:"max_peers_to_send"`
	RecentPeerThreshold uint16            `yaml:"recent_peer_threshold" json:"recent_peer_threshold"`
	Logging             *LoggingConfig    `yaml:"logging" json:"logging"`
	NetworkOverrides    *NetworkOverrides `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork     *string           `yaml:"selected_network" json:"selected_network"`
	DNSServers          []string          `yaml:"dns_servers" json:"dns_servers"`
	SSL                 SSLConfig         `yaml:"ssl" json:"ssl"`
}

// WalletConfig wallet configuration section
type WalletConfig struct {
	UnknownFields                  map[string]any `yaml:",inline" json:",inline"`
	PortConfig                     `yaml:",inline" json:",inline"`
	StartRPCServer                 *bool             `yaml:"start_rpc_server" json:"start_rpc_server"`
	EnableProfiler                 bool              `yaml:"enable_profiler" json:"enable_profiler"`
	EnableMemoryProfiler           bool              `yaml:"enable_memory_profiler" json:"enable_memory_profiler"`
	DBSync                         string            `yaml:"db_sync" json:"db_sync"`
	DBReaders                      uint8             `yaml:"db_readers" json:"db_readers"`
	ConnectToUnknownPeers          bool              `yaml:"connect_to_unknown_peers" json:"connect_to_unknown_peers"`
	InitialNumPublicKeys           uint16            `yaml:"initial_num_public_keys" json:"initial_num_public_keys"`
	ReusePublicKeyForChange        map[string]bool   `yaml:"reuse_public_key_for_change" json:"reuse_public_key_for_change"`
	DNSServers                     []string          `yaml:"dns_servers" json:"dns_servers"`
	FullNodePeers                  []Peer            `yaml:"full_node_peers" json:"full_node_peers"`
	NFTMetadataCachePath           string            `yaml:"nft_cache" json:"nft_cache"`
	NFTMetadataCacheHashLength     uint8             `yaml:"nft_metadata_cache_hash_length" json:"nft_metadata_cache_hash_length"`
	MultiprocessingStartMethod     string            `yaml:"multiprocessing_start_method" json:"multiprocessing_start_method"`
	Testing                        bool              `yaml:"testing" json:"testing"`
	DatabasePath                   string            `yaml:"database_path" json:"database_path"`
	WalletPeersPath                string            `yaml:"wallet_peers_path" json:"wallet_peers_path"`
	WalletPeersFilePath            string            `yaml:"wallet_peers_file_path" json:"wallet_peers_file_path"`
	LogSqliteCmds                  bool              `yaml:"log_sqlite_cmds" json:"log_sqlite_cmds"`
	Logging                        *LoggingConfig    `yaml:"logging" json:"logging"`
	NetworkOverrides               *NetworkOverrides `yaml:"network_overrides" json:"network_overrides"`
	SelectedNetwork                *string           `yaml:"selected_network" json:"selected_network"`
	TargetPeerCount                uint16            `yaml:"target_peer_count" json:"target_peer_count"`
	PeerConnectInterval            uint8             `yaml:"peer_connect_interval" json:"peer_connect_interval"`
	RecentPeerThreshold            uint16            `yaml:"recent_peer_threshold" json:"recent_peer_threshold"`
	IntroducerPeer                 Peer              `yaml:"introducer_peer" json:"introducer_peer"`
	SSL                            SSLConfig         `yaml:"ssl" json:"ssl"`
	TrustedPeers                   map[string]string `yaml:"trusted_peers" json:"trusted_peers"`
	ShortSyncBlocksBehindThreshold uint16            `yaml:"short_sync_blocks_behind_threshold" json:"short_sync_blocks_behind_threshold"`
	InboundRateLimitPercent        uint8             `yaml:"inbound_rate_limit_percent" json:"inbound_rate_limit_percent"`
	OutboundRateLimitPercent       uint8             `yaml:"outbound_rate_limit_percent" json:"outbound_rate_limit_percent"`
	WeightProofTimeout             uint16            `yaml:"weight_proof_timeout" json:"weight_proof_timeout"`
	AutomaticallyAddUnknownCats    bool              `yaml:"automatically_add_unknown_cats" json:"automatically_add_unknown_cats"`
	DIDAutoAddLimit                *int              `yaml:"did_auto_add_limit,omitempty" json:"did_auto_add_limit,omitempty"`
	TxResendTimeoutSecs            uint16            `yaml:"tx_resend_timeout_secs" json:"tx_resend_timeout_secs"`
	ResetSyncForFingerprint        *int              `yaml:"reset_sync_for_fingerprint" json:"reset_sync_for_fingerprint"`
	SpamFilterAfterNTxs            uint16            `yaml:"spam_filter_after_n_txs" json:"spam_filter_after_n_txs"`
	XCHSpamAmount                  uint64            `yaml:"xch_spam_amount" json:"xch_spam_amount"`
	EnableNotifications            *bool             `yaml:"enable_notifications,omitempty" json:"enable_notifications"`
	RequiredNotificationAmount     uint64            `yaml:"required_notification_amount" json:"required_notification_amount"`
	UseDeltaSync                   bool              `yaml:"use_delta_sync" json:"use_delta_sync"`
	// PuzzleDecorators
	AutoClaim   AutoClaim `yaml:"auto_claim" json:"auto_claim"`
	AutoSignTxs *bool     `yaml:"auto_sign_txs,omitempty" json:"auto_sign_txs,omitempty"`
	// trusted_cidrs allows marking certain nodes as "trusted" in the full node and wallet
	// Not in the initial config anywhere, since it's a more advanced option
	TrustedCIDRs []string `yaml:"trusted_cidrs,omitempty" json:"trusted_cidrs,omitempty"`
}

// AutoClaim settings for auto claim in wallet
type AutoClaim struct {
	UnknownFields map[string]any `yaml:",inline" json:",inline"`
	Enabled       bool           `yaml:"enabled" json:"enabled"`
	TxFee         uint64         `yaml:"tx_fee" json:"tx_fee"`
	MinAmount     uint64         `yaml:"min_amount" json:"min_amount"`
	BatchSize     uint16         `yaml:"batch_size" json:"batch_size"`
}

// DataLayerConfig datalayer configuration section
type DataLayerConfig struct {
	UnknownFields               map[string]any `yaml:",inline" json:",inline"`
	WalletPeer                  Peer           `yaml:"wallet_peer" json:"wallet_peer"`
	DatabasePath                string         `yaml:"database_path" json:"database_path"`
	ServerFilesLocation         string         `yaml:"server_files_location" json:"server_files_location"`
	ClientTimeout               uint16         `yaml:"client_timeout" json:"client_timeout"`
	ConnectTimeout              uint16         `yaml:"connect_timeout" json:"connect_timeout"`
	ProxyURL                    string         `yaml:"proxy_url,omitempty" json:"proxy_url,omitempty"`
	HostIP                      string         `yaml:"host_ip" json:"host_ip"`
	HostPort                    uint16         `yaml:"host_port" json:"host_port"`
	ManageDataInterval          uint16         `yaml:"manage_data_interval" json:"manage_data_interval"`
	SelectedNetwork             *string        `yaml:"selected_network" json:"selected_network"`
	StartRPCServer              bool           `yaml:"start_rpc_server" json:"start_rpc_server"`
	RPCServerMaxRequestBodySize uint32         `yaml:"rpc_server_max_request_body_size" json:"rpc_server_max_request_body_size"`
	LogSqliteCmds               bool           `yaml:"log_sqlite_cmds" json:"log_sqlite_cmds"`
	EnableBatchAutoinsert       bool           `yaml:"enable_batch_autoinsert" json:"enable_batch_autoinsert"`
	Logging                     *LoggingConfig `yaml:"logging" json:"logging"`
	PortConfig                  `yaml:",inline" json:",inline"`
	SSL                         SSLConfig        `yaml:"ssl" json:"ssl"`
	Plugins                     DataLayerPlugins `yaml:"plugins" json:"plugins"`
	MaximumFullFileCount        uint16           `yaml:"maximum_full_file_count" json:"maximum_full_file_count"`
	GroupFilesByStore           bool             `yaml:"group_files_by_store" json:"group_files_by_store"` // False is default, so non-ptr is fine here
}

// DataLayerPlugins Settings for data layer plugins
type DataLayerPlugins struct {
	// @TODO
	UnknownFields map[string]any `yaml:",inline" json:",inline"`
}

// SimulatorConfig settings for simulator
type SimulatorConfig struct {
	UnknownFields  map[string]any `yaml:",inline" json:",inline"`
	AutoFarm       bool           `yaml:"auto_farm" json:"auto_farm"`
	KeyFingerprint int            `yaml:"key_fingerprint" json:"key_fingerprint"`
	FarmingAddress string         `yaml:"farming_address" json:"farming_address"`
	PlotDirectory  string         `yaml:"plot_directory" json:"plot_directory"`
	UseCurrentTime bool           `yaml:"use_current_time" json:"use_current_time"`
}
