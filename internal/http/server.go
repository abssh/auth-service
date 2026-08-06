package http

import (
	"log/slog"
	"net"
	stdhttp "net/http"
	"strconv"
)

type Server struct {
	config HttpConfig
	logger *slog.Logger
	mux    *stdhttp.ServeMux
}

func NewServer(logger *slog.Logger, cfg HttpConfig) *Server {
	mux := stdhttp.NewServeMux()

	s := &Server{
		config: cfg,
		logger: logger,
		mux:    mux,
	}

	s.registerRoutes()
	return s
}

func (s *Server) Listen() error {
	addr := s.httpAddress()
	s.logger.Info("starting http server", "address", addr)

	return stdhttp.ListenAndServe(addr, s.mux)
}

func (s *Server) httpAddress() string {
	return net.JoinHostPort(
		s.config.GetHttpHost(),
		strconv.Itoa(s.config.GetHttpPort()),
	)
}
