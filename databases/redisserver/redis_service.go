package redisserver

import (
	"github.com/go-redis/redis"
	"github.com/suriyajaboon/bigquery/structs"
)

type IReDisServer interface {
	RedisClient() *redis.Client
}

type RedisServer struct {
	*structs.RedisServer
}
func RedisNewClient(r *RedisServer) IReDisServer {
	return &RedisServer{&structs.RedisServer{r.Host, r.Port, r.User, r.Pass, r.DB}}
}

func (r *RedisServer) RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options {
		Addr: r.Host + ":" + r.Port,
		Password: r.Pass,
		DB: r.DB,
	})
	return client
}
