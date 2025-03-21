package tls

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

var testCertPEMBlock = []byte(`-----BEGIN CERTIFICATE-----
MIIDVTCCAj2gAwIBAgIRAKo6h63J+hpNFV3CpKReWW8wDQYJKoZIhvcNAQELBQAw
RDENMAsGA1UEChMEQ2hpYTEhMB8GA1UECxMYT3JnYW5pYyBGYXJtaW5nIERpdmlz
aW9uMRAwDgYDVQQDEwdDaGlhIENBMB4XDTI1MDMyMTAwMzAyOFoXDTM1MDMyMTAw
MzAyOFowRDENMAsGA1UEChMEQ2hpYTEhMB8GA1UECxMYT3JnYW5pYyBGYXJtaW5n
IERpdmlzaW9uMRAwDgYDVQQDEwdDaGlhIENBMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAn0SIrAkL6tJH8HKMHVJlrx3hCogh2EBoECETv+I9k1/7Rlq+
1vLK4MV2U+ei2rif18YVXM1gqOZFMnwtgYn9bu0L+aUC4fWYLIoQjaHW+RRC2+yU
XjbmCR+qoYLa0628Kjmlrxq6zIu066sbn9pUOWI2C/AKO0vzD1bl3A5Qixyojl1o
fXkBXW1Zo0Xbdx+dVzMBp17EDTB9vmhfnBseFdE1+OZZnmXbPKfwOxnPR4nvBNQu
cYzPDlcEvjNXImG2Qo2fzy5HMyMkKO1tPgNQ+yGADMyYLiGkfWjYlRgQbaXmuYAD
KZAOhhbW9rlfVJyZfzSmSqR7n2R8QWjM7kMpFwIDAQABo0IwQDAOBgNVHQ8BAf8E
BAMCAoQwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQU8B08DgU4J40PnA6q1p7E
rDL3JRcwDQYJKoZIhvcNAQELBQADggEBAEA6lILY4I8k8+h3qSGCav5wdZFWL2cv
9mUWqROxZwar3qaeD2FizzPguAES4tPPQ5poWA6UjTAV/8DA/YXKK4EwncMaDZLV
iB4NXBcWcmf9XQ/BCIsLsWgwAsNMf8gxnQWdtSKLpPr1aNqYlsqbQ6c25Ro6Nsqd
Op5+Ynd6Jyn+r+wnz4FG/zFz542RPwG2oIBKf+NOMnCKWBGsB4zpOVmLdimQNzp/
HQg1t7/8Xig9/RX0nWs+cj9+iRVa9cxveQKCL9M8CQXnROYkLk35MSlmmVeU0WWJ
R4l85xte60BXTq/nIFc2WRC78wTFzQc+TFCSgtTokCOtu7XpQJaRDWo=
-----END CERTIFICATE-----`)

var testKeyPEMBlock = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCfRIisCQvq0kfw
cowdUmWvHeEKiCHYQGgQIRO/4j2TX/tGWr7W8srgxXZT56LauJ/XxhVczWCo5kUy
fC2Bif1u7Qv5pQLh9ZgsihCNodb5FELb7JReNuYJH6qhgtrTrbwqOaWvGrrMi7Tr
qxuf2lQ5YjYL8Ao7S/MPVuXcDlCLHKiOXWh9eQFdbVmjRdt3H51XMwGnXsQNMH2+
aF+cGx4V0TX45lmeZds8p/A7Gc9Hie8E1C5xjM8OVwS+M1ciYbZCjZ/PLkczIyQo
7W0+A1D7IYAMzJguIaR9aNiVGBBtpea5gAMpkA6GFtb2uV9UnJl/NKZKpHufZHxB
aMzuQykXAgMBAAECggEAfr6RbSa93x98tHLT4jnCRfunLTRsiqWmqr9H8jne+rs1
QiXRHUmV/g3mPptl1F18hsBSG8otE/w8MRL1O9NOZcoq735Lrvo9IaS1y6BxbUKc
elvpLpjNs5EJvwJdlnr59ThvC8xfv4umbK18jFe5Evl/PTzHR60HPrvOrLKPkkP4
1UNiqGGWfIOxdFbPvBiOFbcVJwGHVJ3TE/x/lMrdapW59hJZ5Rtxx5xUDR6EkkJd
KOCI68zCiISUyEdfrKEfJPs/8LG9o9OgaxeBnGOtpxO0LyegC7tuuZ/TvuTaTQ0X
cgrVXNuaCJfGmqPzpdATMdtQJoRVyF0Kt6dog7DgsQKBgQDJWWC2G6KUHNz6Phsp
U3yYJSoKGWvPcZds5Fjfvy6eQEaOaU8aN+0XyWhgs/QetuOIbpfZid5ZgD6z0g9E
4/a6mlgKKTl7CGnanKFxTGG2cXcBE96MMtPubQ0UurTqP++O7ZaC4m6TUCdt2MIy
w+u6/ZLz/kFR73D3f8nRXTCwQwKBgQDKfydw1a2N0yRTdRvA/3CWzRIIosIGiHf0
Tf4WCdCtPaftJ7l9Y1LdfV6BePp3dIyGmCHCACWvneei+qte2LMAky3fplIvjqBA
pnWvBIrd+jZQgQFe6OEkUHiZ1RpPczXaxTwxekhpVTQay1og+3F0qLD4yS4Ae0nT
NgWC7dOwnQKBgQCTRHoF+ER7THkb1t0K5vNUXKpY5KsD+TMmBAY08KJqQNzaQJAI
vyr8oOVlBXniFSZqnWkXRU2J7NDvuQ5N9uZ5KXaHSAuwv0CdEr7KHXHCfU7rTNsT
dAGqe7x7kuvMAaN3yLKzXGY//Po5z7aKZt490EXxi9++zAC2JZM5PI3l/QKBgEkd
NjFwhZSyyufzXc0GrjFU5BEIK0ROm/ky++4bJySWIX7om/nhFfdxH+FhvBXLmD20
ymOQyAqr2gonth6t4ZvwiFy7YetX9RbCw7Uoz7csc9YHbmZFcZ06DQGGR1SuhaBz
HLPEskaOBB00lVtZTnLPwe5iPWDhIxvG4qCOnKOlAoGBALR/JAQysYGjkm8H7UM6
H7ij3Gb3zSd+ziUMa7WxDCe8kBA5delHjzy8m8UrhmijpMJIOKk/5EwPEyCUnREZ
5EFVojPksoBn2LY0ZrS59/xELtLd7Rs9kdKBOoUoBGRtHR5Dr0W9+yzTt3Zl3aJz
Rb7RT3n8oEZxwPDlbW3OAVBH
-----END PRIVATE KEY-----`)

func TestGenerateAllCerts_ProvidedCA(t *testing.T) {
	// Decode the PEM-encoded certificate to get DER format
	certBlock, _ := pem.Decode(testCertPEMBlock)
	if certBlock == nil {
		t.Fatal("Failed to decode PEM block containing certificate")
	}
	// Convert DER to x509 Certificate
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse certificate: %v", err)
	}

	// Decode the PEM-encoded private key
	keyBlock, _ := pem.Decode(testKeyPEMBlock)
	if keyBlock == nil {
		t.Fatal("Failed to decode PEM block containing private key")
	}
	// Parse the private key in PKCS#8 format
	parsedKey, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse private key: %v", err)
	}
	// Type assert to RSA private key
	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		t.Fatal("Parsed key is not an RSA private key")
	}

	allCerts, err := GenerateAllCerts(cert, privateKey)
	if err != nil {
		t.Fatalf("Failed to generate all certificates: %v", err)
	}

	// Test that private CA cert and private key match from the GenerateAllCerts output
	privateCACert, err := x509.ParseCertificate(allCerts.PrivateCA.CertificateDER)
	if err != nil {
		t.Fatalf("Failed to parse private CA certificate: %v", err)
	}
	if !CertMatchesPrivateKey(privateCACert, allCerts.PrivateCA.PrivateKey) {
		t.Fatal("Private CA certificate and private key do not match")
	}

	// Loop through all public certificate-key pairs and verify that the pair is not nil, that the key matches the certificate
	for _, nodeHelpers := range publicNodes {
		crtKey := nodeHelpers.fetch(allCerts)

		// Test pair is not nil
		if crtKey == nil || crtKey.PrivateKey == nil || crtKey.CertificateDER == nil || len(crtKey.CertificateDER) == 0 {
			t.Fatalf("Certificate pair, or one of the certificate or key is empty : %s", nodeHelpers.certKeyBase)
		}

		// Test pair matches
		cert, err := x509.ParseCertificate(crtKey.CertificateDER)
		if err != nil {
			t.Fatalf("Failed to parse certificate for %s: %v", nodeHelpers.certKeyBase, err)
		}
		if !CertMatchesPrivateKey(cert, crtKey.PrivateKey) {
			t.Fatalf("certificate and private key do not match for %s", nodeHelpers.certKeyBase)
		}
	}

	// Loop through all private certificate-key pairs and verify that the pair is not nil, that the key matches the certificate
	for _, nodeHelpers := range privateNodes {
		crtKey := nodeHelpers.fetch(allCerts)

		// Test pair is not nil
		if crtKey == nil || crtKey.PrivateKey == nil || crtKey.CertificateDER == nil || len(crtKey.CertificateDER) == 0 {
			t.Fatalf("Certificate pair, or one of the certificate or key is empty : %s", nodeHelpers.certKeyBase)
		}

		// Test pair matches
		cert, err := x509.ParseCertificate(crtKey.CertificateDER)
		if err != nil {
			t.Fatalf("Failed to parse certificate for %s: %v", nodeHelpers.certKeyBase, err)
		}
		if !CertMatchesPrivateKey(cert, crtKey.PrivateKey) {
			t.Fatalf("certificate and private key do not match for %s", nodeHelpers.certKeyBase)
		}
	}
}

func TestGenerateAllCerts_GeneratedCA(t *testing.T) {
	allCerts, err := GenerateAllCerts(nil, nil)
	if err != nil {
		t.Fatalf("Failed to generate all certificates: %v", err)
	}

	// Test that private CA cert and private key match from the GenerateAllCerts output
	privateCACert, err := x509.ParseCertificate(allCerts.PrivateCA.CertificateDER)
	if err != nil {
		t.Fatalf("Failed to parse private CA certificate: %v", err)
	}
	if !CertMatchesPrivateKey(privateCACert, allCerts.PrivateCA.PrivateKey) {
		t.Fatal("Private CA certificate and private key do not match")
	}

	// Loop through all public certificate-key pairs and verify that the pair is not nil, that the key matches the certificate
	for _, nodeHelpers := range publicNodes {
		crtKey := nodeHelpers.fetch(allCerts)

		// Test pair is not nil
		if crtKey == nil || crtKey.PrivateKey == nil || crtKey.CertificateDER == nil || len(crtKey.CertificateDER) == 0 {
			t.Fatalf("Certificate pair, or one of the certificate or key is empty : %s", nodeHelpers.certKeyBase)
		}

		// Test pair matches
		cert, err := x509.ParseCertificate(crtKey.CertificateDER)
		if err != nil {
			t.Fatalf("Failed to parse certificate for %s: %v", nodeHelpers.certKeyBase, err)
		}
		if !CertMatchesPrivateKey(cert, crtKey.PrivateKey) {
			t.Fatalf("certificate and private key do not match for %s", nodeHelpers.certKeyBase)
		}
	}

	// Loop through all private certificate-key pairs and verify that the pair is not nil, that the key matches the certificate
	for _, nodeHelpers := range privateNodes {
		crtKey := nodeHelpers.fetch(allCerts)

		// Test pair is not nil
		if crtKey == nil || crtKey.PrivateKey == nil || crtKey.CertificateDER == nil || len(crtKey.CertificateDER) == 0 {
			t.Fatalf("Certificate pair, or one of the certificate or key is empty : %s", nodeHelpers.certKeyBase)
		}

		// Test pair matches
		cert, err := x509.ParseCertificate(crtKey.CertificateDER)
		if err != nil {
			t.Fatalf("Failed to parse certificate for %s: %v", nodeHelpers.certKeyBase, err)
		}
		if !CertMatchesPrivateKey(cert, crtKey.PrivateKey) {
			t.Fatalf("certificate and private key do not match for %s", nodeHelpers.certKeyBase)
		}
	}
}
