package store_v1

import (
	"fmt"
	"log"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/scipie28/test-grpc-server/pkg/store_v1"
)

var orderMap = map[string]store_v1.Order{
	"1": {
		Id: "1",
		Items: []string{
			"Apple",
			"Orange",
		},
		Description: "Good items",
		Price:       500.0,
		Destination: "For mother",
	},
	"2": {
		Id: "2",
		Items: []string{
			"PC",
			"Smartphone",
		},
		Description: "Good items",
		Price:       2000.0,
		Destination: "For grandmother",
	},
}

func (s *Server) SearchOrders(searchQuery *wrappers.StringValue, stream store_v1.OrderManagement_SearchOrdersServer) error {
	for key, order := range orderMap {
		log.Print(key, order)

		for _, itemStr := range order.Items {
			log.Print(itemStr)

			if strings.Contains(itemStr, searchQuery.Value) {
				// Отправляем подходящие заказы в поток
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}

				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}

	return nil
}
