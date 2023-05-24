package services

import (
	"product-service/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	models.InitProductModel(db)

	return db
}

func teardownDatabase(db *gorm.DB) {
	_ = db.Migrator().DropTable(&models.Product{})
	sql, _ := db.DB()
	sql.Close()
}

func Test_Service_ShouldCreateProductWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product := &models.Product{
		ID:          1,
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct, err := CreateProduct(product.Name, product.Description, product.Price, product.Quantity)

	assert.NoError(t, err)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, product.Name, createdProduct.Name)
	assert.Equal(t, product.Description, createdProduct.Description)
	assert.Equal(t, product.Price, createdProduct.Price)
	assert.Equal(t, product.Quantity, createdProduct.Quantity)
}

func Test_Service_ShouldCreateProductThrowAnErrorIfFailedToCreateANewProduct(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product := &models.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct1, err1 := CreateProduct(product.Name, product.Description, product.Price, product.Quantity)
	createdProduct2, err2 := CreateProduct(product.Name, product.Description, product.Price, product.Quantity)

	assert.NoError(t, err1)
	assert.NotNil(t, createdProduct1)
	assert.Equal(t, product.Name, createdProduct1.Name)
	assert.Equal(t, product.Description, createdProduct1.Description)
	assert.Equal(t, product.Price, createdProduct1.Price)
	assert.Equal(t, product.Quantity, createdProduct1.Quantity)
	assert.Error(t, err2)
	assert.Nil(t, createdProduct2)
	assert.Equal(t, "error in creating a new product", err2.Error())
}

func Test_Service_GetAllProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product1 := &models.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	product2 := &models.Product{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	CreateProduct(product1.Name, product1.Description, product1.Price, product1.Quantity)
	CreateProduct(product2.Name, product2.Description, product2.Price, product2.Quantity)

	products, err := GetAllProducts()

	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 2)
	assert.NotNil(t, products[0])
	assert.Equal(t, product1.Name, products[0].Name)
	assert.Equal(t, product1.Description, products[0].Description)
	assert.Equal(t, product1.Price, products[0].Price)
	assert.Equal(t, product1.Quantity, products[0].Quantity)
	assert.Equal(t, product2.Name, products[1].Name)
	assert.Equal(t, product2.Description, products[1].Description)
	assert.Equal(t, product2.Price, products[1].Price)
	assert.Equal(t, product2.Quantity, products[1].Quantity)

}

func Test_Service_GetProductShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &models.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct, err1 := CreateProduct(newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Quantity)

	product, err2 := GetProduct(int32(createdProduct.ID))

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotNil(t, createdProduct)
	assert.NotNil(t, product)
	assert.Equal(t, createdProduct.ID, product.ID)
	assert.Equal(t, createdProduct.Name, product.Name)
	assert.Equal(t, createdProduct.Description, product.Description)
	assert.Equal(t, createdProduct.Price, product.Price)
	assert.Equal(t, createdProduct.Quantity, product.Quantity)
}

func Test_Service_GetProductShouldThrowAnErrorIfTheProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := GetProduct(100)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func Test_Service_ShouldDeleteProductWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &models.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct, err1 := CreateProduct(newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Quantity)

	err2 := DeleteProduct(int32(createdProduct.ID))

	deletedProduct, err3 := GetProduct(int32(createdProduct.ID))

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Error(t, err3)
	assert.NotNil(t, createdProduct)
	assert.Nil(t, deletedProduct)
	assert.Equal(t, gorm.ErrRecordNotFound, err3)
}

func Test_Service_ShouldDeleteProductThrowAnErrorIfProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	err := DeleteProduct(1001)

	assert.Error(t, err)
	assert.Equal(t, "product with id 1001 does not exist", err.Error())
}

func Test_Service_AddProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &models.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	CreateProduct(newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Quantity)

	product, err := AddProducts(1, 5)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, int32(15), product.Quantity)
	assert.Equal(t, uint(1), product.ID)
	assert.Equal(t, newProduct.Name, product.Name)
	assert.Equal(t, newProduct.Description, product.Description)
	assert.Equal(t, newProduct.Price, product.Price)
}

func Test_Service_AddProductsShouldThrowAnErrorIfQuantityIsLessThanZero(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := AddProducts(1, -5)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "quantity added cannot be less than 0", err.Error())
}

func Test_Service_AddProductsShouldThrowAnErrorIfProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := AddProducts(1, 5)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func Test_Service_RemoveProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &models.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    15,
	}

	CreateProduct(newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Quantity)

	product, err := RemoveProducts(1, 5)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, int32(10), product.Quantity)
	assert.Equal(t, uint(1), product.ID)
	assert.Equal(t, newProduct.Name, product.Name)
	assert.Equal(t, newProduct.Description, product.Description)
	assert.Equal(t, newProduct.Price, product.Price)
}

func Test_Service_RemoveProductsShouldThrowErrorIfQuantityIsLessThanZero(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := RemoveProducts(1, -5)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "quantity removed cannot be less than 0", err.Error())
}

func Test_Service_RemoveProductsShouldThrowAnErrorIfInsufficentQuantityLeft(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &models.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    5,
	}

	CreateProduct(newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Quantity)

	product, err := RemoveProducts(1, 10)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "too many products to be removed", err.Error())
}

func Test_Service_RemoveProductsShouldThrowAnErrorIfProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := RemoveProducts(1, 5)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
