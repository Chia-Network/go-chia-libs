package rpcinterface

import "net/http"

// Request is a wrapped http.Request that indicates the service we're making the RPC call to
type Request struct {
	Service  ServiceType
	Endpoint Endpoint
	Data     interface{}
	Request  *http.Request
}
