package main

import (
	internalLogger "github.com/abssh/auth-service/internal/logger"
	internalHttp "github.com/abssh/auth-service/internal/http"
)

func main () {
	logger := internalLogger.New()

	server := internalHttp.NewServer(logger)

	if err := server.Listen(":8080"); err != nil {
		logger.Error("server stopped", "error", err)
	}

}