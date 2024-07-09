package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func (c *Client) Connect() error {
	if c.connected {
		return nil
	}
	uri := fmt.Sprintf("%v/connect", c.baseUri)
	conn, _, err := websocket.DefaultDialer.Dial(uri, nil)
	if err != nil {
		log.Fatal(err)
	}

	message, _ := json.Marshal(Message{Command: "connect"})
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err
	}

	_, reply, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	var data struct {
		Id    string `json:"id"`
		Token string `json:"token"`
	}
	_ = json.Unmarshal(reply, &data)

	c.id = data.Id
	c.token = data.Token
	c.connected = true
	c.connection = conn

	return nil
}
