package config

import (
	"crypto/tls"
	"path/filepath"
)

// LoadPrivateKeyPair loads the private key pair for the SSLConfig
func (s *SSLConfig) LoadPrivateKeyPair() (*tls.Certificate, error) {
	rootPath, err := GetChiaRootPath()
	if err != nil {
		return nil, err
	}

	pair, err := tls.LoadX509KeyPair(filepath.Join(rootPath, s.PrivateCRT), filepath.Join(rootPath, s.PrivateKey))
	return &pair, err
}

// LoadPublicKeyPair loads the public key pair for the SSLConfig
func (s *SSLConfig) LoadPublicKeyPair() (*tls.Certificate, error) {
	rootPath, err := GetChiaRootPath()
	if err != nil {
		return nil, err
	}

	pair, err := tls.LoadX509KeyPair(filepath.Join(rootPath, s.PublicCRT), filepath.Join(rootPath, s.PublicKey))
	return &pair, err
}
