package productclient

import (
	proto "api-gateway/proto/product"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ProductServiceClient proto.ProductServiceClient
)

func InitProductClient() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatal("Failed to connect to the product microservice")
	}
	defer conn.Close()

	log.Printf("gRPC Client connected to Product Microservice")
	ProductServiceClient = proto.NewProductServiceClient(conn)
}
