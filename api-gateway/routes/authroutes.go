package routes

import (
	"api-gateway/handlers/authhandler"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/register", authhandler.RegisterUser).Methods("POST")
	router.HandleFunc("/login", authhandler.LoginUser).Methods("POST")
}
