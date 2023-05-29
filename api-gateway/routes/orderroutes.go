package routes

import (
	"api-gateway/handlers/orderhandler"
	"api-gateway/middlewares"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {
	router.HandleFunc("/orders", middlewares.AuthMiddleware(orderhandler.CreateOrder)).Methods("POST")
	router.HandleFunc("/orders/user/{userId}", middlewares.AuthMiddleware(orderhandler.GetAllOrdersByUserId)).Methods("GET")
}
