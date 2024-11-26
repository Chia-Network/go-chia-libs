package rpc

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/require"

	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

func TestGetNFTs(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/nft_get_nfts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("wallet/nft_get_nfts.json"))
		if err != nil {
			return
		}
	})

	want := GetNFTsResponse{
		Response: rpcinterface.Response{
			Success: true,
		},
		WalletID: mo.Some[uint32](4),
		NFTList: mo.Some([]types.NFTInfo{
			{
				ChainInfo:          "((117 \"https://pixnio.com/free-images/2022/06/03/2022-06-03-07-49-44-900x1350.jpg\") (104 . 0x77923fb8d2556cfff3962ec2232259d71c11238c46b6ea2ff5c22c803f4847f2) (28021) (27765) (29550 . 1) (29556 . 1))",
				DataHash:           getBytesFromHexString(t, "0x77923fb8d2556cfff3962ec2232259d71c11238c46b6ea2ff5c22c803f4847f2"),
				DataUris:           []string{"https://pixnio.com/free-images/2022/06/03/2022-06-03-07-49-44-900x1350.jpg"},
				EditionNumber:      1,
				EditionTotal:       1,
				LauncherID:         getBytes32FromHexString(t, "0x13fe99139226c1b2d76351a11affe0887f4d93ddfbfe1d07e1c181dc8ae6dc5f"),
				LauncherPuzhash:    getBytes32FromHexString(t, "0xeff07522495060c066f66f32acc2a77e3a3e737aca8baea4d1a64ea4cdc13da9"),
				LicenseHash:        getBytesFromHexString(t, "0x630c2b0ddf2fb42de6ecff0e2965908a121fd884433153b61ee309cf6b19efd4"),
				LicenseURIs:        []string{"http://fakedata.com/fakenft/license.txt"},
				MetadataHash:       getBytesFromHexString(t, "0x630c2b0ddf2fb42de6ecff0e2965908a121fd884433153b61ee309cf6b19efd4"),
				MetadataURIs:       []string{"http://fakedata.com/fakenft/metadata.txt"},
				MintHeight:         1619167,
				MinterDid:          mo.Some(getBytes32FromHexString(t, "0xc42825559e5bda2bd31b03de428ea871a101ce0301a2a2f79ba5e833b84aa29d")),
				NftCoinID:          getBytes32FromHexString(t, "0x3ca34e2f816a78a87ee0f694f25843d0e0079038bde0e498cf9393b828845ace"),
				OffChainMetadata:   mo.None[string](),
				OwnerDid:           mo.None[types.Bytes32](),
				P2Address:          getBytes32FromHexString(t, "0x3bed5ecaeabea5616bd3ca9657317281f82ac6c277da9a80a296cbb71de12f8c"),
				PendingTransaction: false,
				RoyaltyPercentage:  mo.Some(uint16(0)),
				RoyaltyPuzzleHash:  mo.Some(getBytes32FromHexString(t, "0xc0c9d006a7b1b0aa6f8fd198e95948c35fbafbc2536656de94a76764f554cafc")),
				SupportsDid:        true,
				UpdaterPuzhash:     getBytes32FromHexString(t, "0xfe8a4b4e27a2e29a4d3fc7ce9d527adbcaccbab6ada3903ccf3ba9a769d2d78b"),
			},
		}),
	}

	r, resp, err := client.WalletService.GetNFTs(&GetNFTsOptions{
		WalletID: 4,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}

func TestNFTGetByDid(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/nft_get_by_did", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("wallet/nft_get_by_did.json"))
		if err != nil {
			return
		}
	})

	want := NFTGetByDidResponse{
		Response: rpcinterface.Response{
			Success: true,
		},
		WalletID: mo.Some[uint32](4),
	}

	r, resp, err := client.WalletService.NFTGetByDid(&NFTGetByDidOptions{
		DidID: getBytes32FromHexString(t, "0xc42825559e5bda2bd31b03de428ea871a101ce0301a2a2f79ba5e833b84aa29d"),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}
