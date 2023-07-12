package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

var client *redis.Client
var redisContext = context.Background()

func newRedisClient() *redis.Client {
	redisUrl, ok := os.LookupEnv("REDIS_URL")

	if !ok {
		panic("REDIS_URL is not set")
	}

	opt, err := redis.ParseURL(redisUrl)

	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}

func GetClient() *redis.Client {
	if client == nil {
		client = newRedisClient()
	}
	return client
}

func Get(key string) (string, error) {
	return GetClient().Get(redisContext, key).Result()
}

func Set(key string, value string, exp time.Duration) error {
	return GetClient().Set(redisContext, key, value, exp).Err()
}

func HSet(key string, obj map[string]string) error {
	c := GetClient()
	for k, v := range obj {
		err := c.HSet(redisContext, key, k, v).Err()

		if err != nil {
			return err
		}
	}

	return nil
}

func HGetAll(key string) (map[string]string, error) {
	res, err := GetClient().HGetAll(redisContext, key).Result()
	GetClient().Expire(redisContext, key, ProductTTL)
	return res, err
}
