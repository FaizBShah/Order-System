package services

import (
	"errors"
	"product-service/models"
)

func CreateProduct(name string, description string, price float64, quantity int32) (*models.Product, error) {
	newProduct := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}

	return models.CreateProduct(&newProduct)
}

func GetAllProducts() ([]models.Product, error) {
	return models.GetAllProducts()
}

func GetProduct(id int32) (*models.Product, error) {
	return models.GetProduct(id)
}

func DeleteProduct(id int32) error {
	return models.DeleteProduct(id)
}

func AddProducts(id int32, quantity int32) (*models.Product, error) {
	if quantity <= 0 {
		return nil, errors.New("quantity added cannot be less than 0")
	}

	return models.AddProducts(id, quantity)
}

func RemoveProducts(id int32, quantity int32) (*models.Product, error) {
	if quantity <= 0 {
		return nil, errors.New("quantity removed cannot be less than 0")
	}

	return models.RemoveProducts(id, quantity)
}

func UpdateProducts(ids []int64, quantities []int32) error {
	return models.UpdateProducts(ids, quantities)
}
