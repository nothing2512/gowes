package server

import (
	"github.com/gorilla/websocket"
)

func (s *Server) send(m Message) {
	token := m.Token
	uid := s.decrypt(token)
	if _, exists := s.clients[uid]; exists {
		if to, exists := s.clients[m.To]; exists {
			to.WriteMessage(websocket.TextMessage, []byte(m.Message))
		}
	}
}
