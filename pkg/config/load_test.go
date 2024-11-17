package config_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/config"
)

func TestDealWithAnchors(t *testing.T) {
	cfg, err := config.LoadDefaultConfig()
	assert.NoError(t, err)

	mainnetConstants := cfg.NetworkOverrides.Constants["mainnet"]
	mainnetConstants.GenesisChallenge = "testing123"
	cfg.NetworkOverrides.Constants["mainnet"] = mainnetConstants

	assert.Equal(t, "testing123", cfg.Seeder.NetworkOverrides.Constants["mainnet"].GenesisChallenge)
	assert.Equal(t, "testing123", cfg.Harvester.NetworkOverrides.Constants["mainnet"].GenesisChallenge)
	assert.Equal(t, "testing123", cfg.Pool.NetworkOverrides.Constants["mainnet"].GenesisChallenge)
	assert.Equal(t, "testing123", cfg.Farmer.NetworkOverrides.Constants["mainnet"].GenesisChallenge)
	assert.Equal(t, "testing123", cfg.Timelord.NetworkOverrides.Constants["mainnet"].GenesisChallenge)
	assert.Equal(t, "testing123", cfg.FullNode.NetworkOverrides.Constants["mainnet"].GenesisChallenge)
	assert.Equal(t, "testing123", cfg.UI.NetworkOverrides.Constants["mainnet"].GenesisChallenge)
	assert.Equal(t, "testing123", cfg.Introducer.NetworkOverrides.Constants["mainnet"].GenesisChallenge)
	assert.Equal(t, "testing123", cfg.Wallet.NetworkOverrides.Constants["mainnet"].GenesisChallenge)

	*cfg.SelectedNetwork = "unittestnet"

	assert.Equal(t, "unittestnet", *cfg.Seeder.SelectedNetwork)
	assert.Equal(t, "unittestnet", *cfg.Harvester.SelectedNetwork)
	assert.Equal(t, "unittestnet", *cfg.Pool.SelectedNetwork)
	assert.Equal(t, "unittestnet", *cfg.Farmer.SelectedNetwork)
	assert.Equal(t, "unittestnet", *cfg.Timelord.SelectedNetwork)
	assert.Equal(t, "unittestnet", *cfg.FullNode.SelectedNetwork)
	assert.Equal(t, "unittestnet", *cfg.UI.SelectedNetwork)
	assert.Equal(t, "unittestnet", *cfg.Introducer.SelectedNetwork)
	assert.Equal(t, "unittestnet", *cfg.Wallet.SelectedNetwork)
	assert.Equal(t, "unittestnet", *cfg.DataLayer.SelectedNetwork)
}

func TestFillDatabasePath(t *testing.T) {
	def, err := config.LoadDefaultConfig()
	assert.NoError(t, err)
	assert.Equal(t, "db/blockchain_v2_mainnet.sqlite", def.FullNode.DatabasePath)

	tmpDir, err := os.MkdirTemp("", "testfs")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatalf("Error cleaning up test directory: %v", err)
		}
	}(tmpDir)

	tmpFilePath := tmpDir + "/tempfile.txt"
	err = def.SavePath(tmpFilePath)
	assert.NoError(t, err)
	assert.Equal(t, "db/blockchain_v2_CHALLENGE.sqlite", def.FullNode.DatabasePath)
}

func TestBackCompatSOASerialNumber(t *testing.T) {
	badConfig := `seeder:
  port: 8444
  other_peers_port: 8444
  dns_port: 53
  peer_connect_timeout: 2
  crawler_db_path: "crawler.db"
  bootstrap_peers:
    - "node.chia.net"
  minimum_height: 240000
  minimum_version_count: 100
  domain_name: "seeder.example.com."
  nameserver: "example.com."
  ttl: 300
  soa:
    rname: "hostmaster.example.com"
    serial_number: "1619105223"
    refresh: 10800
    retry: 10800
    expire: 604800
    minimum: 1800`

	cfg := config.FixBackCompat([]byte(badConfig))
	assert.True(t, strings.Contains(string(cfg), "serial_number: 1619105223"))
}
