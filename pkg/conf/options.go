package conf

import "go.uber.org/zap"

var (
	Logger      *zap.Logger
	LoggerLevel string
	RunMode     string
	Host        string
	Port        int
)

type apiConfig struct {
	Port   int
	Prefix string
}

var API = &apiConfig{
	Port:   83,
	Prefix: "/awecloud/bmq/api",
}

type grpcConfig struct {
	Port string
}

var GRPC = &grpcConfig{
	Port: "81",
}

type messageConfig struct {
	Port   int
	Token  string
	Prefix string
}

var Message = &messageConfig{
	Port:   82,
	Token:  "bmq",
	Prefix: "/awecloud/bmq/message",
}

var Server = &ServerConfig{}
