package config

import (
	"crypto/tls"
	"errors"
	"path/filepath"
)

// LoadPrivateKeyPair loads the private key pair for the SSLConfig
func (s *SSLConfig) LoadPrivateKeyPair() (*tls.Certificate, error) {
	rootPath, err := GetChiaRootPath()
	if err != nil {
		return nil, err
	}

	if s.PrivateCRT == "" || s.PrivateKey == "" {
		return nil, errors.New("missing private key or cert. Ensure config.yaml is up to date with the latest changes")
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

	if s.PublicCRT == "" || s.PublicKey == "" {
		return nil, errors.New("missing public key or cert. Ensure config.yaml is up to date with the latest changes")
	}

	pair, err := tls.LoadX509KeyPair(filepath.Join(rootPath, s.PublicCRT), filepath.Join(rootPath, s.PublicKey))
	return &pair, err
}
