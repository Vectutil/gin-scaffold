package redis

import (
	"context"
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	// 初始化配置和日志
	config.InitConfig("/home/breelk/workspace/src/gin-scaffold/config.yaml")
	logger.InitLogger()
	Init()
	GetClient().Set(context.Background(), "key1", "test1", 30*time.Second)
	defer func() {

	}()
}

func TestDefer(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("test panic")
}
