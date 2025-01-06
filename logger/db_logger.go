package logger

import (
	"context"
	"go.uber.org/zap"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"passport/internal/common"
	"time"
)

type DBLogger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  glogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func NewDBLogger() glogger.Interface {
	return &DBLogger{
		ZapLogger:                 logger,
		LogLevel:                  glogger.Warn,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: true,
	}
}

func (log *DBLogger) SetAsDefault() {
	glogger.Default = log
}

func (log *DBLogger) LogMode(level glogger.LogLevel) glogger.Interface {
	return &DBLogger{
		ZapLogger:                 log.ZapLogger,
		SlowThreshold:             log.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          log.SkipCallerLookup,
		IgnoreRecordNotFoundError: log.IgnoreRecordNotFoundError,
	}
}

func (log *DBLogger) Info(ctx context.Context, message string, fields ...interface{}) {
	if log.LogLevel >= glogger.Info {
		log.WithContext(ctx).Sugar().Infof(message, fields...)
	}

}

func (log *DBLogger) Warn(ctx context.Context, message string, fields ...interface{}) {
	if log.LogLevel >= glogger.Warn {
		log.WithContext(ctx).Sugar().Warnf(message, fields...)
	}
}

func (log *DBLogger) Error(ctx context.Context, message string, fields ...interface{}) {
	if log.LogLevel >= glogger.Error {
		log.WithContext(ctx).Sugar().Errorf(message, fields...)
	}
}

func (log *DBLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, row := fc()
	log.WithContext(ctx).Info("trace",
		zap.String("file", utils.FileWithLineNum()),
		zap.Int64("row", row), zap.String("sql", sql),
		zap.Duration("cost", elapsed),
		zap.Error(err),
	)
}

func (log *DBLogger) WithContext(ctx context.Context) *zap.Logger {
	if ctx != nil {
		switch value := ctx.Value(common.ServerContextKey).(type) {
		case common.ServerContextValue:
			return log.ZapLogger.With(zap.Object(common.ServerContextKey, value))
		}
	}
	return log.ZapLogger
}
