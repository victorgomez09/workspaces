package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victin09/workspaces/database"
	"github.com/victin09/workspaces/models"
	"github.com/victin09/workspaces/utils"
)

type RegisterDto struct {
	Email     string `json:"email" binding:"required"`
	Firstname string `json:"firstName" binding:"required"`
	Lastname  string `json:"lastName" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var data RegisterDto
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := utils.Encrpt(data.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Firstname:   data.Firstname,
		Lastname:    data.Lastname,
		Email:       data.Email,
		Password:    string(passwordHash),
		CreatedDate: time.Now(),
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var data LoginDto

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if utils.ValidateEncrypt(data.Password, user.Password) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad credentials"})
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
