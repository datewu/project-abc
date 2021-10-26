package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func Init(ctx context.Context) (func(), error) {
	err := initPool(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	fn := func() {
		pool.Close()
	}
	return fn, nil
}

func initPool(ctx context.Context, url string) error {
	p, err := pgxpool.Connect(ctx, url)
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	pool = p
	return nil
}
