package services

import (
	"auth-service/models"
	"auth-service/utils"
	"errors"
)

func RegisterUser(name string, email string, password string, userType models.UserType) (*models.User, error) {
	if user, _ := models.FindUserByEmail(email); user != nil {
		return nil, errors.New("user is already registered")
	}

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

func LoginUser(email string, password string) (string, error) {
	user, err := models.FindUserByEmail(email)

	if err != nil {
		return "", errors.New("user's account does not exist")
	}

	if !utils.ValidatePassword(user.Password, password) {
		return "", errors.New("email/password is invalid")
	}

	jwtToken, err := utils.GenerateJwtToken(*user)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func AuthenticateUser(token string) (*utils.JwtClaims, error) {
	return utils.ValidateJwtToken(token)
}
