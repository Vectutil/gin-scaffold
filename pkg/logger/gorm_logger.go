package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type ZapGormLogger struct {
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
}

func NewGormLogger(threshold time.Duration) logger.Interface {
	return &ZapGormLogger{
		SlowThreshold: threshold,
		LogLevel:      logger.Info, // 你可以控制 gorm 默认日志输出级别
	}
}

func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	Logger.Sugar().Infof(msg, data...)
}

func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	Logger.Sugar().Warnf(msg, data...)
}

func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	ErrorLogger.Sugar().Errorf(msg, data...)
}

func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []zap.Field{
		zap.String("file", utils.FileWithLineNum()),
		zap.Duration("elapsed", elapsed),
		zap.Int64("rows", rows),
		zap.String("sql", sql),
	}

	// 错误 SQL -> error.log
	if err != nil && l.LogLevel >= logger.Error {
		ErrorLogger.Error("SQL 错误", append(fields, zap.Error(err))...)
	}

	// 慢 SQL -> slow.log
	if elapsed > l.SlowThreshold {
		SlowLogger.Warn("慢 SQL", fields...)
	}

	// 所有 SQL -> app.log
	if l.LogLevel >= logger.Info {
		Logger.Info("SQL 执行", fields...)
	}

}
