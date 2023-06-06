package menurepositorygorm

import (
	"gorm.io/gorm"
	"restaurant/internal/restaurant_menu/models/modelsgorm"
	"restaurant/internal/restaurant_menu/repositories/menurepository"
	"restaurant/internal/restaurant_menu/repositories/productrepository"
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	"restaurant/pkg/data/slice"
	"time"
)

type Repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r Repository) Create(menu *menurepository.Menu) (*menurepository.Menu, error) {
	onDate, err := time.Parse(time.DateOnly, menu.OnDate.Format(time.DateOnly))
	if err != nil {
		return nil, err
	}
	localMenu := &modelsgorm.Menu{
		OnDate:          onDate,
		OpeningRecordAt: menu.OpeningRecordAt,
		ClosingRecordAt: menu.ClosingRecordAt,
		Items: slice.Map[string, *modelsgorm.MenuItem](menu.ProductUuids, func(index int, value string, slice []string) *modelsgorm.MenuItem {
			return &modelsgorm.MenuItem{
				ProductID: value,
			}
		}),
	}
	result := r.DB.Create(&localMenu)
	if result.Error != nil {
		return nil, result.Error
	}
	return nil, nil
}

func (r Repository) GetOnDate(date time.Time) (*menurepository.Menu, error) {
	var localMenu modelsgorm.Menu
	onDate, err := time.Parse(time.DateOnly, date.Format(time.DateOnly))
	if err != nil {
		return nil, err
	}
	result := r.DB.Where("on_date = ?", onDate).Preload("Items.Product").First(&localMenu)
	if result.Error != nil {
		return nil, result.Error
	}
	return &menurepository.Menu{
		Uuid:            localMenu.ID,
		OnDate:          localMenu.OnDate,
		OpeningRecordAt: localMenu.OpeningRecordAt,
		ClosingRecordAt: localMenu.ClosingRecordAt,
		Products: slice.Map[*modelsgorm.MenuItem, *productrepository.Product](localMenu.Items, func(index int, value *modelsgorm.MenuItem, slice []*modelsgorm.MenuItem) *productrepository.Product {
			return &productrepository.Product{
				Uuid:        value.Product.ID,
				Name:        value.Product.Name,
				Description: value.Product.Description,
				Type:        restaurant.ProductType(value.Product.Type),
				Weight:      value.Product.Weight,
				Price:       value.Product.Price,
				CreatedAt:   value.Product.CreatedAt,
			}
		}),
	}, nil
}

func (r Repository) GetRecording() (*menurepository.Menu, error) {
	var localMenu modelsgorm.Menu
	currentTime := time.Now()
	result := r.DB.Where("opening_record_at <= ? AND closing_record_at >= ?", currentTime, currentTime).Preload("Items.Product").First(&localMenu)
	if result.Error != nil {
		return nil, result.Error
	}
	return &menurepository.Menu{
		Uuid:            localMenu.ID,
		OnDate:          localMenu.OnDate,
		OpeningRecordAt: localMenu.OpeningRecordAt,
		ClosingRecordAt: localMenu.ClosingRecordAt,
		Products: slice.Map[*modelsgorm.MenuItem, *productrepository.Product](localMenu.Items, func(index int, value *modelsgorm.MenuItem, slice []*modelsgorm.MenuItem) *productrepository.Product {
			return &productrepository.Product{
				Uuid:        value.Product.ID,
				Name:        value.Product.Name,
				Description: value.Product.Description,
				Type:        restaurant.ProductType(value.Product.Type),
				Weight:      value.Product.Weight,
				Price:       value.Product.Price,
				CreatedAt:   value.Product.CreatedAt,
			}
		}),
	}, nil
}

func (r Repository) GetIntersectingByRecording(startDate, endDate time.Time) (*menurepository.Menu, error) {
	var localMenu modelsgorm.Menu
	result := r.DB.Where("(opening_record_at <= ? AND closing_record_at >= ?) OR (opening_record_at <= ? AND closing_record_at >= ?)", startDate, startDate, endDate, endDate).Preload("Items.Product").First(&localMenu)
	if result.Error != nil {
		return nil, result.Error
	}
	return &menurepository.Menu{
		Uuid:            localMenu.ID,
		OnDate:          localMenu.OnDate,
		OpeningRecordAt: localMenu.OpeningRecordAt,
		ClosingRecordAt: localMenu.ClosingRecordAt,
		Products: slice.Map[*modelsgorm.MenuItem, *productrepository.Product](localMenu.Items, func(index int, value *modelsgorm.MenuItem, slice []*modelsgorm.MenuItem) *productrepository.Product {
			return &productrepository.Product{
				Uuid:        value.Product.ID,
				Name:        value.Product.Name,
				Description: value.Product.Description,
				Type:        restaurant.ProductType(value.Product.Type),
				Weight:      value.Product.Weight,
				Price:       value.Product.Price,
				CreatedAt:   value.Product.CreatedAt,
			}
		}),
	}, nil
}
