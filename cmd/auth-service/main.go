package main

import (
	"fmt"
	"log"

	internalConfig "github.com/abssh/auth-service/internal/config"
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

	server := internalHttp.NewServer(logger, &cfg)

	if err := server.Listen(); err != nil {
		return fmt.Errorf("http server: %w", err)
	}
	return nil
}