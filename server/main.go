package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var manager = ClientManager{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconneted,"})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (client *Client) read() {
	defer func() {
		manager.unregister <- client
		client.socket.Close(websocket.StatusNormalClosure, "")
	}()

	for {
		_, message, err := client.socket.Read(context.Background())
		if err != nil {
			manager.unregister <- client
			client.socket.Close(websocket.StatusInternalError, "Message not found!")
			break
		}

		jsonMessage, _ := json.Marshal(&Message{Sender: client.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

func (client *Client) write() {
	defer client.socket.Close(websocket.StatusNormalClosure, "")

	for message := range client.send {
		client.socket.Write(context.Background(), websocket.MessageText, message)
	}
	client.socket.Write(context.Background(), websocket.MessageText, []byte{})
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	connection, err := websocket.Accept(res, req, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("error ocurred: ", err)
		return
	}
	client := &Client{id: uuid.NewString(), socket: connection, send: make(chan []byte)}

	manager.register <- client

	go client.read()
	go client.write()
}

func main() {
	fmt.Println("Starting application...")
	go manager.start()
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":12345", nil)
}
