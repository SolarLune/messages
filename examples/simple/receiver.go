package main

import (
	"fmt"
	"time"

	"github.com/solarlune/messages"
)

// In this example, the StatusListener is a message receiver.
// It listens for various kinds of messages that indicate a change in program state.
type StatusListener struct{}

func NewStatusListener() *StatusListener { return &StatusListener{} }

// By creating a ReceiveMessage function, this type fulfills messages.IReceiver.
// This means it can be registered with a messages.Dispatcher.
func (s *StatusListener) ReceiveMessage(receivedMessage messages.IMessage) {

	// By switching against the message type, we can access the specifics of various kinds of messages.
	switch msg := receivedMessage.(type) {

	case PlayerStartMessage:
		fmt.Println("The player is setting out! His HP is at: ", msg.HPRemaining)

	case PlayerDamageMessage:
		fmt.Println("The player got hurt! HP now at: ", msg.HPRemaining)

	case PlayerHealMessage:
		fmt.Println("The player has healed! Thank goodness.\nHe was at ", msg.OldHP, ", but now he's at: ", msg.NewHP)

	case PlayerDieMessage:
		fmt.Println("The player has died! Oh no!")

	case PlayerWhineMessage:
		fmt.Println("The player whines. He's kind of weak...")

	}

	time.Sleep(time.Second * 2)

}

// Now let's say we don't want to listen to the Player's whining.
// We can ignore those messages by simply creating a Subscribe() function
// that returns the message types we're interested in.
// This makes the StatusReceiver implement messages.ISubscriber:

// func (s *StatusReceiver) Subscribe() messages.MessageType {
// 	return TypePlayerStartMessage + TypePlayerDamageMessage + TypePlayerDieMessage + TypePlayerHealMessage
// }
