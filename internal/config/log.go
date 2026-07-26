package config

import "github.com/abssh/auth-service/internal/logger"

func (cfg *Config) GetLogLevel() logger.LogLevel {
	return cfg.LogLevel
}