package helper

import (
	"context"
	"crud-echo-postgres-redis/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ctx = context.Background()

func createConnection() *redis.Client {
	env, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading app.env file")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     env.CacheSource,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

func Set(key string, value string, ttl int64) {
	rdb := createConnection()

	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			log.Fatalf("Unable to close Redis connection. %v", err)
		}
	}(rdb)

	err := rdb.Set(ctx, key, value, time.Duration(ttl)).Err()
	if err != nil {
		log.Fatalf("Unable to set value in cache")
	}
}

func Get() any {
	rdb := createConnection()

	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			log.Fatalf("Unable to close Redis connection. %v", err)
		}
	}(rdb)

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Fatalf("Unable to get value from cache")

		return nil
	}

	return val
}
