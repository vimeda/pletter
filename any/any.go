package any

import (
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/vimeda/pletter/pb"
)

// PackAndMarshal packs a proto message into an envelope message and marshal it
func PackAndMarshal(m proto.Message) ([]byte, error) {
	e, err := Pack(m)
	if err != nil {
		return []byte{}, err
	}

	return proto.Marshal(&e)
}

// Pack packs a proto message into an envelope message
func Pack(m proto.Message) (pb.Envelope, error) {
	raw, err := proto.Marshal(m)
	if err != nil {
		return pb.Envelope{}, err
	}

	return pb.Envelope{
		InnerMessage: &any.Any{
			TypeUrl: "github.com/lykon/pletter/" + proto.MessageName(m),
			Value:   raw,
		},
	}, nil
}

// Unpack unpacks a slice of bytes into a proto message.
// The slice of bytes should represent an enveloped proto message
func Unpack(m []byte, t proto.Message) error {
	e, err := getEnvelope(m)
	if err != nil {
		return err
	}

	err = ptypes.UnmarshalAny(e.GetInnerMessage(), t)
	if err != nil {
		return err
	}

	return nil
}

// GetMessageName returns the message name from the wrapped proto message
func GetMessageName(m []byte) (string, error) {
	e, err := getEnvelope(m)
	if err != nil {
		return "", err
	}

	splits := strings.Split(e.GetInnerMessage().GetTypeUrl(), "/")
	return splits[len(splits)-1], nil
}

func getEnvelope(m []byte) (pb.Envelope, error) {
	var receivingEnvelope pb.Envelope
	err := proto.Unmarshal(m, &receivingEnvelope)
	return receivingEnvelope, err
}
