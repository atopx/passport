package interceptor

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
	"template/internal/system"
)

func Context(service, version string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, param interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		value := system.ServerContextValue{
			Service: service,
			Version: version,
		}
		if data, err := jsoniter.Marshal(param); err == nil {
			value.Trace = jsoniter.Get(data, "header").ToString()
		}
		return handler(context.WithValue(ctx, system.SERVER_CONTEXT_KEY, value), param)
	}
}
