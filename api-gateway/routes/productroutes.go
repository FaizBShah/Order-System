package routes

import (
	"api-gateway/handlers/producthandler"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", producthandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", producthandler.CreateProduct).Methods("POST")
}
