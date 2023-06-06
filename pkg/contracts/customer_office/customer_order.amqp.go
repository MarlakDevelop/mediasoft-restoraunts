package customer

import (
	"time"
)

type OrderItemJSON struct {
	ProductUuid string `json:"product_uuid"`
	Count       int32  `json:"count"`
}

type OrderJSON struct {
	Uuid      string           `json:"uuid"`
	UserUuid  string           `json:"user_uuid"`
	Items     []*OrderItemJSON `json:"items"`
	CreatedAt time.Time        `json:"created_at"`
}
