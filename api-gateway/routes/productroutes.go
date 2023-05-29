package routes

import (
	"api-gateway/handlers/producthandler"
	"api-gateway/middlewares"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", middlewares.AuthMiddleware(producthandler.GetAllProducts)).Methods("GET")
	router.HandleFunc("/products", middlewares.AuthMiddleware(producthandler.CreateProduct)).Methods("POST")
	router.HandleFunc("/products/{id}", middlewares.AuthMiddleware(producthandler.GetProduct)).Methods("GET")
	router.HandleFunc("/products/{id}", middlewares.AuthMiddleware(producthandler.DeleteProduct)).Methods("DELETE")
	router.HandleFunc("/products/add-products", middlewares.AuthMiddleware(producthandler.AddProducts)).Methods("PUT")
	router.HandleFunc("/products/remove-products", middlewares.AuthMiddleware(producthandler.RemoveProducts)).Methods("PUT")
}
