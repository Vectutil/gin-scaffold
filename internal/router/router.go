package router

import (
	"gin-scaffold/internal/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter(r *gin.Engine) {
	//apiRouter := r.Group("api/v1")

	//commentRoute := apiRouter.Group("/comment")
	{
		//commentRoute.GET("/list", GetCommentList)
	}

	// 使用慢日志中间件，阈值设置为 2 秒
	r.Use(middleware.SlowLogMiddleware(2 * time.Second))

	// 其他路由配置...
	r.GET("/ping", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
