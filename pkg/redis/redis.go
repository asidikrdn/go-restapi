package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func RedisInit() {
	// Client is setting connection with redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})
}

// SetValue sets the key value pair
func SetValue(key string, value string, expiry time.Duration) error {
	errr := redisClient.Set(key, value, expiry).Err()
	if errr != nil {
		return errr
	}
	return nil
}

// get value from redis
func GetValue(key string) (string, error) {
	value, err := redisClient.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
