package services

import (
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
