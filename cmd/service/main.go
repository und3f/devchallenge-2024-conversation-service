package main

import (
	"context"
	"os"

	"devchallenge.it/conversation/internal/controller"
	"devchallenge.it/conversation/internal/model"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	servConf := model.ServicesConf{
		WhisperUrl: os.Getenv("WHISPER_URL"),
		NlpUrl:     os.Getenv("NLP_URL"),
	}
	controller.New(mux.NewRouter(), NewDao(), servConf).Run()
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
