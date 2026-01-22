package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/chia-network/go-chia-libs/pkg/config"
)

// TestLoggingConfigAnchors verifies that logging is serialized with anchors
func TestLoggingConfigAnchors(t *testing.T) {
	type testStruct struct {
		LoggingConfig1 *config.LoggingConfig `yaml:"logging1"`
		LoggingConfig2 *config.LoggingConfig `yaml:"logging2"`
	}
	loggingCfg := config.LoggingConfig{}
	testInstance := &testStruct{
		LoggingConfig1: &loggingCfg,
		LoggingConfig2: &loggingCfg,
	}

	expected := `logging1: &logging
    log_stdout: false
    log_backcompat: false
    log_filename: ""
    log_level: ""
    log_maxfilesrotation: 0
    log_maxbytesrotation: 0
    log_use_gzip: false
    log_syslog: false
    log_syslog_host: ""
    log_syslog_port: 0
logging2: *logging
`

	out, err := yaml.Marshal(testInstance)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(out))
}

// TestNetworkOverridesAnchors verifies that logging is serialized with anchors
func TestNetworkOverridesAnchors(t *testing.T) {
	type testStruct struct {
		Network1 *config.NetworkOverrides `yaml:"network_overrides1"`
		Network2 *config.NetworkOverrides `yaml:"network_overrides2"`
	}
	no := config.NetworkOverrides{
		Constants: map[string]config.NetworkConstants{
			"mainnet": {},
		},
		Config: map[string]config.NetworkConfig{
			"mainnet": {
				AddressPrefix:       "xch",
				DefaultFullNodePort: 8444,
			},
		},
	}
	testInstance := &testStruct{
		Network1: &no,
		Network2: &no,
	}

	expected := `network_overrides1: &network_overrides
    constants:
        mainnet:
            GENESIS_CHALLENGE: ""
            GENESIS_PRE_FARM_POOL_PUZZLE_HASH: ""
            GENESIS_PRE_FARM_FARMER_PUZZLE_HASH: ""
    config:
        mainnet:
            address_prefix: xch
            default_full_node_port: 8444
network_overrides2: *network_overrides
`

	out, err := yaml.Marshal(testInstance)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(out))
}

// TestLoggingConfigNoAnchor verifies that logging can be serialized without anchors when SetNoAnchor is used
func TestLoggingConfigNoAnchor(t *testing.T) {
	type testStruct struct {
		LoggingConfig1 *config.LoggingConfig `yaml:"logging1"`
		LoggingConfig2 *config.LoggingConfig `yaml:"logging2"`
	}
	loggingCfg := config.LoggingConfig{
		LogLevel: "INFO",
	}
	loggingCfg2 := loggingCfg.CopyWithoutAnchor()
	loggingCfg2.LogLevel = "DEBUG" // Modify the copy independently

	testInstance := &testStruct{
		LoggingConfig1: &loggingCfg,
		LoggingConfig2: loggingCfg2,
	}

	out, err := yaml.Marshal(testInstance)
	assert.NoError(t, err)

	// Verify that logging1 has an anchor
	assert.Contains(t, string(out), "logging1: &logging")
	// Verify that logging2 does NOT use an anchor (should be a full definition, not *logging)
	assert.Contains(t, string(out), "logging2:")
	assert.NotContains(t, string(out), "logging2: *logging")
	// Verify that logging2 has the modified value
	assert.Contains(t, string(out), "log_level: DEBUG")
}

// TestNetworkOverridesNoAnchor verifies that network_overrides can be serialized without anchors when SetNoAnchor is used
func TestNetworkOverridesNoAnchor(t *testing.T) {
	type testStruct struct {
		Network1 *config.NetworkOverrides `yaml:"network_overrides1"`
		Network2 *config.NetworkOverrides `yaml:"network_overrides2"`
	}
	no := config.NetworkOverrides{
		Constants: map[string]config.NetworkConstants{
			"mainnet": {},
		},
		Config: map[string]config.NetworkConfig{
			"mainnet": {
				AddressPrefix:       "xch",
				DefaultFullNodePort: 8444,
			},
		},
	}
	no2 := no.CopyWithoutAnchor()
	// Modify the copy independently
	no2.Config["mainnet"] = config.NetworkConfig{
		AddressPrefix:       "txch",
		DefaultFullNodePort: 8445,
	}

	testInstance := &testStruct{
		Network1: &no,
		Network2: no2,
	}

	out, err := yaml.Marshal(testInstance)
	assert.NoError(t, err)

	// Verify that network_overrides1 has an anchor
	assert.Contains(t, string(out), "network_overrides1: &network_overrides")
	// Verify that network_overrides2 does NOT use an anchor (should be a full definition, not *network_overrides)
	assert.Contains(t, string(out), "network_overrides2:")
	assert.NotContains(t, string(out), "network_overrides2: *network_overrides")
	// Verify that network_overrides2 has the modified value
	assert.Contains(t, string(out), "address_prefix: txch")
}
