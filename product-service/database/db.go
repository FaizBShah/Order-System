package database

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func (d *DB) Create() error {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return errors.New("failed to connect to database")
	}

	d = &DB{db}

	return nil
}

func (d *DB) Close() error {
	sqlDb, err := d.DB.DB()

	if err != nil {
		return errors.New("failed to close the database connection")
	}

	sqlDb.Close()

	return nil
}

