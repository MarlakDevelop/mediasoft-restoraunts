package orderrepositorygorm

import (
	"gorm.io/gorm"
	"restaurant/internal/customer_office/models/modelsgorm"
	"restaurant/internal/customer_office/repositories/officerepository"
	"restaurant/internal/customer_office/repositories/orderrepository"
	"restaurant/pkg/data/slice"
)

type OfficeOrderItem struct {
	Count     int32
	ProductID string
	OfficeID  string
}
type Repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Create(order *orderrepository.Order) (*orderrepository.Order, error) {
	localOrder := &modelsgorm.Order{
		UserID: order.UserUuid,
		Items: slice.Map[*orderrepository.OrderItem, *modelsgorm.OrderItem](order.Items, func(index int, value *orderrepository.OrderItem, slice []*orderrepository.OrderItem) *modelsgorm.OrderItem {
			return &modelsgorm.OrderItem{
				ProductID: value.ProductUuid,
				Count:     value.Count,
			}
		}),
	}
	result := r.DB.Create(localOrder)
	if result.Error != nil {
		return nil, result.Error
	}
	order = &orderrepository.Order{
		Uuid:     localOrder.ID,
		UserUuid: localOrder.UserID,
		Items: slice.Map[*modelsgorm.OrderItem, *orderrepository.OrderItem](localOrder.Items, func(index int, value *modelsgorm.OrderItem, slice []*modelsgorm.OrderItem) *orderrepository.OrderItem {
			return &orderrepository.OrderItem{
				Count:       value.Count,
				ProductUuid: value.ProductID,
			}
		}),
		CreatedAt: localOrder.CreatedAt,
	}
	return order, nil
}

func (r *Repository) GetTotalOrderedProducts() ([]*orderrepository.OrderItem, error) {
	ordersItems := make([]*modelsgorm.OrderItem, 0)
	result := r.DB.Select("product_id", "SUM(count) AS count").Group("product_id").Order("count desc").Find(&ordersItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return slice.Map[*modelsgorm.OrderItem, *orderrepository.OrderItem](ordersItems, func(index int, value *modelsgorm.OrderItem, slice []*modelsgorm.OrderItem) *orderrepository.OrderItem {
		return &orderrepository.OrderItem{
			Count:       value.Count,
			ProductUuid: value.ProductID,
		}
	}), nil
}

func (r *Repository) GetTotalOrderedProductsWithOffices() ([]*orderrepository.OrderItemsWithOffice, error) {
	offices := make([]*modelsgorm.Office, 0)
	result := r.DB.Order("created_at DESC").Find(&offices)
	if result.Error != nil {
		return nil, result.Error
	}
	ordersItems := make([]*OfficeOrderItem, 0)
	result = r.DB.Raw("SELECT SUM(order_items.count) AS count, product_id, users.office_id AS office_id FROM order_items INNER JOIN orders ON order_items.order_id = orders.id INNER JOIN users ON orders.user_id = users.id GROUP BY order_items.product_id, users.office_id ORDER BY users.office_id, SUM(order_items.count) DESC").Scan(&ordersItems)
	if result.Error != nil {
		return nil, result.Error
	}
	var officesWithOrderItems []*orderrepository.OrderItemsWithOffice
	i := 0
	for _, office := range offices {
		var items []*orderrepository.OrderItem
		for true {
			if i >= len(ordersItems) || ordersItems[i].OfficeID != office.ID {
				break
			}
			items = append(items, &orderrepository.OrderItem{
				Count:       ordersItems[i].Count,
				ProductUuid: ordersItems[i].ProductID,
			})
			i++
		}
		officesWithOrderItems = append(officesWithOrderItems, &orderrepository.OrderItemsWithOffice{
			Office: &officerepository.Office{
				Uuid:      office.ID,
				Name:      office.Name,
				Address:   office.Address,
				CreatedAt: office.CreatedAt,
			},
			Items: items,
		})
	}
	return officesWithOrderItems, nil
}
