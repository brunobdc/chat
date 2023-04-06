package client

import (
	"errors"
	"testing"

	"nhooyr.io/websocket"
)

func TestNew(t *testing.T) {
	t.Run("Should return a new client pointer with no error", func(t *testing.T) {
		socket := new(websocket.Conn)

		client, err := New(socket)

		if err != nil {
			t.Error(err)
		}

		if client.socket != socket {
			t.Errorf("Not the expected socket!\nExpected: %p\nGot: %p", socket, client.socket)
		}
	})

	t.Run("Should return the corret error if the socket passed is nil", func(t *testing.T) {
		client, err := New(nil)

		if err == nil {
			t.Error("Should have returned a error informing that the socket can't be nil!")
		}
		if client != nil {
			t.Errorf("Client should be nil: %+v!", *client)
		}
		if !errors.Is(err, errSocketCantBeNil) {
			t.Errorf("Returned not the expected error!\nExpected: %v\nGot: %v", errSocketCantBeNil, err)
		}
	})
}
