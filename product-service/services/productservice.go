package services

import (
	"errors"
	"product-service/models"
)

func CreateProduct(newProduct *models.Product) (*models.Product, error) {
	if newProduct == nil {
		return nil, errors.New("invalid product")
	}

	return models.CreateProduct(newProduct)
}

func GetAllProducts() ([]models.Product, error) {
	return models.GetAllProducts()
}
