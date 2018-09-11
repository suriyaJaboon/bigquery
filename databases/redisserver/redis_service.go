package redisserver

import "github.com/go-redis/redis"

type IReDisServer interface {
	RedisClient() (*redis.Client, error)
}

type RedisServer struct {
	Host string
	Port string
	Password string
	DB int
}