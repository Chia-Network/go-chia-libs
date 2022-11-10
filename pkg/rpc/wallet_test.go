package rpc

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetNFTs(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mux.HandleFunc("/nft_get_nfts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("wallet/nft_get_nfts.json"))
	})

	r, _, err := client.WalletService.GetNFTs(&GetNFTsOptions{
		WalletID: 4,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", r)
	// ... other tests here
}
