package any

import (
	"errors"
	"fmt"
	"strings"

	proto_old "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/vimeda/pletter/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

const pkg = "github.com/lykon/pletter/"

// ErrEmptyMessage when message is nil
var ErrEmptyMessage = errors.New("message is nil")

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
	if m == nil {
		return pb.Envelope{}, ErrEmptyMessage
	}

	raw, err := proto.Marshal(m)
	if err != nil {
		return pb.Envelope{}, fmt.Errorf("error marshaling message: %w", err)
	}

	return pb.Envelope{
		InnerMessage: &anypb.Any{
			TypeUrl: pkg + string(proto.MessageName(m)),
			Value:   raw,
		},
	}, nil
}

// Unpack unpacks a slice of bytes into a proto message.
// The slice of bytes should represent an enveloped proto message
func Unpack(m []byte, t proto.Message) error {
	e, err := getEnvelope(m)
	if err != nil {
		return fmt.Errorf("error getting the envelope: %w", err)
	}

	return ptypes.UnmarshalAny(e.GetInnerMessage(), proto_old.MessageV1(t))
}

// GetMessageName returns the message name from the wrapped proto message
func GetMessageName(m []byte) (string, error) {
	e, err := getEnvelope(m)
	if err != nil {
		return "", fmt.Errorf("error getting the envelope: %w", err)
	}

	splits := strings.Split(e.GetInnerMessage().GetTypeUrl(), "/")

	return splits[len(splits)-1], nil
}

func getEnvelope(m []byte) (pb.Envelope, error) {
	var receivingEnvelope pb.Envelope

	err := proto.Unmarshal(m, &receivingEnvelope)

	return receivingEnvelope, err
}
