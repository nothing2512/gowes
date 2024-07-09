package server

import (
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
}

func Init(secret, iv string) *Server {
	s := &Server{make(map[string]*websocket.Conn), secret, iv, nil}

	http.HandleFunc("/connect", s.connect)

	return s
}

func (s *Server) Start(baseUri string) {
	ln, err := net.Listen("tcp", baseUri)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening On " + baseUri)
	if err = http.Serve(ln, nil); err != nil {
		panic(err)
	}
}
