package http

import (
	stdhttp "net/http"

	"github.com/abssh/auth-service/internal/http/routes"
)

func registerRoutes(mux *stdhttp.ServeMux) {
	mux.HandleFunc("GET /health", routes.Health)
}