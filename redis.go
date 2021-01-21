package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	RunExample()
}

var redisClient *redis.Client
var ctx = context.Background()

func ConnectRedis() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "34.87.152.186:6379",
		Password: "QFrVwg8DU7unQSFQ37ie",
		DB:       1,
	})

	return redisClient
}

func RunExample() {

	rdb := ConnectRedis()

	set, err := rdb.SetNX(ctx, "lock", "value", 10*time.Second).Result()
	fmt.Println(err)
	fmt.Println(set)

	set, err = rdb.SetNX(ctx, "lock", "value", 10*time.Second).Result()
	fmt.Println(set)

	set, err = rdb.SetNX(ctx, "lock", "value", 10*time.Second).Result()
	fmt.Println(set)

	time.Sleep(11 * time.Second)
	set, err = rdb.SetNX(ctx, "lock", "value", 10*time.Second).Result()
	fmt.Println(set)

}
