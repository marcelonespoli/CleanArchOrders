package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type ListOrdersService struct {
	pb.UnimplementedListOrdersServiceServer
	ListOrdersUseCase usecase.ListOrdersUseCase
}

func NewListOrdersService(listOrdersUseCase usecase.ListOrdersUseCase) *ListOrdersService {
	return &ListOrdersService{
		ListOrdersUseCase: listOrdersUseCase,
	}
}

func (s *ListOrdersService) GetOrders(ctx context.Context, in *pb.Blank) (*pb.ListOrdersResponse, error) {
	output, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var pbOrders []*pb.Order
	for _, o := range output {
		order := &pb.Order{
			Id:         o.ID,
			Price:      float32(o.Price),
			Tax:        float32(o.Tax),
			FinalPrice: float32(o.FinalPrice),
		}
		pbOrders = append(pbOrders, order)
	}
	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}
