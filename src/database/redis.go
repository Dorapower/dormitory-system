package database

import (
	"github.com/go-redis/redis"
	"log"
	"os"
)

var RedisDb *redis.Client

func ConnectRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	RedisDb = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	_, err := RedisDb.Ping().Result()
	if err != nil {
		log.Fatalln("Error when trying to connect to the redis server: " + err.Error())
	}
}
