package utils

import (
	"auth-service/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJwtTokenWorkCorrectly(t *testing.T) {
	user := models.User{
		ID:       1,
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: models.Admin,
	}

	token, err1 := GenerateJwtToken(user)
	claims, err2 := ValidateJwtToken(token)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotNil(t, claims)
	assert.Equal(t, user.ID, claims.Id)
	assert.Equal(t, user.Email, claims.Email)
	assert.Equal(t, user.UserType, claims.UserType)
}

func TestValidateTokenWorkCorrectly(t *testing.T) {
	user := models.User{
		ID:       1,
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: models.Admin,
	}

	token, err1 := GenerateJwtToken(user)
	claims, err2 := ValidateJwtToken(token)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotNil(t, claims)
	assert.Equal(t, user.ID, claims.Id)
	assert.Equal(t, user.Email, claims.Email)
	assert.Equal(t, user.UserType, claims.UserType)
}

func TestValidateTokenThrowErrorIfJWTIsInvalid(t *testing.T) {
	token := "random_token"
	claims, err := ValidateJwtToken(token)

	assert.Error(t, err)
	assert.Nil(t, claims)
}
