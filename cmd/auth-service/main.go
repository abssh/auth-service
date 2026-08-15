package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"log/slog"

	"github.com/abssh/auth-service/internal/bootstrap"
	"github.com/abssh/auth-service/internal/di"
	internalGrpc "github.com/abssh/auth-service/internal/grpc"
	internalHttp "github.com/abssh/auth-service/internal/http"
)

func main() {
    c := di.NewContainer()
    
    bootstrapper := bootstrap.NewBootstrapper(c)
    err := bootstrapper.RegisterAll()
    if err != nil {
        log.Fatalf("bootstrap failed: %s", err)
    }
    

    logger, err := di.Resolve[*slog.Logger](c)
    if err != nil {
        log.Fatalf("resolve logger: %s", err)
    }

    httpServer, err := di.Resolve[*internalHttp.Server](c)
    if err != nil {
        log.Fatalf("resolve http server: %s", err)
    }

    grpcServer, err := di.Resolve[*internalGrpc.GrpcServer](c)
    if err != nil {
        log.Fatalf("resolve grpc server: %s", err)
    }

    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    app := NewApp(logger)
    app.AddServer(httpServer)
    app.AddServer(grpcServer)

    for _, err := range app.Run(ctx) {
        logger.Error(err.Error())
    }

    logger.Info("application stopped")
}