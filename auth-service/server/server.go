package server

import (
	"auth-service/models"
	proto "auth-service/proto/auth"
	"auth-service/services"
	"context"
)

type GRPCServer struct {
	proto.UnimplementedAuthServiceServer
}

func (s *GRPCServer) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	user, err := services.RegisterUser(req.Name, req.Email, req.Password, models.UserType(req.UserType.String()))

	if err != nil {
		return nil, err
	}

	var userType proto.UserType

	if user.UserType == models.Admin {
		userType = proto.UserType_ADMIN
	} else {
		userType = proto.UserType_REGULAR
	}

	return &proto.RegisterUserResponse{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		UserType: userType,
	}, nil
}

func (s *GRPCServer) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {
	token, err := services.LoginUser(req.Email, req.Password)

	if err != nil {
		return nil, err
	}

	return &proto.LoginUserResponse{
		Token: token,
	}, nil
}

func (s *GRPCServer) AuthenticateUser(ctx context.Context, req *proto.AuthenticateUserRequest) (*proto.AuthenticateUserResponse, error) {
	claims, err := services.AuthenticateUser(req.Token)

	if err != nil {
		return nil, err
	}

	var userType proto.UserType

	if claims.UserType == models.Admin {
		userType = proto.UserType_ADMIN
	} else {
		userType = proto.UserType_REGULAR
	}

	return &proto.AuthenticateUserResponse{
		Id:       claims.Id,
		Email:    claims.Email,
		UserType: userType,
	}, nil
}
