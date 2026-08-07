package http

import (
	"context"
	"log/slog"
	"net"
	stdhttp "net/http"
	"strconv"
)

type Server struct {
	config HttpConfig
	logger *slog.Logger
	server *stdhttp.Server
}

func NewServer(logger *slog.Logger, cfg HttpConfig) *Server {
	mux := stdhttp.NewServeMux()
	s := &Server{
		config: cfg,
		logger: logger,
	}

	registerRoutes(mux)

	s.server = &stdhttp.Server{
		Addr: net.JoinHostPort(
			s.config.GetHttpHost(),
			strconv.Itoa(s.config.GetHttpPort()),
		),
		Handler: mux,
	}

	return s
}

func (s *Server) Start() error {
	s.logger.Info("starting http server", "address", s.server.Addr)

	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("stopping http server")
	return s.server.Shutdown(ctx)
}
