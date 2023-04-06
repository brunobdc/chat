package services

import (
	"log"

	"github.com/brunobdc/chat/server/src/client"
	"github.com/brunobdc/chat/server/src/manager"
	"github.com/brunobdc/chat/server/src/message"
)

type messageReader struct {
	client  *client.Client
	manager *manager.ClientManager
}

func (reader *messageReader) Start() {
	go func() {
		defer func() {
			reader.manager.Unregister(reader.client)
			reader.client.CloseSocket()
		}()

		for {
			_, bytesMessage, err := reader.client.FindMessage()
			if err != nil {
				log.Printf("Failed to find message. Error: %v", err)
				break
			}

			json, err := message.JsonString(reader.client.ID(), string(bytesMessage))
			if err != nil {
				log.Printf("Failed to marshal json. Error: %v", err)
			} else {
				reader.manager.BroadcastMsg(json)
			}
		}
	}()
}

func MakeMessageReader(client *client.Client, manager *manager.ClientManager) *messageReader {
	return &messageReader{
		client:  client,
		manager: manager,
	}
}
