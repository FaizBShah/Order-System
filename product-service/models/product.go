package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Product struct {
	gorm.Model
	ID          int64   `gorm:"primarykey;AUTO_INCREMENT"`
	Name        string  `gorm:"column:name;unique"`
	Description string  `gorm:"column:description"`
	Price       float64 `gorm:"column:price"`
	Quantity    int32   `gorm:"column:quantity"`
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
		return nil, err
	}

	return products, nil
}

func GetProduct(id int32) (*Product, error) {
	var product *Product

	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteProduct(id int32) error {
	if db.Delete(&Product{}, id).RowsAffected < 1 {
		return fmt.Errorf("product with id %d does not exist", id)
	}

	return nil
}

func AddProducts(id int32, quantity int32) (*Product, error) {
	var product *Product

	if quantity <= 0 {
		return nil, errors.New("quantity added cannot be less than 0")
	}

	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}

	product.Quantity += quantity

	if err := db.Save(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func RemoveProducts(id int32, quantity int32) (*Product, error) {
	var product *Product

	if quantity <= 0 {
		return nil, errors.New("quantity removed cannot be less than 0")
	}

	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}

	if product.Quantity-quantity < 0 {
		return nil, errors.New("too many products to be removed")
	}

	product.Quantity -= quantity

	if err := db.Save(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func UpdateProducts(ids []int64, quantities []int32) error {
	var products []Product

	if len(ids) != len(quantities) {
		return errors.New("items no.s are mismatched")
	}

	if err := db.Where("id IN ?", ids).Find(&products).Error; err != nil {
		return errors.New("failed to update products")
	}

	if len(ids) != len(products) {
		return errors.New("some ids are invalid")
	}

	for idx, quantity := range quantities {
		if quantity > products[idx].Quantity {
			return errors.New("trying to order more items than there is in the inventory")
		}
	}

	for idx, quantity := range quantities {
		products[idx].Quantity -= quantity
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, product := range products {
			if err := db.Save(&product).Error; err != nil {
				return errors.New("failed to update products")
			}
		}

		return nil
	})
}
