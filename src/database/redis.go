package database

import "github.com/go-redis/redis"

var RedisDb *redis.Client

func ConnectRedis() {
	//TODO Add addr and password
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
	_, err := RedisDb.Ping().Result()
	if err != nil {
		panic(err)
	}
}
