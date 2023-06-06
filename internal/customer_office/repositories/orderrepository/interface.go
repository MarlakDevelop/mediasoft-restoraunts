package orderrepository

import (
	"restaurant/internal/customer_office/repositories/officerepository"
	"time"
)

type OrderItem struct {
	Count       int32  `json:"count"`
	ProductUuid string `json:"product_uuid"`
}

type Order struct {
	Uuid      string       `json:"uuid"`
	UserUuid  string       `json:"user_uuid"`
	Items     []*OrderItem `json:"items"`
	CreatedAt time.Time    `json:"created_at"`
}

type OrderItemsWithOffice struct {
	Office *officerepository.Office `json:"office"`
	Items  []*OrderItem             `json:"items"`
}

type OrderRepository interface {
	Create(order *Order) (*Order, error)
	GetTotalOrderedProducts() ([]*OrderItem, error)
	GetTotalOrderedProductsWithOffices() ([]*OrderItemsWithOffice, error)
}
