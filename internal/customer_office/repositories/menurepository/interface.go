package menurepository

import "time"

type Product struct {
	Uuid        string    `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        int32     `json:"type"`
	Weight      int32     `json:"weight"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type Menu struct {
	Uuid            string     `json:"uuid"`
	OnDate          time.Time  `json:"on_date"`
	OpeningRecordAt time.Time  `json:"opening_record_at"`
	ClosingRecordAt time.Time  `json:"closing_record_at"`
	Salads          []*Product `json:"salads"`
	Garnishes       []*Product `json:"garnishes"`
	Meats           []*Product `json:"meats"`
	Soups           []*Product `json:"soups"`
	Drinks          []*Product `json:"drinks"`
	Desserts        []*Product `json:"desserts"`
	CreatedAt       time.Time  `json:"created_at"`
}

type MenuRepository interface {
	GetRecording() (*Menu, error)
}
