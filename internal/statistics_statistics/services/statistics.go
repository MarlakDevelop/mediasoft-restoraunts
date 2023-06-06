package statistics

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"restaurant/internal/statistics_statistics/repositories/orderrepository"
	statistics "restaurant/pkg/contracts/statistics_statistics"
	"restaurant/pkg/data/slice"
)

type StatisticsGRPCService struct {
	orderRepository orderrepository.OrderRepository
}

func NewStatisticsGRPCService(orderRepository orderrepository.OrderRepository) *StatisticsGRPCService {
	return &StatisticsGRPCService{orderRepository: orderRepository}
}

func (s *StatisticsGRPCService) GetAmountOfProfit(ctx context.Context, in *statistics.GetAmountOfProfitRequest) (*statistics.GetAmountOfProfitResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	profit, err := s.orderRepository.GetProfit(in.StartDate.AsTime(), in.EndDate.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &statistics.GetAmountOfProfitResponse{Profit: profit}, nil
}

func (s *StatisticsGRPCService) TopProducts(ctx context.Context, in *statistics.TopProductsRequest) (*statistics.TopProductsResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	products, err := s.orderRepository.GetOrderedProducts(in.StartDate.AsTime(), in.EndDate.AsTime(), int32(in.GetProductType()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &statistics.TopProductsResponse{Result: slice.Map[*orderrepository.OrderedProduct, *statistics.Product](products, func(index int, value *orderrepository.OrderedProduct, slice []*orderrepository.OrderedProduct) *statistics.Product {
		return &statistics.Product{
			Uuid:        value.Uuid,
			Name:        value.Name,
			Count:       value.Count,
			ProductType: statistics.StatisticsProductType(value.ProductType),
		}
	})}, nil
}
