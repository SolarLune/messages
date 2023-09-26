package main

import "github.com/solarlune/messages"

// The Player in this example can take damage, heal, and die (when its HP is 0 after taking damage).
// As it takes these actions, it sends messages through the Dispatcher.
type Player struct {
	HP         int
	Dispatcher *messages.Dispatcher
}

func NewPlayer(dispatcher *messages.Dispatcher) *Player {
	player := &Player{
		HP:         3,
		Dispatcher: dispatcher,
	}
	player.Dispatcher.Send(PlayerStartMessage{HPRemaining: player.HP})
	player.Dispatcher.Send(PlayerWhineMessage{})
	return player
}

func (player *Player) TakeDamage() {
	player.HP--
	player.Dispatcher.Send(PlayerDamageMessage{HPRemaining: player.HP})
	player.Dispatcher.Send(PlayerWhineMessage{})
	if player.HP <= 0 {
		player.Dispatcher.Send(PlayerDieMessage{})
	}
}

func (player *Player) Heal() {
	oldHP := player.HP
	player.HP++
	player.Dispatcher.Send(PlayerHealMessage{OldHP: oldHP, NewHP: player.HP})
	player.Dispatcher.Send(PlayerWhineMessage{})
}
