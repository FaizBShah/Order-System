package services

import (
	"auth-service/models"
	"auth-service/utils"
)

func RegisterUser(name string, email string, password string, userType models.UserType) (*models.User, error) {
	hashedPassword, err := utils.GenerateHashFromPassword(password)

	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		UserType: userType,
	}

	return models.CreateUser(&newUser)
}
