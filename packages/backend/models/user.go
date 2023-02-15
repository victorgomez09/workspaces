package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	Firstname   string         `json:"firstName"`
	Lastname    string         `json:"lastName"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	CreatedDate time.Time      `json:"createdDate"`
	UpdatedDate time.Time      `json:"updatedDate"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
