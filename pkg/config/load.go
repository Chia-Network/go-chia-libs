package config

import (
	// Need to embed the default config into the library
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed initial-config.yml
var initialConfig []byte

// GetChiaConfig returns a struct containing the config.yaml values
func GetChiaConfig() (*ChiaConfig, error) {
	rootPath, err := GetChiaRootPath()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(rootPath, "config", "config.yaml")
	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("chia config file not found at %s. Ensure CHIA_ROOT is set to the correct chia root", configPath)
	}

	return LoadConfigAtRoot(configPath, rootPath)
}

// LoadConfigAtRoot loads the given configPath into a ChiaConfig
// chiaRoot is required to fill the database paths in the config
func LoadConfigAtRoot(configPath, rootPath string) (*ChiaConfig, error) {
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := LoadFromBytes(configBytes, rootPath)
	if err != nil {
		return nil, err
	}
	cfg.configPath = configPath
	return cfg, nil
}

// LoadDefaultConfig loads the initial-config bundled in go-chia-libs
func LoadDefaultConfig() (*ChiaConfig, error) {
	rootPath, err := GetChiaRootPath()
	if err != nil {
		return nil, err
	}
	return LoadFromBytes(initialConfig, rootPath)
}

// LoadFromBytes loads a config from bytes.
// You will typically want to use GetChiaConfig(), LoadConfigAtRoot(), or LoadDefaultConfig() instead
func LoadFromBytes(configBytes []byte, rootPath string) (*ChiaConfig, error) {
	config := &ChiaConfig{}

	configBytes = FixBackCompat(configBytes)
	err := yaml.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}

	config.ChiaRoot = rootPath
	config.fillMissingDefaults()
	config.fillDatabasePath()
	config.dealWithAnchors()

	return config, nil
}

// FixBackCompat fixes any back compat issues with configs that might have been loaded by old versions of this package
func FixBackCompat(configBytes []byte) []byte {
	// soa serial number incorrectly had string as a type for a while, and ended up quoted as a result
	// remove the quotes since it's supposed to be a number
	regex := regexp.MustCompile(`serial_number:\s*["'](\d+)["']`)
	configBytes = regex.ReplaceAll(configBytes, []byte(`serial_number: $1`))

	return configBytes
}

// Save saves the config at the path it was loaded from originally
func (c *ChiaConfig) Save() error {
	if c.configPath == "" {
		return errors.New("configPath is not set on config. Save can only be used with a config that was loaded by this library. Try SavePath(path) instead")
	}

	return c.SavePath(c.configPath)
}

// SavePath saves the config at the given path
func (c *ChiaConfig) SavePath(configPath string) error {
	out, err := c.SaveBytes()
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, out, 0655)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}

// SaveBytes marshalls the config back down to bytes
func (c *ChiaConfig) SaveBytes() ([]byte, error) {
	c.unfillDatabasePath()
	out, err := yaml.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("error marshalling config: %w", err)
	}
	return out, nil
}

// GetChiaRootPath returns the root path for the chia installation
func GetChiaRootPath() (string, error) {
	if root, ok := os.LookupEnv("CHIA_ROOT"); ok {
		return root, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	root := filepath.Join(home, ".chia", "mainnet")

	return root, nil
}

// GetFullPath returns the full path to a particular filename within CHIA_ROOT
func (c *ChiaConfig) GetFullPath(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join(c.ChiaRoot, filename)
}

func (c *ChiaConfig) fillDatabasePath() {
	if c.FullNode.SelectedNetwork != nil {
		c.FullNode.DatabasePath = strings.Replace(c.FullNode.DatabasePath, "CHALLENGE", *c.FullNode.SelectedNetwork, 1)
		c.DataLayer.DatabasePath = strings.Replace(c.DataLayer.DatabasePath, "CHALLENGE", *c.FullNode.SelectedNetwork, 1)
		c.DataLayer.MerkleBlobsPath = strings.Replace(c.DataLayer.MerkleBlobsPath, "CHALLENGE", *c.FullNode.SelectedNetwork, 1)
		c.DataLayer.KeyValueBlobsPath = strings.Replace(c.DataLayer.KeyValueBlobsPath, "CHALLENGE", *c.FullNode.SelectedNetwork, 1)
	}
}

func (c *ChiaConfig) unfillDatabasePath() {
	if c.FullNode.SelectedNetwork != nil {
		c.FullNode.DatabasePath = strings.Replace(c.FullNode.DatabasePath, *c.FullNode.SelectedNetwork, "CHALLENGE", 1)
		c.DataLayer.DatabasePath = strings.Replace(c.DataLayer.DatabasePath, *c.FullNode.SelectedNetwork, "CHALLENGE", 1)
		c.DataLayer.MerkleBlobsPath = strings.Replace(c.DataLayer.MerkleBlobsPath, *c.FullNode.SelectedNetwork, "CHALLENGE", 1)
		c.DataLayer.KeyValueBlobsPath = strings.Replace(c.DataLayer.KeyValueBlobsPath, *c.FullNode.SelectedNetwork, "CHALLENGE", 1)
	}
}

// dealWithAnchors swaps out the distinct sections of the config with pointers to a shared instance
// When loading the config, the anchor definition in the initial-config is the canonical version. The rest will be
// changed to point back to that instance
// .self_hostname
//
//	.harvester.farmer_peers[0].host
//	.farmer.full_node_peers[0].host
//	.timelord_launcher.host
//	.timelord.vdf_clients.ip[0]
//	.timelord.full_node_peers[0].host
//	.timelord.vdf_server.host
//	.ui.daemon_host
//	.introducer.host
//	.full_node_peers[0].host
//
// .selected_network
//
//	.seeder.selected_network
//	.harvester.selected_network
//	.pool.selected_network
//	.farmer.selected_network
//	.timelord.selected_network
//	.full_node.selected_network
//	.ui.selected_network
//	.introducer.selected_network
//	.wallet.selected_network
//	.data_layer.selected_network
//
// .network_overrides
//
//	.seeder.network_overrides
//	.harvester.network_overrides
//	.pool.network_overrides
//	.farmer.network_overrides
//	.timelord.network_overrides
//	.full_node.network_overrides
//	.ui.network_overrides
//	.introducer.network_overrides
//	.wallet.network_overrides
//
// .logging
//
//	.seeder.logging
//	.harvester.logging
//	.pool.logging
//	.farmer.logging
//	.timelord_launcher.logging
//	.timelord.logging
//	.full_node.logging
//	.ui.logging
//	.introducer.logging
//	.wallet.logging
//	.data_layer.logging
func (c *ChiaConfig) dealWithAnchors() {
	// For now, just doing network_overrides and selected_network
	// The rest have some edge case usefulness in not being treated like anchors always
	if c.NetworkOverrides == nil {
		c.NetworkOverrides = &NetworkOverrides{}
	}
	c.Seeder.NetworkOverrides = c.NetworkOverrides
	c.Harvester.NetworkOverrides = c.NetworkOverrides
	c.Pool.NetworkOverrides = c.NetworkOverrides
	c.Farmer.NetworkOverrides = c.NetworkOverrides
	c.Timelord.NetworkOverrides = c.NetworkOverrides
	c.FullNode.NetworkOverrides = c.NetworkOverrides
	c.UI.NetworkOverrides = c.NetworkOverrides
	c.Introducer.NetworkOverrides = c.NetworkOverrides
	c.Wallet.NetworkOverrides = c.NetworkOverrides
	if c.Solver != nil {
		c.Solver.NetworkOverrides = c.NetworkOverrides
	}

	if c.SelectedNetwork == nil {
		mainnet := "mainnet"
		c.SelectedNetwork = &mainnet
	}
	c.Seeder.SelectedNetwork = c.SelectedNetwork
	c.Harvester.SelectedNetwork = c.SelectedNetwork
	c.Pool.SelectedNetwork = c.SelectedNetwork
	c.Farmer.SelectedNetwork = c.SelectedNetwork
	c.Timelord.SelectedNetwork = c.SelectedNetwork
	c.FullNode.SelectedNetwork = c.SelectedNetwork
	c.UI.SelectedNetwork = c.SelectedNetwork
	c.Introducer.SelectedNetwork = c.SelectedNetwork
	c.Wallet.SelectedNetwork = c.SelectedNetwork
	c.DataLayer.SelectedNetwork = c.SelectedNetwork
	if c.Solver != nil {
		c.Solver.SelectedNetwork = c.SelectedNetwork
	}

	if c.Logging == nil {
		c.Logging = &LoggingConfig{}
	}
	c.Seeder.Logging = c.Logging
	c.Harvester.Logging = c.Logging
	c.Pool.Logging = c.Logging
	c.Farmer.Logging = c.Logging
	c.TimelordLauncher.Logging = c.Logging
	c.Timelord.Logging = c.Logging
	c.FullNode.Logging = c.Logging
	c.UI.Logging = c.Logging
	c.Introducer.Logging = c.Logging
	c.Wallet.Logging = c.Logging
	c.DataLayer.Logging = c.Logging
	if c.Solver != nil {
		c.Solver.Logging = c.Logging
	}
}

func (c *ChiaConfig) fillMissingDefaults() {
	if c.Solver == nil {
		defaultConfig, err := LoadDefaultConfig()
		if err == nil {
			c.Solver = defaultConfig.Solver
		}
	}

	if len(c.Farmer.SolverPeers) == 0 {
		c.Farmer.SolverPeers = []Peer{{Host: "localhost", Port: 8666}}
	}
	if c.DataLayer.MerkleBlobsPath == "" {
		c.DataLayer.MerkleBlobsPath = "data_layer/db/merkle_blobs_CHALLENGE"
	}
	if c.DataLayer.KeyValueBlobsPath == "" {
		c.DataLayer.KeyValueBlobsPath = "data_layer/db/key_value_blobs_CHALLENGE"
	}
	if c.DataLayer.MerkleBlobsCacheSize == 0 {
		c.DataLayer.MerkleBlobsCacheSize = 1
	}
}
