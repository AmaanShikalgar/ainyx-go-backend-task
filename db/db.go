package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/AmaanShikalgar/ainyx-users-api/config"
)

func New(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	return db, nil
}
