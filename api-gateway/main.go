package main

import (
	"api-gateway/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}