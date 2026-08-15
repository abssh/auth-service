package di

import "reflect"

type Container struct {
	registry  map[reflect.Type]reflect.Value
}

func NewContainer() *Container {
	return &Container{
		registry:  make(map[reflect.Type]reflect.Value),
	}
}
