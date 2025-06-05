package router

import (
	"gin-scaffold/internal/app/handler"
	"gin-scaffold/internal/middleware"
	"github.com/gin-gonic/gin"
	"time"

	_ "gin-scaffold/docs" // main 文件中导入 docs 包

	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {
	//apiRouter := r.Group("api/v1")

	//commentRoute := apiRouter.Group("/comment")
	{
		//commentRoute.GET("/list", GetCommentList)
	}

	// 使用慢日志中间件，阈值设置为 3 秒
	r.Use(middleware.SlowLogMiddleware(3 * time.Second))
	r.Use(middleware.CorsMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	initSystemRout(r)

	// 其他路由配置...
	testGroup := r.Group("test")
	{
		test := handler.NewTest()
		testGroup.POST("", test.TestHandle)
		testGroup.GET("", test.TestHandleGet)
		testGroup.PUT("", test.TestHandlePut)
		testGroup.DELETE(":id", test.TestHandleDelete)
	}

}
