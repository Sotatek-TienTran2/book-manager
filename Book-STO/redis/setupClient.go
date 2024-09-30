package redis

import "github.com/go-redis/redis"

var RDB *redis.Client

func setupClientRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func NewResdisClient() {
	RDB = setupClientRedis()
}
