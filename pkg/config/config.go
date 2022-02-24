package config

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// ChiaConfig the chia config.yaml
type ChiaConfig struct {
	DaemonPort uint16          `yaml:"daemon_port"`
	DaemonSSL  SSLConfig       `yaml:"daemon_ssl"`
	Farmer     FarmerConfig    `yaml:"farmer"`
	FullNode   FullNodeConfig  `yaml:"full_node"`
	Harvester  HarvesterConfig `yaml:"harvester"`
	Wallet     WalletConfig    `yaml:"wallet"`
	Seeder     SeederConfig    `yaml:"seeder"`
}

// FarmerConfig farmer configuration section
type FarmerConfig struct {
	PortConfig `yaml:",inline"`
	SSL        SSLConfig `yaml:"ssl"`
}

// FullNodeConfig full node configuration section
type FullNodeConfig struct {
	PortConfig `yaml:",inline"`
	SSL        SSLConfig `yaml:"ssl"`
}

// HarvesterConfig harvester configuration section
type HarvesterConfig struct {
	PortConfig `yaml:",inline"`
	SSL        SSLConfig `yaml:"ssl"`
}

// WalletConfig wallet configuration section
type WalletConfig struct {
	PortConfig `yaml:",inline"`
	SSL        SSLConfig `yaml:"ssl"`
}

// SeederConfig seeder configuration section
type SeederConfig struct {
	CrawlerConfig CrawlerConfig `yaml:"crawler"`
}

// CrawlerConfig is the subsection of the seeder config specific to the crawler
type CrawlerConfig struct {
	PortConfig `yaml:",inline"`
	SSL        SSLConfig `yaml:"ssl"`
}

// PortConfig common port settings found in many sections of the config
type PortConfig struct {
	Port    uint16 `yaml:"port"`
	RPCPort uint16 `yaml:"rpc_port"`
}

// SSLConfig common ssl settings found in many sections of the config
type SSLConfig struct {
	PrivateCRT string `yaml:"private_crt"`
	PrivateKey string `yaml:"private_key"`
	PublicCRT  string `yaml:"public_crt"`
	PublicKey  string `yaml:"public_key"`
}

// GetChiaConfig returns a struct containing the config.yaml values
func GetChiaConfig() (*ChiaConfig, error) {
	rootPath, err := GetChiaRootPath()
	if err != nil {
		return nil, err
	}

	configPath := path.Join(rootPath, "config", "config.yaml")
	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found")
	}

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	config := &ChiaConfig{}

	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
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

	root := path.Join(home, ".chia", "mainnet")

	return root, nil
}
