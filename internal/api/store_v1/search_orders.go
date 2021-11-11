package store_v1

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/scipie28/test-grpc-server/pkg/store_v1"
)

func (s *Server) SearchOrders(searchQuery *wrappers.StringValue, stream store_v1.OrderManagement_SearchOrdersServer) error {
	err := s.storeService.SearchOrders(searchQuery, stream)

	return err
}
