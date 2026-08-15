package main

import (
    "context"
    "log/slog"
    "net/http"
    "time"

    commonTypes "github.com/abssh/auth-service/internal/common/types"
)

type App struct {
    logger  *slog.Logger
    servers []commonTypes.Server
}

func NewApp(logger *slog.Logger) *App {
    return &App{logger: logger}
}

func (a *App) AddServer(server commonTypes.Server) {
    a.servers = append(a.servers, server)
}

func (a *App) startServers(errCh chan<- error) {
    for _, server := range a.servers {
        s := server
        go func() { errCh <- s.Start() }()
    }
}

func (a *App) stopServers(errCh chan<- error) {
    for _, server := range a.servers {
        s := server
        go func() {
            ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
            defer cancel()
            errCh <- s.Stop(ctx)
        }()
    }
}

func (a *App) Run(ctx context.Context) []error {
    errs := make([]error, 0)
    startErrCh := make(chan error, len(a.servers))

    a.startServers(startErrCh)

    select {
    case <-ctx.Done():
        a.logger.Info("shutdown signal received")
    case err := <-startErrCh:
        if err != nil && err != http.ErrServerClosed {
            errs = append(errs, err)
            a.logger.Error("server failed", "error", err)
        }
    }

    stopErrCh := make(chan error, len(a.servers))
    a.stopServers(stopErrCh)

    for range a.servers {
        if err := <-stopErrCh; err != nil && err != http.ErrServerClosed {
            errs = append(errs, err)
        }
    }

    return errs
}