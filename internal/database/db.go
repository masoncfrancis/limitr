package database

import (
	"context"
	"github.com/BeehiveBroadband/limitr/internal/config"
	"github.com/redis/go-redis/v9"
)

// CreateDbConn creates a new Redis client and returns the context and client
func CreateDbConn() (context.Context, *redis.Client) {
	dbCtx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetRedisAddr(),
		Password: config.GetRedisPassword(),
		DB:       config.GetRedisDb(),
	})

	return dbCtx, rdb
}
