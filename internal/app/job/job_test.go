package job

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
	"gin-scaffold/pkg/mysql"
	"testing"
)

func Test(t *testing.T) {
	xxxinit()
	addExampleJob()
}

func xxxinit() {
	config.InitConfig("E:\\workspace\\src\\jz-scraw\\config.yaml")
	mysql.InitMysql()
	logger.InitLogger()
}
