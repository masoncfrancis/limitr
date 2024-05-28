package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func GetAndIncrementIPValue(rdb *redis.Client, ip string, dbCtx context.Context) (int, error) {
	// Get the current value
	val, err := rdb.Get(dbCtx, ip).Result()
	if err != nil {
		if err == redis.Nil {
			// Key does not exist, set initial value to 0
			if err := rdb.Set(dbCtx, ip, 0, 0).Err(); err != nil {
				return 0, err
			}
			val = "0"
		} else {
			return 0, err
		}
	}

	// Convert the value to an integer
	currentValue, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	// Increment the value in Redis
	if err := rdb.Incr(dbCtx, ip).Err(); err != nil {
		return 0, err
	}

	// Return the original value
	return currentValue, nil
}
