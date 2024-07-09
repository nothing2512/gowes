package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	clients map[string]*websocket.Conn
	secret  string
	iv      string
	handle  func(m Message)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Init(secret, iv string) *Server {
	s := &Server{make(map[string]*websocket.Conn), secret, iv, nil}

	http.HandleFunc("/connect", s.connect)
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		var m Message
		_ = json.NewDecoder(r.Body).Decode(&m)
		if m.Command == "send" {
			s.send(m)
			w.Write([]byte("Sent"))
		} else if m.Command == "command" {
			token := m.Token
			uid := s.decrypt(token)
			if _, exists := s.clients[uid]; !exists {
				w.Write([]byte("Unsent"))
				return
			}
			if s.handle != nil {
				var cmd Message
				_ = json.Unmarshal([]byte(m.Message), &cmd)
				s.handle(cmd)
			}
			w.Write([]byte("Sent"))
		} else {
			w.Write([]byte("Unsent"))
		}
	})

	return s
}

func (s *Server) Start(baseUri string) {
	ln, err := net.Listen("tcp", baseUri)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening WebSocket On " + baseUri)
	if err = http.Serve(ln, nil); err != nil {
		panic(err)
	}
}
