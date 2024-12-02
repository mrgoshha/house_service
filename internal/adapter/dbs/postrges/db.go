package postrges

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(cfg *Config) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)

	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf(`connection  %w`, err)
	}

	return db, nil
}
