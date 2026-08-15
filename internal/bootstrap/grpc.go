package bootstrap

import (
	"fmt"

	authv1 "github.com/abssh/api-contract-sdk/sdk/go/auth/v1"
	cfg "github.com/abssh/auth-service/internal/config"
	"github.com/abssh/auth-service/internal/di"
	internalGrpc "github.com/abssh/auth-service/internal/grpc"
	grpcHandler "github.com/abssh/auth-service/internal/grpc/handler"
	gogrpc "google.golang.org/grpc"
)

func (b *Bootstrapper) registerGrpcServer() error {
	c := b.c
	if err := di.SingletonReregisterAs[internalGrpc.GrpcConfig, *cfg.Config](c); err != nil {
		return fmt.Errorf("register grpc config: %w", err)
	}
	if err := di.SingletonRegister[*internalGrpc.GrpcServer](c, internalGrpc.NewServer); err != nil {
		return fmt.Errorf("register grpc server: %w", err)
	}
	return nil
}

func (b *Bootstrapper) wireGrpcServer() error {
	c := b.c
	
	grpcServer, err := di.Resolve[*internalGrpc.GrpcServer](c)
	if err != nil {
		return fmt.Errorf("resolve grpc server: %w", err)
	}

	authHandler, err := di.Resolve[*grpcHandler.AuthHandler](c)
	if err != nil {
		return fmt.Errorf("resolve auth handler: %w", err)
	}

	grpcServer.RegisterHandler(func(s *gogrpc.Server) {
		authv1.RegisterAuthenticationServiceServer(s, authHandler)
	})
	
	return nil
}
