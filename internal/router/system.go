package router

import (
	"gin-scaffold/internal/app/handler/system"
	"github.com/gin-gonic/gin"
)

func initSystemRout(r *gin.Engine) {
	systemRoute := r.Group("")
	{

		userHandler := system.NewUserHandler()
		userGroup := systemRoute.Group("/users")
		{
			userGroup.POST("", userHandler.Create)
			userGroup.PUT("/:id", userHandler.Update)
			userGroup.DELETE("/:id", userHandler.Delete)
			userGroup.GET("/:id", userHandler.GetByID)
			userGroup.GET("", userHandler.List)
		}
	}
}
