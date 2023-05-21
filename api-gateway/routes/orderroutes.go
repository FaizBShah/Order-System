package routes

import (
	"api-gateway/handlers"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.GetAllOrders).Methods("GET")
}