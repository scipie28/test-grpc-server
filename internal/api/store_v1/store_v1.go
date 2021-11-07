package store_v1

import (
	"github.com/scipie28/test-grpc-server/pkg/store_v1"
)

// Сервер используется для реализации store_v1/product_info
type server struct{
	productMap map[string]*store_v1.Product
}

