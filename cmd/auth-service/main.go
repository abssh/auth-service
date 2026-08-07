package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	authv1 "github.com/abssh/api-contract-sdk/sdk/go/auth/v1"
	internalCommonTypes "github.com/abssh/auth-service/internal/common/types"
	internalConfig "github.com/abssh/auth-service/internal/config"
	internalGrpc "github.com/abssh/auth-service/internal/grpc"
	internalGrpcHandler "github.com/abssh/auth-service/internal/grpc/handler"
	internalGrpcService "github.com/abssh/auth-service/internal/grpc/service"
	internalHttp "github.com/abssh/auth-service/internal/http"
	internalLogger "github.com/abssh/auth-service/internal/logger"
	gogrpc "google.golang.org/grpc"
)

type App struct {
	config  *internalConfig.Config
	logger  *slog.Logger
	servers []internalCommonTypes.Server
}

func (a *App) AddServer(server internalCommonTypes.Server) {
	a.servers = append(a.servers, server)
}

func (a *App) startServers(errCh chan<- error) {
	for _, server := range a.servers {
		s := server

		go func() {
			errCh <- s.Start()
		}()
	}
}

func (a *App) stopServers(errCh chan<- error) {
	for _, server := range a.servers {
		s := server

		go func() {
			ctx, cancel := context.WithTimeout(
				context.Background(),
				10*time.Second,
			)
			defer cancel()

			errCh <- s.Stop(ctx)
		}()
	}
}

func (a *App) run() []error {
	errs := make([]error, 0)

	startErrCh := make(chan error, len(a.servers))

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// Start servers
	a.startServers(startErrCh)

	// Wait for shutdown condition
	select {
	case <-ctx.Done():
		a.logger.Info(
			"shutdown signal received",
			"signal", ctx.Err(),
		)

	case err := <-startErrCh:
		if err != nil && err != http.ErrServerClosed {
			errs = append(errs, err)

			a.logger.Error(
				"server failed",
				"error", err,
			)
		}
	}

	// Stop servers
	stopErrCh := make(chan error, len(a.servers))

	a.stopServers(stopErrCh)

	for range a.servers {
		err := <-stopErrCh

		if err != nil && err != http.ErrServerClosed {
			errs = append(errs, err)
		}
	}

	return errs
}

func main() {
	cfg := &internalConfig.Config{}

	err := cfg.Load()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	logger := internalLogger.New(cfg)

	logger.Info(
		"log level is set to "+
			cfg.GetLogLevel().Level().String(),
	)

	httpServer := internalHttp.NewServer(
		logger,
		cfg,
	)

	// services
	authSvc := internalGrpcService.NewAuthService()

	// handlers
	authHandler := internalGrpcHandler.NewAuthHandler(authSvc)


	grpcServer := internalGrpc.NewServer(
		logger,
		cfg,
	)
	grpcServer.RegisterHandler(func(s *gogrpc.Server) {
		authv1.RegisterAuthenticationServiceServer(s, authHandler)
	})

	app := App{
		config: cfg,
		logger: logger,
	}

	app.AddServer(httpServer)
	app.AddServer(grpcServer)

	for _, err := range app.run() {
		logger.Error(err.Error())
	}

	logger.Info("application stopped")
}