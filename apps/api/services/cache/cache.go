package cache

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var redisContext = context.Background()

func newRedisClient() *redis.Client {
	redisUrl, ok := os.LookupEnv("REDIS_URL")

	if !ok {
		panic("REDIS_URL is not set")
	}

	options, err := redis.ParseURL(redisUrl)

	if err != nil {
		panic(err)
	}

	return redis.NewClient(options)
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

func Del(key string) error {
	return GetClient().Del(redisContext, key).Err()
}

func HSet(key string, v any, ttl time.Duration) error {
	c := GetClient()
	serialized, err := json.Marshal(v)

	if err != nil {
		return err
	}

	err = c.Set(redisContext, key, serialized, ttl).Err()

	if err != nil {
		return err
	}

	c.Expire(redisContext, key, ttl)

	return nil
}

func HGet[T any](key string) (*T, error) {
	serialized, err := Get(key)

	if err != nil {
		return nil, err
	}

	var v T

	err = json.Unmarshal([]byte(serialized), &v)
	return &v, err
}
