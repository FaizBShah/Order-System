package database

import (
	"errors"
	"product-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Create() error {
	var err error

	dsn := ""
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return errors.New("failed to connect to database")
	}

	initModels()

	return nil
}

func Close() error {
	sqlDb, err := DB.DB()

	if err != nil {
		return errors.New("failed to close the database connection")
	}

	sqlDb.Close()

	return nil
}

func initModels() {
	models.InitProductModel(DB)
}
