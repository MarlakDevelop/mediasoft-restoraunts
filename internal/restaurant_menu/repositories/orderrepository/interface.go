package orderrepository

type Office struct {
	Uuid    string `json:"uuid"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type OrderItem struct {
	Count       int32  `json:"count"`
	ProductUuid string `json:"product_uuid"`
}

type OfficeWithOrderItems struct {
	Office *Office      `json:"office"`
	Items  []*OrderItem `json:"items"`
}

type OrderRepository interface {
	GetTotalOrderedProducts() ([]*OrderItem, error)
	GetOfficesWithTotalOrderedProducts() ([]*OfficeWithOrderItems, error)
}
