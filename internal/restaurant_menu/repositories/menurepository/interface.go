package menurepository

import (
	"restaurant/internal/restaurant_menu/repositories/productrepository"
	"time"
)

type Menu struct {
	Uuid            string                       `json:"uuid"`
	OnDate          time.Time                    `json:"on_date"`
	OpeningRecordAt time.Time                    `json:"opening_record_at"`
	ClosingRecordAt time.Time                    `json:"closing_record_at"`
	ProductUuids    []string                     `json:"product_uuids"`
	Products        []*productrepository.Product `json:"products"`
	CreatedAt       time.Time                    `json:"created_at"`
}

type MenuRepository interface {
	Create(menu *Menu) (*Menu, error)
	GetOnDate(date time.Time) (*Menu, error)
	GetRecording() (*Menu, error)
	GetIntersectingByRecording(startDate, endDate time.Time) (*Menu, error)
}
