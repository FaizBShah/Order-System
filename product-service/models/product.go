package models

import (
	"errors"

	"gorm.io/gorm"
)

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

func CreateProduct(newProduct *Product) (*Product, error) {
	if newProduct == nil {
		return nil, errors.New("invalid product")
	}

	if err := db.Create(newProduct).Error; err != nil {
		return nil, errors.New("error in creating a new product")
	}

	return newProduct, nil
}

func GetAllProducts() ([]Product, error) {
	var products []Product

	if err := db.Find(&products).Error; err != nil {
		return nil, errors.New("error while fetching products")
	}

	return products, nil
}
