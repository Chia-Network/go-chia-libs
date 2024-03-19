package types

import (
	"github.com/samber/mo"
)

// PrivateKey is a chia_rs type that represents a private key
type PrivateKey struct {
	// @TODO CHIA_RS BINDINGS: Add when we have the rust -> go bindings
}

// KeyDataSecrets contains the secret portion of key data
type KeyDataSecrets struct {
	Mnemonic   []string   `json:"mnemonic" streamable:""`
	Entropy    []byte     `json:"entropy" streamable:""`
	PrivateKey PrivateKey `json:"PrivateKey" streamable:""`
}

// KeyData is the KeyData type from chia-blockchain
type KeyData struct {
	Fingerprint uint32                    `json:"fingerprint" streamable:""`
	PublicKey   G1Element                 `json:"public_key" streamable:""`
	Label       mo.Option[string]         `json:"label" streamable:""`
	Secrets     mo.Option[KeyDataSecrets] `json:"secrets" streamable:""`
}
