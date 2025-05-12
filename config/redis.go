package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Check if the connection is successful
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		fmt.Println("Connected to Redis successfully!")
	}
}

// contactRedisClient, err := redis.InitRedisClient(ctx, viper.GetInt("db.redis.contact_index"), viper.GetInt("db.redis.poolsize"))
// 	if err != nil {
// 		log.Fatalf("Failed to initialize Redis client for contact verification: %v", err)
// 	}
