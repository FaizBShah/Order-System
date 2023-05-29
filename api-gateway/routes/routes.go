package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	RegisterProductRoutes(router)
	RegisterOrderRoutes(router)
	RegisterAuthRoutes(router)

	return router
}
