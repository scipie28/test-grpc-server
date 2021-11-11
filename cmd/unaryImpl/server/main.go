package main

import (
	"log"
	"net"

	storeAPI "github.com/scipie28/test-grpc-server/internal/api/store_v1"
	"github.com/scipie28/test-grpc-server/pkg/store_v1"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	store_v1.RegisterOrderManagementServer(s, &storeAPI.Server{})
	log.Printf("Starting gRPC listener on port " + port)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
