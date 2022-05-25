package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func _context() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, param interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = context.WithValue(ctx, "server", system)
		return handler(ctx, param)
	}
}
