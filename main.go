package main

import (
	"fmt"
	"query/databases/redisserver"
	"query/logs"
	"log"
	"query/hashs"
	"github.com/gin-gonic/gin"
	"query/settings"
	"query/routers"
)

func main()  {
	redis := redisserver.RedisNewClient(&redisserver.RedisServer{"localhost", "6379", "", 1})
	clientRedis, eRedis := redis.RedisClient()
	defer clientRedis.Close()
	if eRedis != nil {
		log.Printf("%v\n", eRedis)
	}
	f := logs.NewCount("/tmp/sender_redis_success.log")
	//getResult := clientRedis.Get(f.Count().Name())
	clientRedis.Set(f.Count().Name(), f.Count().Size(), 0)
	fmt.Printf("%v\n", clientRedis.Get(f.Count().Name()))

	//elastic := elasticserver.ElasticNewClient(&elasticserver.ElasticServer{"http://", "localhost", "9200"})
	//clientElastic, eElastic := elastic.CreateIndexing("youtube")
	//if eElastic != nil {
	//	log.Printf("Error %v", eElastic)
	//}
	//
	//fmt.Printf("Success\n %v\n %v\n %v\n", clientElastic.Indexing, clientElastic.State, clientElastic.Message)


	password := []byte("secret")
	hash, err := hashs.HashPassWord(password) // ignore error for the sake of simplicity
	if err != nil {
		fmt.Printf("HashPassWord %v\n", err)
	}
	clientRedis.Set(string(password), hash, 0)
	if hashs.CheckPassWordHash(password, hash) {
		fmt.Println(clientRedis.Get(string(password)))
	}
	route := gin.Default()
	route.Use(settings.Cores())
	v1 := route.Group("/v1")
	routers.InitRoutes(v1)
	route.Run(":8090")
}

