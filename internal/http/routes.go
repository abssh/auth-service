package http

func (s *Server) registerRoutes() {
	s.mux.HandleFunc("GET /health", s.health)
}