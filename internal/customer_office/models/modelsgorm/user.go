package modelsgorm

import gormAddon "restaurant/pkg/database/gorm"

type User struct {
	gormAddon.Model
	Name     string
	OfficeID string
	Office   *Office `gorm:"foreignKey:OfficeID"`
}
