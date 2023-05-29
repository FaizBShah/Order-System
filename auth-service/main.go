package main

import (
	"auth-service/database"
	proto "auth-service/proto/auth"
	"auth-service/server"
	"log"
	"net"

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
	proto.RegisterAuthServiceServer(grpcServer, &server.GRPCServer{})

	log.Printf("Server started at port 9003")
	lis, err := net.Listen("tcp", ":9003")

	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

	defer lis.Close()
	defer grpcServer.Stop()
}
