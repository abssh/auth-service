package logger

import (
	"log/slog"
	"os"
)

func New(cfg LogConfig) *slog.Logger {
	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: cfg.GetLogLevel(),
		}),
	)
}