package server

import (
	"context"
	proto "product-service/proto/product"
	"product-service/services"

	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	proto.UnimplementedProductServiceServer
}

func (s *GRPCServer) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	product, err := services.CreateProduct(req.Name, req.Description, req.Price, req.Quantity)

	if err != nil {
		return nil, err
	}

	return &proto.CreateProductResponse{
		Id:          int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity}, nil
}

func (s *GRPCServer) GetAllProducts(ctx context.Context, _ *emptypb.Empty) (*proto.GetAllProductsResponse, error) {
	products, err := services.GetAllProducts()

	if err != nil {
		return nil, err
	}

	productsResponse := make([]*proto.CreateProductResponse, len(products))

	for idx, product := range products {
		productsResponse[idx] = &proto.CreateProductResponse{
			Id:          int32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity}
	}

	return &proto.GetAllProductsResponse{Products: productsResponse}, nil
}

func (s *GRPCServer) GetProduct(ctx context.Context, req *proto.ProductIdRequest) (*proto.CreateProductResponse, error) {
	product, err := services.GetProduct(req.Id)

	if err != nil {
		return nil, err
	}

	return &proto.CreateProductResponse{
		Id:          int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity}, nil
}

func (s *GRPCServer) DeleteProduct(ctx context.Context, req *proto.ProductIdRequest) (*proto.ProductIdRequest, error) {
	if err := services.DeleteProduct(req.Id); err != nil {
		return nil, err
	}

	return &proto.ProductIdRequest{Id: req.Id}, nil
}

func (s *GRPCServer) AddProducts(ctx context.Context, req *proto.UpdateProductQuantityRequest) (*proto.CreateProductResponse, error) {
	product, err := services.AddProducts(req.Id, req.Quantity)

	if err != nil {
		return nil, err
	}

	return &proto.CreateProductResponse{
		Id:          int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity}, nil
}

func (s *GRPCServer) RemoveProducts(ctx context.Context, req *proto.UpdateProductQuantityRequest) (*proto.CreateProductResponse, error) {
	product, err := services.RemoveProducts(req.Id, req.Quantity)

	if err != nil {
		return nil, err
	}

	return &proto.CreateProductResponse{
		Id:          int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity}, nil
}
