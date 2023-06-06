package modelsgorm

import (
	gormAddon "restaurant/pkg/database/gorm"
)

type Office struct {
	gormAddon.Model
	Name    string
	Address string
	Users   []*User `gorm:"foreignKey:OfficeID"`
}
