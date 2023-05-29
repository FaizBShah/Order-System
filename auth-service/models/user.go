package models

import (
	"errors"

	"gorm.io/gorm"
)

type UserType string

const (
	Admin   UserType = "ADMIN"
	Regular UserType = "REGULAR"
)

var (
	db *gorm.DB
)

type User struct {
	gorm.Model
	ID       int64    `gorm:"primarykey;AUTO_INCREMENT"`
	Name     string   `gorm:"column:name"`
	Email    string   `gorm:"column:email;unique"`
	Password string   `gorm:"column:password"`
	UserType UserType `gorm:"column:user_type"`
}

func InitProductModel(dbInstance *gorm.DB) {
	db = dbInstance
	db.AutoMigrate(&User{})
}

func CreateUser(newUser *User) (*User, error) {
	if newUser == nil {
		return nil, errors.New("invalid product")
	}

	if err := db.Create(newUser).Error; err != nil {
		return nil, errors.New("error in creating a new product")
	}

	return newUser, nil
}

func FindUserByEmail(email string) (*User, error) {
	var user *User

	if len(email) == 0 {
		return nil, errors.New("email is empty")
	}

	if err := db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
