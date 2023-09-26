package main

import "github.com/solarlune/messages"

type Player struct {
	HP         int
	Dispatcher *messages.Dispatcher
}

func NewPlayer(dispatcher *messages.Dispatcher) *Player {
	return &Player{
		HP:         3,
		Dispatcher: dispatcher,
	}
}

func (player *Player) TakeDamage() {
	// When taking damage, the Player will emit messages through the message Dispatcher.
	player.HP--
	if player.HP <= 0 {
		player.Dispatcher.SendMessage(PlayerDieMessage{})
	} else {
		player.Dispatcher.SendMessage(PlayerDamageMessage{HPRemaining: player.HP})
		player.Dispatcher.SendMessage(PlayerWhineMessage{})
	}
}

func (player *Player) Heal() {
	oldHP := player.HP
	player.HP++
	player.Dispatcher.SendMessage(PlayerHealMessage{OldHP: oldHP, NewHP: player.HP})
}
