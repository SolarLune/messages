package main

import (
	"fmt"
	"time"

	"github.com/solarlune/messages"
)

// This example is one that shows how an object can send a message and
// read it after sending to get a reply.

func main() {

	// Create a new Dispatcher.
	dispatcher := messages.NewDispatcher()

	// Create and register our Replier object.
	dispatcher.Register(NewMessageModifier())

	fmt.Println("Let me send a message out.")
	time.Sleep(time.Second)

	// Create and send the message of interest.
	msg := NewMResponse()
	dispatcher.Send(msg)
	fmt.Println("...")
	time.Sleep(time.Second)

	// Examine the contents, as modified by the Replier.
	fmt.Println("The message was modified! The response value is:", msg.Response)
	time.Sleep(time.Second)

}
