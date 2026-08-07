package handler

import (
	"github.com/abssh/auth-service/internal/grpc/service"
    authv1   "github.com/abssh/api-contract-sdk/sdk/go/auth/v1"
)

type AuthHandler struct {
    authv1.UnimplementedAuthenticationServiceServer
    svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
    return &AuthHandler{svc: svc}
}
