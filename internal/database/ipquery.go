package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func CheckIp(rdb *redis.Client, ip string, dbCtx context.Context, timeWindow int, rateLimit int) (bool, error) {
	// Get the current time
	now := time.Now().Unix()

	// Get the start time for the time window
	startTime := now - int64(timeWindow)

	// Get the number of requests made by the IP
	count, err := rdb.ZCount(dbCtx, ip, strconv.FormatInt(startTime, 10), strconv.FormatInt(now, 10)).Result()
	if err != nil {
		return false, err
	}

	// If the number of requests is greater than the rate limit, return true
	if count >= int64(rateLimit) {
		return true, nil
	}

	// Add the current request to the sorted set with the time window for the expiration time
	_, err = rdb.ZAdd(dbCtx, ip, redis.Z{Score: float64(now), Member: strconv.FormatInt(now, 10)}).Result()
	if err != nil {
		return false, err
	}

	return false, nil
}
