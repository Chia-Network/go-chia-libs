package tls

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	// Need to embed the default config into the library
	_ "embed"
	"encoding/pem"
	"errors"
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
// If privateCACert and privateCAKey are both nil, a new private CA will be generated
func GenerateAllCerts(outDir string, privateCACert *x509.Certificate, privateCAKey *rsa.PrivateKey) error {
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

	if privateCACert == nil && privateCAKey == nil {
		// If privateCACert and privateCAKey are both nil, we will generate a new one
		privateCACertDER, privateCAKey, err := GenerateNewCA()
		if err != nil {
			return fmt.Errorf("error creating private ca pair: %w", err)
		}
		privateCACertBytes, _, err := WriteCertAndKey(privateCACertDER, privateCAKey, path.Join(outDir, "ca", "private_ca"))
		if err != nil {
			return fmt.Errorf("error writing private ca: %w", err)
		}
		privateCACert, err = ParsePemCertificate(privateCACertBytes)
		if err != nil {
			return fmt.Errorf("error parsing generated private_ca.crt: %w", err)
		}
	} else if privateCACert == nil || privateCAKey == nil {
		// If only one of them is nil, we can't continue
		return errors.New("you must provide the CA cert and key if providing a CA, or set both to nil and a new CA will be generated")
	} else {
		// Must have non-nil values for both, so ensure the cert and key match
		if !CertMatchesPrivateKey(privateCACert, privateCAKey) {
			return errors.New("provided private CA Cert and Key do not match")
		}
	}

	for _, node := range publicNodeNames {
		cert, key, err := GenerateCASignedCert(chiaCACert, chiaCAKey)
		if err != nil {
			return fmt.Errorf("error generating public pair for %s: %w", node, err)
		}
		_, _, err = WriteCertAndKey(cert, key, path.Join(outDir, node, fmt.Sprintf("public_%s", node)))
		if err != nil {
			return fmt.Errorf("error writing public pair for %s: %w", node, err)
		}
	}

	for _, node := range privateNodeNames {
		cert, key, err := GenerateCASignedCert(privateCACert, privateCAKey)
		if err != nil {
			return fmt.Errorf("error generating private pair for %s: %w", node, err)
		}
		_, _, err = WriteCertAndKey(cert, key, path.Join(outDir, node, fmt.Sprintf("private_%s", node)))
		if err != nil {
			return fmt.Errorf("error writing private pair for %s: %w", node, err)
		}
	}

	return nil
}

// GetChiaCACertAndKey returns the cert and key bytes for chia_ca.crt and chia_ca.key
func GetChiaCACertAndKey() ([]byte, []byte) {
	return chiaCACrtBytes, chiaCAKeyBytes
}

// CertMatchesPrivateKey tests to make the sure cert and private key match
func CertMatchesPrivateKey(cert *x509.Certificate, privateKey *rsa.PrivateKey) bool {
	publicKey := &privateKey.PublicKey

	certPublicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		fmt.Println("Certificate public key is not of type RSA")
		return false
	}

	if publicKey.N.Cmp(certPublicKey.N) == 0 && publicKey.E == certPublicKey.E {
		return true
	}
	return false
}

// ParsePemCertificate parses a certificate
func ParsePemCertificate(certPem []byte) (*x509.Certificate, error) {
	// Load CA certificate
	caCertBlock, rest := pem.Decode(certPem)
	if caCertBlock == nil || caCertBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("failed to decode CA certificate PEM")
	}
	if len(rest) != 0 {
		return nil, fmt.Errorf("cert file had extra data at the end")
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
	if caKeyBlock == nil || caKeyBlock.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode CA private key PEM")
	}
	if len(rest) != 0 {
		return nil, fmt.Errorf("key file had extra data at the end")
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

// EncodeCertAndKeyToPEM encodes the cert and key to PEM
func EncodeCertAndKeyToPEM(certDER []byte, certKey *rsa.PrivateKey) ([]byte, []byte, error) {
	certPemBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	keyBytes, err := x509.MarshalPKCS8PrivateKey(certKey)
	if err != nil {
		return nil, nil, fmt.Errorf("error encoding private key to PKCS8: %w", err)
	}
	keyPemBytes := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes})

	return certPemBytes, keyPemBytes, nil
}

// WriteCertAndKey Returns the written cert bytes, key bytes, and error
func WriteCertAndKey(certDER []byte, certKey *rsa.PrivateKey, certKeyBase string) ([]byte, []byte, error) {
	certPemBytes, keyPemBytes, err := EncodeCertAndKeyToPEM(certDER, certKey)
	if err != nil {
		return nil, nil, fmt.Errorf("error encoding certificates: %w", err)
	}

	// Write the new certificate to file
	certOut := fmt.Sprintf("%s.crt", certKeyBase)
	if err := os.WriteFile(certOut, certPemBytes, 0600); err != nil {
		return nil, nil, fmt.Errorf("failed to write cert PEM: %w", err)
	}

	// Write the new private key to file in PKCS#8 format
	keyOut := fmt.Sprintf("%s.key", certKeyBase)
	if err := os.WriteFile(keyOut, keyPemBytes, 0600); err != nil {
		return nil, nil, fmt.Errorf("failed to write key PEM: %w", err)
	}

	return certPemBytes, keyPemBytes, nil
}

// GenerateNewCA generates a new CA
func GenerateNewCA() ([]byte, *rsa.PrivateKey, error) {
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

	return certDER, privateKey, nil
}

// GenerateCASignedCert generates a new key/cert signed by the given CA
func GenerateCASignedCert(caCert *x509.Certificate, caKey *rsa.PrivateKey) ([]byte, *rsa.PrivateKey, error) {
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

	return certDER, certKey, nil
}
