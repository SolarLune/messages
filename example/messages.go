package main

import "github.com/solarlune/messages"

const (
	TypePlayerDamageMessage = 1 << iota
	TypePlayerHealMessage
	TypePlayerDieMessage
	TypePlayerWhineMessage
)

type PlayerDamageMessage struct{ HPRemaining int }

func (msg PlayerDamageMessage) Type() messages.MessageType { return TypePlayerDamageMessage }

type PlayerHealMessage struct{ OldHP, NewHP int }

func (msg PlayerHealMessage) Type() messages.MessageType { return TypePlayerHealMessage }

type PlayerDieMessage struct{}

func (msg PlayerDieMessage) Type() messages.MessageType { return TypePlayerDieMessage }

type PlayerWhineMessage struct{}

func (msg PlayerWhineMessage) Type() messages.MessageType { return TypePlayerWhineMessage }
