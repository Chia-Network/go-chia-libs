package protocols

// ProtocolMessageType corresponds to ProtocolMessageTypes in Chia
type ProtocolMessageType uint8

const (
	// ProtocolMessageTypeHandshake Handshake
	ProtocolMessageTypeHandshake ProtocolMessageType = 1

	// there are many more of these in Chia - only listing the ones current is use for now

	// ProtocolMessageTypeRequestPeers request_peers
	ProtocolMessageTypeRequestPeers ProtocolMessageType = 43

	// ProtocolMessageTypeRespondPeers respond_peers
	ProtocolMessageTypeRespondPeers ProtocolMessageType = 44
)
