package protocols

import (
	"github.com/chia-network/go-chia-libs/pkg/streamable"

	"github.com/samber/mo"
)

// Message is a protocol message
type Message struct {
	ProtocolMessageType ProtocolMessageType `streamable:""`
	ID                  mo.Option[uint16]   `streamable:""`
	Data                []byte              `streamable:""`
}

// DecodeData decodes the data in the message to the provided type
func (m *Message) DecodeData(v interface{}) error {
	return streamable.Unmarshal(m.Data, v)
}

// MakeMessage makes a new Message with the given data
func MakeMessage(messageType ProtocolMessageType, data interface{}) (*Message, error) {
	msg := &Message{
		ProtocolMessageType: messageType,
	}

	var dataBytes []byte
	var err error
	if data != nil {
		dataBytes, err = streamable.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	msg.Data = dataBytes

	return msg, nil
}

// MakeMessageBytes calls MakeMessage and converts everything down to bytes
func MakeMessageBytes(messageType ProtocolMessageType, data interface{}) ([]byte, error) {
	msg, err := MakeMessage(messageType, data)
	if err != nil {
		return nil, err
	}

	return streamable.Marshal(msg)
}

// DecodeMessage is a helper function to quickly decode bytes to Message
func DecodeMessage(bytes []byte) (*Message, error) {
	msg := &Message{}

	err := streamable.Unmarshal(bytes, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// DecodeMessageData decodes a message.data into the given interface
func DecodeMessageData(bytes []byte, v interface{}) error {
	msg, err := DecodeMessage(bytes)
	if err != nil {
		return err
	}

	return msg.DecodeData(v)
}
