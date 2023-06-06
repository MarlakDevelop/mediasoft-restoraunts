package userrepositorygorm

import (
	"gorm.io/gorm"
	"restaurant/internal/customer_office/models/modelsgorm"
	"restaurant/internal/customer_office/repositories/officerepository"
	"restaurant/internal/customer_office/repositories/userrepository"
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

func (r *Repository) List(officeId string) ([]*userrepository.User, error) {
	users := make([]*modelsgorm.User, 0)
	result := r.DB.Where(&modelsgorm.User{OfficeID: officeId}).Preload("Office").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return slice.Map[*modelsgorm.User, *userrepository.User](users, func(index int, value *modelsgorm.User, slice []*modelsgorm.User) *userrepository.User {
		return &userrepository.User{
			Uuid: value.ID,
			Name: value.Name,
			Office: &officerepository.Office{
				Uuid:      value.Office.ID,
				Name:      value.Office.Name,
				Address:   value.Office.Address,
				CreatedAt: value.Office.CreatedAt,
			},
			CreatedAt: value.CreatedAt,
		}
	}), nil
}

func (r *Repository) Get(id string) (*userrepository.User, error) {
	var user modelsgorm.User
	result := r.DB.Preload("Office").First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userrepository.User{
		Uuid: user.ID,
		Name: user.Name,
		Office: &officerepository.Office{
			Uuid:      user.Office.ID,
			Name:      user.Office.Name,
			Address:   user.Office.Address,
			CreatedAt: user.Office.CreatedAt,
		},
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *Repository) Create(name, officeId string) (*userrepository.User, error) {
	user := &modelsgorm.User{
		Name:     name,
		OfficeID: officeId,
	}
	result := r.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	result = r.DB.Preload("Office").Find(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userrepository.User{
		Uuid: user.ID,
		Name: user.Name,
		Office: &officerepository.Office{
			Uuid:      user.Office.ID,
			Name:      user.Office.Name,
			Address:   user.Office.Address,
			CreatedAt: user.Office.CreatedAt,
		},
		CreatedAt: user.CreatedAt,
	}, nil
}
