package config

import (
	// Need to embed the default config into the library
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
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

	return commonLoad(configBytes, rootPath)
}

// LoadDefaultConfig loads the initial-config bundled in go-chia-libs
func LoadDefaultConfig() (*ChiaConfig, error) {
	rootPath, err := GetChiaRootPath()
	if err != nil {
		return nil, err
	}
	return commonLoad(initialConfig, rootPath)
}

func commonLoad(configBytes []byte, rootPath string) (*ChiaConfig, error) {
	config := &ChiaConfig{}

	err := yaml.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}

	config.ChiaRoot = rootPath
	config.fillDatabasePath()

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
	c.FullNode.DatabasePath = strings.Replace(c.FullNode.DatabasePath, "CHALLENGE", c.FullNode.SelectedNetwork, 1)
}
