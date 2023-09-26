package main

import (
	"github.com/solarlune/messages"
)

func main() {

	// Create a new dispatcher.
	dispatcher := messages.NewDispatcher()

	// Create and register a Receiver object.
	// In this example, the Receiver will print messages and pause for a bit as it hears about what happens to the Player.
	dispatcher.Register(NewReceiver())

	// Create an object that emits messages; in this case, a Player character.
	// We'll pass the dispatcher to the Player so he can emit messages through it.
	player := NewPlayer(dispatcher)

	player.TakeDamage()
	player.Heal()
	player.TakeDamage()
	player.TakeDamage()
	player.TakeDamage()

}
