package database

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/sb-cloud/player-ms-api/internal/config"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func New(cfg *config.Config) (*DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{db}, nil
}

func (db *DB) Close(ctx context.Context) error {
	return db.DB.Close()
}

func (db *DB) Ping(ctx context.Context) error {
	return db.DB.PingContext(ctx)
}
