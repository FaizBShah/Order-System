package routes

import (
	"api-gateway/handlers/orderhandler"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {
	router.HandleFunc("/orders", orderhandler.GetAllOrders).Methods("GET")
}
