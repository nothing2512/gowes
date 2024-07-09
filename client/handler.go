package client

import "encoding/json"

type Message struct {
	Command string `json:"command"`
	Message string `json:"message"`
	Token   string `json:"token"`
	To      string `json:"to"`
}

func (c *Client) OnMessage(handle func(m Message)) {
	for {
		_, message, _ := c.connection.ReadMessage()
		var m Message
		_ = json.Unmarshal(message, &m)
		handle(m)
	}
}
