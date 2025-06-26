package logger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

const (
	colorDebug = "\033[36m" // Cyan
	colorInfo  = "\033[32m" // Green
	colorWarn  = "\033[33m" // Yellow
	colorError = "\033[31m" // Red
	colorReset = "\033[0m"  // Reset
)

var Logger *zap.Logger

//var SlowLogger *zap.Logger
//var ErrorLogger *zap.Logger

func InitLogger() {

	// app.log: 记录所有日志
	writer, _ := rotatelogs.New(
		"./cache/logs/info/app_%Y-%m-%d.log", // 自动按天分割
		//rotatelogs.WithLinkName("./cache/logs/info/app_current.log"), // 软链接最新日志
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 最大保留30天
		rotatelogs.WithRotationTime(24*time.Hour), // 每天滚动一次
	)
	appWriter := zapcore.AddSync(writer)

	// 错误日志单独输出
	ewriter, _ := rotatelogs.New(
		"./cache/logs/err/err_%Y-%m-%d.log", // 自动按天分割
		//rotatelogs.WithLinkName("./cache/logs/err/err_current.log"), // 软链接最新日志
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 最大保留30天
		rotatelogs.WithRotationTime(24*time.Hour), // 每天滚动一次
	)
	errorWriter := zapcore.AddSync(ewriter)

	// 慢日志单独输出
	swriter, _ := rotatelogs.New(
		"./cache/logs/info/slow_%Y-%m-%d.log", // 自动按天分割
		//rotatelogs.WithLinkName("./cache/logs/slow/slow_current.log"), // 软链接最新日志
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 最大保留30天
		rotatelogs.WithRotationTime(24*time.Hour), // 每天滚动一次
	)
	slowWriter := zapcore.AddSync(swriter)

	consoleWriter := zapcore.AddSync(os.Stdout)

	fileEncoder := getFileEncoder()
	consoleEncoder := getConsoleEncoder()

	warnOnly := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	// 核心组合
	core := zapcore.NewTee(
		// 文件输出：JSON 格式
		zapcore.NewCore(fileEncoder, appWriter, zapcore.DebugLevel),   // app.log 记录所有日志
		zapcore.NewCore(fileEncoder, errorWriter, zapcore.ErrorLevel), // error.log 记录错误日志
		zapcore.NewCore(fileEncoder, slowWriter, warnOnly),            // slow.log 慢日志

		// 控制台输出（只对 consoleWriter 使用 consoleEncoder）
		zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel),
	)

	Logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Logger)
}

func getFileEncoder() zapcore.Encoder {
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

func getConsoleEncoder() zapcore.Encoder {
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

	// 带颜色的级别输出
	encoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		var color string
		switch level {
		case zapcore.DebugLevel:
			color = colorDebug
		case zapcore.InfoLevel:
			color = colorInfo
		case zapcore.WarnLevel:
			color = colorWarn
		case zapcore.ErrorLevel:
			color = colorError
		default:
			color = colorReset
		}
		enc.AppendString(color + level.CapitalString() + colorReset)
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

//func getEncoder() zapcore.Encoder {
//	cfg := zap.NewProductionEncoderConfig()
//	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
//	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
//	cfg.EncodeCaller = zapcore.ShortCallerEncoder
//	return zapcore.NewJSONEncoder(cfg)
//}
