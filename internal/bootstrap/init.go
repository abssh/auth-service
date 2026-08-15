package bootstrap

import (
	"fmt"
	"log/slog"

	cfg "github.com/abssh/auth-service/internal/config"
	"github.com/abssh/auth-service/internal/di"
	"github.com/abssh/auth-service/internal/logger"
)

func (b *Bootstrapper) registerConfig() error {
	return b.c.Register(di.Singleton(func() (*cfg.Config, error) {
        config := &cfg.Config{}
        if err := config.Load(); err != nil {
            return nil, fmt.Errorf("load config: %w", err)
        }
        return config, nil
    }));

}

func (b *Bootstrapper) registerLogger() error {
	c := b.c
	if err := di.SingletonReregisterAs[logger.LogConfig, *cfg.Config](c); err != nil {
        return fmt.Errorf("register log config: %w", err)
    }
	if err := di.SingletonRegister[*slog.Logger](c, logger.New); err != nil {
        return fmt.Errorf("register logger: %w", err)
    }
	return nil
}