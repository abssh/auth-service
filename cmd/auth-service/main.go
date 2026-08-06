package main

import (
	"fmt"
	"log"

	internalConfig "github.com/abssh/auth-service/internal/config"
	internalGrpc "github.com/abssh/auth-service/internal/grpc"
	internalHttp "github.com/abssh/auth-service/internal/http"
	internalLogger "github.com/abssh/auth-service/internal/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Startup failed: %s", err.Error())
	}
}

func run () error {

	cfg := internalConfig.Config{}
	err := cfg.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	
	logger := internalLogger.New(&cfg)
	logger.Info("log level is set to " + cfg.GetLogLevel().Level().String())

	errCh := make(chan error, 2)

	httpServer := internalHttp.NewServer(logger, &cfg)
	go func () {
		errCh <- httpServer.Listen()
	}()
	
	grpcServer := internalGrpc.NewServer(logger, &cfg)
	go func() {
		errCh <- grpcServer.Start()
	}()
	

	if err := <- errCh; err != nil {
		return fmt.Errorf("server: %w", err)
	}
	return nil
}