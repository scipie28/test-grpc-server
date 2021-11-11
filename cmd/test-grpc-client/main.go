package main

import (
	"context"
	"log"
	"time"

	"github.com/scipie28/test-grpc-server/pkg/store_v1"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	//test := []int64{2: 3, 8, 0: 9}
	//fmt.Println(test)
	//return

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := store_v1.NewProductInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.AddProduct(ctx, &store_v1.Product{
		Name:        "Apple iPhone 11",
		Description: "Meet Apple iPhone 11. All-new dual-camera system with\n Ultra Wide and Night mode.",
		Price:       1000.0,
	})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
		return
	}

	log.Printf("Product ID: %s added successfully", res.Value)

	product, err := c.GetProduct(ctx, &store_v1.ProductID{
		Value: res.Value,
	})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
		return
	}

	log.Printf("Product: %v", product.String())

}
