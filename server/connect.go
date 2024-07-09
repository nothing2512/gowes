package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Message struct {
	Command string `json:"command"`
	Message string `json:"message"`
	Token   string `json:"token"`
	To      string `json:"to"`
}

func (s *Server) connect(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		t, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		switch t {
		case websocket.TextMessage:

			var m Message
			_ = json.Unmarshal(message, &m)

			if m.Command == "connect" {
				uid := uuid.New().String()
				fmt.Println("Connected: ", uid)
				b, _ := json.Marshal(map[string]interface{}{
					"id":      uid,
					"token":   s.encrypt(uid),
					"command": "connect",
				})
				fmt.Println(uid)
				fmt.Println(s.encrypt(uid))
				err := conn.WriteMessage(t, b)
				s.clients[uid] = conn
				if err != nil {
					log.Println(err)
					return
				}
			} else if m.Command == "send" {
				s.send(m)
			} else if m.Command == "disconnect" {
				s.disconnect(m)
			} else if m.Command == "command" {
				if s.handle != nil {
					var cmd Message
					_ = json.Unmarshal([]byte(m.Message), &cmd)
					s.handle(cmd)
				}
			}
		}
	}
}
