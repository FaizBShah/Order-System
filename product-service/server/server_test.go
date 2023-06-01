package server

import (
	"context"
	"product-service/models"
	proto "product-service/proto/product"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var server GRPCServer

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:yourDbName?mode=memory&cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	server = GRPCServer{}
	models.InitProductModel(db)

	return db
}

func teardownDatabase(db *gorm.DB) {
	_ = db.Migrator().DropTable(&models.Product{})
	sql, _ := db.DB()
	sql.Close()
}

func Test_Server_ShouldCreateProductWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct, err := server.CreateProduct(context.Background(), product)

	assert.NoError(t, err)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, product.Name, createdProduct.Name)
	assert.Equal(t, product.Description, createdProduct.Description)
	assert.Equal(t, product.Price, createdProduct.Price)
	assert.Equal(t, product.Quantity, createdProduct.Quantity)
}

func Test_Server_ShouldCreateProductThrowAnErrorIfFailedToCreateANewProduct(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct1, err1 := server.CreateProduct(context.Background(), product)
	createdProduct2, err2 := server.CreateProduct(context.Background(), product)

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

func Test_Server_GetAllProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product1 := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	product2 := &proto.CreateProductRequest{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	server.CreateProduct(context.Background(), product1)
	server.CreateProduct(context.Background(), product2)

	products, err := server.GetAllProducts(context.Background(), &emptypb.Empty{})

	assert.NoError(t, err)
	assert.NotEmpty(t, products.Products)
	assert.Len(t, products.Products, 2)
	assert.NotNil(t, products.Products[0])
	assert.Equal(t, product1.Name, products.Products[0].Name)
	assert.Equal(t, product1.Description, products.Products[0].Description)
	assert.Equal(t, product1.Price, products.Products[0].Price)
	assert.Equal(t, product1.Quantity, products.Products[0].Quantity)
	assert.Equal(t, product2.Name, products.Products[1].Name)
	assert.Equal(t, product2.Description, products.Products[1].Description)
	assert.Equal(t, product2.Price, products.Products[1].Price)
	assert.Equal(t, product2.Quantity, products.Products[1].Quantity)
}

func Test_Server_GetProductShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct, err1 := server.CreateProduct(context.Background(), newProduct)

	product, err2 := server.GetProduct(context.Background(), &proto.ProductIdRequest{Id: createdProduct.Id})

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotNil(t, createdProduct)
	assert.NotNil(t, product)
	assert.Equal(t, createdProduct.Id, product.Id)
	assert.Equal(t, createdProduct.Name, product.Name)
	assert.Equal(t, createdProduct.Description, product.Description)
	assert.Equal(t, createdProduct.Price, product.Price)
	assert.Equal(t, createdProduct.Quantity, product.Quantity)
}

func Test_Server_GetProductShouldThrowAnErrorIfTheProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := server.GetProduct(context.Background(), &proto.ProductIdRequest{Id: 100})

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func Test_Server_ShouldDeleteProductWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	createdProduct, err1 := server.CreateProduct(context.Background(), newProduct)

	res, err2 := server.DeleteProduct(context.Background(), &proto.ProductIdRequest{Id: createdProduct.Id})

	deletedProduct, err3 := server.GetProduct(context.Background(), &proto.ProductIdRequest{Id: createdProduct.Id})

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Error(t, err3)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, createdProduct.Id, res.Id)
	assert.Nil(t, deletedProduct)
	assert.Equal(t, gorm.ErrRecordNotFound, err3)
}

func Test_Server_ShouldDeleteProductThrowAnErrorIfProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	res, err := server.DeleteProduct(context.Background(), &proto.ProductIdRequest{Id: 1001})

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, "product with id 1001 does not exist", err.Error())
}

func Test_Server_AddProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	server.CreateProduct(context.Background(), newProduct)

	product, err := server.AddProducts(context.Background(), &proto.UpdateProductQuantityRequest{Id: 1, Quantity: 5})

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, int32(15), product.Quantity)
	assert.Equal(t, int32(1), product.Id)
	assert.Equal(t, newProduct.Name, product.Name)
	assert.Equal(t, newProduct.Description, product.Description)
	assert.Equal(t, newProduct.Price, product.Price)
}

func Test_Server_AddProductsShouldThrowAnErrorIfQuantityIsLessThanZero(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := server.AddProducts(context.Background(), &proto.UpdateProductQuantityRequest{Id: 1, Quantity: -5})

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "quantity added cannot be less than 0", err.Error())
}

func Test_Server_AddProductsShouldThrowAnErrorIfProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := server.AddProducts(context.Background(), &proto.UpdateProductQuantityRequest{Id: 1, Quantity: 5})

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func Test_Server_RemoveProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    15,
	}

	server.CreateProduct(context.Background(), newProduct)

	product, err := server.RemoveProducts(context.Background(), &proto.UpdateProductQuantityRequest{Id: 1, Quantity: 5})

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, int32(10), product.Quantity)
	assert.Equal(t, int32(1), product.Id)
	assert.Equal(t, newProduct.Name, product.Name)
	assert.Equal(t, newProduct.Description, product.Description)
	assert.Equal(t, newProduct.Price, product.Price)
}

func Test_Server_RemoveProductsShouldThrowErrorIfQuantityIsLessThanZero(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := server.RemoveProducts(context.Background(), &proto.UpdateProductQuantityRequest{Id: 1, Quantity: -5})

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "quantity removed cannot be less than 0", err.Error())
}

func Test_Server_RemoveProductsShouldThrowAnErrorIfInsufficentQuantityLeft(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    5,
	}

	server.CreateProduct(context.Background(), newProduct)

	product, err := server.RemoveProducts(context.Background(), &proto.UpdateProductQuantityRequest{Id: 1, Quantity: 10})

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "too many products to be removed", err.Error())
}

func Test_Server_RemoveProductsShouldThrowAnErrorIfProductDoesNotExist(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	product, err := server.RemoveProducts(context.Background(), &proto.UpdateProductQuantityRequest{Id: 1, Quantity: 5})

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUpdateProductsShouldWorkCorrectly(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct1 := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	newProduct2 := &proto.CreateProductRequest{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	server.CreateProduct(context.Background(), newProduct1)
	server.CreateProduct(context.Background(), newProduct2)

	_, err1 := server.UpdateProducts(context.Background(), &proto.UpdateProductRequest{
		Products: []*proto.UpdateProduct{
			{
				Id:       1,
				Quantity: 6,
			},
			{
				Id:       2,
				Quantity: 2,
			},
		},
	})

	products, err2 := server.GetAllProducts(context.Background(), &emptypb.Empty{})

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotEmpty(t, products.Products)
	assert.Len(t, products.Products, 2)
	assert.NotNil(t, products.Products[0])
	assert.Equal(t, newProduct1.Name, products.Products[0].Name)
	assert.Equal(t, newProduct1.Description, products.Products[0].Description)
	assert.Equal(t, newProduct1.Price, products.Products[0].Price)
	assert.Equal(t, int32(4), products.Products[0].Quantity)
	assert.NotNil(t, products.Products[1])
	assert.Equal(t, newProduct2.Name, products.Products[1].Name)
	assert.Equal(t, newProduct2.Description, products.Products[1].Description)
	assert.Equal(t, newProduct2.Price, products.Products[1].Price)
	assert.Equal(t, int32(2), products.Products[1].Quantity)
}

func TestUpdateProductsShouldThrowErrorIfSomeIdsAreInvalid(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct1 := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	newProduct2 := &proto.CreateProductRequest{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	server.CreateProduct(context.Background(), newProduct1)
	server.CreateProduct(context.Background(), newProduct2)

	res, err := server.UpdateProducts(context.Background(), &proto.UpdateProductRequest{
		Products: []*proto.UpdateProduct{
			{
				Id:       1,
				Quantity: 6,
			},
			{
				Id:       3,
				Quantity: 2,
			},
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, res.Response)
	assert.Nil(t, res.GetSuccessResponse())
	assert.NotNil(t, res.GetErrorResponse())
	assert.Equal(t, "some ids are invalid", res.GetErrorResponse().Message)
}

func TestUpdateProductsShouldBuyMoreItemsThanPrsentInTheInventory(t *testing.T) {
	db := setupDatabase(t)
	defer teardownDatabase(db)

	newProduct1 := &proto.CreateProductRequest{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Quantity:    10,
	}

	newProduct2 := &proto.CreateProductRequest{
		Name:        "Test Product 2",
		Description: "This is a test product 2",
		Price:       8.75,
		Quantity:    4,
	}

	server.CreateProduct(context.Background(), newProduct1)
	server.CreateProduct(context.Background(), newProduct2)

	res, err := server.UpdateProducts(context.Background(), &proto.UpdateProductRequest{
		Products: []*proto.UpdateProduct{
			{
				Id:       1,
				Quantity: 11,
			},
			{
				Id:       2,
				Quantity: 2,
			},
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, res.Response)
	assert.Nil(t, res.GetSuccessResponse())
	assert.NotNil(t, res.GetErrorResponse())
	assert.Equal(t, "trying to order more items than there is in the inventory", res.GetErrorResponse().Message)
}
