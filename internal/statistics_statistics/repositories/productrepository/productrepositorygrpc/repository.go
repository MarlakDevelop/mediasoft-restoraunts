package productrepositorygrpc

import (
	"context"
	"restaurant/internal/statistics_statistics/repositories/productrepository"
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	statistics "restaurant/pkg/contracts/statistics_statistics"
	"restaurant/pkg/data/slice"
)

type Repository struct {
	client restaurant.ProductServiceClient
}

func New(client restaurant.ProductServiceClient) *Repository {
	return &Repository{
		client: client,
	}
}
func (r *Repository) List(uuids []string) ([]*productrepository.Product, error) {
	result, err := r.client.GetProductListByUuids(context.Background(), &restaurant.GetProductListByUuidsRequest{Uuids: uuids})
	if err != nil {
		return nil, err
	}
	return slice.Map[*restaurant.Product, *productrepository.Product](result.Result, func(index int, value *restaurant.Product, slice []*restaurant.Product) *productrepository.Product {
		return &productrepository.Product{
			Uuid:        value.Uuid,
			Name:        value.Name,
			Description: value.Description,
			Type:        statistics.StatisticsProductType(value.Type),
			Weight:      value.Weight,
			Price:       value.Price,
			CreatedAt:   value.CreatedAt.AsTime(),
		}
	}), nil
}
