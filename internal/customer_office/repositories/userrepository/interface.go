package userrepository

import (
	"restaurant/internal/customer_office/repositories/officerepository"
	"time"
)

type User struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Office    *officerepository.Office
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	List(officeId string) ([]*User, error)
	Get(id string) (*User, error)
	Create(name, officeId string) (*User, error)
}
