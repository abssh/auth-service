package config

func (cfg *Config) GetGrpcHost() string {
	return cfg.GrpcHost
}

func (cfg *Config) GetGrpcPort() int {
	return cfg.GrpcPort
}