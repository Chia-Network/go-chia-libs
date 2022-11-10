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

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup(t *testing.T) func() {
	mux = http.NewServeMux()
	server = httptest.NewTLSServer(mux)

	u, err := url.Parse(server.URL)
	if err != nil {
		log.Fatal(err)
	}
	p, err := strconv.ParseUint(u.Port(), 10, 16)
	if err != nil {
		t.Fatal(err)
	}
	client, err = NewClient(ConnectionModeHTTP, PortOptions{
		Wallet: uint16(p),
	})
	if err != nil {
		t.Fatal(err)
	}

	return func() {
		server.Close()
	}
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
