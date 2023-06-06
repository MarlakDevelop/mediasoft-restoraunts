package orderrepositorygrpc

import (
	"context"
	"restaurant/internal/restaurant_menu/repositories/orderrepository"
	customer "restaurant/pkg/contracts/customer_office"
	"restaurant/pkg/data/slice"
)

type Repository struct {
	client customer.OrderServiceClient
}

func New(client customer.OrderServiceClient) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) GetTotalOrderedProducts() ([]*orderrepository.OrderItem, error) {
	result, err := r.client.GetTotalOrderedProducts(context.Background(), &customer.GetTotalOrderedProductsRequest{})
	if err != nil {
		return nil, err
	}
	return slice.Map[*customer.OrderItem, *orderrepository.OrderItem](result.Products, func(index int, value *customer.OrderItem, slice []*customer.OrderItem) *orderrepository.OrderItem {
		return &orderrepository.OrderItem{
			Count:       value.Count,
			ProductUuid: value.ProductUuid,
		}
	}), nil
}

func (r *Repository) GetOfficesWithTotalOrderedProducts() ([]*orderrepository.OfficeWithOrderItems, error) {
	result, err := r.client.GetTotalOrdersByOffices(context.Background(), &customer.GetTotalOrdersByOfficesRequest{})
	if err != nil {
		return nil, err
	}
	return slice.Map[*customer.OfficeOrderItems, *orderrepository.OfficeWithOrderItems](result.OrderedItemsByOffices, func(index int, value *customer.OfficeOrderItems, officesSlice []*customer.OfficeOrderItems) *orderrepository.OfficeWithOrderItems {
		return &orderrepository.OfficeWithOrderItems{
			Office: &orderrepository.Office{
				Uuid:    value.Office.Uuid,
				Name:    value.Office.Name,
				Address: value.Office.Address,
			},
			Items: slice.Map[*customer.OrderItem, *orderrepository.OrderItem](value.Products, func(index int, item *customer.OrderItem, itemsSlice []*customer.OrderItem) *orderrepository.OrderItem {
				return &orderrepository.OrderItem{
					Count:       item.Count,
					ProductUuid: item.ProductUuid,
				}
			}),
		}
	}), nil
}
