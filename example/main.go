package main

import (
	"fmt"
	"time"

	"github.com/solarlune/messages"
)

func main() {

	// Create a new dispatcher.
	dispatcher := messages.NewDispatcher()

	// Create an object that emits messages; in this case, a Player character.
	// We'll pass the dispatcher to the Player so he can emit messages through it.
	player := NewPlayer(dispatcher)

	// Create and register a Listener object.
	// In this example, the Listener will print messages as it hears about what happens to the Player.
	dispatcher.Register(NewListener())

	fmt.Println("The player is alive and feeling good. HP is at: ", player.HP)
	time.Sleep(time.Second * 2)

	player.TakeDamage()
	player.Heal()
	player.TakeDamage()
	player.TakeDamage()
	player.TakeDamage()

}
