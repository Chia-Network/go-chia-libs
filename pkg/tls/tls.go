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

type fieldHelpers struct {
	assign      func(certs *ChiaCertificates, pair *CertificateKeyPair)
	fetch       func(certs *ChiaCertificates) *CertificateKeyPair
	certKeyBase string
}

var (
	privateNodes = map[string]fieldHelpers{
		"full_node": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PrivateFullNode = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PrivateFullNode },
			certKeyBase: "private_full_node",
		},
		"wallet": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PrivateWallet = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PrivateWallet },
			certKeyBase: "private_wallet",
		},
		"farmer": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PrivateFarmer = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PrivateFarmer },
			certKeyBase: "private_farmer",
		},
		"harvester": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PrivateHarvester = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PrivateHarvester },
			certKeyBase: "private_harvester",
		},
		"timelord": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PrivateTimelord = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PrivateTimelord },
			certKeyBase: "private_timelord",
		},
		"crawler": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PrivateCrawler = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PrivateCrawler },
			certKeyBase: "private_crawler",
		},
		"data_layer": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PrivateDatalayer = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PrivateDatalayer },
			certKeyBase: "private_data_layer",
		},
		"daemon": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PrivateDaemon = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PrivateDaemon },
			certKeyBase: "private_daemon",
		},
	}
	publicNodes = map[string]fieldHelpers{
		"full_node": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PublicFullNode = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PublicFullNode },
			certKeyBase: "public_full_node",
		},
		"wallet": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PublicWallet = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PublicWallet },
			certKeyBase: "public_wallet",
		},
		"farmer": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PublicFarmer = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PublicFarmer },
			certKeyBase: "public_farmer",
		},
		"introducer": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PublicIntroducer = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PublicIntroducer },
			certKeyBase: "public_introducer",
		},
		"timelord": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PublicTimelord = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PublicTimelord },
			certKeyBase: "public_timelord",
		},
		"data_layer": {
			assign:      func(c *ChiaCertificates, p *CertificateKeyPair) { c.PublicDatalayer = p },
			fetch:       func(c *ChiaCertificates) *CertificateKeyPair { return c.PublicDatalayer },
			certKeyBase: "public_data_layer",
		},
	}

	//go:embed chia_ca.crt
	chiaCACrtBytes []byte

	//go:embed chia_ca.key
	chiaCAKeyBytes []byte
)

// ChiaCertificates contains the data for all Chia TLS certificate-key pairs
type ChiaCertificates struct {
	PrivateCA        *CertificateKeyPair
	PrivateCrawler   *CertificateKeyPair
	PrivateDaemon    *CertificateKeyPair
	PrivateDatalayer *CertificateKeyPair
	PublicDatalayer  *CertificateKeyPair
	PrivateFarmer    *CertificateKeyPair
	PublicFarmer     *CertificateKeyPair
	PrivateFullNode  *CertificateKeyPair
	PublicFullNode   *CertificateKeyPair
	PrivateHarvester *CertificateKeyPair
	PublicIntroducer *CertificateKeyPair
	PrivateTimelord  *CertificateKeyPair
	PublicTimelord   *CertificateKeyPair
	PrivateWallet    *CertificateKeyPair
	PublicWallet     *CertificateKeyPair
}

// CertificateKeyPair represents a TLS certificate and its corresponding private key used for secure communications.
// This pair can be encoded to PEM with the EncodeCertAndKeyToPEM function.
type CertificateKeyPair struct {
	// CertificateDER contains the X.509 certificate in ASN.1 DER binary format.
	// This is the raw binary representation of the certificate, not PEM encoded.
	CertificateDER []byte

	// PrivateKey contains the RSA private key corresponding to the public key
	// embedded in the certificate.
	PrivateKey *rsa.PrivateKey
}

// GenerateAllCerts  generates the full set of required certs for chia blockchain
// If privateCACert and privateCAKey are both nil, a new private CA will be generated
func GenerateAllCerts(privateCACert *x509.Certificate, privateCAKey *rsa.PrivateKey) (*ChiaCertificates, error) {
	var chiaCerts = &ChiaCertificates{}

	if privateCACert == nil && privateCAKey == nil {
		// If privateCACert and privateCAKey are both nil, we will generate a new one
		var err error
		var privateCACertDER []byte
		privateCACertDER, privateCAKey, err = GenerateNewCA()
		if err != nil {
			return nil, fmt.Errorf("error creating private ca pair: %w", err)
		}
		privateCACertPEMBytes, _, err := EncodeCertAndKeyToPEM(privateCACertDER, privateCAKey)
		if err != nil {
			return nil, fmt.Errorf("error encoding private ca certificates: %w", err)
		}
		privateCACert, err = ParsePemCertificate(privateCACertPEMBytes)
		if err != nil {
			return nil, fmt.Errorf("error parsing generated private_ca.crt: %w", err)
		}
		chiaCerts.PrivateCA = &CertificateKeyPair{
			CertificateDER: privateCACertDER,
			PrivateKey:     privateCAKey,
		}
	} else if privateCACert == nil || privateCAKey == nil {
		// If only one of them is nil, we can't continue
		return nil, errors.New("you must provide the CA cert and key if providing a CA, or set both to nil and a new CA will be generated")
	} else {
		// Must have non-nil values for both, so ensure the cert and key match
		if !CertMatchesPrivateKey(privateCACert, privateCAKey) {
			return nil, errors.New("provided private CA Cert and Key do not match")
		}
		chiaCerts.PrivateCA = &CertificateKeyPair{
			CertificateDER: privateCACert.Raw,
			PrivateKey:     privateCAKey,
		}
	}

	// Parse public CA cert and key bytes
	chiaCACert, err := ParsePemCertificate(chiaCACrtBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing chia_ca.crt")
	}
	chiaCAKey, err := ParsePemKey(chiaCAKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing chia_ca.key")
	}

	// Create all Certificate-Key pairs from public CA
	for node, nodeData := range publicNodes {
		cert, key, err := GenerateCASignedCert(chiaCACert, chiaCAKey)
		if err != nil {
			return nil, fmt.Errorf("error generating public pair for %s: %w", node, err)
		}
		nodeData.assign(chiaCerts, &CertificateKeyPair{
			CertificateDER: cert,
			PrivateKey:     key,
		})
	}

	// Create all Certificate-Key pairs from private CA
	for node, nodeData := range privateNodes {
		cert, key, err := GenerateCASignedCert(privateCACert, privateCAKey)
		if err != nil {
			return nil, fmt.Errorf("error generating private pair for %s: %w", node, err)
		}
		nodeData.assign(chiaCerts, &CertificateKeyPair{
			CertificateDER: cert,
			PrivateKey:     key,
		})
	}

	return chiaCerts, nil
}

// GenerateAndWriteAllCerts generates the full set of required certs for chia blockchain and writes them to a given directory
// If privateCACert and privateCAKey are both nil, a new private CA will be generated
func GenerateAndWriteAllCerts(outDir string, privateCACert *x509.Certificate, privateCAKey *rsa.PrivateKey) error {
	// First, ensure that all output directories exist
	allNodes := make(map[string]bool)
	for k := range privateNodes {
		allNodes[k] = true
	}
	for k := range publicNodes {
		allNodes[k] = true
	}
	allNodes["ca"] = true
	for subdir := range allNodes {
		err := os.MkdirAll(path.Join(outDir, subdir), 0700)
		if err != nil {
			return fmt.Errorf("error making output directory for certs: %w", err)
		}
	}

	// Generate all the certificates
	allCerts, err := GenerateAllCerts(privateCACert, privateCAKey)
	if err != nil {
		return fmt.Errorf("error generating certificates: %w", err)
	}

	// Write the private CA cert/key
	_, _, err = WriteCertAndKey(allCerts.PrivateCA.CertificateDER, allCerts.PrivateCA.PrivateKey, path.Join(outDir, "ca", "private_ca"))
	if err != nil {
		return fmt.Errorf("error writing private ca: %w", err)
	}

	// Next, write the chia_ca cert/key
	err = os.WriteFile(path.Join(outDir, "ca", "chia_ca.crt"), chiaCACrtBytes, 0600)
	if err != nil {
		return fmt.Errorf("error copying chia_ca.crt: %w", err)
	}
	err = os.WriteFile(path.Join(outDir, "ca", "chia_ca.key"), chiaCAKeyBytes, 0600)
	if err != nil {
		return fmt.Errorf("error copying chia_ca.key: %w", err)
	}

	for node, nodeHelpers := range publicNodes {
		crtKey := nodeHelpers.fetch(allCerts)
		_, _, err = WriteCertAndKey(crtKey.CertificateDER, crtKey.PrivateKey, path.Join(outDir, node, nodeHelpers.certKeyBase))
		if err != nil {
			return fmt.Errorf("error writing public pair for %s: %w", node, err)
		}
	}

	for node, nodeHelpers := range privateNodes {
		crtKey := nodeHelpers.fetch(allCerts)
		_, _, err = WriteCertAndKey(crtKey.CertificateDER, crtKey.PrivateKey, path.Join(outDir, node, nodeHelpers.certKeyBase))
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
