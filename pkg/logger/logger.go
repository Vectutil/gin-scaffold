package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Logger *zap.Logger
var SlowLogger *zap.Logger
var ErrorLogger *zap.Logger

func InitLogger() {
	encoder := getEncoder()

	// app.log: 记录所有日志
	appWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./cache/logs/info/app_" + time.Now().Format("2006-01-02") + ".log",
		MaxSize:    20,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   true,
	})
	// 错误日志单独输出
	errorWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./cache/logs/error/error_" + time.Now().Format("2006-01-02") + ".log",
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   true,
	})

	// 慢日志单独输出
	slowWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./cache/logs/slow/slow_" + time.Now().Format("2006-01-02") + ".log",
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   true,
	})

	consoleWriter := zapcore.AddSync(os.Stdout)

	// 核心组合
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(appWriter, consoleWriter), zapcore.DebugLevel), // 记录全部日志
		zapcore.NewCore(encoder, errorWriter, zapcore.ErrorLevel),                                           // 错误日志单独
	)

	Logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Logger)

	ErrorLogger = zap.New(zapcore.NewCore(encoder, errorWriter, zapcore.ErrorLevel))
	SlowLogger = zap.New(zapcore.NewCore(encoder, slowWriter, zapcore.InfoLevel)) // 慢日志你手动打点

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

//func getEncoder() zapcore.Encoder {
//	cfg := zap.NewProductionEncoderConfig()
//	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
//	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
//	cfg.EncodeCaller = zapcore.ShortCallerEncoder
//	return zapcore.NewJSONEncoder(cfg)
//}
