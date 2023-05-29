package dto

import (
	proto "api-gateway/proto/auth"
)

type RegisterUserRequest struct {
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	UserType proto.UserType `json:"user_type"`
}

type RegisterUserResponse struct {
	Id       int64          `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	UserType proto.UserType `json:"user_type"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}
