package modelsgorm

import (
	gormAddon "restaurant/pkg/database/gorm"
)

type Order struct {
	gormAddon.Model
	UserID string
	User   *User        `gorm:"foreignKey:UserID"`
	Items  []*OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gormAddon.Model
	Count     int32
	ProductID string
	OrderID   string
}
