package routes

import (
	"api-gateway/handlers/producthandler"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", producthandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", producthandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", producthandler.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", producthandler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products/add-products", producthandler.AddProducts).Methods("PUT")
	router.HandleFunc("/products/remove-products", producthandler.RemoveProducts).Methods("PUT")
}
