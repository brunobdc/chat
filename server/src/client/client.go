package client

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
	ticker *time.Ticker
}

func New(socket *websocket.Conn) (*Client, error) {
	if socket == nil {
		return nil, errSocketCantBeNil
	}

	client := &Client{
		id:     uuid.NewString(),
		socket: socket,
		send:   make(chan []byte),
		ticker: time.NewTicker(30 * time.Second),
	}

	go client.keepConnAlive()

	go client.messageWriter()

	return client, nil
}

func (client Client) ID() string {
	return client.id
}

func (client *Client) SendMessage(message []byte) bool {
	select {
	case client.send <- message:
		return true
	default:
		return false
	}
}

func (client *Client) keepConnAlive() {
	for range client.ticker.C {
		client.socket.Ping(context.Background())
	}
}

func (client *Client) messageWriter() {
	for message := range client.send {
		err := client.socket.Write(context.Background(), websocket.MessageText, message)
		if err != nil {
			log.Printf("Error trying to write a message: %v\nError: %v", string(message), err)
		}
	}
}

func (client *Client) FindMessage() (websocket.MessageType, []byte, error) {
	msgType, message, err := client.socket.Read(context.Background())
	if err != nil {
		err = errors.Join(err, errFailedToReadAMessage)
		closeSocketError := client.socket.Close(websocket.StatusInternalError, "Error reading the message!")
		if closeSocketError != nil {
			err = errors.Join(err, errFailedToCloseTheSocket)
		}
		return 0, nil, err
	}

	return msgType, message, nil
}

func (client *Client) CloseSocket() error {
	client.ticker.Stop()
	close(client.send)
	err := client.socket.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		errors.Join(err, errFailedToCloseTheSocket)
	}
	return nil
}
