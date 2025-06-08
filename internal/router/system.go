package router

import (
	"gin-scaffold/internal/app/handler/system"
	"gin-scaffold/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initSystemRout(r *gin.Engine) {
	systemRoute := r.Group("")
	authRouter := systemRoute
	authRouter.Use(middleware.AuthMiddleware())

	{

		userHandler := system.NewUserHandler()

		{
			systemRoute.POST("/login", system.NewAuthHandler().Login)
		}

		userGroup := authRouter.Group("/user")
		{
			userGroup.POST("", userHandler.Create)
			userGroup.PUT("/:id", userHandler.Update)
			userGroup.DELETE("/:id", userHandler.Delete)
			userGroup.GET("/:id", userHandler.GetByID)
			userGroup.GET("", userHandler.List)
		}
	}
}
