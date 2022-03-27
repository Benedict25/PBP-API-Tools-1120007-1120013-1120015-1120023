package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		// fmt.Println(err)
	}
}

func SetRedis(rdb *redis.Client, key string, value string, expiration int) {
	err := rdb.Set(ctx, key, value, 0).Err()
	CheckError(err)
}

func GetRedis(rdb *redis.Client, key string) {
	val, err := rdb.Get(ctx, key).Result()

	// if err == redis.Nil {
	// 	fmt.Println(key, "does not exist")
	// }

	CheckError(err)
	fmt.Println("Value:", val)
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	SetRedis(rdb, "key1", "test", 0)

	GetRedis(rdb, "key1")

	// Erase All Keys
	// rdb.FlushDB(ctx)
}
