package service

import (
	"context"
	"os"

	"devchallenge.it/conversation/internal/model"
)

func NewDao() *model.Dao {

	conn, err := pgx.Connect(context.Background())
	return model.NewDao(rdb)
}

func ConnectPostgres(ctx context.Context) {
	pgUrl := os.Getenv("POSTGRES_URL")

	conn, err := pgx.Connect(ctx, pgx.ParseConfig(pgUrl))
}
