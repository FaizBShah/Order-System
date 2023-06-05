package services

import (
	"auth-service/models"
	"auth-service/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:yourDbName?mode=memory&cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

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

	registeredUser, err := RegisterUser(newUser.Name, newUser.Email, newUser.Password, newUser.UserType)

	assert.NoError(t, err)
	assert.NotNil(t, registeredUser)
	assert.Equal(t, int64(1), registeredUser.ID)
	assert.Equal(t, newUser.Name, registeredUser.Name)
	assert.Equal(t, newUser.Email, registeredUser.Email)
	assert.NotEqual(t, newUser.Password, registeredUser.Password)
	assert.True(t, utils.ValidatePassword(registeredUser.Password, newUser.Password))
	assert.Equal(t, newUser.UserType, registeredUser.UserType)
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

	registeredUser1, err1 := RegisterUser(newUser.Name, newUser.Email, newUser.Password, newUser.UserType)
	registeredUser2, err2 := RegisterUser(newUser.Name, newUser.Email, newUser.Password, newUser.UserType)

	assert.NoError(t, err1)
	assert.NotNil(t, registeredUser1)
	assert.Equal(t, newUser.Name, registeredUser1.Name)
	assert.Equal(t, newUser.Email, registeredUser1.Email)
	assert.True(t, utils.ValidatePassword(registeredUser1.Password, newUser.Password))
	assert.Equal(t, newUser.UserType, registeredUser1.UserType)
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

	registeredUser, err1 := RegisterUser(newUser.Name, newUser.Email, newUser.Password, newUser.UserType)

	token, err2 := LoginUser(newUser.Email, newUser.Password)
	claims, err3 := utils.ValidateJwtToken(token)

	assert.NoError(t, err1)
	assert.NotNil(t, registeredUser)
	assert.NoError(t, err2)
	assert.NotEmpty(t, token)
	assert.NoError(t, err3)
	assert.NotNil(t, claims)
	assert.Equal(t, registeredUser.ID, claims.Id)
	assert.Equal(t, newUser.Email, claims.Email)
	assert.Equal(t, newUser.UserType, claims.UserType)
}

func TestLoginUserThrowAnErrorUserDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	token, err := LoginUser("test@example.com", "testPassword")

	assert.Error(t, err)
	assert.Empty(t, token)
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

	RegisterUser(newUser.Name, newUser.Email, newUser.Password, newUser.UserType)

	token, err := LoginUser(newUser.Email, "wrongPassword")

	assert.Empty(t, token)
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

	registeredUser, err1 := RegisterUser(newUser.Name, newUser.Email, newUser.Password, newUser.UserType)

	token, err2 := LoginUser(newUser.Email, newUser.Password)
	claims, err3 := AuthenticateUser(token)

	assert.NoError(t, err1)
	assert.NotNil(t, registeredUser)
	assert.NoError(t, err2)
	assert.NotEmpty(t, token)
	assert.NoError(t, err3)
	assert.NotNil(t, claims)
	assert.Equal(t, registeredUser.ID, claims.Id)
	assert.Equal(t, newUser.Email, claims.Email)
	assert.Equal(t, newUser.UserType, claims.UserType)
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

	RegisterUser(newUser.Name, newUser.Email, newUser.Password, newUser.UserType)
	LoginUser(newUser.Email, newUser.Password)

	claims, err := AuthenticateUser("invalid_token")

	assert.Nil(t, claims)
	assert.Error(t, err)
}
