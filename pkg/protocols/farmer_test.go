package protocols_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/protocols"
	"github.com/chia-network/go-chia-libs/pkg/streamable"
)

func TestNewSignagePoint(t *testing.T) {
	type Result struct {
		ChallengeHash     string
		ChallengeChainSP  string
		RewardChainSP     string
		Difficulty        uint64
		SubSlotIters      uint64
		SignagePointIndex uint8
		PeakHeight        uint32
		SPSourceDataNil   bool
	}
	type Test struct {
		Hex    string
		Result Result
	}

	tests := []Test{
		{
			"5c999e5ace11e34e1ee7441faee83ef296d48b5ca7585607a04dae3468532486ce3a1e781d279cb8d9e04dbb9c34ff37df1ec5a52b4326682c4d0f10d3f7e751d1cbc28e2a295b115c2f2b3d2099ca1b79f776384f294d126ac4b6562863c9aa0000000000002e0000000000228000002c004f58ae0100010300f1fd0f6d05218d619538d82d6002fb07d8a2022ff445fc8433f028d44b539e7c43178e232f4168e70ea7dd983c31287c434a069b38bfc5cfa2f264c75625333b8982f2906da92bf7ac48228d0efe294f7f5336e24f6eb16045fe505bde4cf8590100000042cbd82db6c7cf1ee006486393360c23e3bc618c1112dc61754edcb0a433c3e0ecc6a9d51e969364c662ea2722845df9e5c68fb661593464311695f42887b7233b029122af6634fc79c6cb719566888bac0db1aa04ddaf39d20d2bdb745929380201",
			Result{
				"5c999e5ace11e34e1ee7441faee83ef296d48b5ca7585607a04dae3468532486",
				"ce3a1e781d279cb8d9e04dbb9c34ff37df1ec5a52b4326682c4d0f10d3f7e751",
				"d1cbc28e2a295b115c2f2b3d2099ca1b79f776384f294d126ac4b6562863c9aa",
				11776,
				578813952,
				44,
				5200046,
				false,
			},
		},
		{
			"69171fb97a11a983e1c45f01393a0755e3b65016be6e92ea776ab8ee5b24b66a73165326a79bf653220e33573ef2cace709b31e8f9939cd9e7e7df7f009d08815c0f452f044d025bf0bb4d51744737e09723ab323731a9bbeed9c687e369ec170000000000002e0000000000228000001a004f511d0100010200661db3afacf0463587a92a3661bae6921424cbecfa6323975b102f6aa264df9e4b1c3b03bd678b5182e468d11b5a9134c8eb6534c144f9b7e56a31f3a63d8b1d8593acf1818abead01caf180c3bbf1027a32f56f441f52768f0a86ff66391a1c010000008d5d8a08d92ed62cf61288aa7043cb694cf94e49a96fe03f0d30cc01fbb10a7b5f10632cbab5fedd21e4c1c5bae92c7425fb73a5da71ec5fd426864622538b326da5bcf8f290971c1a95044e54d7fccbd9ccba2a494f6d5131221e18508c81030100",
			Result{
				"69171fb97a11a983e1c45f01393a0755e3b65016be6e92ea776ab8ee5b24b66a",
				"73165326a79bf653220e33573ef2cace709b31e8f9939cd9e7e7df7f009d0881",
				"5c0f452f044d025bf0bb4d51744737e09723ab323731a9bbeed9c687e369ec17",
				11776,
				578813952,
				26,
				5198109,
				false,
			},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			// Hex to bytes
			encodedBytes, err := hex.DecodeString(test.Hex)
			assert.NoError(t, err)

			rp := &protocols.NewSignagePoint{}

			err = streamable.Unmarshal(encodedBytes, rp)
			assert.NoError(t, err)

			assert.Equal(t, test.Result.ChallengeHash, hex.EncodeToString(rp.ChallengeHash[:]))
			assert.Equal(t, test.Result.ChallengeChainSP, hex.EncodeToString(rp.ChallengeChainSP[:]))
			assert.Equal(t, test.Result.RewardChainSP, hex.EncodeToString(rp.RewardChainSP[:]))
			assert.Equal(t, test.Result.Difficulty, rp.Difficulty)
			assert.Equal(t, test.Result.SubSlotIters, rp.SubSlotIters)
			assert.Equal(t, test.Result.SignagePointIndex, rp.SignagePointIndex)
			assert.Equal(t, test.Result.PeakHeight, rp.PeakHeight)
			if test.Result.SPSourceDataNil {
				assert.Nil(t, rp.SPSourceData)
			} else {
				assert.NotNil(t, rp.SPSourceData)
			}
		})
	}
	//hexStr := ""
	//
	//// Hex to bytes
	//encodedBytes, err := hex.DecodeString(hexStr)
	//assert.NoError(t, err)
	//
	//rp := &protocols.NewSignagePoint{}
	//
	//err = streamable.Unmarshal(encodedBytes, rp)
	//assert.NoError(t, err)

	//assert.Equal(t, "69171fb97a11a983e1c45f01393a0755e3b65016be6e92ea776ab8ee5b24b66a", hex.EncodeToString(rp.ChallengeHash[:]))
	//assert.Equal(t, "73165326a79bf653220e33573ef2cace709b31e8f9939cd9e7e7df7f009d0881", hex.EncodeToString(rp.ChallengeChainSP[:]))
	//assert.Equal(t, "5c0f452f044d025bf0bb4d51744737e09723ab323731a9bbeed9c687e369ec17", hex.EncodeToString(rp.RewardChainSP[:]))
	//assert.Equal(t, uint64(11776), rp.Difficulty)
	//assert.Equal(t, uint64(578813952), rp.SubSlotIters)
	//assert.Equal(t, uint8(26), rp.SignagePointIndex)
	//assert.Equal(t, uint32(5198109), rp.PeakHeight)
	//
	//assert.NotNil(t, rp.SPSourceData)

	// todo test SPSourceData
}
