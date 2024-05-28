package database

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func CreateDbConn() (context.Context, *redis.Client) {
	dbCtx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return dbCtx, rdb
}
