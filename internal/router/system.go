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

	userHandler := system.NewUserHandler()

	{
		systemRoute.POST("/login", system.NewAuthHandler().Login)
	}

	userGroup := authRouter.Group("/user")
	{
		userGroup.POST("", userHandler.Create)
		userGroup.PUT("/:id", userHandler.Update)
		userGroup.DELETE("/:id", userHandler.Delete)
		userGroup.GET("/:id", userHandler.GetById)
		userGroup.GET("", userHandler.List)
	}

	department := authRouter.Group("/department")
	{
		h := system.NewDepartmentHandler()
		department.POST("", h.Create)
		department.PUT("/:id", h.Update)
		department.DELETE("/:id", h.Delete)
		department.GET("/:id", h.GetById)
		department.GET("", h.List)
		department.GET("/tree", h.GetTree)
	}

	role := authRouter.Group("/role")
	{
		h := system.NewRoleHandler()
		role.POST("", h.Create)
		role.PUT("/:id", h.Update)
		role.DELETE("/:id", h.Delete)
		role.GET("/:id", h.GetById)
		role.GET("", h.List)
	}
}
