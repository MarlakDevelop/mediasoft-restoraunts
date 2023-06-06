package gorm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m Model) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}
