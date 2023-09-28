# ✉️ Messages ✉️

## What's Messages?

Messages is a simple pure-Go message-passing repo for gamedev. It's made to make it possible to pass messages in an ambiguous, but still safely abstracted and Go-idiomatic way.

## How do I get it?

```go get github.com/solarlune/messages```

## How do I use it? 

Check the example directory for a bit more of an in-depth example, but it's pretty straightforward overall - you create a Dispatcher, register receiving objects to the dispatcher, and send messages through the dispatcher. Registered listeners will receive the message as necessary.

Here's a simplified example:

```go

package main

import (
	"fmt"

	"github.com/solarlune/messages"
)

// Receivers can be anything; they just need to implement messages.IReceiver.
// That means that the object has a ReceiveMessage function and can receive an
// arbitrary message.
type Receiver struct{}

func NewReceiver() *Receiver { return &Receiver{} }

func (receiver *Receiver) ReceiveMessage(msg messages.IMessage) {
    // You can type switch the message against your pre-defined message struct
    // types to get specific contents easily.
	fmt.Println("Received a message!")
}

// If you don't want a Receiver to receive all messages, then you can implement
// ISubscriber, which means adding a Subscribe() function that returns the types
// of Messages the Receiver subscribes to (accepts).

////////

// A Message can be anything, and only needs to implement messages.IMessage.
// That means it has a Type() function that returns the type of the message
// as a messages.MessageType, which is a uint64 that is made to do some simple 
// bitwise operations (so the types should be a base-2 numeral, i.e. 1, 2, 4, 8, etc).
// This is primarily done to determine if a Receiver subscribes to a specific 
// type of message.

type MyMessage struct{}

func NewMyMessage() *MyMessage { return &MyMessage{} }

func (msg *MyMessage) Type() messages.MessageType { return 1 }

////////

func main() {

    // Create a Dispatcher.
    dispatcher := messages.NewDispatcher()

    // Register your receiver.
    dispatcher.Register(NewReceiver())

    // Send a message; by default, it goes to all registered receivers.
    dispatcher.Send(NewMyMessage())

}

```

# To-do

[ ] - Add the ability to consume a message so other receivers don't receive it (This can actually easily be done by just designing the message so that it is skippable and then altering it as it goes, but this would still be a nice option)