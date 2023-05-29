package utils

import (
	"auth-service/models"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	jwt.StandardClaims
	Id       int64
	Email    string
	UserType models.UserType
}

func GenerateJwtToken(user models.User) (string, error) {
	expirationTime, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRATION"), 10, 64)

	claims := &JwtClaims{
		Id:       user.ID,
		Email:    user.Email,
		UserType: user.UserType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(expirationTime)).Unix(),
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJwtToken(signedToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}
