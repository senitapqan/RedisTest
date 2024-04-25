package app

import "github.com/redis/go-redis/v9"

func RedisConnectionBuilders() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return rdb
} 