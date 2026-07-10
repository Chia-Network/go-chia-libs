package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/ptr"
)

func TestRateLimitsOptional(t *testing.T) {
	t.Run("default config omits rate_limits", func(t *testing.T) {
		cfg, err := config.LoadDefaultConfig()
		require.NoError(t, err)
		assert.Equal(t, uint8(0), cfg.RateLimits)

		out, err := cfg.SaveBytes()
		require.NoError(t, err)
		assert.NotContains(t, string(out), "rate_limits")
	})

	t.Run("marshal omits zero, includes non-zero", func(t *testing.T) {
		cfg := config.ChiaConfig{}

		out, err := yaml.Marshal(cfg)
		require.NoError(t, err)
		assert.NotContains(t, string(out), "rate_limits")

		cfg.RateLimits = 3
		out, err = yaml.Marshal(cfg)
		require.NoError(t, err)
		assert.Contains(t, string(out), "rate_limits: 3")
	})

	t.Run("unmarshal absent yields zero and is omitted on remarshal", func(t *testing.T) {
		cfg, err := config.LoadFromBytes([]byte(`inbound_rate_limit_percent: 100
outbound_rate_limit_percent: 30
`), "/test/path")
		require.NoError(t, err)
		assert.Equal(t, uint8(0), cfg.RateLimits)

		out, err := cfg.SaveBytes()
		require.NoError(t, err)
		assert.NotContains(t, string(out), "rate_limits")
	})

	t.Run("unmarshal present retains value on remarshal", func(t *testing.T) {
		cfg, err := config.LoadFromBytes([]byte(`inbound_rate_limit_percent: 100
outbound_rate_limit_percent: 30
rate_limits: 3
`), "/test/path")
		require.NoError(t, err)
		assert.Equal(t, uint8(3), cfg.RateLimits)

		out, err := cfg.SaveBytes()
		require.NoError(t, err)
		assert.Contains(t, string(out), "rate_limits: 3")
	})
}

func TestNFTAutoAddLimitPointer(t *testing.T) {
	t.Run("marshal omits nil, includes zero and non-zero", func(t *testing.T) {
		wallet := config.WalletConfig{}

		out, err := yaml.Marshal(wallet)
		require.NoError(t, err)
		assert.NotContains(t, string(out), "nft_auto_add_limit")

		wallet.NFTAutoAddLimit = ptr.Pointer(0)
		out, err = yaml.Marshal(wallet)
		require.NoError(t, err)
		assert.Contains(t, string(out), "nft_auto_add_limit: 0")

		wallet.NFTAutoAddLimit = ptr.Pointer(100)
		out, err = yaml.Marshal(wallet)
		require.NoError(t, err)
		assert.Contains(t, string(out), "nft_auto_add_limit: 100")
	})

	t.Run("unmarshal absent yields nil and is omitted on remarshal", func(t *testing.T) {
		var wallet config.WalletConfig
		require.NoError(t, yaml.Unmarshal([]byte(`db_sync: auto`), &wallet))
		assert.Nil(t, wallet.NFTAutoAddLimit)

		out, err := yaml.Marshal(wallet)
		require.NoError(t, err)
		assert.NotContains(t, string(out), "nft_auto_add_limit")
	})

	t.Run("unmarshal zero yields non-nil pointer to zero and is retained", func(t *testing.T) {
		var wallet config.WalletConfig
		require.NoError(t, yaml.Unmarshal([]byte(`nft_auto_add_limit: 0`), &wallet))
		require.NotNil(t, wallet.NFTAutoAddLimit)
		assert.Equal(t, 0, *wallet.NFTAutoAddLimit)

		out, err := yaml.Marshal(wallet)
		require.NoError(t, err)
		assert.Contains(t, string(out), "nft_auto_add_limit: 0")
	})

	t.Run("unmarshal non-zero retains value on remarshal", func(t *testing.T) {
		var wallet config.WalletConfig
		require.NoError(t, yaml.Unmarshal([]byte(`nft_auto_add_limit: 100`), &wallet))
		require.NotNil(t, wallet.NFTAutoAddLimit)
		assert.Equal(t, 100, *wallet.NFTAutoAddLimit)

		out, err := yaml.Marshal(wallet)
		require.NoError(t, err)
		assert.Contains(t, string(out), "nft_auto_add_limit: 100")
	})
}

func TestBlockCreationPointer(t *testing.T) {
	t.Run("default config has nil BlockCreation", func(t *testing.T) {
		cfg, err := config.LoadDefaultConfig()
		require.NoError(t, err)
		assert.Nil(t, cfg.FullNode.BlockCreation)
	})

	t.Run("marshal omits nil, includes zero and non-zero", func(t *testing.T) {
		fn := config.FullNodeConfig{}

		out, err := yaml.Marshal(fn)
		require.NoError(t, err)
		assert.NotContains(t, string(out), "block_creation")

		fn.BlockCreation = ptr.Pointer(int64(0))
		out, err = yaml.Marshal(fn)
		require.NoError(t, err)
		assert.Contains(t, string(out), "block_creation: 0")

		fn.BlockCreation = ptr.Pointer(int64(1))
		out, err = yaml.Marshal(fn)
		require.NoError(t, err)
		assert.Contains(t, string(out), "block_creation: 1")
	})

	t.Run("unmarshal absent yields nil", func(t *testing.T) {
		var fn config.FullNodeConfig
		require.NoError(t, yaml.Unmarshal([]byte(`db_sync: auto`), &fn))
		assert.Nil(t, fn.BlockCreation)
	})

	t.Run("unmarshal zero yields non-nil pointer to zero", func(t *testing.T) {
		var fn config.FullNodeConfig
		require.NoError(t, yaml.Unmarshal([]byte(`block_creation: 0`), &fn))
		require.NotNil(t, fn.BlockCreation)
		assert.Equal(t, int64(0), *fn.BlockCreation)
	})

	t.Run("unmarshal one yields non-nil pointer to one", func(t *testing.T) {
		var fn config.FullNodeConfig
		require.NoError(t, yaml.Unmarshal([]byte(`block_creation: 1`), &fn))
		require.NotNil(t, fn.BlockCreation)
		assert.Equal(t, int64(1), *fn.BlockCreation)
	})
}

func TestNetworkConstantsOptionalPointersOmitted(t *testing.T) {
	nc := config.NetworkConstants{
		GenesisChallenge:               "",
		GenesisPreFarmPoolPuzzleHash:   "",
		GenesisPreFarmFarmerPuzzleHash: "",
	}

	out, err := yaml.Marshal(nc)
	assert.NoError(t, err)
	outStr := string(out)

	for _, field := range []string{
		"NUMBER_ZERO_BITS_PLOT_FILTER_V1",
		"NUMBER_ZERO_BITS_PLOT_FILTER_V2",
		"HARD_FORK_HEIGHT",
		"HARD_FORK2_HEIGHT",
		"SOFT_FORK4_HEIGHT",
		"SOFT_FORK5_HEIGHT",
		"SOFT_FORK6_HEIGHT",
		"SOFT_FORK8_HEIGHT",
		"PLOT_FILTER_128_HEIGHT",
		"PLOT_FILTER_64_HEIGHT",
		"PLOT_FILTER_32_HEIGHT",
		"PLOT_V1_PHASE_OUT_EPOCH_BITS",
		"MIN_PLOT_STRENGTH",
		"MAX_PLOT_STRENGTH",
		"QUALITY_PROOF_SCAN_FILTER",
		"PLOT_FILTER_V2_FIRST_ADJUSTMENT_HEIGHT",
		"PLOT_FILTER_V2_SECOND_ADJUSTMENT_HEIGHT",
		"PLOT_FILTER_V2_THIRD_ADJUSTMENT_HEIGHT",
	} {
		assert.NotContains(t, outStr, field)
	}
}

func TestNetworkConstantsOptionalPointersIncludedWhenSet(t *testing.T) {
	uint8Zero := func() *uint8 {
		val := uint8(0)
		return &val
	}
	uint32Zero := func() *uint32 {
		val := uint32(0)
		return &val
	}

	nc := config.NetworkConstants{
		GenesisChallenge:                   "",
		GenesisPreFarmPoolPuzzleHash:       "",
		GenesisPreFarmFarmerPuzzleHash:     "",
		NumberZeroBitsPlotFilterV1:         uint8Zero(),
		NumberZeroBitsPlotFilterV2:         uint8Zero(),
		HardForkHeight:                     uint32Zero(),
		HardFork2Height:                    uint32Zero(),
		SoftFork4Height:                    uint32Zero(),
		SoftFork5Height:                    uint32Zero(),
		SoftFork6Height:                    uint32Zero(),
		SoftFork8Height:                    uint32Zero(),
		PlotFilter128Height:                uint32Zero(),
		PlotFilter64Height:                 uint32Zero(),
		PlotFilter32Height:                 uint32Zero(),
		PlotV1PhaseOutEpochBits:            uint8Zero(),
		MinPlotStrength:                    uint8Zero(),
		MaxPlotStrength:                    uint8Zero(),
		QualityProofScanFilter:             uint8Zero(),
		PlotFilterV2FirstAdjustmentHeight:  uint32Zero(),
		PlotFilterV2SecondAdjustmentHeight: uint32Zero(),
		PlotFilterV2ThirdAdjustmentHeight:  uint32Zero(),
	}

	out, err := yaml.Marshal(nc)
	assert.NoError(t, err)
	outStr := string(out)

	for _, field := range []string{
		"NUMBER_ZERO_BITS_PLOT_FILTER_V1: 0",
		"NUMBER_ZERO_BITS_PLOT_FILTER_V2: 0",
		"HARD_FORK_HEIGHT: 0",
		"HARD_FORK2_HEIGHT: 0",
		"SOFT_FORK4_HEIGHT: 0",
		"SOFT_FORK5_HEIGHT: 0",
		"SOFT_FORK6_HEIGHT: 0",
		"SOFT_FORK8_HEIGHT: 0",
		"PLOT_FILTER_128_HEIGHT: 0",
		"PLOT_FILTER_64_HEIGHT: 0",
		"PLOT_FILTER_32_HEIGHT: 0",
		"PLOT_V1_PHASE_OUT_EPOCH_BITS: 0",
		"MIN_PLOT_STRENGTH: 0",
		"MAX_PLOT_STRENGTH: 0",
		"QUALITY_PROOF_SCAN_FILTER: 0",
		"PLOT_FILTER_V2_FIRST_ADJUSTMENT_HEIGHT: 0",
		"PLOT_FILTER_V2_SECOND_ADJUSTMENT_HEIGHT: 0",
		"PLOT_FILTER_V2_THIRD_ADJUSTMENT_HEIGHT: 0",
	} {
		assert.Contains(t, outStr, field)
	}
}
