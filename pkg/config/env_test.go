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
	assert.Equal(t, defaultConfig.SelectedNetwork, "mainnet")
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
	assert.Equal(t, defaultConfig.SelectedNetwork, "unittestnet")

	err = defaultConfig.SetFieldByPath([]string{"logging", "log_level"}, "INFO")
	assert.NoError(t, err)
	assert.Equal(t, defaultConfig.Logging.LogLevel, "INFO")
}

func TestChiaConfig_FillValuesFromEnvironment(t *testing.T) {
	defaultConfig, err := config.LoadDefaultConfig()
	assert.NoError(t, err)
	// Make assertions about the default state, to ensure the assumed initial values are correct
	assert.Equal(t, uint16(8444), defaultConfig.FullNode.Port)
	assert.Equal(t, uint16(8555), defaultConfig.FullNode.RPCPort)
	assert.NotNil(t, defaultConfig.NetworkOverrides.Constants["mainnet"])
	assert.Equal(t, defaultConfig.NetworkOverrides.Constants["mainnet"].DifficultyConstantFactor, types.Uint128{})
	assert.Equal(t, defaultConfig.SelectedNetwork, "mainnet")
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
	assert.Equal(t, defaultConfig.SelectedNetwork, "unittestnet")
	assert.Equal(t, defaultConfig.Logging.LogLevel, "INFO")
}
