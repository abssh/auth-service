package main

import (
	internalConfig "github.com/abssh/auth-service/internal/config"
	internalHttp "github.com/abssh/auth-service/internal/http"
	internalLogger "github.com/abssh/auth-service/internal/logger"
)

func main () {
	cfg := internalConfig.Config{}
	cfg.Load()
	
	logger := internalLogger.New()

	server := internalHttp.NewServer(logger, &cfg)

	if err := server.Listen(); err != nil {
		logger.Error("server stopped", "error", err)
	}

}