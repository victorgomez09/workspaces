package database

import "github.com/victin09/workspaces/models"

func MigrateDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Workspace{})
	DB.AutoMigrate(&models.Template{})
}
