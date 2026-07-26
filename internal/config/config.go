package config

import "reflect"

type Config struct {
	HttpHost string `env:"HTTP_HOST" default:""`
	HttpPort int    `env:"HTTP_PORT" default:"8080"`
}

func (cfg *Config) Load() error {
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

		parsed, err := parseValue(*tag, fieldType)
		if err != nil {
			return err
		}

		value.Set(parsed)
	}

	return nil
}
