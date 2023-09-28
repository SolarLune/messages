package main

import "github.com/solarlune/messages"

const (
	TypeResponse = 1 << iota
)

// In this example, MResponse is a message whose purpose is to return a response from a third party.
type MResponse struct {
	Response int
}

func NewMResponse() *MResponse {
	return &MResponse{}
}

func (m *MResponse) Type() messages.MessageType {
	return TypeResponse
}
