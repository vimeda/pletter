package any_test

import (
	"testing"

	"github.com/vimeda/pletter/any"
	"github.com/vimeda/pletter/pb"
)

func TestAny(t *testing.T) {
	tests := []struct {
		scenario string
		function func(*testing.T)
	}{
		{
			scenario: "pack proto message successfully",
			function: testPackProtoMessage,
		},
		{
			scenario: "pack and marshal proto message successfully",
			function: testPackAndMarshallProtoMessage,
		},
		{
			scenario: "pack and marshal proto message fail when nil message given",
			function: testPackAndMarshallProtoMessageFail,
		},
		{
			scenario: "unpack a proto message successfully",
			function: testUnpackProtoMessage,
		},
		{
			scenario: "unpack a proto message fail",
			function: testUnpackProtoMessageFail,
		},
		{
			scenario: "get proto message name successfully",
			function: testGetProtoMessageName,
		},
		{
			scenario: "get proto message name fail",
			function: testGetProtoMessageNameFail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, tt.function)
	}
}

func testPackProtoMessage(t *testing.T) {
	ac := pb.Example{
		ID: "1231231312",
	}

	e, err := any.Pack(&ac)
	if err != nil {
		t.Errorf("an error was not expected when packing a message: %s", err)
		t.Fail()
	}

	if e.GetInnerMessage() == nil {
		t.Error("a message was expected to be in the inner message")
	}
}

func testPackAndMarshallProtoMessage(t *testing.T) {
	ac := pb.Example{
		ID: "1231231312",
	}

	raw, err := any.PackAndMarshal(&ac)
	if err != nil {
		t.Errorf("an error was not expected when packing a message: %s", err)
		t.Fail()
	}

	if len(raw) <= 0 {
		t.Error("a alice of bytes was expected when packing and marshalling a message", err)
	}
}

func testPackAndMarshallProtoMessageFail(t *testing.T) {
	_, err := any.PackAndMarshal(nil)
	if err == nil {
		t.Error("an error was expected when packing a message")
	}
}

func testUnpackProtoMessage(t *testing.T) {
	ac := pb.Example{
		ID: "1231231312",
	}

	raw, err := any.PackAndMarshal(&ac)
	if err != nil {
		t.Errorf("an error was not expected when packing a message: %s", err)
		t.Fail()
	}

	var expectedExample pb.Example
	err = any.Unpack(raw, &expectedExample)
	if err != nil {
		t.Errorf("an error was not expected when unpacking a message: %s", err)
		t.Fail()
	}

	if expectedExample.ID != ac.ID {
		t.Error("the IDs of the messages were supposed to be equal")
	}
}

func testUnpackProtoMessageFail(t *testing.T) {
	var expectedExample pb.Example
	err := any.Unpack([]byte(`123`), &expectedExample)
	if err == nil {
		t.Error("an error was expected when unpacking a message with random byte value")
	}

	err = any.Unpack(nil, &expectedExample)
	if err == nil {
		t.Error("an error was expected when unpacking a nil message")
	}
}

func testGetProtoMessageName(t *testing.T) {
	ac := pb.Example{
		ID: "1231231312",
	}

	raw, err := any.PackAndMarshal(&ac)
	if err != nil {
		t.Errorf("an error was not expected when packing a message: %s", err)
		t.Fail()
	}

	name, err := any.GetMessageName(raw)
	if err != nil {
		t.Errorf("an error was not expected when getting the message name: %s", err)
		t.Fail()
	}

	expectedName := "pb.Example"
	if name != expectedName {
		t.Errorf("expected name %q, got %q", expectedName, name)
	}
}

func testGetProtoMessageNameFail(t *testing.T) {
	_, err := any.GetMessageName([]byte(`123`))
	if err == nil {
		t.Error("an error was expected when getting a message name")
	}
}
