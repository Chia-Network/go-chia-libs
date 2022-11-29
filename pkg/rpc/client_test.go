package rpc

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func setup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewTLSServer(mux)

	u, err := url.Parse(server.URL)
	if err != nil {
		log.Fatal(err)
	}
	p, err := strconv.ParseUint(u.Port(), 10, 16)
	if err != nil {
		t.Fatal(err)
	}
	client, err := NewClient(ConnectionModeHTTP,
		WithDaemonPort(uint16(p)),
		WithNodePort(uint16(p)),
		WithFarmerPort(uint16(p)),
		WithHarvesterPort(uint16(p)),
		WithWalletPort(uint16(p)),
		WithCrawlerPort(uint16(p)))
	if err != nil {
		t.Fatal(err)
	}

	return mux, server, client
}

func teardown(server *httptest.Server) {
	server.Close()
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
