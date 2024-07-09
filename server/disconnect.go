package server

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func (s *Server) disconnect(m Message) {
	token := m.Token
	uid := s.decrypt(token)
	c := s.clients[uid]
	fmt.Println("Disconnected")
	c.WriteMessage(websocket.TextMessage, []byte("Disconnected"))
	c.Close()
	delete(s.clients, uid)
}
