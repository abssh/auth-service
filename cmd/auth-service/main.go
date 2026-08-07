package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		if err := httpServer.Start(); err != nil && !errors.Is(err, stdhttp.ErrServerClosed) {
			errCh <- err
		}
		
	}()
	
	grpcServer := internalGrpc.NewServer(logger, &cfg)
	go func() {
		errCh <- grpcServer.Start()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		logger.Error("server error, shutting down", "error", err)
	case sig := <-quit:
		logger.Info("shutting down", "signal", sig)
	}

	httpContext, httpCancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer httpCancel()

	go func() {
		if err:= httpServer.Stop(httpContext); err != nil {
			logger.Error("grpc shutdown error", "error", err)
		}
	}()

	return nil
}