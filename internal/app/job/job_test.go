package job

import (
	"jz-scraw/internal/config"
	"jz-scraw/pkg/logger"
	"jz-scraw/pkg/mysql"
	"testing"
)

func Test(t *testing.T) {
	xxxinit()
	addRunStartSync()
}

func Test2(t *testing.T) {
	xxxinit()
	searchJobMqOne()
}

func xxxinit() {
	config.InitConfig("E:\\workspace\\src\\jz-scraw\\config.yaml")
	mysql.InitMysql()
	logger.InitLogger()
}
