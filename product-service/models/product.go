package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `gorm:"column:name;unique"`
	Description string `gorm:"column:description"`
	Price float64 `gorm:"column:price"`
	Quantity int `gorm:"column:quantity"`
}