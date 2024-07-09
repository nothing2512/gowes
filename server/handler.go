package server

func (s *Server) OnCommand(handle func(m Message)) {
	s.handle = handle
}
