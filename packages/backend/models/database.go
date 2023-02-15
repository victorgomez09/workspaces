package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
	db, error := gorm.Open(sqlite.Open("./database.db"), &gorm.C)
}
