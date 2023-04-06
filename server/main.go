package main

import (
	"log"
	"net/http"

	"github.com/brunobdc/chat/server/src/client"
	"github.com/brunobdc/chat/server/src/manager"
	"github.com/brunobdc/chat/server/src/services"
	"nhooyr.io/websocket"
)

func main() {
	manager := manager.New()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		connection, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
		if err != nil {
			log.Println("error ocurred: ", err)
			return
		}
		client, err := client.New(connection)
		if err != nil {
			log.Println("error ocurred: ", err)
		}

		manager.Register(client)
		services.MakeMessageReader(client, manager).Start()
	})
	log.Fatal(http.ListenAndServe(":12345", nil))
}
