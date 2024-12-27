package modules

import (
	"context"

	"four/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB() (*DB, error) {
	cfg := config.GetConfig()
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.DBConnString)
	if err != nil {
		return nil, err
	}

	return &DB{
		pool: pool,
	}, nil

}