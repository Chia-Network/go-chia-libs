package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

func TestChiaConfig_SetFieldByPath(t *testing.T) {
	defaultConfig, err := config.LoadDefaultConfig()
	assert.NoError(t, err)
	// Make assertions about the default state, to ensure the assumed initial values are correct
	assert.Equal(t, uint16(8444), defaultConfig.FullNode.Port)
	assert.Equal(t, uint16(8555), defaultConfig.FullNode.RPCPort)
	assert.NotNil(t, defaultConfig.NetworkOverrides.Constants["mainnet"])
	assert.Equal(t, defaultConfig.NetworkOverrides.Constants["mainnet"].DifficultyConstantFactor, types.Uint128{})
	assert.Equal(t, *defaultConfig.SelectedNetwork, "mainnet")
	assert.Equal(t, defaultConfig.Logging.LogLevel, "WARNING")

	err = defaultConfig.SetFieldByPath([]string{"full_node", "port"}, "1234")
	assert.NoError(t, err)
	assert.Equal(t, uint16(1234), defaultConfig.FullNode.Port)

	err = defaultConfig.SetFieldByPath([]string{"full_node", "rpc_port"}, "5678")
	assert.NoError(t, err)
	assert.Equal(t, uint16(5678), defaultConfig.FullNode.RPCPort)

	err = defaultConfig.SetFieldByPath([]string{"network_overrides", "constants", "mainnet", "DIFFICULTY_CONSTANT_FACTOR"}, "44445555")
	assert.NoError(t, err)
	assert.NotNil(t, defaultConfig.NetworkOverrides.Constants["mainnet"])
	assert.Equal(t, types.Uint128From64(44445555), defaultConfig.NetworkOverrides.Constants["mainnet"].DifficultyConstantFactor)

	err = defaultConfig.SetFieldByPath([]string{"selected_network"}, "unittestnet")
	assert.NoError(t, err)
	assert.Equal(t, *defaultConfig.SelectedNetwork, "unittestnet")
	// Ensure all the other selected networks got set too
	assert.Equal(t, *defaultConfig.Seeder.SelectedNetwork, "unittestnet")
	assert.Equal(t, *defaultConfig.Harvester.SelectedNetwork, "unittestnet")
	assert.Equal(t, *defaultConfig.Pool.SelectedNetwork, "unittestnet")
	assert.Equal(t, *defaultConfig.Farmer.SelectedNetwork, "unittestnet")
	assert.Equal(t, *defaultConfig.Timelord.SelectedNetwork, "unittestnet")
	assert.Equal(t, *defaultConfig.FullNode.SelectedNetwork, "unittestnet")
	assert.Equal(t, *defaultConfig.UI.SelectedNetwork, "unittestnet")
	assert.Equal(t, *defaultConfig.Introducer.SelectedNetwork, "unittestnet")
	assert.Equal(t, *defaultConfig.Wallet.SelectedNetwork, "unittestnet")
	assert.Equal(t, *defaultConfig.DataLayer.SelectedNetwork, "unittestnet")

	err = defaultConfig.SetFieldByPath([]string{"logging", "log_level"}, "INFO")
	assert.NoError(t, err)
	assert.Equal(t, defaultConfig.Logging.LogLevel, "INFO")
}

// TestChiaConfig_SetFieldByPath_FullObjects Tests that we can pass in and correctly parse a whole section of config
// as json or yaml and that it gets set properly
func TestChiaConfig_SetFieldByPath_FullObjects(t *testing.T) {
	defaultConfig, err := config.LoadDefaultConfig()
	assert.NoError(t, err)
	// Make assertions about the default state, to ensure the assumed initial values are correct
	assert.Equal(t, uint16(8444), defaultConfig.FullNode.Port)
	assert.Equal(t, uint16(8555), defaultConfig.FullNode.RPCPort)
	assert.NotNil(t, defaultConfig.NetworkOverrides.Constants["mainnet"])
	assert.Equal(t, defaultConfig.NetworkOverrides.Constants["mainnet"].DifficultyConstantFactor, types.Uint128{})
	assert.Equal(t, *defaultConfig.SelectedNetwork, "mainnet")
	assert.Equal(t, defaultConfig.Logging.LogLevel, "WARNING")

	// Test passing in json blobs
	err = defaultConfig.SetFieldByPath([]string{"network_overrides", "constants"}, `{"jsonnet":{"DIFFICULTY_CONSTANT_FACTOR":44445555,"GENESIS_CHALLENGE":e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806}}`)
	assert.NoError(t, err)
	assert.NotNil(t, defaultConfig.NetworkOverrides.Constants["jsonnet"])
	assert.Equal(t, types.Uint128From64(44445555), defaultConfig.NetworkOverrides.Constants["jsonnet"].DifficultyConstantFactor)
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)
	// Ensure this applied to the other areas of config as well
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.Seeder.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.Harvester.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.Pool.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.Farmer.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.Timelord.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.FullNode.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.UI.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.Introducer.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)
	assert.Equal(t, "e739da31bcc4ab1767d9f1ca99eb3cec765fb3b3508f82e090374d5913d24806", defaultConfig.Wallet.NetworkOverrides.Constants["jsonnet"].GenesisChallenge)

	// Test passing in yaml blobs
	err = defaultConfig.SetFieldByPath([]string{"network_overrides", "constants"}, `yamlnet:
  DIFFICULTY_CONSTANT_FACTOR: 44445555
  GENESIS_CHALLENGE: 9eb3cec765fb3b3508f82e090374d5913d24806e739da31bcc4ab1767d9f1ca9`)
	assert.NoError(t, err)
	assert.NotNil(t, defaultConfig.NetworkOverrides.Constants["yamlnet"])
	assert.Equal(t, types.Uint128From64(44445555), defaultConfig.NetworkOverrides.Constants["yamlnet"].DifficultyConstantFactor)
	assert.Equal(t, "9eb3cec765fb3b3508f82e090374d5913d24806e739da31bcc4ab1767d9f1ca9", defaultConfig.NetworkOverrides.Constants["yamlnet"].GenesisChallenge)
}

func TestChiaConfig_FillValuesFromEnvironment(t *testing.T) {
	defaultConfig, err := config.LoadDefaultConfig()
	assert.NoError(t, err)
	// Make assertions about the default state, to ensure the assumed initial values are correct
	assert.Equal(t, uint16(8444), defaultConfig.FullNode.Port)
	assert.Equal(t, uint16(8555), defaultConfig.FullNode.RPCPort)
	assert.NotNil(t, defaultConfig.NetworkOverrides.Constants["mainnet"])
	assert.Equal(t, defaultConfig.NetworkOverrides.Constants["mainnet"].DifficultyConstantFactor, types.Uint128{})
	assert.Equal(t, *defaultConfig.SelectedNetwork, "mainnet")
	assert.Equal(t, defaultConfig.Logging.LogLevel, "WARNING")

	assert.NoError(t, os.Setenv("chia.full_node.port", "1234"))
	assert.NoError(t, os.Setenv("chia__full_node__rpc_port", "5678"))
	assert.NoError(t, os.Setenv("chia.network_overrides.constants.mainnet.DIFFICULTY_CONSTANT_FACTOR", "44445555"))
	assert.NoError(t, os.Setenv("chia.selected_network", "unittestnet"))
	assert.NoError(t, os.Setenv("chia__logging__log_level", "INFO"))

	assert.NoError(t, defaultConfig.FillValuesFromEnvironment())
	assert.Equal(t, uint16(1234), defaultConfig.FullNode.Port)
	assert.Equal(t, uint16(5678), defaultConfig.FullNode.RPCPort)
	assert.Equal(t, types.Uint128From64(44445555), defaultConfig.NetworkOverrides.Constants["mainnet"].DifficultyConstantFactor)
	assert.Equal(t, *defaultConfig.SelectedNetwork, "unittestnet")
	assert.Equal(t, defaultConfig.Logging.LogLevel, "INFO")
}

func TestChiaConfig_ParsePathsAndValuesFromStrings(t *testing.T) {
	// A mix of paths with and without prefixes with both separators
	strings := []string{
		"chia.full_node.port=8444",
		"chia__full_node__db_sync=auto",
		"full_node.db_readers=4",
		"full_node__database_path=testing.db",
	}

	// Test that both strings with prefixes are matched with requirePrefix
	result := config.ParsePathsAndValuesFromStrings(strings, true)
	assert.Len(t, result, 2)
	assert.Contains(t, result, "full_node.port")
	assert.Equal(t, []string{"full_node", "port"}, result["full_node.port"].Path)
	assert.Equal(t, "8444", result["full_node.port"].Value)
	assert.Contains(t, result, "full_node__db_sync")
	assert.Equal(t, []string{"full_node", "db_sync"}, result["full_node__db_sync"].Path)
	assert.Equal(t, "auto", result["full_node__db_sync"].Value)
	assert.NotContains(t, result, "full_node.db_readers")
	assert.NotContains(t, result, "full_node__database_path")

	// Test that both strings with prefixes are matched with requirePrefix
	result = config.ParsePathsAndValuesFromStrings(strings, false)
	assert.Len(t, result, 4) // 4 because it won't strip chia prefix if its found
	assert.Contains(t, result, "chia.full_node.port")
	assert.Equal(t, []string{"chia", "full_node", "port"}, result["chia.full_node.port"].Path)
	assert.Equal(t, "8444", result["chia.full_node.port"].Value)
	assert.Contains(t, result, "chia__full_node__db_sync")
	assert.Equal(t, []string{"chia", "full_node", "db_sync"}, result["chia__full_node__db_sync"].Path)
	assert.Equal(t, "auto", result["chia__full_node__db_sync"].Value)
	assert.Contains(t, result, "full_node.db_readers")
	assert.Equal(t, []string{"full_node", "db_readers"}, result["full_node.db_readers"].Path)
	assert.Equal(t, "4", result["full_node.db_readers"].Value)
	assert.Contains(t, result, "full_node__database_path")
	assert.Equal(t, []string{"full_node", "database_path"}, result["full_node__database_path"].Path)
	assert.Equal(t, "testing.db", result["full_node__database_path"].Value)
}

func TestChiaConfig_ParsePathsFromStrings(t *testing.T) {
	// A mix of paths with and without prefixes with both separators
	strings := []string{
		"chia.full_node.port",
		"chia__full_node__db_sync",
		"full_node.db_readers",
		"full_node__database_path",
	}

	// Test that both strings with prefixes are matched with requirePrefix
	result := config.ParsePathsFromStrings(strings, true)
	assert.Len(t, result, 2)
	assert.Contains(t, result, "full_node.port")
	assert.Equal(t, []string{"full_node", "port"}, result["full_node.port"])
	assert.Contains(t, result, "full_node__db_sync")
	assert.Equal(t, []string{"full_node", "db_sync"}, result["full_node__db_sync"])
	assert.NotContains(t, result, "full_node.db_readers")
	assert.NotContains(t, result, "full_node__database_path")

	// Test that both strings with prefixes are matched with requirePrefix
	result = config.ParsePathsFromStrings(strings, false)
	assert.Len(t, result, 4) // 4 because it won't strip chia prefix if its found
	assert.Contains(t, result, "chia.full_node.port")
	assert.Equal(t, []string{"chia", "full_node", "port"}, result["chia.full_node.port"])
	assert.Contains(t, result, "chia__full_node__db_sync")
	assert.Equal(t, []string{"chia", "full_node", "db_sync"}, result["chia__full_node__db_sync"])
	assert.Contains(t, result, "full_node.db_readers")
	assert.Equal(t, []string{"full_node", "db_readers"}, result["full_node.db_readers"])
	assert.Contains(t, result, "full_node__database_path")
	assert.Equal(t, []string{"full_node", "database_path"}, result["full_node__database_path"])
}
