package any_test

import (
	"fmt"

	"github.com/vimeda/pletter/any"
	"github.com/vimeda/pletter/pb"
)

func ExamplePack() {
	// Create your proto message
	ac := pb.Example{
		ID: "1231231312",
	}

	// call the Pack function to wrap your message. This returns an envelop type
	_, err := any.Pack(&ac)
	if err != nil {
		fmt.Errorf("an error ocurred while packing the message: %s", err)
	}
}

func ExamplePackAndMarshal() {
	// Create your proto message
	ac := pb.Example{
		ID: "1231231312",
	}

	// call the PackAndMarshal function to wrap your message and already proto.Marshal it
	_, err := any.PackAndMarshal(&ac)
	if err != nil {
		fmt.Errorf("an error ocurred while packing and marshalling the message: %s", err)
	}
}

func ExampleUnpack() {
	// Create your proto message
	ac := pb.Example{
		ID: "1231231312",
	}

	// call the PackAndMarshal function to wrap your message and already proto.Marshal it
	raw, err := any.PackAndMarshal(&ac)
	if err != nil {
		fmt.Errorf("an error ocurred while packing and marshalling the message: %s", err)
	}

	// declare your expected type
	var expectedExample pb.Example

	// call the Unpack function that will unwrap your message from the envelop
	err = any.Unpack(raw, &expectedExample)
	if err != nil {
		fmt.Errorf("an error ocurred while unpacking the message: %s", err)
	}
}

func ExampleGetMessageName() {
	// Create your proto message
	ac := pb.Example{
		ID: "1231231312",
	}

	// call the PackAndMarshal function to wrap your message and already proto.Marshal it
	raw, err := any.PackAndMarshal(&ac)
	if err != nil {
		fmt.Errorf("an error ocurred while packing and marshalling the message: %s", err)
	}

	// call the Unpack function that will unwrap your message from the envelop
	name, err := any.GetMessageName(raw)
	if err != nil {
		fmt.Errorf("an error ocurred while getting the message name: %s", err)
	}

	fmt.Printf("message name: %s", name) // should print pb.Example
}
