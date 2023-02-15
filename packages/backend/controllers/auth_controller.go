package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/victin09/workspaces/services"
)

func InitAuthRoutes(r *gin.Engine) *gin.Engine {
	authGroup := r.Group("/auth")
	authGroup.POST("/login", services.Login)
	authGroup.POST("/register", services.Register)

	return r
}
