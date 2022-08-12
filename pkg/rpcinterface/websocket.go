package rpcinterface

import "github.com/chia-network/go-chia-libs/pkg/types"

// WebsocketResponseHandler is a function that is called to process a received websocket response
type WebsocketResponseHandler func(*types.WebsocketResponse, error)

// DisconnectHandler the function to call when the client is disconnected
type DisconnectHandler func()

// ReconnectHandler the function to call when the client is reconnected
type ReconnectHandler func()
