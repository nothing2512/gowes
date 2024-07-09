package client

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

func (c *Client) Send(to string, data Message) error {
	if !c.connected {
		return nil
	}

	messageData, _ := json.Marshal(data)

	message, _ := json.Marshal(Message{
		Token:   c.token,
		To:      to,
		Command: "send",
		Message: string(messageData),
	})
	err := c.connection.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err
	}

	return nil
}
