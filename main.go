package main

import (
	"flag"
	"log"
	"passport/internal/handler"
	"passport/internal/service"
	"passport/logger"
	"passport/protocol"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	// 配置初始化
	config := flag.String("c", "conf/config.yaml", "config file path.")
	flag.Parse()
	viper.SetConfigFile(*config)
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("load config error: %s", err)
	}
	// 日志初始化
	var loglevel zapcore.Level
	switch viper.GetString("server.env") {
	case "dev", "debug":
		loglevel = zap.DebugLevel
	case "prod", "production", "release":
		loglevel = zap.WarnLevel
	default:
		loglevel = zap.InfoLevel
	}
	if err := logger.Setup(loglevel); err != nil {
		panic(err)
	}
}

func main() {
	// 初始化基础组件 TODO: mysql, redis, kafka, rabbitmq等基础设施
	db, err := handler.NewMySQLConnect()
	if err != nil {
		logger.Fatal(nil, "connect mysql error", zap.Error(err))
	}

	// 初始化服务
	server := handler.NewServer()
	srv := service.New(db)
	protocol.RegisterPassportServiceServer(server, srv)

	// 启动服务
	handler.StartServer(server)
}
