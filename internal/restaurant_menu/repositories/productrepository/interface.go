package productrepository

import (
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	"time"
)

type Product struct {
	Uuid        string                 `json:"uuid"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Type        restaurant.ProductType `json:"type"`
	Weight      int32                  `json:"weight"`
	Price       float64                `json:"price"`
	CreatedAt   time.Time              `json:"created_at"`
}

type ProductRepository interface {
	List(uuids []string) ([]*Product, error)
	Create(product *Product) (*Product, error)
}
