# ✉️ Messages ✉️

## What's Messages?

Messages is a simple pure-Go message-passing repo for gamedev. It's made to make it possible to pass messages in an ambiguous, but still safely abstracted and Go-idiomatic way.

## How do I get it?

```go get github.com/solarlune/messages```

## How do I use it? 

Check the example. It's pretty straightforward - you simply create a Dispatcher, register listener objects to the dispatcher, and send pre-defined messages through the dispatcher. Registered listeners will receive the messages as necessary.

Here's a simplified example:

```go

package main

import (
	"fmt"

	"github.com/solarlune/messages"
)



type Receiver struct{}

func NewReceiver() *Receiver { return &Receiver{} }

// Receivers can be anything; they just need to implement messages.IReceiver, which means
// being able to receive an arbitrary message. You can type switch to get the specific contents more easily.
func (receiver *Receiver) ReceiveMessage(msg messages.IMessage) {
	fmt.Println("Received a message!")
}



type MyMessage struct{}

func NewMyMessage() *MyMessage { return &MyMessage{} }

// A Message can be anything, and only needs to implement messages.IMessage, which means
// returning the type of the message as a uint64 (which can be a constant). 
// This is primarily done to identify message types to determine if a specific type of message goes
// to a Receiver.
func (msg *MyMessage) Type() messages.MessageType { return 1 }



func main() {

    // Create a Dispatcher.
	dispatcher := messages.NewDispatcher()

    // Register your receiver.
	dispatcher.Register(NewReceiver())

    // Send a message; by default, it goes to all registered receivers.
	dispatcher.SendMessage(NewMyMessage())

}

```
