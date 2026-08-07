package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	gogrpc "google.golang.org/grpc"
)

type GrpcServer struct {
	server *gogrpc.Server
	config GrpcConfig
	logger *slog.Logger
}

func NewServer(logger *slog.Logger, cfg GrpcConfig) *GrpcServer {
	s := &GrpcServer{
		server: gogrpc.NewServer(),
		config: cfg,
		logger: logger,
	}

	s.registerService()

	return s
}

func (s *GrpcServer) Start() error {
	addr := fmt.Sprintf(
		"%s:%d",
		s.config.GetGrpcHost(),
		s.config.GetGrpcPort(),
	)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.logger.Info("starting grpc server", "address", addr)
	return s.server.Serve(listener)
}

func (s *GrpcServer) Stop(ctx context.Context) error {
	s.logger.Info("stopping grpc server")

	grpcDone := make(chan struct{})

	go func() {
		s.server.GracefulStop()
		close(grpcDone)
	}()

	select {
	case <-grpcDone:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("grpc graceful stop timed out, forced stop")
	}
}