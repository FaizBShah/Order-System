package models

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey;AUTO_INCREMENT"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email;unique"`
	Password string `gorm:"column:password"`
	UserType string `gorm:"column:user_type"`
}

func InitProductModel(dbInstance *gorm.DB) {
	db = dbInstance
	db.AutoMigrate(&User{})
}
