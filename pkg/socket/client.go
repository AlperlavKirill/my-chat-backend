package socket

import (
	"ChatProgramming/pkg/models"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	MessageRepo models.MessageRepository
	wsConn      *websocket.Conn
}

func (c *Client) processMessages() {
	defer func(wsConn *websocket.Conn) {
		err := wsConn.Close()
		if err != nil {

		}
	}(c.wsConn)
	for {
		var message models.Message

		err := c.wsConn.ReadJSON(&message)
		if err != nil {
			log.Printf("Failed to read JSON: %s", err.Error())
			break
		}
		_, err = c.MessageRepo.Create(message.Content, message.Author)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Message Received: %s\n", message.Content)
	}

}

func WsHandler(client *Client, w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	client.wsConn = conn

	client.processMessages()

}
