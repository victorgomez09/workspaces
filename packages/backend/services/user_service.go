package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victin09/workspaces/database"
	"github.com/victin09/workspaces/models"
)

type UpdateUserDto struct {
	Email     string `json:"email"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Password  string `json:"password"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")

	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Params.ByName("email")

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var data UpdateUserDto
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(models.User{
		Email: data.Email, Firstname: data.Firstname, Lastname: data.Lastname, UpdatedDate: time.Now(),
	})
	c.JSON(http.StatusOK, user)
}
