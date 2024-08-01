package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/chia-network/go-chia-libs/pkg/config"
)

// TestWonkySet_UnmarshalYAML ensures that this will unmarshal both empty list [] and dict {} into empty map[string]string
// And also ensures an actal !!set as it would show up in the yaml unmarshals correctly
func TestWonkySet_UnmarshalYAML(t *testing.T) {
	var yamlWithList = []byte(`
farmer:
  pool_public_keys: []
`)
	var yamlWithDict = []byte(`
farmer:
  pool_public_keys: {}
`)

	var yamlWithData = []byte(`
farmer:
  pool_public_keys: !!set
    abc123: null
    456xyz: null
`)

	cfg := &config.ChiaConfig{}
	err := yaml.Unmarshal(yamlWithList, cfg)
	assert.NoError(t, err)
	assert.Len(t, cfg.Farmer.PoolPublicKeys, 0)

	cfg = &config.ChiaConfig{}
	err = yaml.Unmarshal(yamlWithDict, cfg)
	assert.NoError(t, err)
	assert.Len(t, cfg.Farmer.PoolPublicKeys, 0)

	cfg = &config.ChiaConfig{}
	err = yaml.Unmarshal(yamlWithData, cfg)
	assert.NoError(t, err)
	assert.Len(t, cfg.Farmer.PoolPublicKeys, 2)
}
