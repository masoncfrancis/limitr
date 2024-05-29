package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func GetAndIncrementIPValue(rdb *redis.Client, ip string, dbCtx context.Context, timeWindow int) (int, error) {
	// Get the current timestamp
	now := time.Now().Unix()

	// Remove all requests older than the time window
	if _, err := rdb.ZRemRangeByScore(dbCtx, ip, "0", strconv.FormatInt(now-int64(timeWindow), 10)).Result(); err != nil {
		return 0, err
	}

	// Add the new request with the current timestamp
	if _, err := rdb.ZAdd(dbCtx, ip, redis.Z{Score: float64(now), Member: now}).Result(); err != nil {
		return 0, err
	}

	// Get the count of requests in the last time window
	count, err := rdb.ZCount(dbCtx, ip, strconv.FormatInt(now-int64(timeWindow), 10), strconv.FormatInt(now, 10)).Result()
	if err != nil {
		return 0, err
	}

	// Return the count
	return int(count), nil
}
