package store_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/scipie28/test-grpc-server/pkg/store_v1"
)

func (s *Server) GetOrder(ctx context.Context, searchQuery *wrappers.StringValue) (*store_v1.Order, error) {

	return &store_v1.Order{}, nil
}
