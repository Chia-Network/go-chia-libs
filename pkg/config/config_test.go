package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/chia-network/go-chia-libs/pkg/config"
)

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
