package productrepositorygorm

import (
	"gorm.io/gorm"
	"restaurant/internal/restaurant_menu/models/modelsgorm"
	"restaurant/internal/restaurant_menu/repositories/productrepository"
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	"restaurant/pkg/data/slice"
)

type Repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) List(uuids []string) ([]*productrepository.Product, error) {
	var result *gorm.DB
	products := make([]*modelsgorm.Product, 0)
	if len(uuids) > 0 {
		result = r.DB.Where("id IN ?", uuids).Order("created_at DESC").Find(&products)
	} else {
		result = r.DB.Order("created_at DESC").Find(&products)
	}
	if result.Error != nil {
		return []*productrepository.Product{}, result.Error
	}
	return slice.Map[*modelsgorm.Product, *productrepository.Product](products, func(index int, value *modelsgorm.Product, slice []*modelsgorm.Product) *productrepository.Product {
		return &productrepository.Product{
			Uuid:        value.ID,
			Name:        value.Name,
			Description: value.Description,
			Type:        restaurant.ProductType(value.Type),
			Weight:      value.Weight,
			Price:       value.Price,
			CreatedAt:   value.CreatedAt,
		}
	}), nil
}

func (r Repository) Create(product *productrepository.Product) (*productrepository.Product, error) {
	localProduct := &modelsgorm.Product{
		Name:        product.Name,
		Description: product.Description,
		Type:        int32(product.Type),
		Weight:      product.Weight,
		Price:       product.Price,
	}
	result := r.DB.Create(&localProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	product = &productrepository.Product{
		Uuid:        localProduct.ID,
		Name:        localProduct.Name,
		Description: localProduct.Description,
		Type:        restaurant.ProductType(localProduct.Type),
		Weight:      localProduct.Weight,
		Price:       localProduct.Price,
		CreatedAt:   localProduct.CreatedAt,
	}
	return product, nil
}
