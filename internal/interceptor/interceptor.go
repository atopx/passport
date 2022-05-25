package interceptor

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"template/internal/model"
)

var system model.System

func New(server, version string) grpc.ServerOption {
	system.Server = server
	system.Version = version
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(_context(), _recovery(), _logger()))
}
