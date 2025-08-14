package script

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
	"math/rand"
	"os/exec"
	"time"
)

func openChrome(urlReq string) {
	// 初始化随机种子
	randomSec := 30 + rand.Intn(5)
	// 获取一个随机数（1-20）
	// 随机等待1-20秒

	logger.Logger.Info("开启了窗口")
	cmd := exec.Command(`cmd`, `/c`, `start`, config.Cfg.System.Chrome, "https://cn.bing.com/search?q="+urlReq)
	cmd.Start()

	// 确保至少等待40秒
	time.Sleep(time.Duration(randomSec) * time.Second)

}
