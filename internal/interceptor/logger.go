package interceptor

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"template/logger"
	"time"
)

func _logger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, param interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		start := time.Now()
		reply, err := handler(ctx, param)
		logger.Info("Interchange", zap.Error(err),
			zap.Object("server", system),
			zap.Duration("cost", time.Since(start)),
			zap.Any("param", param),
		)
		return reply, err
	}
}
