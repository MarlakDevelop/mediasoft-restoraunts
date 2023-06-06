package orderrepositorysqlx

import (
	"github.com/jmoiron/sqlx"
	"restaurant/internal/statistics_statistics/repositories/orderrepository"
	"restaurant/pkg/data/slice"
	"time"
)

type OrderedProduct struct {
	Uuid        string    `db:"product_uuid"`
	Name        string    `db:"product_name"`
	Price       float64   `db:"product_price"`
	Count       int64     `db:"order_item_count"`
	ProductType int32     `db:"product_type"`
	CreatedAt   time.Time `db:"order_created_at"`
}

type Repository struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetProfit(startTime time.Time, endTime time.Time) (float64, error) {
	var profit float64
	r.DB.QueryRow("SELECT SUM(product_price * order_item_count) FROM ordered_products WHERE order_created_at BETWEEN $1 AND $2", startTime, endTime).Scan(&profit)
	return profit, nil
}

func (r *Repository) GetOrderedProducts(startTime time.Time, endTime time.Time, productType int32) ([]*orderrepository.OrderedProduct, error) {
	orderedProducts := make([]*OrderedProduct, 0)
	err := r.DB.Select(&orderedProducts,
		`
		SELECT product_uuid, product_name, SUM(order_item_count) AS order_item_count, product_type
			FROM ordered_products
		WHERE (order_created_at BETWEEN $1 AND $2) AND product_type = $3
		GROUP BY product_uuid, product_name, product_type
		ORDER BY SUM(order_item_count) DESC
		`, startTime, endTime, productType)
	if err != nil {
		return nil, err
	}
	return slice.Map[*OrderedProduct, *orderrepository.OrderedProduct](orderedProducts, func(index int, value *OrderedProduct, slice []*OrderedProduct) *orderrepository.OrderedProduct {
		return &orderrepository.OrderedProduct{
			Uuid:        value.Uuid,
			Name:        value.Name,
			Count:       value.Count,
			ProductType: value.ProductType,
		}
	}), nil
}

func (r *Repository) SaveOrderedProducts(orderedProducts []*orderrepository.OrderedProduct) error {
	orderedProductsStructs := slice.Map[*orderrepository.OrderedProduct, *OrderedProduct](orderedProducts, func(index int, value *orderrepository.OrderedProduct, slice []*orderrepository.OrderedProduct) *OrderedProduct {
		return &OrderedProduct{
			Uuid:        value.Uuid,
			Name:        value.Name,
			Price:       value.Price,
			Count:       value.Count,
			ProductType: value.ProductType,
			CreatedAt:   value.CreatedAt,
		}
	})
	_, err := r.DB.NamedExec(`
	INSERT INTO ordered_products (
	    product_uuid,
		product_name,
		product_price,
		order_item_count,
		product_type,
		order_created_at
	)
	VALUES (
        :product_uuid,
	    :product_name,
		:product_price,
		:order_item_count,
		:product_type,
		:order_created_at
	)`, orderedProductsStructs)
	if err != nil {
		return err
	}
	return nil
}
