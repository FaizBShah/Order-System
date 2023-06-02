package models

import (
	"errors"

	"gorm.io/gorm"
)

type UserType int32

const (
	Admin UserType = iota
	Regular
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

func InitUserModel(dbInstance *gorm.DB) {
	db = dbInstance
	db.AutoMigrate(&User{})
}

func CreateUser(newUser *User) (*User, error) {
	if newUser == nil {
		return nil, errors.New("invalid user")
	}

	if err := db.Create(newUser).Error; err != nil {
		return nil, errors.New("error in creating a new user")
	}

	return newUser, nil
}

func FindUserByEmail(email string) (*User, error) {
	var user *User

	if len(email) == 0 {
		return nil, errors.New("email is empty")
	}

	result := db.Where("email = ?", email).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return user, nil
}
