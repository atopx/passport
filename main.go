package main

import (
	"flag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"template/internal/handler"
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
		logger.Warn(nil, "run mode is develop.")
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
	// 初始化基础组件 TODO: mysql, redis, kafka, rabbitmq等基础设施
	db, err := handler.NewMySQLConnect()
	if err != nil {
		logger.Fatal(nil, "new mysql error", zap.Error(err))
	}

	// 初始化服务
	server := handler.NewServer()
	srv := service.NewTemplateService(db)
	protocol.RegisterTemplateServiceServer(server, srv)

	// 启动服务
	handler.StartServer(server)
}
