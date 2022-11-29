package rpc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/chia-network/go-chia-libs/pkg/config"
)

func setup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewTLSServer(mux)

	// Get port from server
	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	p, err := strconv.ParseUint(u.Port(), 10, 16)
	if err != nil {
		t.Fatal(err)
	}

	// Get user home dir for CHIA_ROOT
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	client, err := NewClient(ConnectionModeHTTP,
		WithManualConfig(config.ChiaConfig{
			ChiaRoot: fmt.Sprintf("%s/.chia/mainnet", home),
			Wallet: config.WalletConfig{
				PortConfig: config.PortConfig{
					RPCPort: uint16(p),
				},
				SSL: config.SSLConfig{
					PrivateCRT: "config/ssl/wallet/private_wallet.crt",
					PrivateKey: "config/ssl/wallet/private_wallet.key",
				},
			},
		}))
	if err != nil {
		t.Fatal(err)
	}

	return mux, server, client
}

func teardown(server *httptest.Server) {
	server.Close()
}

func fixture(path string) string {
	b, err := os.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
