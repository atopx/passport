package handler

import (
	"fmt"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"passport/internal/interceptor"
	"passport/logger"
	"syscall"
)

func NewServer() *grpc.Server {
	name := viper.GetString("server.name")
	version := viper.GetString("VERSION")
	return grpc.NewServer(
		grpc.UnaryInterceptor(
			middleware.ChainUnaryServer(
				interceptor.Context(name, version),
				interceptor.Recovery(),
				interceptor.Logger(),
			),
		),
	)
}

// StartServer 优雅的启动服务
func StartServer(server *grpc.Server) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("server.port")))
	if err != nil {
		logger.Panic(nil, "start server", zap.Error(err))
	}

	errors := make(chan error)
	go func() {
		if err = server.Serve(listener); err != nil {
			logger.Panic(nil, "start server", zap.Error(err))
			errors <- err
		}
	}()

	controller := make(chan os.Signal)
	signal.Notify(controller, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	select {
	case err = <-errors:
		panic(err)
	case <-controller:
		server.GracefulStop()
		fmt.Println("server stopped, bye!")
	}
}
