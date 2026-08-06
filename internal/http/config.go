package http

type HttpConfig interface {
	GetHttpHost() string
	GetHttpPort() int
}
