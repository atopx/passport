package handler

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"template/logger"
	"time"
)

// NewMySQLConnect 创建MySQL连接
func NewMySQLConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.dbname"),
	)
	logger.Debug(nil, "new mysql connect", zap.String("dsn", dsn))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.NewDBLogger()})
	if err != nil {
		return db, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}
	sqlDB.SetMaxIdleConns(viper.GetInt("database.max_idle"))
	sqlDB.SetMaxOpenConns(viper.GetInt("database.max_open"))
	sqlDB.SetConnMaxLifetime(viper.GetDuration("database.max_life") * time.Second)
	return db, err
}
