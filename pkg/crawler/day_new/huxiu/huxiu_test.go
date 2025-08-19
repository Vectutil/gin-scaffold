package huxiu

import (
	"context"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
	"testing"
)

func TestHuxiu(t *testing.T) {
	logger.InitLogger()
	config.InitConfig("E:\\workspace\\src\\gin-scaffold\\config.yaml")
	//config.InitConfig("E:\\code\\workspace\\project\\gin-scaffold\\config.yaml")
	ctx := context.Background()
	HuxiuNews(ctx)
}
