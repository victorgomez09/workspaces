package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victin09/workspaces/database"
	"github.com/victin09/workspaces/models"
	"github.com/victin09/workspaces/utils"
)

type CreateWorkspaceInput struct {
	Name     string `json:"name" binding:"required"`
	Template string `json:"template" binding:"required"`
}

func FindAllWorkspaces(c *gin.Context) {
	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var workspaces []models.Workspace
	if err := database.DB.Preload("User").Where("user_id = ?", user.ID).Find(&workspaces).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workspaces)
}

func FindWorkspaceById(c *gin.Context) {
	id := c.Params.ByName("id")

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var workspace models.Workspace
	if err := database.DB.Preload("User").Where("user_id = ?", user.ID).Where("id = ?", id).Find(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workspace)
}

func CreateWorkspace(c *gin.Context) {
	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var data CreateWorkspaceInput
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var exists bool
	var workspace models.Workspace
	if err := database.DB.Model(&workspace).Select("count(name) > 0").Where("name = ?", data.Name).
		Where("user_id = ?", user.ID).Find(&exists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Workspace already exists"})
		return
	}

	workspaceNew := models.Workspace{
		Name:        data.Name,
		Template:    data.Template,
		User:        user,
		CreatedDate: time.Now(),
	}

	database.DB.Create(&workspaceNew)
	c.JSON(http.StatusOK, workspaceNew)
}
