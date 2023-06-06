package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"restaurant/internal/customer_office/producers/orderproducer"
	"restaurant/internal/customer_office/repositories/menurepository"
	"restaurant/internal/customer_office/repositories/orderrepository"
	"restaurant/internal/customer_office/repositories/userrepository"
	customer "restaurant/pkg/contracts/customer_office"
	"restaurant/pkg/data/slice"
)

type OrderServiceGRPCServer struct {
	orderRepository orderrepository.OrderRepository
	menuRepository  menurepository.MenuRepository
	userRepository  userrepository.UserRepository
	orderProducer   orderproducer.OrderProducer
}

func NewOrderServiceGRPCServer(orderRepository orderrepository.OrderRepository, menuRepository menurepository.MenuRepository, userRepository userrepository.UserRepository, orderProducer orderproducer.OrderProducer) *OrderServiceGRPCServer {
	return &OrderServiceGRPCServer{
		orderRepository: orderRepository,
		menuRepository:  menuRepository,
		userRepository:  userRepository,
		orderProducer:   orderProducer,
	}
}

func (s *OrderServiceGRPCServer) CreateOrder(ctx context.Context, in *customer.CreateOrderRequest) (*customer.CreateOrderResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	saladsUuids := slice.Map[*customer.OrderItem, string](in.Salads, func(index int, value *customer.OrderItem, slice []*customer.OrderItem) string {
		return value.ProductUuid
	})
	drinksUuids := slice.Map[*customer.OrderItem, string](in.Drinks, func(index int, value *customer.OrderItem, slice []*customer.OrderItem) string {
		return value.ProductUuid
	})
	soupsUuids := slice.Map[*customer.OrderItem, string](in.Soups, func(index int, value *customer.OrderItem, slice []*customer.OrderItem) string {
		return value.ProductUuid
	})
	meatsUuids := slice.Map[*customer.OrderItem, string](in.Meats, func(index int, value *customer.OrderItem, slice []*customer.OrderItem) string {
		return value.ProductUuid
	})
	garnishesUuids := slice.Map[*customer.OrderItem, string](in.Garnishes, func(index int, value *customer.OrderItem, slice []*customer.OrderItem) string {
		return value.ProductUuid
	})
	dessertsUuids := slice.Map[*customer.OrderItem, string](in.Desserts, func(index int, value *customer.OrderItem, slice []*customer.OrderItem) string {
		return value.ProductUuid
	})
	if slice.HasDuplicates[string](saladsUuids) ||
		slice.HasDuplicates[string](drinksUuids) ||
		slice.HasDuplicates[string](soupsUuids) ||
		slice.HasDuplicates[string](meatsUuids) ||
		slice.HasDuplicates[string](garnishesUuids) ||
		slice.HasDuplicates[string](dessertsUuids) {
		return nil, status.Error(codes.InvalidArgument, "Order has duplicates")
	}
	menu, err := s.menuRepository.GetRecording()
	if err != nil {
		return nil, status.Error(codes.NotFound, "Menu not found")
	}
	if !slice.IsOneInOther[string](saladsUuids, slice.Map[*menurepository.Product, string](menu.Salads, func(index int, value *menurepository.Product, slice []*menurepository.Product) string {
		return value.Uuid
	})) ||
		!slice.IsOneInOther[string](drinksUuids, slice.Map[*menurepository.Product, string](menu.Drinks, func(index int, value *menurepository.Product, slice []*menurepository.Product) string {
			return value.Uuid
		})) ||
		!slice.IsOneInOther[string](soupsUuids, slice.Map[*menurepository.Product, string](menu.Soups, func(index int, value *menurepository.Product, slice []*menurepository.Product) string {
			return value.Uuid
		})) ||
		!slice.IsOneInOther[string](meatsUuids, slice.Map[*menurepository.Product, string](menu.Meats, func(index int, value *menurepository.Product, slice []*menurepository.Product) string {
			return value.Uuid
		})) ||
		!slice.IsOneInOther[string](garnishesUuids, slice.Map[*menurepository.Product, string](menu.Garnishes, func(index int, value *menurepository.Product, slice []*menurepository.Product) string {
			return value.Uuid
		})) ||
		!slice.IsOneInOther[string](dessertsUuids, slice.Map[*menurepository.Product, string](menu.Desserts, func(index int, value *menurepository.Product, slice []*menurepository.Product) string {
			return value.Uuid
		})) {
		return nil, status.Error(codes.InvalidArgument, "Ordered products is not according to menu")
	}
	_, err = s.userRepository.Get(in.UserUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "User not found")
	}
	order, err := s.orderRepository.Create(&orderrepository.Order{
		UserUuid: in.UserUuid,
		Items: slice.Map[*customer.OrderItem, *orderrepository.OrderItem](
			slice.Merge[*customer.OrderItem](in.Salads, in.Drinks, in.Soups, in.Meats, in.Garnishes, in.Desserts), func(index int, value *customer.OrderItem, slice []*customer.OrderItem) *orderrepository.OrderItem {
				return &orderrepository.OrderItem{
					Count:       value.Count,
					ProductUuid: value.ProductUuid,
				}
			},
		),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.orderProducer.SendOrder(&customer.OrderJSON{
		Uuid:     order.Uuid,
		UserUuid: order.UserUuid,
		Items: slice.Map[*orderrepository.OrderItem, *customer.OrderItemJSON](order.Items, func(index int, value *orderrepository.OrderItem, slice []*orderrepository.OrderItem) *customer.OrderItemJSON {
			return &customer.OrderItemJSON{
				ProductUuid: value.ProductUuid,
				Count:       value.Count,
			}
		}),
		CreatedAt: order.CreatedAt,
	})
	return &customer.CreateOrderResponse{}, nil
}

func (s *OrderServiceGRPCServer) GetActualMenu(ctx context.Context, in *customer.GetActualMenuRequest) (*customer.GetActualMenuResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	menu, err := s.menuRepository.GetRecording()
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &customer.GetActualMenuResponse{
		Salads:    s.convertFromMenuProductsToCustomerProducts(menu.Salads),
		Garnishes: s.convertFromMenuProductsToCustomerProducts(menu.Garnishes),
		Meats:     s.convertFromMenuProductsToCustomerProducts(menu.Meats),
		Soups:     s.convertFromMenuProductsToCustomerProducts(menu.Soups),
		Drinks:    s.convertFromMenuProductsToCustomerProducts(menu.Drinks),
		Desserts:  s.convertFromMenuProductsToCustomerProducts(menu.Desserts),
	}, nil
}

func (s *OrderServiceGRPCServer) GetTotalOrderedProducts(ctx context.Context, in *customer.GetTotalOrderedProductsRequest) (*customer.GetTotalOrderedProductsResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	items, err := s.orderRepository.GetTotalOrderedProducts()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &customer.GetTotalOrderedProductsResponse{
		Products: slice.Map[*orderrepository.OrderItem, *customer.OrderItem](items, func(index int, value *orderrepository.OrderItem, slice []*orderrepository.OrderItem) *customer.OrderItem {
			return &customer.OrderItem{
				Count:       value.Count,
				ProductUuid: value.ProductUuid,
			}
		}),
	}, nil
}

func (s *OrderServiceGRPCServer) GetTotalOrdersByOffices(ctx context.Context, in *customer.GetTotalOrdersByOfficesRequest) (*customer.GetTotalOrdersByOfficesResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	officesWithItems, err := s.orderRepository.GetTotalOrderedProductsWithOffices()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &customer.GetTotalOrdersByOfficesResponse{
		OrderedItemsByOffices: slice.Map[*orderrepository.OrderItemsWithOffice, *customer.OfficeOrderItems](officesWithItems, func(index int, value *orderrepository.OrderItemsWithOffice, officesSlice []*orderrepository.OrderItemsWithOffice) *customer.OfficeOrderItems {
			return &customer.OfficeOrderItems{
				Office: &customer.Office{
					Uuid:      value.Office.Uuid,
					Name:      value.Office.Name,
					Address:   value.Office.Address,
					CreatedAt: timestamppb.New(value.Office.CreatedAt),
				},
				Products: slice.Map[*orderrepository.OrderItem, *customer.OrderItem](value.Items, func(index int, item *orderrepository.OrderItem, itemsSlice []*orderrepository.OrderItem) *customer.OrderItem {
					return &customer.OrderItem{
						Count:       item.Count,
						ProductUuid: item.ProductUuid,
					}
				}),
			}
		}),
	}, nil
}

func (s *OrderServiceGRPCServer) convertFromMenuProductsToCustomerProducts(products []*menurepository.Product) []*customer.Product {
	return slice.Map[*menurepository.Product, *customer.Product](products, func(index int, value *menurepository.Product, slice []*menurepository.Product) *customer.Product {
		return &customer.Product{
			Uuid:        value.Uuid,
			Name:        value.Name,
			Description: value.Description,
			Type:        customer.CustomerProductType(value.Type),
			Weight:      value.Weight,
			Price:       value.Price,
			CreatedAt:   timestamppb.New(value.CreatedAt),
		}
	})
}
