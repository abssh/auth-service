package bootstrap

import (
	"fmt"

	"github.com/abssh/auth-service/internal/di"
	grpcHandler "github.com/abssh/auth-service/internal/grpc/handler"
)

func (b *Bootstrapper) registerHandler() error {
	c := b.c
	if err := di.SingletonRegister[*grpcHandler.AuthHandler](c, grpcHandler.NewAuthHandler); err != nil {
		return fmt.Errorf("register auth handler: %w", err)
	}
	return nil
}
