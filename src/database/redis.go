package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Cache *redis.Client
var CachChannel chan string


func SetupRedis()  {
	Cache = redis.NewClient(&redis.Options{
		Addr : "redis:6379",
		DB : 0,
	})
}

func SetupCacheChannel() {
	CachChannel = make(chan string)

	go func(ch chan string) {
		for {
			time.Sleep(3 * time.Second)

			key := <-ch

			Cache.Del(context.Background(), key)

			fmt.Printf("Cache Deleted %s", key)
		}
	}(CachChannel)
}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CachChannel <- key
	}
}
