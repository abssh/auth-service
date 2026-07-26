package http

type ServerConfig interface {
	GetHttpHost() string
	GetHttpPort() int
}