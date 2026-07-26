package config

type EnvUnmarshaller interface {
	UnmarshalEnv(string) error
}