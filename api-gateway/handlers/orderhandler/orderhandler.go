package orderhandler

import (
	"api-gateway/clients/orderclient"
	"api-gateway/dto"
	"api-gateway/middlewares"
	proto "api-gateway/proto/order"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateOrder(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var newOrder dto.CreateOrderRequest
	var products []*proto.Product

	if err := json.NewDecoder(req.Body).Decode(&newOrder); err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	for _, product := range newOrder.Products {
		products = append(products, &proto.Product{
			Id:          int64(product.Id),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
		})
	}

	createdOrder, err := orderclient.OrderServiceClient.CreateOrder(req.Context(), &proto.CreateOrderRequest{
		UserId: req.Context().Value(middlewares.USER_ID).(int64),
		Cart: &proto.Cart{
			Products: products,
		},
	})

	if err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	var dtoProducts []dto.Product

	for _, product := range createdOrder.Cart.Products {
		dtoProducts = append(dtoProducts, dto.Product{
			Id:          int32(product.Id),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
		})
	}

	res := dto.CreateOrderResponse{
		Id:     createdOrder.Id,
		UserId: createdOrder.UserId,
		Cart: dto.Cart{
			Products: dtoProducts,
		},
	}

	respWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(respWriter).Encode(res)
}

func GetAllOrdersByUserId(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var res []dto.CreateOrderResponse

	params := mux.Vars(req)
	userId, err := strconv.Atoi(params["userId"])

	if err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	orders, err := orderclient.OrderServiceClient.GetAllOrdersByUserId(req.Context(), &proto.GetAllOrdersByUserIdRequest{UserId: int64(userId)})

	if err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	for _, order := range orders.Orders {
		var products []dto.Product

		for _, product := range order.Cart.Products {
			products = append(products, dto.Product{
				Id:          int32(product.Id),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Quantity:    product.Quantity,
			})
		}

		res = append(res, dto.CreateOrderResponse{
			Id:     order.Id,
			UserId: order.UserId,
			Cart: dto.Cart{
				Products: products,
			},
		})
	}

	respWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(respWriter).Encode(res)
}
