package main

import (
	"fmt"
	"time"

	"github.com/solarlune/messages"
)

type Listener struct{}

func NewListener() *Listener { return &Listener{} }

// By creating a ReceiveMessage function, the Listener fulfills messages.IReceiver, meaning it can be
// registered with a messages.Dispatcher.
func (listener *Listener) ReceiveMessage(receivedMessage messages.IMessage) {

	// By switching against the message type, we can access the specifics of various kinds of messages.
	switch msg := receivedMessage.(type) {

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

// Now let's say we don't want to listen to the Player's whining. We can ignore those messages by simply creating a Subscribe() function
// that returns the message types we're interested in.

func (listener *Listener) Subscribe() messages.MessageType {
	return TypePlayerDamageMessage + TypePlayerDieMessage + TypePlayerHealMessage
}
