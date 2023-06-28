package protocols_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/protocols"
)

func TestMakeMessage(t *testing.T) {

}

func TestMakeMessage_NilData(t *testing.T) {
	msg, err := protocols.MakeMessage(protocols.ProtocolMessageTypeHandshake, nil)
	assert.NoError(t, err)
	assert.Equal(t, protocols.ProtocolMessageTypeHandshake, msg.ProtocolMessageType)
	assert.False(t, msg.ID.IsPresent())
	assert.Equal(t, []byte(nil), msg.Data)
}

func TestDecodeMessage(t *testing.T) {
	//Message(
	//	uint8(ProtocolMessageTypes.handshake.value),
	//	None,
	//	bytes("This is a sample message to decode".encode(encoding = 'UTF-8', errors = 'string'))
	//)
	encodedHex := "0100000000225468697320697320612073616d706c65206d65737361676520746f206465636f6465"

	messageBytes, err := hex.DecodeString(encodedHex)
	assert.NoError(t, err)

	msg, err := protocols.DecodeMessage(messageBytes)
	assert.NoError(t, err)

	assert.NoError(t, err)
	assert.Equal(t, protocols.ProtocolMessageTypeHandshake, msg.ProtocolMessageType)
	assert.False(t, msg.ID.IsPresent())
	assert.Equal(t, []byte("This is a sample message to decode"), msg.Data)
}

func TestDecodeMessageData(t *testing.T) {
	//Message(
	//	uint8(ProtocolMessageTypes.handshake.value),
	//	None,
	//	Handshake(
	//      "mainnet",
	//      "0.0.33",
	//      "1.2.11",
	//      uint16(8444),
	//      uint8(1),
	//      [(uint16(Capability.BASE.value), "1")],
	//  )
	//)
	encodedHex := "01000000002d000000076d61696e6e657400000006302e302e333300000006312e322e313120fc010000000100010000000131"

	messageBytes, err := hex.DecodeString(encodedHex)
	assert.NoError(t, err)

	handshake := &protocols.Handshake{}
	err = protocols.DecodeMessageData(messageBytes, handshake)
	assert.NoError(t, err)
	assert.Equal(t, "mainnet", handshake.NetworkID)
	assert.Equal(t, "0.0.33", handshake.ProtocolVersion)
	assert.Equal(t, "1.2.11", handshake.SoftwareVersion)
	assert.Equal(t, uint16(8444), handshake.ServerPort)
	assert.Equal(t, protocols.NodeTypeFullNode, handshake.NodeType)
	assert.IsType(t, []protocols.Capability{}, handshake.Capabilities)
	assert.Len(t, handshake.Capabilities, 1)
}
