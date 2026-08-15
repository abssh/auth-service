package bootstrap

import (
	"fmt"

	cfg "github.com/abssh/auth-service/internal/config"
	"github.com/abssh/auth-service/internal/di"
	internalHttp "github.com/abssh/auth-service/internal/http"
)

func (b *Bootstrapper) registerHttpServer() error {
	c := b.c
	if err := di.SingletonReregisterAs[internalHttp.HttpConfig, *cfg.Config](c); err != nil {
		return fmt.Errorf("register http config: %w", err)
	}
	if err := di.SingletonRegister[*internalHttp.Server](c, internalHttp.NewServer); err != nil {
		return fmt.Errorf("register http server: %w", err)
	}
	return nil
}
