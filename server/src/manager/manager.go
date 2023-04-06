package manager

import (
	"github.com/brunobdc/chat/server/src/client"
	"github.com/brunobdc/chat/server/src/message"
)

type ClientManager struct {
	clients    map[*client.Client]bool
	broadcast  chan []byte
	register   chan *client.Client
	unregister chan *client.Client
}

func New() *ClientManager {
	manager := &ClientManager{
		clients:    make(map[*client.Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *client.Client),
		unregister: make(chan *client.Client),
	}

	go manager.start()

	return manager
}

func (manager *ClientManager) Unregister(client *client.Client) {
	manager.unregister <- client
}

func (manager *ClientManager) Register(client *client.Client) {
	manager.register <- client
}

func (manager *ClientManager) BroadcastMsg(msg []byte) {
	manager.broadcast <- msg
}

func (manager *ClientManager) send(message []byte, ignore *client.Client) {
	for conn := range manager.clients {
		if conn != ignore {
			if !conn.SendMessage(message) {
				delete(manager.clients, conn)
			}
		}
	}
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := message.JsonString("", "/A new socket has connected.")
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				delete(manager.clients, conn)
				jsonMessage, _ := message.JsonString("", "/A socket has disconneted.")
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				if !conn.SendMessage(message) {
					delete(manager.clients, conn)
				}
			}
		}
	}
}
