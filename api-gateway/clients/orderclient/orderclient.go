package orderclient

import (
	proto "api-gateway/proto/order"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	OrderServiceClient proto.OrderServiceClient
)

func InitOrderClient() {
	conn, err := grpc.Dial("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatal("Failed to connect to the Order Microservice")
	}

	log.Printf("gRPC Client connected to Order Microservice")
	OrderServiceClient = proto.NewOrderServiceClient(conn)
}
