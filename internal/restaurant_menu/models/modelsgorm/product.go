package modelsgorm

import gormAddon "restaurant/pkg/database/gorm"

type Product struct {
	gormAddon.Model
	Name        string
	Description string
	Type        int32
	Weight      int32
	Price       float64
}
