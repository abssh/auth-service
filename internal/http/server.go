package http

import (
	"log/slog"
	stdhttp "net/http"
)

type Server struct {
	logger *slog.Logger
	mux    *stdhttp.ServeMux
}

func NewServer(logger *slog.Logger) *Server {
	mux := stdhttp.NewServeMux()

	s := &Server{
		logger: logger,
		mux: mux,
	}

	s.registerRoutes()
	return s
}

func (s *Server) Listen(addr string) error {
	s.logger.Info("starting http server", "address", addr)
	
	return stdhttp.ListenAndServe(addr, s.mux)
}
