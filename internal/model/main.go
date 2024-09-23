package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Dao struct {
	pgConn *pgx.Conn
}

func NewDao(pgConn *pgx.Conn) *Dao {
	return &Dao{
		pgConn: pgConn,
	}
}

var ctx = context.Background()
