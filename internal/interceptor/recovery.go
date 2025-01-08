package interceptor

import (
	"context"
	"passport/internal/common"
	"passport/logger"

	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Recovery() grpc.UnaryServerInterceptor {
	handler := func(ctx context.Context, param interface{}) (err error) {
		track := make([]byte, 1<<16)
		common.GetRuntimeStack(&track)
		logger.Error(ctx, "recovery", zap.Error(err),
			zap.Any("param", param),
			zap.ByteString("track", track),
		)
		return status.Errorf(codes.Internal, "Server Internal Error")
	}
	return recovery.UnaryServerInterceptor(recovery.WithRecoveryHandlerContext(handler))
}
