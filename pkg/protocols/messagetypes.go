package protocols

// ProtocolMessageType corresponds to ProtocolMessageTypes in Chia
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/protocol_message_types.py
type ProtocolMessageType uint8

const (
	// there are many more of these in Chia - only listing the ones current is use for now

	// ProtocolMessageTypeHandshake handshake
	ProtocolMessageTypeHandshake ProtocolMessageType = 1

	// ProtocolMessageTypeHarvesterHandshake harvester_handshake
	ProtocolMessageTypeHarvesterHandshake ProtocolMessageType = 3

	// ProtocolMessageTypeNewProofOfSpace new_proof_of_space
	ProtocolMessageTypeNewProofOfSpace ProtocolMessageType = 5

	// ProtocolMessageTypeRequestSignatures request_signatures
	ProtocolMessageTypeRequestSignatures ProtocolMessageType = 6

	// ProtocolMessageTypeRespondSignatures respond_signatures
	ProtocolMessageTypeRespondSignatures ProtocolMessageType = 7

	// ProtocolMessageTypeNewSignagePoint new_signage_point
	ProtocolMessageTypeNewSignagePoint ProtocolMessageType = 8

	// ProtocolMessageTypeDeclareProofOfSpace declare_proof_of_space
	ProtocolMessageTypeDeclareProofOfSpace ProtocolMessageType = 9

	// ProtocolMessageTypeRequestSignedValues request_signed_values
	ProtocolMessageTypeRequestSignedValues ProtocolMessageType = 10

	// ProtocolMessageTypeSignedValues signed_values
	ProtocolMessageTypeSignedValues ProtocolMessageType = 11

	// ProtocolMessageTypeFarmingInfo farming_info
	ProtocolMessageTypeFarmingInfo ProtocolMessageType = 12

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

	// ProtocolMessageTypePlotSyncStart plot_sync_start
	ProtocolMessageTypePlotSyncStart ProtocolMessageType = 78
)
