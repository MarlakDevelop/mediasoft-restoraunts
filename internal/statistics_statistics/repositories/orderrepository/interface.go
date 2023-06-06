package orderrepository

import (
	"time"
)

type OrderedProduct struct {
	Uuid        string    `json:"uuid"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Count       int64     `json:"count"`
	ProductType int32     `json:"product_type"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderRepository interface {
	GetProfit(startTime time.Time, endTime time.Time) (float64, error)
	GetOrderedProducts(startTime time.Time, endTime time.Time, productType int32) ([]*OrderedProduct, error)
	SaveOrderedProducts([]*OrderedProduct) error
}
