package restaurant

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"restaurant/internal/restaurant_menu/repositories/productrepository"
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	"restaurant/pkg/data/slice"
)

type ProductServiceGRPCServer struct {
	productRepository productrepository.ProductRepository
}

func NewProductServiceGRPCServer(productRepository productrepository.ProductRepository) *ProductServiceGRPCServer {
	return &ProductServiceGRPCServer{
		productRepository: productRepository,
	}
}

func (s *ProductServiceGRPCServer) CreateProduct(ctx context.Context, in *restaurant.CreateProductRequest) (*restaurant.CreateProductResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err := s.productRepository.Create(&productrepository.Product{
		Name:        in.Name,
		Description: in.Description,
		Type:        in.Type,
		Weight:      in.Weight,
		Price:       in.Price,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &restaurant.CreateProductResponse{}, nil
}

func (s *ProductServiceGRPCServer) GetProductList(ctx context.Context, in *restaurant.GetProductListRequest) (*restaurant.GetProductListResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	products, err := s.productRepository.List([]string{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &restaurant.GetProductListResponse{
		Result: slice.Map[*productrepository.Product, *restaurant.Product](products, func(index int, value *productrepository.Product, slice []*productrepository.Product) *restaurant.Product {
			return &restaurant.Product{
				Uuid:        value.Uuid,
				Name:        value.Name,
				Description: value.Description,
				Type:        value.Type,
				Weight:      value.Weight,
				Price:       value.Price,
				CreatedAt:   timestamppb.New(value.CreatedAt),
			}
		}),
	}, nil
}

func (s *ProductServiceGRPCServer) GetProductListByUuids(ctx context.Context, in *restaurant.GetProductListByUuidsRequest) (*restaurant.GetProductListByUuidsResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	products, err := s.productRepository.List(in.Uuids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &restaurant.GetProductListByUuidsResponse{
		Result: slice.Map[*productrepository.Product, *restaurant.Product](products, func(index int, value *productrepository.Product, slice []*productrepository.Product) *restaurant.Product {
			return &restaurant.Product{
				Uuid:        value.Uuid,
				Name:        value.Name,
				Description: value.Description,
				Type:        value.Type,
				Weight:      value.Weight,
				Price:       value.Price,
				CreatedAt:   timestamppb.New(value.CreatedAt),
			}
		}),
	}, nil
}
