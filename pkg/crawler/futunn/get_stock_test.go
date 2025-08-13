package futunn

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
	"testing"
)

func TestConnectHtml(t *testing.T) {
	logger.InitLogger()
	config.InitConfig("E:\\code\\workspace\\project\\gin-scaffold\\config.yaml")
	ConnectHtml()
}
