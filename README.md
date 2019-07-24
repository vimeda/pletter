<h1 align="center">Welcome to Pletter 👋</h1>
<p></p>

[![Build Status](https://travis-ci.com/vimeda/pletter.svg?branch=master)](https://travis-ci.com/vimeda/pletter)
[![Coverage Status](https://coveralls.io/repos/github/vimeda/pletter/badge.svg)](https://coveralls.io/github/vimeda/pletter)
[![Go Report Card](https://goreportcard.com/badge/github.com/vimeda/pletter)](https://goreportcard.com/report/github.com/vimeda/pletter)
[![GoDoc](https://godoc.org/github.com/vimeda/pletter?status.svg)](https://godoc.org/github.com/vimeda/pletter)

> A standard way to wrap a proto message

Pletter was born with a single mission: `To standardize wrapping protocol buffer messages`. This is normally needed when you use protobuf as your messaging protocol. 

## The Problem

Imagine that you have an event driven architecture. In this system you chose to use protocol buffers to ensure message contract and to transit information.
Let's assume we use Kafka as our message broker. In this Kafka we have one topic called `accounts`, this means all account messages will go to the `accounts` topic.

On the consumer side, we will receive multiple messages in the same topic. If an application wants to read this messages it needs to identify the message type/name
to either enrich, fan out or handle it. To solve that we have a few options:

1. If you use something like AMQP instead of a message broker you can set the name of the message in the header and deal with it in the consumer side
2. If you use Kafka/Kinesis or any message broker that streams messages, you will need to envelop the message.

Option 2 is the option that Pletter tries to solve in Go.

An interesting article about streams architecture can be found [here](https://docs.confluent.io/current/streams/architecture.html)

## Install

```sh
go get github.com/vimeda/pletter
```

## Use

Pletter is simple to use, we provide you a few functions to deal with the message.

When producing messages:

```go
// Create your proto message
ac := pb.Example{
    ID: "1231231312",
}

// call the PackAndMarshal function to wrap your message and already proto.Marshal it
raw, err := any.PackAndMarshal(&ac)
if err != nil {
    fmt.Errorf("an error ocurred while packing and marshalling the message: %s", err)
}

// send the slice of byte to your message broker
```

When consuming messages:

```go
// declare your expected type
var expectedExample pb.Example

// call the Unpack function that will unwrap your message from the envelop
err = any.Unpack(raw, &expectedExample)
if err != nil {
    fmt.Errorf("an error ocurred while unpacking the message: %s", err)
}
```

You can also filter out messages when consuming them:

```go
// call the Unpack function that will unwrap your message from the envelop
name, err := any.GetMessageName(raw)
if err != nil {
    fmt.Errorf("an error ocurred while getting the message name: %s", err)
}

switch name {
    case "pb.Example":
        // declare your expected type
        var expectedExample pb.Example

        // call the Unpack function that will unwrap your message from the envelop
        err = any.Unpack(raw, &expectedExample)
        if err != nil {
            fmt.Errorf("an error ocurred while unpacking the message: %s", err)
        }
        // do somthing
    default;
        // ignore the other messages
}
```

## Run tests

```sh
go test ./...
```

## Author

👤 **Italo Vietro**

* Github: [@italolelis](https://github.com/italolelis)

👤 **Felipe Umpierre**

* Github: [@felipeumpierre](https://github.com/felipeumpierre)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/vimeda/pletter/issues).
