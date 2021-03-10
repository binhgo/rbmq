package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"time"
)

//func main() {
//	RunExample()
//}

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

func RedisLock() {
	rdClient := ConnectRedis()
	pool := goredis.NewPool(rdClient)
	rs := redsync.New(pool)
	mutexKey := "mutex-key"
	mutex := rs.NewMutex(mutexKey)

	if err := mutex.Lock(); err != nil {
		panic(err)
	}

	// Do your work that requires the lock.

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}

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
