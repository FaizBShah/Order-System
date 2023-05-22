package main

import (
	"api-gateway/clients"
	"api-gateway/routes"
	"log"
	"net/http"
)

func main() {
	clients.InitClients()

	router := routes.NewRouter()

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
