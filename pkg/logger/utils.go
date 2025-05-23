package logger

import (
	"go.uber.org/zap"
	"time"
)

func Error(err error, text string) {
	ErrorLogger.Error(text, zap.Error(err))                     // 记录到 error.log
	SlowLogger.Warn("慢接口", zap.Duration("cost", 2*time.Second)) // 记录到 slow.log
}

func Slow() {

}
