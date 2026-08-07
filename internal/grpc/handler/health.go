package handler

import (
	"context"

	commonv1 "github.com/abssh/api-contract-sdk/sdk/go/common/v1"
)

func (h *AuthHandler) Health(ctx context.Context, req *commonv1.HealthRequest) (*commonv1.HealthResponse, error) {
    status, err := h.svc.Health(ctx)
    if err != nil {
        return nil, err
    }
    return &commonv1.HealthResponse{Status: status}, nil
}