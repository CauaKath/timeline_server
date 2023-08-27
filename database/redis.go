package database

import (
	"github.com/go-redis/redis/v8"

	"github.com/cauakath/timeline-server/config"
)

func ConnectionRedisDb(config *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})

	return client
}
