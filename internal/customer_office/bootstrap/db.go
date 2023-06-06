package bootstrap

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"restaurant/internal/customer_office/config"
	"restaurant/internal/customer_office/models/modelsgorm"
)

func InitGormDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func MigrateSchema(db *gorm.DB) error {
	err := db.AutoMigrate(&modelsgorm.Office{}, &modelsgorm.User{}, &modelsgorm.Order{}, &modelsgorm.OrderItem{})
	if err != nil {
		return err
	}
	return nil
}
