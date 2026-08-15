package di

import (
	"fmt"
	"reflect"
)

func (c *Container) Register(provider any) error {
	providerValue, err := validateProvider(provider)
	if err != nil {
		return err
	}

	targetType := providerValue.Type().Out(0)
	if _, exists := c.registry[targetType]; exists {
		return fmt.Errorf(
			"type %s is already registered",
			targetType,
		)
	}

	c.registry[targetType] = providerValue

	return nil
}

func RegisterAs[T any](c *Container, provider any) error {
	providerValue, err := validateProvider(provider)
	if err != nil {
		return err
	}

	targetType := reflect.TypeOf((*T)(nil)).Elem()

	if _, exists := c.registry[targetType]; exists {
		return fmt.Errorf(
			"type %s is already registered",
			targetType,
		)
	}

	if providedType := providerValue.Type().Out(0); !providedType.AssignableTo(targetType) {
		return fmt.Errorf(
			"provided type (%s) can't be assigned to the target type (%s)",
			providedType,
			targetType,
		)
	}

	c.registry[targetType] = providerValue

	return nil
}

func SingletonRegister[T any](c *Container, factory any) error {
	return c.Register(Singleton(
		func() (T, error) {
			return InvokeFactory[T](c, factory)
		},
	))
}

func SingletonReregisterAs[I any, T any](c *Container) error {
	return RegisterAs[I](c, Singleton(
		func() (T, error) {
			return Resolve[T](c)
		},
	))
}

func validateProvider(provider any) (reflect.Value, error) {
	if provider == nil {
		return reflect.Value{}, fmt.Errorf("can't register nil provider")
	}

	providerValue := reflect.ValueOf(provider)
	providerType := providerValue.Type()

	if providerType.Kind() != reflect.Func {
		return reflect.Value{}, fmt.Errorf(
			"provider must be a function, got: %s",
			providerType,
		)
	}

	switch providerType.NumOut() {
	case 0:
		return reflect.Value{}, fmt.Errorf(
			"provider %s must return a value",
			providerType,
		)

	case 1:
		break

	case 2:
		errorType := reflect.TypeOf((*error)(nil)).Elem()
		if !providerType.Out(1).Implements(errorType) {
			return reflect.Value{}, fmt.Errorf(
				"provider %s second return value must implement error",
				providerType,
			)
		}

	default:
		return reflect.Value{}, fmt.Errorf(
			"provider %s must return one value or one value and an error",
			providerType,
		)
	}

	return providerValue, nil
}
