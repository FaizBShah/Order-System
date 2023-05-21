package routes

import (
	"api-gateway/handlers/orderhandler"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {
	router.HandleFunc("/", orderhandler.GetAllOrders).Methods("GET")
}