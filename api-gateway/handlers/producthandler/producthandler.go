package producthandler

import (
	"api-gateway/clients/productclient"
	"api-gateway/dto"
	proto "api-gateway/proto/product"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
)

func GetAllProducts(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var res []dto.Product

	products, err := productclient.ProductServiceClient.GetAllProducts(req.Context(), &empty.Empty{})

	if err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	for _, product := range products.Products {
		res = append(res, dto.Product{Id: product.Id, Name: product.Name, Description: product.Description, Price: product.Price, Quantity: product.Quantity})
	}

	respWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(respWriter).Encode(res)
}

func CreateProduct(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var newProduct dto.CreateProductRequest

	if err := json.NewDecoder(req.Body).Decode(&newProduct); err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	createdProduct, err := productclient.ProductServiceClient.CreateProduct(req.Context(), &proto.CreateProductRequest{
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		Quantity:    newProduct.Quantity})

	if err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	res := dto.Product{Id: createdProduct.Id, Name: createdProduct.Name, Description: createdProduct.Description, Price: createdProduct.Price, Quantity: createdProduct.Quantity}

	respWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(respWriter).Encode(res)
}
