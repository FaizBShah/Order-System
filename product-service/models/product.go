package models

import "gorm.io/gorm"

var (
	db *gorm.DB
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"column:name;unique"`
	Description string  `gorm:"column:description"`
	Price       float64 `gorm:"column:price"`
	Quantity    int     `gorm:"column:quantity"`
}

func InitProductModel(dbInstance *gorm.DB) {
	db = dbInstance
	db.AutoMigrate(&Product{})
}
