package interceptor

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
	"template/internal/system"
	"template/logger"
)

func Recovery() grpc.UnaryServerInterceptor {
	handler := func(ctx context.Context, param interface{}) (err error) {
		track := make([]byte, 1<<16)
		runtime.Stack(track, false)
		logger.Error(ctx, "recovery", zap.Error(err),
			zap.Object("service", ctx.Value(system.SERVER_CONTEXT_KEY).(system.ServerContextValue)),
			zap.Any("param", param),
			zap.ByteString("track", track),
		)
		return status.Errorf(codes.Unknown, "Server Internal Error")
	}
	return grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(handler))
}
