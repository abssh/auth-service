package grpc

import (
	authv1 "github.com/abssh/api-contract-sdk/sdk/go/auth/v1"
)

func (s *GrpcServer) registerService() {
	authService := NewAuthService()

	authv1.RegisterAuthenticationServiceServer(
		s.server,
		authService,
	)
}