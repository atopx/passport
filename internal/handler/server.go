package handler

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"template/internal/interceptor"
	"template/logger"
)

func NewServer() *grpc.Server {
	return grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptor.Context(viper.GetString("server.name"), os.Getenv("VERSION")),
		interceptor.Recovery(),
		interceptor.Logger(),
	)))
}

// StartServer 优雅的启动服务
func StartServer(server *grpc.Server) {
	linstener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("server.port")))
	if err != nil {
		logger.Panic(nil, "start server", zap.Error(err))
	}
	go func() {
		if err = server.Serve(linstener); err != nil {
			logger.Panic(nil, "start server", zap.Error(err))
		}
	}()

	controller := make(chan os.Signal, 1)
	signal.Notify(controller, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-controller
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			server.Stop()
		}
	}
}
