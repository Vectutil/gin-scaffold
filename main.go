package main

import (
	"context"
	"errors"
	"gin-scaffold/internal/app/job"
	"gin-scaffold/internal/config"
	"gin-scaffold/internal/router"
	"gin-scaffold/pkg/logger"
	"gin-scaffold/pkg/mysql"
	"gin-scaffold/pkg/redis"
	"gin-scaffold/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if config.Cfg.System.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// 6. 创建 Gin 路由引擎
	r := gin.New()
	r.Use(gin.Recovery()) // 可选：防止 panic 导致服务崩溃
	router.InitRouter(r)

	// 7. 构建 HTTP Server 实例
	srv := &http.Server{
		Addr:    ":" + config.Cfg.System.Port,
		Handler: r,
	}
	// 8. 启动 HTTP 服务
	go func() {
		utils.RunInfo()
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Info("------------------项目启动失败------------------")
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 9. 监听系统信号，优雅退出
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.Logger.Info("------------------项目关闭------------------")

	// 10. 执行优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Info("Server forced to shutdown:" + err.Error())
	}
}

func init() {

	// 1. 初始化配置（必须最早）
	config.InitConfig("")

	// 2. 初始化日志（依赖配置，尽早初始化以便后续使用日志）
	logger.InitLogger()
	defer zap.L().Sync()

	// 3. 初始化数据库连接（如 MySQL）
	mysql.InitMysql()
	if config.Cfg.System.Migration {
		mysql.Migration()
	}

	// 4. 初始化 Redis
	redis.Init()

	// 5. 启动定时任务
	job.StartCronJob()

	//defer zap.L().Sync()
}
