package client

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	id         string
	token      string
	baseUri    string
	connected  bool
	connection *websocket.Conn
}

func Init(baseUri string) (*Client, error) {
	return &Client{
		"",
		"",
		fmt.Sprintf("ws://%v", baseUri),
		false,
		nil,
	}, nil
}
