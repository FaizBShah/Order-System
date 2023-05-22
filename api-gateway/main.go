package main

import (
	"api-gateway/clients"
	"api-gateway/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	clients.InitClients()

	router := routes.NewRouter()

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
