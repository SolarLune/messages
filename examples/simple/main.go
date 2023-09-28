package main

import (
	"github.com/solarlune/messages"
)

// The goal of this example is to show how simple message sending works.
// The player object does things, and our StatusListener will listen to
// his messages and print out the status of the current program state.

func main() {

	// Create a new dispatcher.
	dispatcher := messages.NewDispatcher()

	// Create and register a message receiving object.
	// In this example, the StatusPrinter will print messages and pause for a bit as it
	// hears about what happens to the Player.
	dispatcher.Register(NewStatusListener())

	// Create an object that emits messages; in this case, a Player character.
	// We'll pass the dispatcher to the Player so he can emit messages through it.
	player := NewPlayer(dispatcher)

	player.TakeDamage()
	player.Heal()
	player.TakeDamage()
	player.TakeDamage()
	player.TakeDamage()

}
