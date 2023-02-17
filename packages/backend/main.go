package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/victin09/workspaces/controllers"
	"github.com/victin09/workspaces/database"
	"github.com/victin09/workspaces/utils"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitSqlite()
	database.MigrateDatabase()

	utils.RunInstaller()

	controllers.InitAuthRoutes(router)
	controllers.InitUserRoutes(router)
	controllers.InitTemplateRoutes(router)
	controllers.InitWorkspaceRoutes(router)

	router.Run()
}
