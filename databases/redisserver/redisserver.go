package redisserver

import (
	"github.com/go-redis/redis"
)

func RedisNewClient(rn *RedisServer) IReDisServer {
	return &RedisServer{rn.Host, rn.Port, rn.Password, rn.DB}
}

func (rs *RedisServer) RedisClient() (*redis.Client, error) {
	client, err := rs.connectToRedis()
	if err != nil {
		defer client.Close()
		return nil, err
	}
	return client, nil
}

func (rs *RedisServer) connectToRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options {
		Addr: rs.Host + ":" + rs.Port,
		Password: rs.Password,
		DB: rs.DB,
	})
	return client, nil
}
