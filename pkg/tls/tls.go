package tls

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	// Need to embed the default config into the library
	_ "embed"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path"
	"time"
)

var (
	privateNodeNames = []string{
		"full_node",
		"wallet",
		"farmer",
		"harvester",
		"timelord",
		"crawler",
		"data_layer",
		"daemon",
	}
	publicNodeNames = []string{
		"full_node",
		"wallet",
		"farmer",
		"introducer",
		"timelord",
		"data_layer",
	}

	//go:embed chia_ca.crt
	chiaCACrtBytes []byte

	//go:embed chia_ca.key
	chiaCAKeyBytes []byte
)

// GenerateAllCerts generates the full set of required certs for chia blockchain
func GenerateAllCerts(outDir string) error {
	// First, ensure that all output directories exist
	allNodes := append(privateNodeNames, publicNodeNames...)
	for _, subdir := range append(allNodes, "ca") {
		err := os.MkdirAll(path.Join(outDir, subdir), 0700)
		if err != nil {
			return fmt.Errorf("error making output directory for certs: %w", err)
		}
	}

	// Next, copy the chia_ca cert/key
	err := os.WriteFile(path.Join(outDir, "ca", "chia_ca.crt"), chiaCACrtBytes, 0600)
	if err != nil {
		return fmt.Errorf("error copying chia_ca.crt: %w", err)
	}
	err = os.WriteFile(path.Join(outDir, "ca", "chia_ca.key"), chiaCAKeyBytes, 0600)
	if err != nil {
		return fmt.Errorf("error copying chia_ca.key: %w", err)
	}

	chiaCACert, err := ParsePemCertificate(chiaCACrtBytes)
	if err != nil {
		return fmt.Errorf("error parsing chia_ca.crt")
	}

	chiaCAKey, err := ParsePemKey(chiaCAKeyBytes)
	if err != nil {
		return fmt.Errorf("error parsing chia_ca.key")
	}

	privateCACertBytes, privateCAKeyBytes, err := GenerateNewCA(path.Join(outDir, "ca", "private_ca"))
	if err != nil {
		return fmt.Errorf("error creating private ca pair: %w", err)
	}
	privateCACert, err := ParsePemCertificate(privateCACertBytes)
	if err != nil {
		return fmt.Errorf("error parsing generated private_ca.crt: %w", err)
	}
	privateCAKey, err := ParsePemKey(privateCAKeyBytes)
	if err != nil {
		return fmt.Errorf("error parsing generated private_ca.key: %w", err)
	}

	for _, node := range publicNodeNames {
		_, _, err = GenerateCASignedCert(chiaCACert, chiaCAKey, path.Join(outDir, node, fmt.Sprintf("public_%s", node)))
		if err != nil {
			return fmt.Errorf("error generating public pair for %s: %w", node, err)
		}
	}

	for _, node := range privateNodeNames {
		_, _, err = GenerateCASignedCert(privateCACert, privateCAKey, path.Join(outDir, node, fmt.Sprintf("private_%s", node)))
		if err != nil {
			return fmt.Errorf("error generating private pair for %s: %w", node, err)
		}
	}

	return nil
}

// ParsePemCertificate parses a certificate
func ParsePemCertificate(certPem []byte) (*x509.Certificate, error) {
	// Load CA certificate
	caCertBlock, rest := pem.Decode(certPem)
	if len(rest) != 0 {
		return nil, fmt.Errorf("cert file had extra data at the end")
	}
	if caCertBlock == nil || caCertBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("failed to decode CA certificate PEM")
	}
	caCert, err := x509.ParseCertificate(caCertBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CA certificate: %v", err)
	}

	return caCert, nil
}

// ParsePemKey parses a key
func ParsePemKey(keyPem []byte) (*rsa.PrivateKey, error) {
	// Load CA private key
	caKeyBlock, rest := pem.Decode(keyPem)
	if len(rest) != 0 {
		return nil, fmt.Errorf("cert file had extra data at the end")
	}
	if caKeyBlock == nil || caKeyBlock.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode CA private key PEM")
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(caKeyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CA private key: %v", err)
	}

	caKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("unexpected key type: %T", parsedKey)
	}

	return caKey, nil
}

// WriteCertAndKey Returns the written cert bytes, key bytes, and error
func WriteCertAndKey(certDER []byte, certKey *rsa.PrivateKey, certKeyBase string) ([]byte, []byte, error) {
	// Write the new certificate to file
	certOut := fmt.Sprintf("%s.crt", certKeyBase)
	certPemBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	if err := os.WriteFile(certOut, certPemBytes, 0600); err != nil {
		return nil, nil, fmt.Errorf("failed to write cert PEM: %v", err)
	}

	// Marshal private key to PKCS#8
	keyBytes, err := x509.MarshalPKCS8PrivateKey(certKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal private key to PKCS#8: %v", err)
	}

	// Write the new private key to file in PKCS#8 format
	keyOut := fmt.Sprintf("%s.key", certKeyBase)
	keyPemBytes := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes})
	if err := os.WriteFile(keyOut, keyPemBytes, 0600); err != nil {
		return nil, nil, fmt.Errorf("failed to write key PEM: %v", err)
	}

	return certPemBytes, keyPemBytes, nil
}

// GenerateNewCA generates a new CA
func GenerateNewCA(certKeyBase string) ([]byte, []byte, error) {
	// Generate a new RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Create new certificate template
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate serial number: %v", err)
	}

	// Define the certificate template
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization:       []string{"Chia"},
			OrganizationalUnit: []string{"Organic Farming Division"},
			CommonName:         "Chia CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // 10 years
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Create the self-signed certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	return WriteCertAndKey(certDER, privateKey, certKeyBase)
}

// GenerateCASignedCert generates a new key/cert signed by the given CA
func GenerateCASignedCert(caCert *x509.Certificate, caKey *rsa.PrivateKey, certKeyBase string) ([]byte, []byte, error) {
	// Generate new private key
	certKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	// Create new certificate template
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate serial number: %v", err)
	}
	certTemplate := x509.Certificate{
		Subject: pkix.Name{
			CommonName:         "Chia",
			Organization:       []string{"Chia"},
			OrganizationalUnit: []string{"Organic Farming Division"},
		},
		SerialNumber:          serialNumber,
		NotBefore:             time.Now().Add(-24 * time.Hour),
		NotAfter:              time.Date(2100, 8, 2, 0, 0, 0, 0, time.UTC),
		SubjectKeyId:          []byte{1, 2, 3, 4, 6},
		BasicConstraintsValid: true,
		DNSNames:              []string{"chia.net"},
	}

	// Sign the new certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &certTemplate, caCert, &certKey.PublicKey, caKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create certificate: %v", err)
	}

	return WriteCertAndKey(certDER, certKey, certKeyBase)
}
