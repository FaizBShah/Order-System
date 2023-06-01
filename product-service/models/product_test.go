package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:yourDbName?mode=memory&cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	InitProductModel(db)

	return db
}

func teardownDatabase(db *gorm.DB) {
	_ = db.Migrator().DropTable(&Product{})
	sql, _ := db.DB()
	sql.Close()
}

func TestShouldInitProductModelWork(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	err := db.AutoMigrate(&Product{})
	assert.NoError(t, err)
}

func TestShouldCreateProductWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct, err := CreateProduct(newProduct)

	assert.NoError(t, err)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, newProduct.Name, createdProduct.Name)
	assert.Equal(t, newProduct.Description, createdProduct.Description)
	assert.Equal(t, newProduct.Price, createdProduct.Price)
	assert.Equal(t, newProduct.Quantity, createdProduct.Quantity)
}

func TestShouldCreateProductThrowAnErrorIfTryingToPassAnInvalidProuct(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	createdProduct, err := CreateProduct(nil)

	assert.Error(t, err)
	assert.Nil(t, createdProduct)
	assert.Equal(t, "invalid product", err.Error())
}

func TestShouldCreateProductThrowAnErrorIfFailedToCreateANewProduct(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct1, err1 := CreateProduct(newProduct)
	createdProduct2, err2 := CreateProduct(newProduct)

	assert.NoError(t, err1)
	assert.NotNil(t, createdProduct1)
	assert.Equal(t, newProduct.Name, createdProduct1.Name)
	assert.Equal(t, newProduct.Description, createdProduct1.Description)
	assert.Equal(t, newProduct.Price, createdProduct1.Price)
	assert.Equal(t, newProduct.Quantity, createdProduct1.Quantity)
	assert.Error(t, err2)
	assert.Nil(t, createdProduct2)
	assert.Equal(t, "error in creating a new product", err2.Error())
}

func TestGetAllProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct1 := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	newProduct2 := &Product{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	CreateProduct(newProduct1)
	CreateProduct(newProduct2)

	products, err := GetAllProducts()

	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 2)
	assert.NotNil(t, products[0])
	assert.Equal(t, newProduct1.Name, products[0].Name)
	assert.Equal(t, newProduct1.Description, products[0].Description)
	assert.Equal(t, newProduct1.Price, products[0].Price)
	assert.Equal(t, newProduct1.Quantity, products[0].Quantity)
	assert.NotNil(t, products[1])
	assert.Equal(t, newProduct2.Name, products[1].Name)
	assert.Equal(t, newProduct2.Description, products[1].Description)
	assert.Equal(t, newProduct2.Price, products[1].Price)
	assert.Equal(t, newProduct2.Quantity, products[1].Quantity)
}

func TestGetProductShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct, err1 := CreateProduct(newProduct)

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

func TestGetProductShouldThrowAnErrorIfTheProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := GetProduct(100)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestShouldDeleteProductWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct, err1 := CreateProduct(newProduct)

	err2 := DeleteProduct(int32(createdProduct.ID))

	deletedProduct, err3 := GetProduct(int32(createdProduct.ID))

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Error(t, err3)
	assert.NotNil(t, createdProduct)
	assert.Nil(t, deletedProduct)
	assert.Equal(t, gorm.ErrRecordNotFound, err3)
}

func TestShouldDeleteProductThrowAnErrorIfProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	err := DeleteProduct(1001)

	assert.Error(t, err)
	assert.Equal(t, "product with id 1001 does not exist", err.Error())
}

func TestAddProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	CreateProduct(newProduct)

	product, err := AddProducts(1, 5)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, int32(15), product.Quantity)
	assert.Equal(t, int64(1), product.ID)
	assert.Equal(t, newProduct.Name, product.Name)
	assert.Equal(t, newProduct.Description, product.Description)
	assert.Equal(t, newProduct.Price, product.Price)
}

func TestAddProductsShouldThrowAnErrorIfQuantityIsLessThanZero(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := AddProducts(1, -5)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "quantity added cannot be less than 0", err.Error())
}

func TestAddProductsShouldThrowAnErrorIfProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := AddProducts(1, 5)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestRemoveProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    15,
	}

	CreateProduct(newProduct)

	product, err := RemoveProducts(1, 5)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, int32(10), product.Quantity)
	assert.Equal(t, int64(1), product.ID)
	assert.Equal(t, newProduct.Name, product.Name)
	assert.Equal(t, newProduct.Description, product.Description)
	assert.Equal(t, newProduct.Price, product.Price)
}

func TestRemoveProductsShouldThrowErrorIfQuantityIsLessThanZero(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := RemoveProducts(1, -5)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "quantity removed cannot be less than 0", err.Error())
}

func TestRemoveProductsShouldThrowAnErrorIfInsufficentQuantityLeft(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    5,
	}

	CreateProduct(newProduct)

	product, err := RemoveProducts(1, 10)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "too many products to be removed", err.Error())
}

func TestRemoveProductsShouldThrowAnErrorIfProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := RemoveProducts(1, 5)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUpdateProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct1 := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	newProduct2 := &Product{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	ids := []int64{1, 2}
	quantities := []int32{6, 2}

	CreateProduct(newProduct1)
	CreateProduct(newProduct2)

	err1 := UpdateProducts(ids, quantities)
	products, err2 := GetAllProducts()

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 2)
	assert.NotNil(t, products[0])
	assert.Equal(t, newProduct1.Name, products[0].Name)
	assert.Equal(t, newProduct1.Description, products[0].Description)
	assert.Equal(t, newProduct1.Price, products[0].Price)
	assert.Equal(t, int32(4), products[0].Quantity)
	assert.NotNil(t, products[1])
	assert.Equal(t, newProduct2.Name, products[1].Name)
	assert.Equal(t, newProduct2.Description, products[1].Description)
	assert.Equal(t, newProduct2.Price, products[1].Price)
	assert.Equal(t, int32(2), products[1].Quantity)
}

func TestUpdateProductsShouldThrowErrorIfNotAllIdsAreValid(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct1 := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	newProduct2 := &Product{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	ids := []int64{}
	quantities := []int32{}

	CreateProduct(newProduct1)
	CreateProduct(newProduct2)

	err := UpdateProducts(ids, quantities)

	assert.Error(t, err)
	assert.Equal(t, "empty id set passed", err.Error())
}

func TestUpdateProductsShouldThrowErrorIfIdsAndQuantitiesAreOfDifferentLength(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct1 := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	newProduct2 := &Product{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	ids := []int64{1, 2}
	quantities := []int32{6}

	CreateProduct(newProduct1)
	CreateProduct(newProduct2)

	err := UpdateProducts(ids, quantities)

	assert.Error(t, err)
	assert.Equal(t, "items no.s are mismatched", err.Error())
}

func TestUpdateProductsShouldThrowErrorIfSomeIdsAreInvalid(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct1 := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	newProduct2 := &Product{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	ids := []int64{1, 3}
	quantities := []int32{6, 2}

	CreateProduct(newProduct1)
	CreateProduct(newProduct2)

	err := UpdateProducts(ids, quantities)

	assert.Error(t, err)
	assert.Equal(t, "some ids are invalid", err.Error())
}

func TestUpdateProductsShouldBuyMoreItemsThanPrsentInTheInventory(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct1 := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	newProduct2 := &Product{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	ids := []int64{1, 2}
	quantities := []int32{11, 2}

	CreateProduct(newProduct1)
	CreateProduct(newProduct2)

	err := UpdateProducts(ids, quantities)

	assert.Error(t, err)
	assert.Equal(t, "trying to order more items than there is in the inventory", err.Error())
}
