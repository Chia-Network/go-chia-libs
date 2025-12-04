package tls

import (
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

var testOldCertPEMBlock = []byte(`-----BEGIN CERTIFICATE-----
MIIDyzCCArOgAwIBAgIUXIpxI5MoZQ65/vhc7DK/d5ymoMUwDQYJKoZIhvcNAQEL
BQAwRDENMAsGA1UECgwEQ2hpYTEQMA4GA1UEAwwHQ2hpYSBDQTEhMB8GA1UECwwY
T3JnYW5pYyBGYXJtaW5nIERpdmlzaW9uMB4XDTI1MTExOTIxMDg1MFoXDTM3MTIz
MTIxMDg1MFowRDENMAsGA1UECgwEQ2hpYTEQMA4GA1UEAwwHQ2hpYSBDQTEhMB8G
A1UECwwYT3JnYW5pYyBGYXJtaW5nIERpdmlzaW9uMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAzz/L219Zjb5CIKnUkpd2julGC+j3E97KUiuOalCH9wdq
gpJi9nBqLccwPCSFXFew6CNBIBM+CW2jT3UVwgzjdXJ7pgtu8gWj0NQ6NqSLiXV2
WbpZovfrVh3x7Z4bjPgI3ouWjyehUfmK1GPIld4BfUSQtPlUJ53+XT32GRizUy+b
0CcJ84jp1XvyZAMajYnclFRNNJSw9WXtTlMUu+Z1M4K7c4ZPwEqgEnCgRc0TCaXj
180vo7mCHJQoDiNSCRATwfH+kWxOOK/nePkq2t4mPSFaX8xAS4yILISIOWYn7sNg
dy9D6gGNFo2SZ0FR3x9hjUjYEV3cPqg3BmNE3DDynQIDAQABo4G0MIGxMA8GA1Ud
EwEB/wQFMAMBAf8wHQYDVR0OBBYEFD4KTvuce45Yfu6qWASlekSwUBZaMH8GA1Ud
IwR4MHaAFD4KTvuce45Yfu6qWASlekSwUBZaoUikRjBEMQ0wCwYDVQQKDARDaGlh
MRAwDgYDVQQDDAdDaGlhIENBMSEwHwYDVQQLDBhPcmdhbmljIEZhcm1pbmcgRGl2
aXNpb26CFFyKcSOTKGUOuf74XOwyv3ecpqDFMA0GCSqGSIb3DQEBCwUAA4IBAQAI
vF2RAVyN5H5BU/OK6icEw8SGGfoaqpSMCQRuSwz9qg3QLDQj8yd5FYs9VECuigdw
i+mVnycT3vAObx+KfsPFlyUpw3RJ0+F53aDy+tW6ykTU94GQeLF5LN528NYLo2vt
YZ/4cGW+wjf+a+qBYjd3bW6HI9QGdJ1/BQhfwxexy4T+x4cleB++/vUFj6ZM1WhM
XRRTC3QzFBOCmTY4y21Mxvh5dOwY9HxK2qvwiO45zRWQ8Z0V95y4rytRqiDbNvFG
Jbs9Iwm8nj6c0Gm6rLkFNZ6OaYhX/5KhX35xE7YYIdorKwi+EKqhDD1jAwDA3q6T
Z4wumksKcqmUZXw9+pYX
-----END CERTIFICATE-----`)

var testOldKeyPEMBlock = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzz/L219Zjb5CIKnUkpd2julGC+j3E97KUiuOalCH9wdqgpJi
9nBqLccwPCSFXFew6CNBIBM+CW2jT3UVwgzjdXJ7pgtu8gWj0NQ6NqSLiXV2WbpZ
ovfrVh3x7Z4bjPgI3ouWjyehUfmK1GPIld4BfUSQtPlUJ53+XT32GRizUy+b0CcJ
84jp1XvyZAMajYnclFRNNJSw9WXtTlMUu+Z1M4K7c4ZPwEqgEnCgRc0TCaXj180v
o7mCHJQoDiNSCRATwfH+kWxOOK/nePkq2t4mPSFaX8xAS4yILISIOWYn7sNgdy9D
6gGNFo2SZ0FR3x9hjUjYEV3cPqg3BmNE3DDynQIDAQABAoIBAGupS4BJdx8gEAAh
2VDRqAAzhHTZb8j9uoKXJ+NotEkKrDTqUMiOu0nOqOsFWdYPo9HjxoggFuEU+Hpl
a4kj4uF3OG6Yj+jgLypjpV4PeoFM6M9R9BCp07In2i7DLLK9gvYA85SoVLBd/tW4
hFH+Qy3M+ZNZ1nLCK4pKjtaYs0dpi5zLoVvpEcEem2O+aRpUPCZqkNwU0umATCfg
ZGfFzgXI/XPJr8Uy+LVZOFp3PXXHfnZZD9T5AjO/ViBeqbMFuWQ8BpVOqapNPKj8
xDY3ovw3uiAYPC7eLib3u/WoFelMc2OMX0QljLp5Y+FScFHAMxoco3AQdWSYvSQw
b5xZmg0CgYEA6zKASfrw3EtPthkLR5NBmesI4RbbY6iFVhS5loLbzTtStvsus8EI
6RQgLgAFF14H21YSHxb6dB1Mbo45BN83gmDpUvKPREslqD3YPMKFo5GXMmv+JhNo
5Y9fhiOEnxzLJGtBB1HeGmg5NXp9mr2Ch9u8w/slfuCHckbA9AYvdxMCgYEA4ZR5
zg73+UA1a6Pm93bLYZGj+hf7OaB/6Hiw9YxCBgDfWM9dJ48iz382nojT5ui0rClV
5YAo8UCLh01Np9AbBZHuBdYm9IziuKNzTeK31UW+Tvbz+dEx7+PlYQffNOhcIgd+
9SXjoZorQksImKdMGZld1lEReHuBawq92JQvtY8CgYEAtNwUws7xQLW5CjKf9d5K
5+1Q2qYU9sG0JsmxHQhrtZoUtRjahOe/zlvnkvf48ksgh43cSYQF/Bw7lhhPyGtN
6DhVs69KdB3FS2ajTbXXxjxCpEdfHDB4zW4+6ouNhD1ECTFgxBw0SuIye+lBhSiN
o6NZuOr7nmFSRpIZ9ox7G3kCgYA4pvxMNtAqJekEpn4cChab42LGLX2nhFp7PMxc
bqQqM8/j0vg3Nihs6isCd6SYKjstvZfX8m7V3/rquQxWp9oRdQvNJXJVGojaDBqq
JdU7V6+qzzSIufQLpjV2P+7br7trxGwrDx/y9vAETynShLmE+FJrv6Jems3u3xy8
psKwmwKBgG5uLzCyMvMB2KwI+f3np2LYVGG0Pl1jq6yNXSaBosAiF0y+IgUjtWY5
EejO8oPWcb9AbqgPtrWaiJi17KiKv4Oyba5+y36IEtyjolWt0AB6F3oDK0X+Etw8
j/xlvBNuzDL6gRJHQg1+d4dO8Lz54NDUbKW8jGl+N/7afGVpGmX9
-----END RSA PRIVATE KEY-----`)

func TestGenerateAllCerts_ProvidedCA(t *testing.T) {
	cert, err := ParsePemCertificate(testCertPEMBlock)
	require.NoError(t, err)
	if err != nil {
		t.Fatalf("Failed to parse certificate: %v", err)
	}

	// Decode the PEM-encoded private key
	privateKey, err := ParsePemKey(testKeyPEMBlock)
	require.NoError(t, err)

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

func TestGenerateAllCerts_ProvidedOldCA(t *testing.T) {
	cert, err := ParsePemCertificate(testOldCertPEMBlock)
	require.NoError(t, err)
	if err != nil {
		t.Fatalf("Failed to parse certificate: %v", err)
	}

	// Decode the PEM-encoded private key
	privateKey, err := ParsePemKey(testOldKeyPEMBlock)
	require.NoError(t, err)

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

func TestParsePemKeyNewFormat(t *testing.T) {
	key, err := ParsePemKey(testKeyPEMBlock)
	require.NoError(t, err)
	err = key.Validate()
	assert.NoError(t, err)
}

func TestParsePemKeyOldFormat(t *testing.T) {
	key, err := ParsePemKey(testOldKeyPEMBlock)
	require.NoError(t, err)
	err = key.Validate()
	assert.NoError(t, err)
}
