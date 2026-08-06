package config

import "github.com/abssh/auth-service/internal/logger"

//

// LogConfig
func (cfg *Config) GetLogLevel() logger.LogLevel {
	return cfg.LogLevel
}

// HttpConfig
func (cfg *Config) GetHttpHost() string {
	return cfg.HttpHost
}

func (cfg *Config) GetHttpPort() int {
	return cfg.HttpPort
}

// GrpcConfig
func (cfg *Config) GetGrpcHost() string {
	return cfg.GrpcHost
}

func (cfg *Config) GetGrpcPort() int {
	return cfg.GrpcPort
}

