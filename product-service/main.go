package main

import (
	"fmt"
	"log"
	"net"
	"product-service/database"
	proto "product-service/proto/product"
	"product-service/server"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadEnv()

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	grpcServer := grpc.NewServer()
	proto.RegisterProductServiceServer(grpcServer, &server.GRPCServer{})

	fmt.Printf("Server started at port 9001")
	lis, err := net.Listen("tcp", ":9001")

	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

	defer lis.Close()
	defer grpcServer.Stop()
}
