package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func parseValue(tag Tag, t reflect.Type) (reflect.Value, error) {
	raw, exist := os.LookupEnv(tag.Name)
	if !exist {
		if tag.Required {
			return reflect.Value{}, errors.New("config: couldn't find required env " + tag.Name)
		}
		if tag.Default != nil {
			raw = *tag.Default
		}
	}

	switch t.Kind() {
	case reflect.String:
		return reflect.ValueOf(raw).Convert(t), nil

	case reflect.Bool:
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(v).Convert(t), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	v, err := strconv.ParseInt(raw, 10, t.Bits())
	if err != nil {
		return reflect.Value{}, err
	}

	return reflect.ValueOf(v).Convert(t), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	v, err := strconv.ParseUint(raw, 10, t.Bits())
	if err != nil {
		return reflect.Value{}, err
	}

	return reflect.ValueOf(v).Convert(t), nil

	case reflect.Float32, reflect.Float64:
	v, err := strconv.ParseFloat(raw, t.Bits())
	if err != nil {
		return reflect.Value{}, err
	}

	return reflect.ValueOf(v).Convert(t), nil

	default:
		return reflect.Value{}, fmt.Errorf(
			"unsupported config type: %s",
			t.Kind(),
		)
	}
}