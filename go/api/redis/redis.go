package redis

import (
	"lib/logger"
	"context"
	"github.com/go-redis/redis/v8"
)
////////////////////////////////////////////////
//Databases
// 0 : testing
// 1 : user session token

////////////////////////////////////////////////
//Structures and Variables

//Redis client
var redisClient *redis.Client

//context of the program for instance of redis running
var ctx = context.Background()

func GetRedisClient() redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "archepassword1234",
			DB:       0, // use default DB
		})
		err := redisClient.Set(ctx, "verify", true, 0).Err()
		if err != nil {
			logger.Error.Fatal(err)
		}
		val, err := redisClient.Get(ctx, "verify").Result()
		if err != nil {
			logger.Error.Fatal(err)
		}
		_ = val
		return *redisClient

	}
	val, err := redisClient.Get(ctx, "verify").Result()
	if err != nil {
		redisClient = nil
		GetRedisClient()
	}
	_ = val
	return *redisClient
}

