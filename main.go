package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func main() {
	// Initialize the Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "SkyLab@2020",    // No password set
		DB:       0,                // Default DB
	})

	// Test the connection
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Redis connected:", pong)

	// Setting a key in Redis
	err = rdb.Set(context.Background(), "mykey", "Hello Redis!", 0).Err()
	if err != nil {
		log.Fatalf("Could not set key in Redis: %v", err)
	}
	fmt.Println("Key set successfully")

	// Getting the value of a key from Redis
	val, err := rdb.Get(context.Background(), "mykey").Result()
	if err != nil {
		log.Fatalf("Could not get key from Redis: %v", err)
	}
	fmt.Println("Key value:", val)
}
