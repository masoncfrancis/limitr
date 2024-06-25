package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

// CheckIp checks if the IP has made more requests than the rate limit in the time window
// Returns true if the IP has made more requests than the rate limit, false otherwise
// Returns an error if there was an error checking the IP
func CheckIp(rdb *redis.Client, ip string, dbCtx context.Context, timeWindow int, rateLimit int) (bool, error) {

	// Get the current time
	now := time.Now().Unix()

	// Get the start time for the time window
	startTime := now - int64(timeWindow)

	// Record the request in the sorted set with the IP as the key and the timestamp as the score
	_, err := rdb.ZAdd(dbCtx, ip, redis.Z{Score: float64(now), Member: strconv.FormatInt(now, 10)}).Result()

	// Get the number of requests made by the IP with a timestamp between the start time and now
	count, err := rdb.ZCount(dbCtx, ip, strconv.FormatInt(startTime, 10), strconv.FormatInt(now, 10)).Result()
	if err != nil {
		return false, err
	}

	// Remove requests that are outside the time window + 2 seconds
	_, err = rdb.ZRemRangeByScore(dbCtx, ip, "-inf", strconv.FormatInt(startTime-2, 10)).Result()

	// If the number of requests is greater than the rate limit, return true
	if count > int64(rateLimit) {
		return true, nil
	}

	// If the number of requests is less than or equal to the rate limit, return false
	return false, nil
}
