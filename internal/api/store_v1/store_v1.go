package store_v1

import "github.com/scipie28/test-grpc-server/internal/api/service/store"

// Server Сервер используется для реализации store_v1/product_info ...
type Server struct {
	//productMap map[string]*store_v1.Product
	storeService store.Server
}
