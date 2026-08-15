package bootstrap

import (
	"fmt"

	"github.com/abssh/auth-service/internal/di"
	grpcService "github.com/abssh/auth-service/internal/grpc/service"

)

func (b *Bootstrapper) registerService() error {
	c := b.c

	if err := di.SingletonRegister[*grpcService.AuthService](c, grpcService.NewAuthService); err != nil {
        return fmt.Errorf("register auth service: %w", err)
    }
	
	return nil
}