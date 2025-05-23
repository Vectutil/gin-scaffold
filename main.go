package main

import (
	"context"
	"gin-scaffold/internal/app/job"
	"gin-scaffold/internal/config"
	"gin-scaffold/internal/router"
	"gin-scaffold/pkg/logger"
	"gin-scaffold/pkg/mysql"
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
	// 1.创建路由
	r := gin.Default()
	router.InitRouter(r)
	// 优雅退出
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + config.Cfg.System.Port,
		Handler: r,
	}

	if config.Cfg.System.Migration {
		mysql.Migration()
	}

	go func() {
		zap.L().Info("项目启动")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	job.StartCronJob()
	<-stop
	zap.L().Info("项目启动失败")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")

}

func init() {
	config.InitConfig("")
	logger.InitLogger()
	mysql.InitMysql()

	migration()

	defer zap.L().Sync()
	logger.Logger.Info("日志初始化完成")
}

func migration() {
	if config.Cfg.System.Migration {
		mysql.Migration()
	}
}
