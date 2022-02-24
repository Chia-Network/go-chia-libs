package rpcinterface

// ServiceType is a type that refers to a particular service
type ServiceType uint8

const (
	// ServiceDaemon the daemon
	ServiceDaemon ServiceType = iota

	// ServiceFullNode the full node service
	ServiceFullNode

	// ServiceFarmer the farmer service
	ServiceFarmer

	// ServiceHarvester the harvester service
	ServiceHarvester

	// ServiceWallet the wallet service
	ServiceWallet

	// ServiceTimelord is the timelord service
	ServiceTimelord

	// ServicePeer full node service, but for communicating with full nodes using the public protocol
	ServicePeer

	// ServiceCrawler crawler service
	ServiceCrawler
)
