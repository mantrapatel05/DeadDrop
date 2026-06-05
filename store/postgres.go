package store

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func Open(ctx context.Context, databaseURL string) (*DB, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}

func (db *DB) Migrate(ctx context.Context) error {
	query := `
	    CREATE TABLE IF NOT EXISTS drops(
        id TEXT PRIMARY KEY,
        ciphertext TEXT NOT NULL,
        reveal_at TIMESTAMP,
        knock_target INT,
        knock_count INT DEFAULT 0,
        burned BOOLEAN NOT NULL DEFAULT false,
        created_at TIMESTAMP
       );`

	if _, err := db.Pool.Exec(ctx, query); err != nil {
		return fmt.Errorf("create table: %w", err)
	}

	log.Println("database migration complete - table ready")
	return nil
}
