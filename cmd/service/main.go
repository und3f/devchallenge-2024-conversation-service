package main

import (
	"context"
	"os"

	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/service"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	whisperUrl := os.Getenv("WHISPER_URL")
	service.New(mux.NewRouter(), NewDao(), whisperUrl).Run()
}

func NewDao() *model.Dao {
	pgConn := ConnectPostgres(context.Background())

	return model.NewDao(pgConn)
}

func ConnectPostgres(ctx context.Context) *pgxpool.Pool {
	pgUrl := os.Getenv("POSTGRES_URL")

	conn, err := pgxpool.New(ctx, pgUrl)
	if err != nil {
		panic(err)
	}

	return conn
}
