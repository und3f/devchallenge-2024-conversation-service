package service

import (
	"os"

	"devchallenge.it/conversation/internal/model"
	"github.com/redis/go-redis/v9"
)

func NewDao() *model.Dao {
	redisAddr := os.Getenv("REDIS_ADDR")

	rdb := redis.NewClient(&redis.Options{Addr: redisAddr})
	return model.NewDao(rdb)
}
