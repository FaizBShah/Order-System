package routes

import (
	"api-gateway/handlers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.GetAllProducts).Methods("GET")
}