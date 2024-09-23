package model

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Dao struct {
	rdb *redis.Client
}

func NewDao(rdb *redis.Client) *Dao {
	return &Dao{
		rdb: rdb,
	}
}

var ctx = context.Background()
