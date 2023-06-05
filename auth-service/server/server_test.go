package server

import (
	"auth-service/models"
	proto "auth-service/proto/auth"
	"auth-service/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var server GRPCServer

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:yourDbName?mode=memory&cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	server = GRPCServer{}
	models.InitUserModel(db)

	return db
}

func teardownDatabase(db *gorm.DB) {
	_ = db.Migrator().DropTable(&models.User{})
	sql, _ := db.DB()
	sql.Close()
}

func TestShouldInitUserModelWork(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	err := db.AutoMigrate(&models.User{})
	assert.NoError(t, err)
}

func TestShouldRegisterUserWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newUser := &models.User{
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: models.Admin,
	}

	registeredUser, err := server.RegisterUser(context.Background(), &proto.RegisterUserRequest{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		UserType: proto.UserType(newUser.UserType),
	})

	assert.NoError(t, err)
	assert.NotNil(t, registeredUser)
	assert.Equal(t, int64(1), registeredUser.Id)
	assert.Equal(t, newUser.Name, registeredUser.Name)
	assert.Equal(t, newUser.Email, registeredUser.Email)
}

func TestShouldRegisterUserThrowAnErrorIfUserIsAlreadyRegistered(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newUser := &models.User{
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: models.Admin,
	}

	registeredUser1, err1 := server.RegisterUser(context.Background(), &proto.RegisterUserRequest{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		UserType: proto.UserType(newUser.UserType),
	})
	registeredUser2, err2 := server.RegisterUser(context.Background(), &proto.RegisterUserRequest{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		UserType: proto.UserType(newUser.UserType),
	})

	assert.NoError(t, err1)
	assert.NotNil(t, registeredUser1)
	assert.Equal(t, newUser.Name, registeredUser1.Name)
	assert.Equal(t, newUser.Email, registeredUser1.Email)
	assert.Nil(t, registeredUser2)
	assert.Error(t, err2)
	assert.Equal(t, "user is already registered", err2.Error())
}

func TestShouldLoginUserWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newUser := &models.User{
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: models.Admin,
	}

	registeredUser, err1 := server.RegisterUser(context.Background(), &proto.RegisterUserRequest{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		UserType: proto.UserType(newUser.UserType),
	})

	res, err2 := server.LoginUser(context.Background(), &proto.LoginUserRequest{
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	claims, err3 := utils.ValidateJwtToken(res.Token)

	assert.NoError(t, err1)
	assert.NotNil(t, registeredUser)
	assert.NoError(t, err2)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Token)
	assert.NoError(t, err3)
	assert.NotNil(t, claims)
	assert.Equal(t, registeredUser.Id, claims.Id)
	assert.Equal(t, newUser.Email, claims.Email)
	assert.Equal(t, newUser.UserType, claims.UserType)
}

func TestLoginUserThrowAnErrorUserDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	res, err := server.LoginUser(context.Background(), &proto.LoginUserRequest{
		Email:    "test@test.com",
		Password: "testPassword",
	})

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, "user's account does not exist", err.Error())
}

func TestLoginUserThrowAnErrorIfPasswordIsIncorrect(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newUser := &models.User{
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: models.Admin,
	}

	server.RegisterUser(context.Background(), &proto.RegisterUserRequest{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		UserType: proto.UserType(newUser.UserType),
	})

	res, err := server.LoginUser(context.Background(), &proto.LoginUserRequest{
		Email:    newUser.Email,
		Password: "wrongPassword",
	})

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Equal(t, "email/password is invalid", err.Error())
}

func TestShouldAuthenticateUserWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newUser := &models.User{
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: models.Admin,
	}

	registeredUser, err1 := server.RegisterUser(context.Background(), &proto.RegisterUserRequest{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		UserType: proto.UserType(newUser.UserType),
	})

	res, err2 := server.LoginUser(context.Background(), &proto.LoginUserRequest{
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	claims, err3 := server.AuthenticateUser(context.Background(), &proto.AuthenticateUserRequest{
		Token: res.Token,
	})

	assert.NoError(t, err1)
	assert.NotNil(t, registeredUser)
	assert.NoError(t, err2)
	assert.NotEmpty(t, res.Token)
	assert.NoError(t, err3)
	assert.NotNil(t, claims)
	assert.Equal(t, registeredUser.Id, claims.Id)
	assert.Equal(t, newUser.Email, claims.Email)
}

func TestShouldAuthenticateUserThrowErrorIfInvalidAuthentication(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newUser := &models.User{
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: models.Admin,
	}

	server.RegisterUser(context.Background(), &proto.RegisterUserRequest{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		UserType: proto.UserType(newUser.UserType),
	})
	server.LoginUser(context.Background(), &proto.LoginUserRequest{
		Email:    newUser.Email,
		Password: newUser.Password,
	})

	claims, err := server.AuthenticateUser(context.Background(), &proto.AuthenticateUserRequest{
		Token: "invalid_token",
	})

	assert.Nil(t, claims)
	assert.Error(t, err)
}
