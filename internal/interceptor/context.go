package interceptor

import (
	"context"
	"passport/internal/common"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
)

type args any

func Context(service, version string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, param interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		value := common.ServerContextValue{
			Service: service,
			Version: version,
		}
		if data, err := jsoniter.Marshal(param); err == nil {
			value.Trace = jsoniter.Get(data, "header").ToString()
		}
		return handler(context.WithValue(ctx, common.ServerContextKey, value), param)
	}
}
