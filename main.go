package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"template/internal/interceptor"
	"template/internal/service"
	"template/logger"
	"template/protocol"
)

func init() {
	// 配置初始化
	config := flag.String("c", "configs/prod.yaml", "config file path.")
	flag.Parse()
	viper.SetConfigFile(*config)
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("load config error: %s", err)
	}
	// 日志初始化
	var loglevel string
	switch viper.GetString("server.env") {
	case "dev":
		loglevel = zap.DebugLevel.String()
	case "prod":
		loglevel = zap.InfoLevel.String()
	default:
		loglevel = zap.InfoLevel.String()
	}
	if err := logger.Setup(loglevel); err != nil {
		panic(err)
	}
}
func main() {
	// 系统初始化
	linstener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("server.port")))
	if err != nil {
		logger.Panic("start server", zap.Error(err))
	}
	server := grpc.NewServer(interceptor.New(viper.GetString("server.name"), os.Getenv("VERSION")))
	// 注册服务
	protocol.RegisterTemplateServiceServer(server, service.NewTemplateService(nil))
	go func() {
		if err = server.Serve(linstener); err != nil {
			logger.Panic("start server", zap.Error(err))
		}
	}()
	// 优雅启停
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
