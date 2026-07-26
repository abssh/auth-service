package logger

import (
	"fmt"
	"log/slog"
	"strings"
)

type LogLevel slog.Level 

func (l *LogLevel) UnmarshalEnv(value string) error {
	switch strings.ToUpper(value) {
	case "DEBUG":
		*l = LogLevel(slog.LevelDebug)
	case "INFO":
		*l = LogLevel(slog.LevelInfo)
	case "WARN":
		*l = LogLevel(slog.LevelWarn)
	case "ERROR":
		*l = LogLevel(slog.LevelError)
	default:
		return fmt.Errorf("invalid log level")
	}

	return nil
}

func (l LogLevel) Level() slog.Level {
	return slog.Level(l);
}