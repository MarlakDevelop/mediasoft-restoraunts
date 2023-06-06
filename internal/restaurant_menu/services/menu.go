package restaurant

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"restaurant/internal/restaurant_menu/repositories/menurepository"
	"restaurant/internal/restaurant_menu/repositories/productrepository"
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	"restaurant/pkg/data/slice"
	"time"
)

type MenuServiceGRPCServer struct {
	menuRepository    menurepository.MenuRepository
	productRepository productrepository.ProductRepository
}

func NewMenuServiceGRPCServer(menuRepository menurepository.MenuRepository, productRepository productrepository.ProductRepository) *MenuServiceGRPCServer {
	return &MenuServiceGRPCServer{menuRepository: menuRepository, productRepository: productRepository}
}

func (s *MenuServiceGRPCServer) CreateMenu(ctx context.Context, in *restaurant.CreateMenuRequest) (*restaurant.CreateMenuResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if (in.OpeningRecordAt.Seconds*int64(time.Second) + int64(in.OpeningRecordAt.Nanos)*int64(time.Nanosecond)) > (in.ClosingRecordAt.Seconds*int64(time.Second) + int64(in.ClosingRecordAt.Nanos)*int64(time.Nanosecond)) {
		return nil, status.Error(codes.InvalidArgument, "Menu closing record time must be later than opening record time")
	}
	day := 24 * int64(time.Hour)
	if (in.OnDate.Seconds*int64(time.Second)+int64(in.OnDate.Nanos)*int64(time.Nanosecond))/day <= (in.ClosingRecordAt.Seconds*int64(time.Second)+int64(in.ClosingRecordAt.Nanos)*int64(time.Nanosecond))/day {
		return nil, status.Error(codes.InvalidArgument, "Menu delivery day must be later than day of closing record day")
	}
	_, err := s.menuRepository.GetIntersectingByRecording(in.OpeningRecordAt.AsTime(), in.ClosingRecordAt.AsTime())
	if err == nil {
		return nil, status.Error(codes.InvalidArgument, "Menu intersecting by recording time with other menu")
	}
	_, err = s.menuRepository.GetOnDate(in.OnDate.AsTime())
	if err == nil {
		return nil, status.Error(codes.InvalidArgument, "Menu on this date already exists")
	}
	productUuids := slice.Merge[string](in.Salads, in.Garnishes, in.Meats, in.Soups, in.Drinks, in.Desserts)
	if slice.HasDuplicates[string](productUuids) {
		return nil, status.Error(codes.InvalidArgument, "Menu has duplicates")
	}
	products, err := s.productRepository.List(productUuids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := s.validateMenuSection(in.Salads, products, restaurant.ProductType_PRODUCT_TYPE_SALAD); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.validateMenuSection(in.Garnishes, products, restaurant.ProductType_PRODUCT_TYPE_GARNISH); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.validateMenuSection(in.Meats, products, restaurant.ProductType_PRODUCT_TYPE_MEAT); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.validateMenuSection(in.Soups, products, restaurant.ProductType_PRODUCT_TYPE_SOUP); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.validateMenuSection(in.Drinks, products, restaurant.ProductType_PRODUCT_TYPE_DRINK); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.validateMenuSection(in.Desserts, products, restaurant.ProductType_PRODUCT_TYPE_DESSERT); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err = s.menuRepository.Create(&menurepository.Menu{
		OnDate:          in.OnDate.AsTime(),
		OpeningRecordAt: in.OpeningRecordAt.AsTime(),
		ClosingRecordAt: in.ClosingRecordAt.AsTime(),
		ProductUuids:    productUuids,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &restaurant.CreateMenuResponse{}, nil
}

func (s *MenuServiceGRPCServer) GetMenu(ctx context.Context, in *restaurant.GetMenuRequest) (*restaurant.GetMenuResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	menu, err := s.menuRepository.GetOnDate(in.OnDate.AsTime())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &restaurant.GetMenuResponse{
		Menu: s.fromMenuToProto(menu),
	}, nil
}

func (s *MenuServiceGRPCServer) GetRecordingMenu(ctx context.Context, in *restaurant.GetRecordingMenuRequest) (*restaurant.GetRecordingMenuResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	menu, err := s.menuRepository.GetRecording()
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &restaurant.GetRecordingMenuResponse{
		Menu: s.fromMenuToProto(menu),
	}, nil
}

func (s *MenuServiceGRPCServer) validateMenuSection(uuids []string, products []*productrepository.Product, requiredType restaurant.ProductType) error {
	for _, productUuid := range uuids {
		productI := slice.IndexOf[*productrepository.Product](products, func(index int, value *productrepository.Product, slice []*productrepository.Product) bool {
			return value.Uuid == productUuid
		})
		if productI == -1 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Product(%s) not found", productUuid))
		}
		product := products[productI]
		if product.Type != requiredType {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Product(%s) not according to type %s", productUuid, product.Type.String()))
		}
	}
	return nil
}

func (s *MenuServiceGRPCServer) fromMenuToProto(menu *menurepository.Menu) *restaurant.Menu {
	menuPb := &restaurant.Menu{
		Uuid:            menu.Uuid,
		OnDate:          timestamppb.New(menu.OnDate),
		OpeningRecordAt: timestamppb.New(menu.OpeningRecordAt),
		ClosingRecordAt: timestamppb.New(menu.ClosingRecordAt),
		Salads:          make([]*restaurant.Product, 0),
		Garnishes:       make([]*restaurant.Product, 0),
		Meats:           make([]*restaurant.Product, 0),
		Soups:           make([]*restaurant.Product, 0),
		Drinks:          make([]*restaurant.Product, 0),
		Desserts:        make([]*restaurant.Product, 0),
		CreatedAt:       timestamppb.New(menu.CreatedAt),
	}
	for _, product := range menu.Products {
		switch product.Type {
		case restaurant.ProductType_PRODUCT_TYPE_SALAD:
			menuPb.Salads = append(menuPb.Salads, &restaurant.Product{
				Uuid:        product.Uuid,
				Name:        product.Name,
				Description: product.Description,
				Type:        product.Type,
				Weight:      product.Weight,
				Price:       product.Price,
				CreatedAt:   timestamppb.New(product.CreatedAt),
			})
		case restaurant.ProductType_PRODUCT_TYPE_GARNISH:
			menuPb.Garnishes = append(menuPb.Garnishes, &restaurant.Product{
				Uuid:        product.Uuid,
				Name:        product.Name,
				Description: product.Description,
				Type:        product.Type,
				Weight:      product.Weight,
				Price:       product.Price,
				CreatedAt:   timestamppb.New(product.CreatedAt),
			})
		case restaurant.ProductType_PRODUCT_TYPE_MEAT:
			menuPb.Meats = append(menuPb.Meats, &restaurant.Product{
				Uuid:        product.Uuid,
				Name:        product.Name,
				Description: product.Description,
				Type:        product.Type,
				Weight:      product.Weight,
				Price:       product.Price,
				CreatedAt:   timestamppb.New(product.CreatedAt),
			})
		case restaurant.ProductType_PRODUCT_TYPE_SOUP:
			menuPb.Soups = append(menuPb.Soups, &restaurant.Product{
				Uuid:        product.Uuid,
				Name:        product.Name,
				Description: product.Description,
				Type:        product.Type,
				Weight:      product.Weight,
				Price:       product.Price,
				CreatedAt:   timestamppb.New(product.CreatedAt),
			})
		case restaurant.ProductType_PRODUCT_TYPE_DRINK:
			menuPb.Drinks = append(menuPb.Drinks, &restaurant.Product{
				Uuid:        product.Uuid,
				Name:        product.Name,
				Description: product.Description,
				Type:        product.Type,
				Weight:      product.Weight,
				Price:       product.Price,
				CreatedAt:   timestamppb.New(product.CreatedAt),
			})
		case restaurant.ProductType_PRODUCT_TYPE_DESSERT:
			menuPb.Desserts = append(menuPb.Desserts, &restaurant.Product{
				Uuid:        product.Uuid,
				Name:        product.Name,
				Description: product.Description,
				Type:        product.Type,
				Weight:      product.Weight,
				Price:       product.Price,
				CreatedAt:   timestamppb.New(product.CreatedAt),
			})
		}
	}
	return menuPb
}
