package http

import (
	"encoding/json"
	stdhttp "net/http"
)

type HealthResponse struct {
	Status string  `json:"status"`
}

func(s *Server) health(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(stdhttp.StatusOK)

	_ = json.NewEncoder(w).Encode(HealthResponse{
		Status: "up",
	})
}