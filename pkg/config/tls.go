package config

import (
	"crypto/tls"
	"path"
)

// LoadPrivateKeyPair loads the private key pair for the SSLConfig
func (s *SSLConfig) LoadPrivateKeyPair() (*tls.Certificate, error) {
	rootPath, err := GetChiaRootPath()
	if err != nil {
		return nil, err
	}

	pair, err := tls.LoadX509KeyPair(path.Join(rootPath, s.PrivateCRT), path.Join(rootPath, s.PrivateKey))
	return &pair, err
}

// LoadPublicKeyPair loads the public key pair for the SSLConfig
func (s *SSLConfig) LoadPublicKeyPair() (*tls.Certificate, error) {
	rootPath, err := GetChiaRootPath()
	if err != nil {
		return nil, err
	}

	pair, err := tls.LoadX509KeyPair(path.Join(rootPath, s.PublicCRT), path.Join(rootPath, s.PublicKey))
	return &pair, err
}
