package store_v1

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/scipie28/test-grpc-server/pkg/store_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddProduct реализует store_v1.AddProduct
func (s *Server) AddProduct(_ context.Context, in *store_v1.Product) (*store_v1.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating Product ID", err)
	}

	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*store_v1.Product)
	}

	s.productMap[in.Id] = in
	return &store_v1.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}
