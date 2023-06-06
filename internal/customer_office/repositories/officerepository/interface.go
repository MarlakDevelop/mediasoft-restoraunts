package officerepository

import "time"

type Office struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type OfficeRepository interface {
	Get(id string) (*Office, error)
	List() ([]*Office, error)
	Create(name, address string) error
}
