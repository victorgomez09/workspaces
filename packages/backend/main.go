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
	r := gin.Default()

	r.Use(cors.Default())

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.RunInstaller()

	database.InitSqlite()
	database.MigrateDatabase()

	controllers.InitAuthRoutes(r)
	controllers.InitUserRoutes(r)
	controllers.InitTemplateRoutes(r)

	r.Run()
}
