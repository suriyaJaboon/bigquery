package redisserver

import (
	"github.com/go-redis/redis"
)

type IReDisServer interface {
	RedisClient() *redis.Client
}

type RedisServer struct {
	Host string
	Port string
	Password string
	DB int
}
func RedisNewClient(rn *RedisServer) IReDisServer {
	return &RedisServer{rn.Host, rn.Port, rn.Password, rn.DB}
}

func (rs *RedisServer) RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options {
		Addr: rs.Host + ":" + rs.Port,
		Password: rs.Password,
		DB: rs.DB,
	})
	return client
}
