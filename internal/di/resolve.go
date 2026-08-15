package di

import (
	"fmt"
	"reflect"
)

func Resolve[T any](c *Container) (T, error) {
	var zero T
	targetType := reflect.TypeOf((*T)(nil)).Elem()

	result, err := c.resolve(targetType, make(map[reflect.Type]bool))
	if err != nil {
		return zero, err
	}

	return result.Interface().(T), nil
}

func InvokeFactory[T any](c *Container, factory any) (T, error) {
	var zero T

	providerValue, err := validateProvider(factory)
	if err != nil {
		return zero, err
	}

	result, err := c.callProvider(providerValue, make(map[reflect.Type]bool))
	if err != nil {
		return zero, err
	}

	return result.Interface().(T), nil
}

func (c *Container) resolve(t reflect.Type, resolving map[reflect.Type]bool) (reflect.Value, error) {
	provider, exists := c.registry[t]
	if !exists {
		return reflect.Value{}, fmt.Errorf(
			"type %s is not registered",
			t,
		)
	}

	if resolving[t] {
		return reflect.Value{}, fmt.Errorf(
			"circular dependency detected involving %s",
			t,
		)
	}
	resolving[t] = true
	defer delete(resolving, t)

	result, err := c.callProvider(provider, resolving)
	if err != nil {
		return reflect.Value{}, fmt.Errorf(
			"resolve %s: %w",
			t,
			err,
		)
	}

	return result, nil
}

func (c *Container) callProvider(provider reflect.Value, resolving map[reflect.Type]bool) (reflect.Value, error) {
	providerType := provider.Type()

	args := make([]reflect.Value, providerType.NumIn())

	for i := 0; i < providerType.NumIn(); i++ {
		dependencyType := providerType.In(i)

		dependency, err := c.resolve(dependencyType, resolving)
		if err != nil {
			return reflect.Value{}, fmt.Errorf(
				"resolve argument %d (%s): %w",
				i,
				dependencyType,
				err,
			)
		}

		args[i] = dependency
	}

	results := provider.Call(args)
	if len(results) == 2 {
		errValue := results[1]

		if !errValue.IsNil() {
			return reflect.Value{}, errValue.Interface().(error)
		}
	}

	return results[0], nil
}
