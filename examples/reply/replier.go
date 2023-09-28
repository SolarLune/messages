package main

import (
	"math/rand"

	"github.com/solarlune/messages"
)

// The MessageModifier in this example simply receives the message from the sender,
// and modifies it.
type MessageModifier struct{}

func NewMessageModifier() *MessageModifier { return &MessageModifier{} }

func (r *MessageModifier) ReceiveMessage(message messages.IMessage) {

	switch msg := message.(type) {
	case *MResponse:
		msg.Response = rand.Int()
	}

}
