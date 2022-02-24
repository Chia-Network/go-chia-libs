package rpcinterface

// ClientOptionFunc can be used to customize a new RPC client.
type ClientOptionFunc func(client Client) error
