package rpcinterface

import "github.com/chia-network/go-chia-libs/pkg/types"

// WebsocketResponseHandler is a function that is called to process a received websocket response
type WebsocketResponseHandler func(*types.WebsocketResponse, error)
