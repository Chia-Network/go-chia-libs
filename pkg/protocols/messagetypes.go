package protocols

// ProtocolMessageType corresponds to ProtocolMessageTypes in Chia
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/protocol_message_types.py
type ProtocolMessageType uint8

const (
	// ProtocolMessageTypeHandshake handshake
	ProtocolMessageTypeHandshake ProtocolMessageType = 1

	// ProtocolMessageTypeHarvesterHandshake harvester_handshake
	ProtocolMessageTypeHarvesterHandshake ProtocolMessageType = 3

	// there are many more of these in Chia - only listing the ones current is use for now

	// ProtocolMessageTypeNewSignagePoint new_signage_point
	ProtocolMessageTypeNewSignagePoint ProtocolMessageType = 8

	// ProtocolMessageTypeNewPeak new_peak
	ProtocolMessageTypeNewPeak ProtocolMessageType = 20

	// ProtocolMessageTypeRequestBlock request_block
	ProtocolMessageTypeRequestBlock ProtocolMessageType = 26

	// ProtocolMessageTypeRespondBlock respond_block
	ProtocolMessageTypeRespondBlock ProtocolMessageType = 27

	// ProtocolMessageTypeRequestPeers request_peers
	ProtocolMessageTypeRequestPeers ProtocolMessageType = 43

	// ProtocolMessageTypeRespondPeers respond_peers
	ProtocolMessageTypeRespondPeers ProtocolMessageType = 44

	// ProtocolMessageTypeNewSignagePointHarvester new_signage_point_harvester
	ProtocolMessageTypeNewSignagePointHarvester ProtocolMessageType = 66
)
