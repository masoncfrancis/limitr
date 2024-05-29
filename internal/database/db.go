package database

import (
	"context"
	"github.com/masoncfrancis/limit/internal/config"
	"github.com/redis/go-redis/v9"
)

func CreateDbConn() (context.Context, *redis.Client) {
	dbCtx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetRedisAddr(),
		Password: config.GetRedisPassword(),
		DB:       0,
	})

	return dbCtx, rdb
}
