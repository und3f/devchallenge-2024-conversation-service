package main

import (
	"context"
	"os"

	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/service"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func main() {
	service.New(mux.NewRouter(), NewDao()).Run()
}

func NewDao() *model.Dao {
	pgConn := ConnectPostgres(context.Background())

	return model.NewDao(pgConn)
}

func ConnectPostgres(ctx context.Context) *pgx.Conn {
	pgUrl := os.Getenv("POSTGRES_URL")

	conn, err := pgx.Connect(ctx, pgUrl)
	if err != nil {
		panic(err)
	}

	return conn
}
