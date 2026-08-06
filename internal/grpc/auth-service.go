package grpc

import (
	"context"

	authv1 "github.com/abssh/api-contract-sdk/sdk/go/auth/v1"
	commonv1 "github.com/abssh/api-contract-sdk/sdk/go/common/v1"
)

type AuthService struct {
	authv1.UnimplementedAuthenticationServiceServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Health(ctx context.Context, req *commonv1.HealthRequest) (resp *commonv1.HealthResponse, err error) {
	resp = &commonv1.HealthResponse{
		Status: "up",
	}
	return resp, nil
}