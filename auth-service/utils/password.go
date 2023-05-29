package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 11)

	if err != nil {
		return "", errors.New("error in generating hash of password")
	}

	return string(bytes), nil
}

func ValidatePassword(hashedPassword string, rawPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)) == nil
}
