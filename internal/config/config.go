package config

import (
	"reflect"
	"github.com/abssh/auth-service/internal/logger"
)

type Config struct {
	HttpHost string `env:"HTTP_HOST" default:""`
	HttpPort int    `env:"HTTP_PORT" default:"8080"`

	LogLevel logger.LogLevel `env:"LOG_LEVEL" default:"INFO"`
}

func (cfg *Config) Load() error {
	var unmarshalerType = reflect.TypeOf((*EnvUnmarshaller)(nil)).Elem()

	t := reflect.TypeOf(*cfg)
	v := reflect.ValueOf(cfg).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fieldType := field.Type

		tag, err := parseTag(field.Tag)
		if err != nil {
			return err
		} 

		raw, err := getRawValue(*tag)

		if reflect.PointerTo(fieldType).Implements(unmarshalerType) {
			
			ptr := reflect.New(fieldType)

			u := ptr.Interface().(EnvUnmarshaller)

			if err := u.UnmarshalEnv(raw); err != nil {
				return err
			}

			value.Set(ptr.Elem())
		
		} else {
			parsed, err := parseValue(raw, fieldType)
			if err != nil {
				return err
			}
			value.Set(parsed)

		}		
	}

	return nil
}
