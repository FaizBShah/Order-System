package routes

import (
	"api-gateway/handlers/orderhandler"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {
	router.HandleFunc("/orders", orderhandler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/user/{userId}", orderhandler.GetAllOrdersByUserId).Methods("GET")
}
