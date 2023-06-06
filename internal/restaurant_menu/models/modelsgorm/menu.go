package modelsgorm

import (
	gormAddon "restaurant/pkg/database/gorm"
	"time"
)

type Menu struct {
	gormAddon.Model
	OnDate          time.Time
	OpeningRecordAt time.Time
	ClosingRecordAt time.Time
	Items           []*MenuItem `gorm:"foreignKey:MenuID"`
}

type MenuItem struct {
	gormAddon.Model
	MenuID    string
	ProductID string
	Product   *Product `gorm:"foreignKey:ProductID"`
}
