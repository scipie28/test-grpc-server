package store_v1

import (
	"context"

	"github.com/scipie28/test-grpc-server/pkg/store_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetProduct реализует ecommerce.GetProduct
func (s *Server) GetProduct(ctx context.Context, in *store_v1.ProductID) (*store_v1.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}
