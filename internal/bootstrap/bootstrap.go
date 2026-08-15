package bootstrap

import "github.com/abssh/auth-service/internal/di"

type Bootstrapper struct {
	c *di.Container
}

func NewBootstrapper(container *di.Container) *Bootstrapper {
	return &Bootstrapper{
		c: container,
	}
}

func (b Bootstrapper) GetContainer() *di.Container {
	return b.c
}

func (b *Bootstrapper) RegisterAll() error {
	var err error

	err = b.registerConfig()
	if err != nil {
		return err
	}

	err = b.registerLogger()
	if err != nil {
		return err
	}

	err = b.registerHttpServer()
	if err != nil {
		return err
	}

	err = b.registerGrpcServer()
	if err != nil {
		return err
	}
	// repositories go here
	err = b.registerService()
	if err != nil {
		return err
	}
	err = b.registerHandler()
	if err != nil {
		return err
	}

	err = b.wireGrpcServer()
	if err != nil {
		return err
	}

	return nil
}