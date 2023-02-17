package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/victin09/workspaces/middlewares"
	"github.com/victin09/workspaces/services"
)

func InitWorkspaceRoutes(r *gin.Engine) *gin.Engine {
	workspaceGroup := r.Group("/workspaces")
	workspaceGroup.Use(middlewares.JWTAuthMiddleware())
	workspaceGroup.GET("", services.FindAllWorkspaces)
	workspaceGroup.GET("/:id", services.FindWorkspaceById)
	workspaceGroup.POST("", services.CreateWorkspace)

	return r
}
