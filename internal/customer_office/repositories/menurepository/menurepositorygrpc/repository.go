package menurepositorygrpc

import (
	"context"
	"restaurant/internal/customer_office/repositories/menurepository"
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	"restaurant/pkg/data/slice"
	"time"
)

type Repository struct {
	client restaurant.MenuServiceClient
}

func New(client restaurant.MenuServiceClient) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) GetRecording() (*menurepository.Menu, error) {
	response, err := r.client.GetRecordingMenu(context.Background(), &restaurant.GetRecordingMenuRequest{})
	if err != nil {
		return nil, err
	}
	return &menurepository.Menu{
		Uuid:            response.Menu.Uuid,
		OnDate:          response.Menu.OnDate.AsTime(),
		OpeningRecordAt: response.Menu.OpeningRecordAt.AsTime(),
		ClosingRecordAt: response.Menu.ClosingRecordAt.AsTime(),
		Salads:          r.convertProductsFromProtoToModel(response.Menu.Salads),
		Garnishes:       r.convertProductsFromProtoToModel(response.Menu.Garnishes),
		Meats:           r.convertProductsFromProtoToModel(response.Menu.Meats),
		Soups:           r.convertProductsFromProtoToModel(response.Menu.Soups),
		Drinks:          r.convertProductsFromProtoToModel(response.Menu.Drinks),
		Desserts:        r.convertProductsFromProtoToModel(response.Menu.Desserts),
		CreatedAt:       response.Menu.CreatedAt.AsTime(),
	}, nil
}

func (r *Repository) convertProductsFromProtoToModel(products []*restaurant.Product) []*menurepository.Product {
	return slice.Map[*restaurant.Product, *menurepository.Product](products, func(index int, value *restaurant.Product, slice []*restaurant.Product) *menurepository.Product {
		return &menurepository.Product{
			Uuid:        value.Uuid,
			Name:        value.Name,
			Description: value.Description,
			Type:        int32(value.Type),
			Weight:      value.Weight,
			Price:       value.Price,
			CreatedAt:   time.Time{},
		}
	})
}
