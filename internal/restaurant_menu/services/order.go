package restaurant

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"restaurant/internal/restaurant_menu/repositories/orderrepository"
	"restaurant/internal/restaurant_menu/repositories/productrepository"
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	"restaurant/pkg/data/slice"
)

type OrderServiceGRPCServer struct {
	orderRepository   orderrepository.OrderRepository
	productRepository productrepository.ProductRepository
}

func NewOrderServiceGRPCServer(orderRepository orderrepository.OrderRepository, productRepository productrepository.ProductRepository) *OrderServiceGRPCServer {
	return &OrderServiceGRPCServer{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (s *OrderServiceGRPCServer) GetUpToDateOrderList(ctx context.Context, in *restaurant.GetUpToDateOrderListRequest) (*restaurant.GetUpToDateOrderListResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	orderedProducts, err := s.orderRepository.GetTotalOrderedProducts()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	officesWithOrderedProducts, err := s.orderRepository.GetOfficesWithTotalOrderedProducts()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	productUuids := slice.Merge[string](
		slice.Map[*orderrepository.OrderItem, string](orderedProducts, func(index int, value *orderrepository.OrderItem, slice []*orderrepository.OrderItem) string {
			return value.ProductUuid
		}),
		slice.MergeArray[string](slice.Map[*orderrepository.OfficeWithOrderItems, []string](officesWithOrderedProducts, func(index int, office *orderrepository.OfficeWithOrderItems, offices []*orderrepository.OfficeWithOrderItems) []string {
			return slice.Map[*orderrepository.OrderItem, string](office.Items, func(index int, item *orderrepository.OrderItem, items []*orderrepository.OrderItem) string {
				return item.ProductUuid
			})
		})),
	)
	products, err := s.productRepository.List(productUuids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &restaurant.GetUpToDateOrderListResponse{
		TotalOrders: s.convertOrderedProductsToProto(orderedProducts, products),
		TotalOrdersByCompany: slice.Map[*orderrepository.OfficeWithOrderItems, *restaurant.OrdersByOffice](officesWithOrderedProducts, func(index int, office *orderrepository.OfficeWithOrderItems, offices []*orderrepository.OfficeWithOrderItems) *restaurant.OrdersByOffice {
			return &restaurant.OrdersByOffice{
				CompanyId:     office.Office.Uuid,
				OfficeName:    office.Office.Name,
				OfficeAddress: office.Office.Address,
				Result:        s.convertOrderedProductsToProto(office.Items, products),
			}
		}),
	}, nil
}

func (s *OrderServiceGRPCServer) convertOrderedProductsToProto(orderedItems []*orderrepository.OrderItem, products []*productrepository.Product) []*restaurant.Order {
	return slice.Map[*orderrepository.OrderItem, *restaurant.Order](orderedItems, func(index int, item *orderrepository.OrderItem, items []*orderrepository.OrderItem) *restaurant.Order {
		productI := slice.IndexOf[*productrepository.Product](products, func(index int, value *productrepository.Product, slice []*productrepository.Product) bool {
			return item.ProductUuid == value.Uuid
		})
		var product *productrepository.Product
		if productI != -1 {
			product = products[productI]
		}
		return &restaurant.Order{
			ProductId:   item.ProductUuid,
			ProductName: product.Name,
			Count:       int64(item.Count),
		}
	})
}
