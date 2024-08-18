package repository

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	client    *redis.Client
	redisonce sync.Once
)

func InitRedis() {
	redisonce.Do(func() {
		log.Println("Initializing Redis client...")
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		i, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Could not connect to Redis: %v", err)
		}
		log.Println(i)
		log.Println("Redis Connected!")
	})
}

func GetRedis() *redis.Client {
	if client == nil {
		log.Fatal("Redis connection is not initialized")
	}
	return client
}
