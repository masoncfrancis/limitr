package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

// CheckIp checks if the IP has made more requests than the rate limit in the time window
// Returns true if the IP has made more requests than the rate limit, false otherwise
// Returns an error if there was an error checking the IP
func CheckIp(rdb *redis.Client, ip string, dbCtx context.Context, timeWindow int, rateLimit int) (bool, error) {
	now := time.Now().Unix()
	startTime := now - int64(timeWindow)

	// Generate a unique identifier for the request
	uid := uuid.New().String()

	// Use the unique identifier as the member in the sorted set
	_, err := rdb.ZAdd(dbCtx, ip, redis.Z{Score: float64(now), Member: strconv.FormatInt(now, 10) + ":" + uid}).Result()

	count, err := rdb.ZCount(dbCtx, ip, strconv.FormatInt(startTime, 10), strconv.FormatInt(now, 10)).Result()
	if err != nil {
		return false, err
	}

	_, err = rdb.ZRemRangeByScore(dbCtx, ip, "-inf", strconv.FormatInt(startTime-1, 10)).Result()

	if count > int64(rateLimit) {
		return true, nil
	}

	return false, nil
}
