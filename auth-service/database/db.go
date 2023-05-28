package database

import (
	"auth-service/models"
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() error {
	var err error

	dsn := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return errors.New("failed to connect to database")
	}

	initModels()

	log.Printf("Database connected...")

	return nil
}

func Close() error {
	sqlDb, err := DB.DB()

	if err != nil {
		return errors.New("failed to close the database connection")
	}

	sqlDb.Close()

	log.Printf("Database closed...")

	return nil
}

func initModels() {
	log.Printf("Initializing models")
	models.InitProductModel(DB)
}
