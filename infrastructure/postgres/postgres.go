package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateConnection(connectionString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	pool.Config().MaxConns = 50
	return pool, nil
}
