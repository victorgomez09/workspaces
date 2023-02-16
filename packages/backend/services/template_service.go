package services

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/victin09/workspaces/utils"
)

type TemplateDto struct {
	Name string `json:"templateName" binding:"required"`
}

func InitTemplate(c *gin.Context) {

}

func GetAllTemplates(c *gin.Context) {
	terraformDir := filepath.Join(os.Getenv("HOME"), ".workspaces/terraform")
	entries, err := utils.ReadDir(terraformDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	templates := []string{}
	for _, e := range entries {
		templates = append(templates, e.Name())
	}

	c.JSON(http.StatusOK, templates)
}

func RunTemplate(c *gin.Context) {
	var data TemplateDto

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	templateDir := filepath.Join(os.Getenv("HOME"), ".workspaces/terraform/"+data.Name)
	version, err := utils.RunTemplate(templateDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, version)
}
