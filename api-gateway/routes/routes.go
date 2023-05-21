package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	RegisterProductRoutes(router.PathPrefix("/products").Subrouter())
	RegisterOrderRoutes(router.PathPrefix("/orders").Subrouter())

	return router
}