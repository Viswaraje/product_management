package services

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func CacheProduct(id string, data string) {
	err := rdb.Set(ctx, id, data, 0).Err()
	if err != nil {
		log.Printf("Error caching product: %v", err)
	}
}

func GetCachedProduct(id string) (string, error) {
	return rdb.Get(ctx, id).Result()
}
