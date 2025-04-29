package rpc

import (
	"fmt"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
	"github.com/samber/mo"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestGetKeysValues(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/get_keys_values", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("datalayer/get_keys_values.json"))
		if err != nil {
			return
		}
	})

	want := DatalayerGetKeysValuesResponse{
		Response: rpcinterface.Response{
			Success: true,
		},
		KeysValues: []struct {
			Atom  mo.Option[string] `json:"atom"`
			Hash  string            `json:"hash"`
			Key   string            `json:"key"`
			Value string            `json:"value"`
		}{
			{
				Atom:  mo.None[string](),
				Hash:  "0xc543f6377e3600563f3aa9f7a9e6ccba8379172352e277cdc175f6ab3017a567",
				Key:   "0x7631",
				Value: "0x62303239646333613166636361636464393361623131656462343936376639326663333962646261643037326161313930623133333136386231616662316462",
			},
		},
	}

	r, resp, err := client.DataLayerService.GetKeysValues(&DatalayerGetKeysValuesOptions{
		ID: "607b73c0f7c1edf42281509ac06a76c833e1e79e7bfc5b94b988f2d450ed4bbd",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}
