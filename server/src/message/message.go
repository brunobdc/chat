package message

import (
	"encoding/json"
)

type Message struct {
	Sender  string `json:"sender,omitempty"`
	Content string `json:"content,omitempty"`
}

func New(sender string, msg string) *Message {
	return &Message{Sender: sender, Content: msg}
}

func (msg *Message) JsonString() ([]byte, error) {
	return json.Marshal(msg)
}

func JsonString(sender string, msg string) ([]byte, error) {
	return New(sender, msg).JsonString()
}
