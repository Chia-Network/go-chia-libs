package rpcinterface

// Service defines the interface for a service
type Service interface {
	NewRequest(rpcEndpoint Endpoint, opt interface{}) (*Request, error)
	GetClient() Client
}
