package routes

import (
	"api-gateway/handlers/producthandler"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/", producthandler.GetAllProducts).Methods("GET")
}