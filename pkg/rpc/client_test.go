package rpc

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/types"
	"github.com/stretchr/testify/require"
)

var (
	tmpDir      string
	crtFilename = "host.crt"
	keyFilename = "host.key"
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

	// Get temp directory used for this test back
	tmpDir, err = generateSSLFiles()
	if err != nil {
		t.Fatal(err)
	}

	// Set port config to the port of this test server
	portConf := config.PortConfig{
		RPCPort: uint16(p),
	}
	// Set SSL config for the fake cert and key we just made
	sslConf := config.SSLConfig{
		PrivateCRT: crtFilename,
		PrivateKey: keyFilename,
	}

	client, err := NewClient(ConnectionModeHTTP,
		WithManualConfig(config.ChiaConfig{
			ChiaRoot:   tmpDir,
			DaemonPort: portConf.RPCPort,
			DaemonSSL:  sslConf,
			FullNode: config.FullNodeConfig{
				PortConfig: portConf,
				SSL:        sslConf,
			},
			Farmer: config.FarmerConfig{
				PortConfig: portConf,
				SSL:        sslConf,
			},
			Harvester: config.HarvesterConfig{
				PortConfig: portConf,
				SSL:        sslConf,
			},
			Wallet: config.WalletConfig{
				PortConfig: portConf,
				SSL:        sslConf,
			},
			Seeder: config.SeederConfig{
				CrawlerConfig: config.CrawlerConfig{
					PortConfig: portConf,
					SSL:        sslConf,
				},
			},
		}))
	if err != nil {
		t.Fatal(err)
	}

	return mux, server, client
}

func generateSSLFiles() (string, error) {
	dir, err := os.MkdirTemp("", "*-chia")
	if err != nil {
		return "", err
	}

	key, crt, err := generateSSL()
	if err != nil {
		return dir, err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s", dir, keyFilename), key.Bytes(), 0755)
	if err != nil {
		return dir, err
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s", dir, crtFilename), crt.Bytes(), 0755)
	if err != nil {
		return dir, err
	}

	return dir, nil
}

func generateSSL() (*bytes.Buffer, *bytes.Buffer, error) {
	crtTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{"Testing, Inc."},
			Country:       []string{"US"},
			Province:      []string{"Missouri"},
			Locality:      []string{"St Louis"},
			StreetAddress: []string{"111 St Louis St"},
			PostalCode:    []string{"69420"},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}
	keyPEM := new(bytes.Buffer)
	err = pem.Encode(keyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	if err != nil {
		return nil, nil, err
	}

	crt, err := x509.CreateCertificate(rand.Reader, crtTemplate, crtTemplate, &key.PublicKey, key)
	if err != nil {
		return nil, nil, err
	}
	crtPEM := new(bytes.Buffer)
	err = pem.Encode(crtPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: crt,
	})
	if err != nil {
		return nil, nil, err
	}

	return keyPEM, crtPEM, nil
}

func teardown(server *httptest.Server) {
	server.Close()
	err := os.RemoveAll(tmpDir)
	if err != nil {
		panic(err)
	}
}

func fixture(path string) string {
	b, err := os.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func getBytesFromHexString(t *testing.T, hex string) types.Bytes {
	b, err := types.BytesFromHexString(hex)
	require.NoError(t, err)
	return b
}

func getBytes32FromHexString(t *testing.T, hex string) types.Bytes32 {
	b, err := types.Bytes32FromHexString(hex)
	require.NoError(t, err)
	return b
}

func ExampleNewClient() {
	client, err := NewClient(ConnectionModeHTTP, WithAutoConfig())
	if err != nil {
		panic(err)
	}

	state, _, err := client.FullNodeService.GetBlockchainState()
	if err != nil {
		panic(err)
	}

	if state.BlockchainState.IsPresent() {
		fmt.Println(state.BlockchainState.MustGet().Space)
	}
}
