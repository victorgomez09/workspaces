package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/victin09/workspaces/middlewares"
	"github.com/victin09/workspaces/services"
)

func InitUserRoutes(r *gin.Engine) *gin.Engine {
	userGroup := r.Group("/users")
	userGroup.Use(middlewares.JWTAuthMiddleware())
	userGroup.GET("/", services.GetUsers)
	userGroup.GET("/me", services.GetMe)
	userGroup.GET("/:id", services.GetUserById)
	userGroup.GET("/email/:email", services.GetUserByEmail)
	userGroup.PUT("/:id", services.UpdateUser)

	return r
}
