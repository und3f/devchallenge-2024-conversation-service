package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Dao struct {
	pg *pgxpool.Pool
}

func NewDao(pgPool *pgxpool.Pool) *Dao {
	return &Dao{
		pg: pgPool,
	}
}

var ctx = context.Background()
