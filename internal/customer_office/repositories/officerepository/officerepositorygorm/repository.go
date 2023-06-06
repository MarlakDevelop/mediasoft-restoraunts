package officerepositorygorm

import (
	"gorm.io/gorm"
	"restaurant/internal/customer_office/models/modelsgorm"
	"restaurant/internal/customer_office/repositories/officerepository"
	"restaurant/pkg/data/slice"
	gormAddon "restaurant/pkg/database/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Get(id string) (*officerepository.Office, error) {
	office := modelsgorm.Office{
		Model: gormAddon.Model{
			ID: id,
		},
	}
	result := r.DB.Take(&office)
	if result.Error != nil {
		return nil, result.Error
	}
	return &officerepository.Office{
		Uuid:      office.ID,
		Name:      office.Name,
		Address:   office.Address,
		CreatedAt: office.CreatedAt,
	}, nil
}

func (r *Repository) List() ([]*officerepository.Office, error) {
	offices := make([]*modelsgorm.Office, 0)
	result := r.DB.Find(&offices)
	if result.Error != nil {
		return nil, result.Error
	}
	return slice.Map[*modelsgorm.Office, *officerepository.Office](offices, func(index int, office *modelsgorm.Office, offices []*modelsgorm.Office) *officerepository.Office {
		return &officerepository.Office{
			Uuid:      office.ID,
			Name:      office.Name,
			Address:   office.Address,
			CreatedAt: office.CreatedAt,
		}
	}), nil
}

func (r *Repository) Create(name, address string) error {
	result := r.DB.Create(&modelsgorm.Office{
		Name:    name,
		Address: address,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
