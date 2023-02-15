package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitSqlite() {
	var error error
	DB, error = gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if error != nil {
		log.Fatal(error.Error())
	}
}
