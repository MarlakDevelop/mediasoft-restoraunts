package bootstrap

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"restaurant/internal/statistics_statistics/config"
)

const dbSetUpSQL = `
	CREATE TABLE IF NOT EXISTS ordered_products (
	    id BIGSERIAL PRIMARY KEY,
	    product_uuid TEXT NOT NULL,
		product_name TEXT NOT NULL,
		product_price DOUBLE PRECISION NOT NULL,
		product_type INTEGER NOT NULL,
		order_item_count BIGINT NOT NULL,
		order_created_at TIMESTAMP NOT NULL
	);
`

func InitSqlxDB(cfg config.Config) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", formatConnect(cfg))
}

func MigrateSchema(db *sqlx.DB) error {
	_, err := db.Exec(dbSetUpSQL)
	if err != nil {
		return err
	}
	return nil
}

func formatConnect(cfg config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
}
