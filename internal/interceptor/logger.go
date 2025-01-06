package interceptor

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"passport/logger"
	"time"
)

func Logger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, param interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		start := time.Now()
		reply, err := handler(ctx, param)
		logger.Info(ctx, "interchange", zap.Error(err),
			zap.Duration("cost", time.Since(start)),
			zap.Any("param", param),
			zap.Any("reply", reply),
		)
		return reply, err
	}
}
