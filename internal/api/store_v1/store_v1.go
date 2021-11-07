package store_v1

import (
	"github.com/scipie28/test-grpc-server/pkg/store_v1"
)

// Server Сервер используется для реализации store_v1/product_info ...
type Server struct {
	productMap map[string]*store_v1.Product
}
