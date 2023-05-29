package authclient

import (
	proto "api-gateway/proto/auth"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	AuthServiceClient proto.AuthServiceClient
)

func InitAuthClient() {
	conn, err := grpc.Dial("localhost:9003", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatal("Failed to connect to the Auth Microservice")
	}

	log.Printf("gRPC Client connected to Auth Microservice")
	AuthServiceClient = proto.NewAuthServiceClient(conn)
}
