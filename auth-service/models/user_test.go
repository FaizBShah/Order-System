package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:yourDbName?mode=memory&cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	InitProductModel(db)

	return db
}

func teardownDatabase(db *gorm.DB) {
	_ = db.Migrator().DropTable(&User{})
	sql, _ := db.DB()
	sql.Close()
}

func TestShouldInitProductModelWork(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	err := db.AutoMigrate(&User{})
	assert.NoError(t, err)
}

func TestShouldCreateUserWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newUser := &User{
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: Admin,
	}

	createdUser, err := CreateUser(newUser)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, int64(1), createdUser.ID)
	assert.Equal(t, newUser.Name, createdUser.Name)
	assert.Equal(t, newUser.Email, createdUser.Email)
	assert.Equal(t, newUser.Password, createdUser.Password)
	assert.Equal(t, newUser.UserType, createdUser.UserType)
}

func TestShouldCreateUserThrowAnErrorIfUserIsNil(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	createdUser, err := CreateUser(nil)

	assert.Error(t, err)
	assert.Nil(t, createdUser)
	assert.Equal(t, "invalid user", err.Error())
}

func TestShouldCreateUserThrowAnErrorIfFailedToCreateANewUser(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newUser := &User{
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: Admin,
	}

	createdUser1, err1 := CreateUser(newUser)
	createdUser2, err2 := CreateUser(newUser)

	assert.NoError(t, err1)
	assert.NotNil(t, createdUser1)
	assert.Equal(t, newUser.Name, createdUser1.Name)
	assert.Equal(t, newUser.Email, createdUser1.Email)
	assert.Equal(t, newUser.Password, createdUser1.Password)
	assert.Equal(t, newUser.UserType, createdUser1.UserType)
	assert.Error(t, err2)
	assert.Nil(t, createdUser2)
	assert.Equal(t, "error in creating a new user", err2.Error())
}

func TestShouldFindUserByEmailWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newUser := &User{
		Name:     "Faiz Bachoo Shah",
		Email:    "test@example.com",
		Password: "testPassword",
		UserType: Admin,
	}

	CreateUser(newUser)

	user, err := FindUserByEmail("test@example.com")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, newUser.Name, user.Name)
	assert.Equal(t, newUser.Email, user.Email)
	assert.Equal(t, newUser.Password, user.Password)
	assert.Equal(t, newUser.UserType, user.UserType)
}

func TestShouldFindUserByEmailThrowErrorIfEmailIsEmpty(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	user, err := FindUserByEmail("")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "email is empty", err.Error())
}

func TestShouldFindUserByEmailReturnEmptyIfUserDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	user, err := FindUserByEmail("test@example.com")

	assert.NoError(t, err)
	assert.Nil(t, user)
}
