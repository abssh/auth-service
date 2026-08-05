package grpc

type GrpcConfig interface {
	GetGrpcHost() string
	GetGrpcPort() int
}