package main

import (
	"context"
	"io"
	"log"

	"github.com/scipie28/test-grpc-server/pkg/store_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:50051"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := store_v1.NewOrderManagementClient(conn)

	searchStream, _ := c.SearchOrders(ctx, &wrapperspb.StringValue{Value: "PC"})

	for {
		searchOrder, errOrder := searchStream.Recv()
		if errOrder == io.EOF {
			break
		}

		// обрабатываем другие потенциальные ошибки
		log.Print("Search Result : ", searchOrder)
	}

}
