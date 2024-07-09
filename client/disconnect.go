package client

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

func (c *Client) Disconnect() error {
	if !c.connected {
		return nil
	}

	message, _ := json.Marshal(Message{
		Token:   c.token,
		Command: "disconnect",
	})
	err := c.connection.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err
	}

	c.connection.Close()
	c.connected = false

	return nil
}
