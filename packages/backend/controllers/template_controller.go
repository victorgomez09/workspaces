package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/victin09/workspaces/middlewares"
	"github.com/victin09/workspaces/services"
)

func InitTemplateRoutes(r *gin.Engine) *gin.Engine {
	templateGroup := r.Group("templates")
	templateGroup.Use(middlewares.JWTAuthMiddleware())
	templateGroup.GET("", services.GetAllTemplates)
	templateGroup.POST("/run", services.RunTemplate)

	return r
}
